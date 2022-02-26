package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tharsis/evmos/x/synapse/types"
)

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.GetKmsAddress(ctx),
		k.GetKmsPendingAddress(ctx),
		k.GetEvmMiddlewareAddress(ctx),
	)
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetKmsAddress(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyKmsAddress, &res)
	return res
}

func (k Keeper) SetKmsAddress(ctx sdk.Context, addr string) (err error) {
	if err := sdk.VerifyAddressFormat([]byte(addr)); err != nil {
		return err
	}
	k.paramstore.Set(ctx, types.KeyKmsAddress, addr)
	return nil
}

func (k Keeper) GetKmsPendingAddress(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyKmsPendingAddress, &res)
	return res
}

func (k Keeper) SetKmsPendingAddress(ctx sdk.Context, addr string) (err error) {
	if err := sdk.VerifyAddressFormat([]byte(addr)); err != nil {
		return err
	}
	k.paramstore.Set(ctx, types.KeyKmsPendingAddress, addr)
	return nil
}

func (k Keeper) SetEvmMiddlewareAddress(ctx sdk.Context, addr string) (err error) {
	if !common.IsHexAddress(addr) {
		return fmt.Errorf("addr is not in Hex format")
	}
	k.paramstore.Set(ctx, types.KeyEvmMiddlewareAddress, addr)
	return nil
}

func (k Keeper) GetEvmMiddlewareAddress(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyEvmMiddlewareAddress, &res)
	return res
}
