package types

import "fmt"

// NewGenesisState creates a new genesis state.
func NewGenesisState(params Params, pairs []TokenPair) *GenesisState {
	return &GenesisState{
		Params:     params,
		TokenPairs: pairs,
	}
}

// DefaultGenesisState sets default token genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate perform genesis state validation for token module.
func (gs *GenesisState) Validate() error {
	seenContract := make(map[string]bool)
	seenDenom := make(map[string]bool)

	for _, b := range gs.TokenPairs {
		if seenContract[b.ContractAddress] {
			return fmt.Errorf("erc20 contract duplicated on genesis '%s'", b.ContractAddress)
		}

		if seenDenom[b.Denom] {
			return fmt.Errorf("coin denomination duplicated on genesis: '%s'", b.Denom)
		}

		if err := b.Validate(); err != nil {
			return err
		}
	}

	return nil
}
