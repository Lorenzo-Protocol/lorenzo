package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: &Params{
			BtcConfirmationsDepth: 6,
			// 130 vbytes * 1e7
			// native token 's precision is 1e18, but the fee rate value is increased by 1000 times, so finally multiply by 1e7.
			BurnFeeFactor: 130 * 1e7,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return nil
}
