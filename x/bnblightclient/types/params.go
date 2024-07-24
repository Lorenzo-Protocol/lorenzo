package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// Validate validates the Params struct.
//
// It does not take any parameters.
// It returns an error if the validation fails.
func (p Params) Validate() error {
	if len(p.StakePlanHubAddress) == 0 {
		return fmt.Errorf("stake plan hub address cannot be empty")
	}

	if !common.IsHexAddress(p.StakePlanHubAddress) {
		return fmt.Errorf("invalid stake plan hub address: %s", p.StakePlanHubAddress)
	}

	if len(p.EventName) == 0 {
		return fmt.Errorf("event name cannot be empty")
	}

	if err := ValidateAllowList(p.AllowList); err != nil {
		return err
	}
	return nil
}

// ValidateAllowList validates a list of addresses.
//
// It takes a list of strings representing addresses as a parameter.
// It returns an error if any address is invalid.
func ValidateAllowList(allowList []string) error {
	seenAddr := make(map[string]bool)
	for _, addr := range allowList {
		if seenAddr[addr] {
			return fmt.Errorf("duplicate address: %s", addr)
		}

		// check that the address is valid
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return fmt.Errorf("invalid address: %s", addr)
		}
		seenAddr[addr] = true
	}
	return nil
}
