package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgUpdateHeader)(nil)
	_ sdk.Msg = (*MsgUploadHeaders)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

// ValidateBasic implements sdk.Msg interface
func (m *MsgUpdateHeader) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Signer); err != nil {
		return err
	}
	return VerifyHeaders([]*Header{m.Header})
}

// GetSigners implements sdk.Msg interface
func (m *MsgUpdateHeader) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Signer)}
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgUploadHeaders) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Signer); err != nil {
		return err
	}
	return VerifyHeaders(m.Headers)
}

// GetSigners implements sdk.Msg interface
func (m *MsgUploadHeaders) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Signer)}
}

// ValidateBasic implements sdk.Msg interface
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}
	return m.Params.Validate()
}

// GetSigners implements sdk.Msg interface
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}
