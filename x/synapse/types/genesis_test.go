package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tharsis/evmos/x/synapse/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				// PortId: types.PortID,
				Params: types.Params{
					KmsAddress:           "bera10jmp6sgh4cc6zt3e8gw05wavvejgr5pw3per66",
					KmsPendingAddress:    "",
					EvmMiddlewareAddress: "0x0000000000000000000000000000000000000001",
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
