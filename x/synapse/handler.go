package synapse

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/tharsis/evmos/x/synapse/keeper"
	"github.com/tharsis/evmos/x/synapse/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {

		txSigner := msg.GetSigners()[0]

		switch msg := msg.(type) {
		case *types.MsgBridgeIn:
			if err := verifySigner(&ctx, &k, &txSigner, false); err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unauthorized")
			}
			res, err := msgServer.BridgeIn(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBridgeOut:
			if err := verifySigner(&ctx, &k, &txSigner, false); err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unauthorized")
			}
			res, err := msgServer.BridgeOut(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetPendingKmsAddress:
			if err := verifySigner(&ctx, &k, &txSigner, false); err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unauthorized")
			}
			res, err := msgServer.SetPendingKmsAddress(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetKmsAddress:
			if err := verifySigner(&ctx, &k, &txSigner, true); err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Unauthorized")
			}
			res, err := msgServer.UpdateKmsAddress(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func verifySigner(ctx *sdk.Context, k *keeper.Keeper, signer *sdk.AccAddress, isPendingKms bool) error {
	// Verify Authentication

	// Do we want to be auth'd by the pending KMS or the existing KMS
	addr := ""
	if isPendingKms {
		addr = k.GetKmsAddress(*ctx)
	} else {
		addr = k.GetKmsPendingAddress(*ctx)
	}
	requiredSigner, err := sdk.AccAddressFromBech32(addr)

	if err != nil {
		return err
	}

	if !sdk.Address.Equals(requiredSigner, signer) {
		return sdkerrors.Wrap(
			sdkerrors.Error{},
			fmt.Sprintf(
				"%s is not Authenticated", signer,
			),
		)
	}

	// Store Hash of TxBytes for double sign protection
	keccakHash := ethcrypto.Keccak256(ctx.TxBytes())

	if foundExistingHash := k.GetSafetyHash(
		*ctx,
		keccakHash,
	); foundExistingHash != nil {
		return sdkerrors.Wrap(sdkerrors.Error{}, "Txn Hash already exists")
	}

	k.SetSafetyHash(
		*ctx,
		keccakHash,
	)
	return nil
}
