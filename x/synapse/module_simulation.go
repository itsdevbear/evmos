package synapse

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/tharsis/evmos/testutil/sample"
	synapsesimulation "github.com/tharsis/evmos/x/synapse/simulation"
	"github.com/tharsis/evmos/x/synapse/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = synapsesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgBridgeIn = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBridgeIn int = 100

	opWeightMsgBridgeOut = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBridgeOut int = 100

	opWeightMsgSetPendingKmsAddress = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetPendingKmsAddress int = 100

	opWeightMsgSetKmsAddress = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetKmsAddress int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	synapseGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&synapseGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgBridgeIn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBridgeIn, &weightMsgBridgeIn, nil,
		func(_ *rand.Rand) {
			weightMsgBridgeIn = defaultWeightMsgBridgeIn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBridgeIn,
		synapsesimulation.SimulateMsgBridgeIn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBridgeOut int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBridgeOut, &weightMsgBridgeOut, nil,
		func(_ *rand.Rand) {
			weightMsgBridgeOut = defaultWeightMsgBridgeOut
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBridgeOut,
		synapsesimulation.SimulateMsgBridgeOut(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetPendingKmsAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetPendingKmsAddress, &weightMsgSetPendingKmsAddress, nil,
		func(_ *rand.Rand) {
			weightMsgSetPendingKmsAddress = defaultWeightMsgSetPendingKmsAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetPendingKmsAddress,
		synapsesimulation.SimulateMsgSetPendingKmsAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetKmsAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetKmsAddress, &weightMsgSetKmsAddress, nil,
		func(_ *rand.Rand) {
			weightMsgSetKmsAddress = defaultWeightMsgSetKmsAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetKmsAddress,
		synapsesimulation.SimulateMsgSetKmsAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
