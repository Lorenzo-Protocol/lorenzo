package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// DefaultGenesisState returns the default genesis state.
//
// No parameters.
// Returns a pointer to GenesisState.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
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

	for _, agent := range data.Agents {
		if len(strings.TrimSpace(agent.Name)) == 0 {
			return ErrNameEmpty
		}

		if len(strings.TrimSpace(agent.BtcReceivingAddress)) == 0 {
			return ErrBtcReceivingAddressEmpty
		}

		if len(agent.EthAddr) != 0 && !common.IsHexAddress(agent.EthAddr) {
			return errorsmod.Wrap(ErrInvalidEthAddress, "EthAddr must be empty or a valid eth addr")
		}
	}
	return nil
}
