package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgUpdateParams)(nil)
	_ sdk.Msg = (*MsgAddAgent)(nil)
	_ sdk.Msg = (*MsgRemoveAgent)(nil)
	_ sdk.Msg = (*MsgEditAgent)(nil)
)

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

// ValidateBasic executes sanity validation on the provided data
func (m *MsgAddAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(m.Name)) == 0 {
		return ErrNameEmpty
	}

	if len(strings.TrimSpace(m.BtcReceivingAddress)) == 0 {
		return ErrBtcReceivingAddressEmpty
	}

	if len(m.EthAddr) != 0 && !common.IsHexAddress(m.EthAddr) {
		return errorsmod.Wrap(ErrInvalidEthAddress, "EthAddr must be empty or a valid eth addr")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgAddAgent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgEditAgent) ValidateBasic() error {
	if m.Id <= 0 {
		return ErrInvalidID
	}
	if len(strings.TrimSpace(m.Name)) == 0 {
		return ErrNameEmpty
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	return err
}

// GetSigners returns the expected signers for a MsgEditAgent message
func (m *MsgEditAgent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgRemoveAgent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}
	if m.Id <= 0 {
		return ErrInvalidID
	}
	return err
}

// GetSigners returns the expected signers for a MsgRemoveAgent message
func (m *MsgRemoveAgent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}
