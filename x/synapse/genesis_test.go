package synapse_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tharsis/evmos/testutil/keeper"
	"github.com/tharsis/evmos/testutil/nullify"
	"github.com/tharsis/evmos/x/synapse"
	"github.com/tharsis/evmos/x/synapse/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SynapseKeeper(t)
	synapse.InitGenesis(ctx, *k, genesisState)
	got := synapse.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// require.Equal(t, genesisState.PortId, got.PortId)
	// this line is used by starport scaffolding # genesis/test/assert
}
