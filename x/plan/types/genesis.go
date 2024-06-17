package types

// Validate validates the GenesisState.
//
// It returns an error if the GenesisState parameters are invalid.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// NewGenesisState returns a new GenesisState.
//
// It takes a Params parameter and returns a pointer to GenesisState.
func NewGenesisState(params Params, nextNumber uint64, plans []Plan) *GenesisState {
	return &GenesisState{
		Params:     params,
		NextNumber: nextNumber,
		Plans:      plans,
	}
}

// DefaultGenesisState gets raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:     Params{},
		NextNumber: 0,
	}
}
