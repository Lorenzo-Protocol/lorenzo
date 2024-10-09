package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validate validates the Params struct.
//
// It does not take any parameters.
// It returns an error if the validation fails.
func (p Params) Validate() error {
	return ValidateAllowList(p.AllowList)
}

// ValidateAllowList validates a list of addresses.
//
// It takes a list of strings representing addresses as a parameter.
// It returns an error if any address is invalid.
func ValidateAllowList(allowList []string) error {
	seenAddr := make(map[string]bool)
	for _, addr := range allowList {
		if seenAddr[addr] {
			return errorsmod.Wrapf(ErrDuplicateAddress, "duplicate address: %s", addr)
		}

		// check that the address is valid
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address: %s", addr)
		}
		seenAddr[addr] = true
	}
	return nil
}
