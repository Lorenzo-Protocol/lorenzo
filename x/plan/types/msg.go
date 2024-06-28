package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = (*MsgUpdateParams)(nil)
	_ sdk.Msg = (*MsgUpgradeYAT)(nil)
	_ sdk.Msg = (*MsgCreatePlan)(nil)
	_ sdk.Msg = (*MsgClaims)(nil)
	_ sdk.Msg = (*MsgCreateYAT)(nil)
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
func (m *MsgUpgradeYAT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	if !common.IsHexAddress(m.Implementation) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "expecting a hex address, got %s", m.Implementation)
	}
	return nil
}

// GetSigners returns the expected signers for a MsgUpdateParams message
func (m *MsgUpgradeYAT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgCreatePlan) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgCreatePlan) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgClaims) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.Receiver) {
		return errorsmod.Wrap(ErrReceiver, "invalid receiver address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgClaims) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgCreateYAT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.Sender) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgCreateYAT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}
