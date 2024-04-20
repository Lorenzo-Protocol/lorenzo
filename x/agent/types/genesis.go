package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesisState returns the default genesis state.
//
// No parameters.
// Returns a pointer to GenesisState.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		NextNumber: 1,
	}
}

// Validate checks if the given GenesisState is valid.
//
// It validates the sender address in the GenesisState by converting it from Bech32 format to sdk.AccAddress.
// If the conversion fails, it returns the error. Otherwise, it returns nil.
//
// Parameters:
// - data: the GenesisState to be validated.
//
// Returns:
// - error: an error if the sender address is invalid, otherwise nil.
func (data GenesisState) Validate() error {
	_, err := sdk.AccAddressFromBech32(data.Admin)
	if err != nil {
		return err
	}

	if data.NextNumber <= 0 {
		return ErrInvalidID
	}

	for _, agent := range data.Agents {
		if len(strings.TrimSpace(agent.Name)) == 0 {
			return ErrNameEmpty
		}

		if len(strings.TrimSpace(agent.BtcReceivingAddress)) == 0 {
			return ErrBtcReceivingAddressEmpty
		}
	}

	return nil
}
