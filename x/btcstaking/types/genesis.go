package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: &Params{
			BtcConfirmationsDepth: 6,
		},
	}
}

func (receiver Receiver) Validate() error {
	if len(receiver.Name) == 0 {
		return fmt.Errorf("receiver name cannot be empty")
	}
	if len(receiver.Addr) == 0 {
		return fmt.Errorf("receiver addr cannot be empty")
	}
	if !common.IsHexAddress(receiver.EthAddr) && len(receiver.EthAddr) != 0 {
		return fmt.Errorf("receiver's eth addr must be empty or a valid eth addr")
	}
	return nil
}

func ValidateAddressList(allowList []string) error {
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

func (params Params) Validate() error {
	receivers := map[string]bool{}
	for _, receiver := range params.Receivers {
		if err := receiver.Validate(); err != nil {
			return err
		}
		if _, receiverExists := receivers[receiver.Name]; receiverExists {
			return fmt.Errorf("duplicate receiver name: %s", receiver)
		}
		receivers[receiver.Name] = true
	}
	if err := ValidateAddressList(params.MinterAllowList); err != nil {
		return fmt.Errorf("invalid minter allow list")
	}
	if !common.IsHexAddress(params.BridgeAddr) {
		return fmt.Errorf("invalid cross chain mint address")
	}
	return nil
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if gs.Params == nil {
		return fmt.Errorf("params cannot be nil")
	}
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}
