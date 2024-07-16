package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{}
}

func NewParams(allowList []string, beacon string) Params {
	return Params{
		AllowList: allowList,
		Beacon:    beacon,
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

	if len(p.Beacon) != 0 && !common.IsHexAddress(p.Beacon) {
		return fmt.Errorf("invalid beacon address: %s", p.Beacon)
	}

	return nil
}

func ValidateAddressList(i interface{}) error {
	allowList, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T, expected []string", i)
	}
	for _, a := range allowList {
		if _, err := sdk.AccAddressFromBech32(a); err != nil {
			return fmt.Errorf("invalid address: %s", a)
		}
	}
	return nil
}
