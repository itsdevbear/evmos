package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// PortId:          PortID,
		SeenKeccaks: nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// if err := host.PortIdentifierValidator(gs.PortId); err != nil {
	// 	return err
	// }
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
