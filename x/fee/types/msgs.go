package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgUpdateParams)(nil)

// ValidateBasic executes sanity validation on the provided data
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	return m.Params.Validate()
}

// GetSigners returns the expected signers for a MsgUpdateParams message
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// Validate validates the Params struct.
//
// It does not take any parameters.
// It returns an error if the validation fails.
func (p Params) Validate() error {
	seen := make(map[string]bool)
	for _,msg := range p.NonFeeMsgs {
		if seen[msg] {
			return errorsmod.Wrapf(ErrDuplicateMsg, "duplicate msg %s", msg)
		}
		seen[msg] = true
	}
	return nil
}


// Validate validates the GenesisState.
//
// It returns an error if the GenesisState parameters are invalid.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// NewGenesisState returns a new GenesisState.
//
// It takes a Params parameter and returns a pointer to GenesisState.
func NewGenesisState(params Params) *GenesisState {
	return &GenesisState{
		Params: params,
	}
}

// DefaultGenesisState gets raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: Params{},
	}
}