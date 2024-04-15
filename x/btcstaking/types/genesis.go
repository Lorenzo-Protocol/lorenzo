package types

// 130 vbytes * 1e7
// native token 's precision is 1e18, but the fee rate value is increased by 1000 times, so finally multiply by 1e7.
const initBurnFeeFactor = 130 * 1e7

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: &Params{
			BtcConfirmationsDepth: 6,
			BurnFeeFactor:         initBurnFeeFactor,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return nil
}
