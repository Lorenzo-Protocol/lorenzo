package types

// Validate validates the GenesisState.
//
// It returns an error if the GenesisState parameters are invalid.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	for _, chain := range gs.ChainStates {
		if err := ValidateClient(chain.Client); err != nil {
			return err
		}

		for _, header := range chain.Headers {
			if err := ValidateHeader(header); err != nil {
				return err
			}
		}

		for _, contract := range chain.Contracts {
			if err := ValidateContract(contract.Contract); err != nil {
				return err
			}
		}
	}
	return nil
}

// DefaultGenesisState returns the default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: &Params{},
	}
}
