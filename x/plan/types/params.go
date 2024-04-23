package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates the Params struct.
//
// It does not take any parameters.
// It returns an error if the validation fails.
func (p Params) Validate() error {
	if err := ValidateAddressList(p.AllowList); err != nil {
		return err
	}

	return nil
}

func ValidateAddressList(i interface{}) error {
	allowList, ok := i.([]string)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, a := range allowList {
		if _, err := sdk.AccAddressFromBech32(a); err != nil {
			return fmt.Errorf("invalid address")
		}
	}

	return nil
}
