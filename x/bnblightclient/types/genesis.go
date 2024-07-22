package types

// Validate validates the GenesisState.
//
// It returns an error if the GenesisState parameters are invalid.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return VeryHeaders(gs.Headers)
}

// DefaultGenesisState returns the default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: &Params{},
		Headers: []*Header{},
		Events: []*EvmEvent{},
	}
}
