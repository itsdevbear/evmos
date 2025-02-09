package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/tharsis/ethermint/types"
	erc20types "github.com/tharsis/evmos/x/erc20/types"
	"github.com/tharsis/evmos/x/synapse/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

const (
	EVM    = "evm"
	COSMOS = "cosmos" // TODO put these somewhere nicer
)

/*
 * Handle Bridging Assets onto the Chain: Called by the Synapse Protocol
 */
func (k msgServer) BridgeIn(goCtx context.Context, msg *types.MsgBridgeIn) (*types.MsgBridgeInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//Bridge the Assets in
	if err := k.HandleBridgeIn(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgBridgeInResponse{}, nil
}

func (k Keeper) HandleBridgeIn(ctx sdk.Context, msg *types.MsgBridgeIn) error {
	logger := k.Logger(ctx)
	// For each Bridge Request
	for _, element := range msg.Data {
		logger.Info("incoming bridge", "coins", element.Coin.String(), "dest_addr", element.DestAddr, "dest_env", element.DestEnv)
		// Mint Coins to facilitate
		k.bankKeeper.MintCoins(ctx, types.ModuleName, []sdk.Coin{*element.Coin})
		switch element.DestEnv {
		case EVM:
			// If Destination is Bech32 convert to Hex
			acc := element.DestAddr
			if !common.IsHexAddress(element.DestEnv) {
				addr, err := sdk.AccAddressFromBech32(element.DestAddr)
				if err != nil {
					return err
				}
				acc = k.accountKeeper.GetAccount(ctx, addr).(ethermint.EthAccountI).EthAddress().Hex()
			}
			// Verify Address
			if err := ethermint.ValidateAddress(acc); err != nil {
				return err
			}
			// Send to Hex Account in the EVM
			if _, err := k.erc20Keeper.ConvertCoin(sdk.WrapSDKContext(ctx), &erc20types.MsgConvertCoin{
				Coin:     *element.Coin,
				Receiver: acc,
				Sender:   k.accountKeeper.GetModuleAddress(types.ModuleName).String(),
			}); err != nil {
				return err
			}
		case COSMOS:
			addr, err := sdk.AccAddressFromBech32(element.DestAddr)
			if err != nil {
				return err
			}
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, []sdk.Coin{*element.Coin}); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid destEnv")
		}
	}
	return nil
}

/*
 * Handle Bridging Assets off the Chain: Called by a User who wants to bridge out
 */

func (k msgServer) BridgeOut(goCtx context.Context, msg *types.MsgBridgeOut) (*types.MsgBridgeOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Bridge the Assets Out
	err := k.HandleBridgeOut(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &types.MsgBridgeOutResponse{}, nil
}

func (k Keeper) HandleBridgeOutFromEVM(ctx sdk.Context, msg *types.MsgBridgeOut) (*types.MsgBridgeOutResponse, error) {
	err := k.HandleBridgeOut(ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &types.MsgBridgeOutResponse{}, nil
}

func (k Keeper) HandleBridgeOut(ctx sdk.Context, msg *types.MsgBridgeOut) error {
	// For each Bridge Request
	logger := k.Logger(ctx)
	for _, element := range msg.Data {
		coins := sdk.NewCoins(sdk.NewCoin(element.Coin.Denom, element.Coin.Amount))
		logger.Info("Bridging Out: ", element.Coin.String(), " via ", types.ModuleName)
		k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, coins)

		// if IBC token escrow, else burn
		// TODO: Talk to Caeser to see how we want to handle ibc tokens,
		// maybe its best to transfer it back to its home chain or something
		if !strings.HasPrefix(element.Coin.Denom, "ibc") {
			k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
		}

		// Emit events for Synapse Protocol to watch for
		// TODO, move from event based system => submit txn on Synapse
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeBridgeOut,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.TokenDenom, element.Coin.Denom),
				sdk.NewAttribute(types.TokenAmount, element.Coin.Amount.String()),
				sdk.NewAttribute(types.DestAddr, element.DestAddr),
				sdk.NewAttribute(types.DestChain, element.DestChain),
			),
		)
	}

	return nil
}
