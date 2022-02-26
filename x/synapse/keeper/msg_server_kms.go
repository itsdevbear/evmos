package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/evmos/x/synapse/types"
)

func (k msgServer) SetPendingKmsAddress(goCtx context.Context, msg *types.MsgSetPendingKmsAddress) (*types.MsgSetPendingKmsAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.SetKmsPendingAddress(ctx, msg.PendingKmsAddress); err != nil {
		return nil, err
	}

	return &types.MsgSetPendingKmsAddressResponse{}, nil
}

func (k msgServer) UpdateKmsAddress(goCtx context.Context, msg *types.MsgSetKmsAddress) (*types.MsgSetKmsAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	kmsPendingAddress := k.GetKmsPendingAddress(ctx)

	// Set the active Address to the pending Address
	if err := k.SetKmsAddress(ctx, kmsPendingAddress); err != nil {
		return nil, err
	}

	// Clear Old Pending Address
	if err := k.SetKmsPendingAddress(ctx, ""); err != nil {
		return nil, err
	}

	return &types.MsgSetKmsAddressResponse{}, nil
}
