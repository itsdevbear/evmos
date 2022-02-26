package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tharsis/evmos/x/synapse/types"
)

func (k Keeper) GetSafetyHash(ctx sdk.Context, keccakHash []byte) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixSafetyHash)
	return store.Get(keccakHash)
}

func (k Keeper) SetSafetyHash(ctx sdk.Context, keccakHash []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixSafetyHash)
	store.Set(keccakHash, []byte("1"))
}

func (k Keeper) GetAllSafetyHash(ctx sdk.Context) [][]byte {
	hashes := [][]byte{}

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixSafetyHash)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var keccakHash []byte
		hashes = append(hashes, keccakHash)
	}

	return hashes
}
