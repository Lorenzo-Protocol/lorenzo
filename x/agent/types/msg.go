package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgAddAgent)(nil)
	_ sdk.Msg = (*MsgRemoveAgent)(nil)
	_ sdk.Msg = (*MsgEditAgent)(nil)
)

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

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	if m.BtcReceivingAddress != DoNotModifyDesc && len(strings.TrimSpace(m.BtcReceivingAddress)) == 0 {
		return ErrBtcReceivingAddressEmpty
	}
	return nil
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
	return nil
}

// GetSigners returns the expected signers for a MsgRemoveAgent message
func (m *MsgRemoveAgent) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}
