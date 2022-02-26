package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	"github.com/tharsis/evmos/contracts"
	"github.com/tharsis/evmos/x/synapse/types"
)

// Hooks wrapper struct for erc20 keeper
type Hooks struct {
	k Keeper
}

var _ evmtypes.EvmHooks = Hooks{}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// PostTxProcessing implements EvmHooks.PostTxProcessing
// This hook must be fired off AFTER the ERC20 hooks
// to ensure the token the user is trying to bridge out has been converted to
// a comsos denom token
func (h Hooks) PostTxProcessing(
	ctx sdk.Context,
	from common.Address,
	to *common.Address,
	receipt *ethtypes.Receipt,
) error {
	params := h.k.GetParams(ctx)

	synapseMiddleware := contracts.CosmosRelayedERC20Contract.ABI
	for _, log := range receipt.Logs {

		// If event was not emitted by the middleware contract => ignore
		if log.Address != common.HexToAddress(params.GetEvmMiddlewareAddress()) {
			continue
		}

		if len(log.Topics) < 3 { // TODO: Is this check required?
			continue
		}

		eventID := log.Topics[0] // event ID

		event, err := synapseMiddleware.EventByID(eventID)
		if err != nil {
			// invalid event for ERC20
			continue
		}

		// Ensure we saw a bridge out request from the contract
		if event.Name != types.EvmMiddlewareEventBridgeOut {
			h.k.Logger(ctx).Info("emitted event", "name", event.Name, "signature", event.Sig)
			continue
		}

		// Unpack the event ABI
		bridgeOutEvent, err := synapseMiddleware.Unpack(event.Name, log.Data)
		if err != nil {
			h.k.Logger(ctx).Error("failed to unpack EvmMiddlewareEventBridgeOut event", "error", err.Error())
			continue
		}

		if len(bridgeOutEvent) == 0 {
			continue
		}

		// Get and format data from event
		tokenContractAddress := bridgeOutEvent[0].(common.Address)
		amount, ok := bridgeOutEvent[1].(sdk.Int)
		if !ok {
			panic("revert")
		}
		destChain := bridgeOutEvent[2].(string)
		senderAcc, err := sdk.AccAddressFromHex(bridgeOutEvent[3].(common.Address).Hex())
		if err != nil {
			panic(err)
		}
		receiver := bridgeOutEvent[4].(common.Address).Hex()

		id := h.k.erc20Keeper.GetERC20Map(ctx, tokenContractAddress)

		if len(id) == 0 {
			// no token is registered for the caller contract
			panic("revert to prevent user funds from being trapped in the synapse module")
		}

		pair, found := h.k.erc20Keeper.GetTokenPair(ctx, id)
		if !found {
			panic("revert to prevent user funds from being trapped in the synapse module")
		}

		coin := sdk.NewCoin(pair.Denom, amount)
		// Treansfer coins from the users account to the Synapse Module
		h.k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAcc, types.ModuleName, sdk.NewCoins(coin))

		// Execute the Transfer out
		bridgeOutMsg := &types.MsgBridgeOut{
			Creator: string(h.k.accountKeeper.GetModuleAddress(types.ModuleName)),
			Data: []*types.OutBridgeData{
				{
					Coin:      &coin,
					DestAddr:  receiver,
					DestChain: destChain,
				},
			},
		}
		h.k.HandleBridgeOutFromEVM(ctx, bridgeOutMsg)
	}
	return nil
}
