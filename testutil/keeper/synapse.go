package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	"github.com/tharsis/evmos/x/synapse/keeper"
	"github.com/tharsis/evmos/x/synapse/types"
)

func SynapseKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	appCodec := codec.NewProtoCodec(registry)
	// capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, storeKey, memStoreKey)

	// ss := typesparams.NewSubspace(appCodec,
	// 	types.Amino,
	// 	storeKey,
	// 	memStoreKey,
	// 	"SynapseSubSpace",
	// )
	// IBCKeeper := ibckeeper.NewKeeper(
	// 	appCodec,
	// 	storeKey,
	// 	ss,
	// 	nil,
	// 	nil,
	// 	capabilityKeeper.ScopeToModule("SynapseIBCKeeper"),
	// )

	paramsSubspace := typesparams.NewSubspace(appCodec,
		types.Amino,
		storeKey,
		memStoreKey,
		"SynapseParams",
	)
	k := keeper.NewKeeper(
		appCodec,
		storeKey,
		memStoreKey,
		paramsSubspace,
		// IBCKeeper.ChannelKeeper,
		// &IBCKeeper.PortKeeper,
		// capabilityKeeper.ScopeToModule("SynapseScopedKeeper"),
		nil,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
