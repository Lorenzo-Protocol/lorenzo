package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		// TODO: Add default genesis state
	}
}

func (gs *GenesisState) Validate() error {
	panic("implement me")
}
