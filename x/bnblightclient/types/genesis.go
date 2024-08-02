package types

// Validate validates the GenesisState.
//
// It returns an error if the GenesisState parameters are invalid.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return VerifyHeaders(gs.Headers)
}

// DefaultGenesisState returns the default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:  &Params{
			StakePlanHubAddress: "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
			EventName:           "StakeBTC2JoinStakePlan",
			RetainedBlocks:      100,
			AllowList:           []string{
				"lrz1v7vnrdvhwy99s2u825jnuac6tfxzpjch8m9e4n",
			},
			ChainId:             56,
		},
		Headers: []*Header{},
		Events:  []*EvmEvent{},
	}
}
