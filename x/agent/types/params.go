package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{}
}

func NewParams(allowList []string) Params {
	return Params{
		AllowList: allowList,
	}
}

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
		return fmt.Errorf("invalid parameter type: %T, expected []string", i)
	}
	seenMap := make(map[string]bool)
	for _, a := range allowList {
		if seenMap[a] {
			return fmt.Errorf("duplicate address: %s", a)
		}
		if _, err := sdk.AccAddressFromBech32(a); err != nil {
			return fmt.Errorf("invalid address: %s", a)
		}
		seenMap[a] = true
	}
	return nil
}
