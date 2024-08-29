package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const (
	TypeMsgUpdateHeader  = "update_header"
	TypeMsgUploadHeaders = "upload_headers"
	TypeMsgUpdateParams  = "update_params"
)

var (
	_ sdk.Msg = (*MsgUpdateHeader)(nil)
	_ sdk.Msg = (*MsgUploadHeaders)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

var (
	_ legacytx.LegacyMsg = (*MsgUpdateHeader)(nil)
	_ legacytx.LegacyMsg = (*MsgUploadHeaders)(nil)
	_ legacytx.LegacyMsg = (*MsgUpdateParams)(nil)
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

func (m *MsgUpdateHeader) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateHeader) Route() string {
	return RouterKey
}

func (m *MsgUpdateHeader) Type() string {
	return TypeMsgUpdateHeader
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

func (m *MsgUploadHeaders) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUploadHeaders) Route() string {
	return RouterKey
}

func (m *MsgUploadHeaders) Type() string {
	return TypeMsgUploadHeaders
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

func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateParams) Route() string {
	return RouterKey
}

func (m *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}
