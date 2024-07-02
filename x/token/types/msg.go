package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_ sdk.Msg = &MsgConvertCoin{}
	_ sdk.Msg = &MsgConvertERC20{}
	_ sdk.Msg = &MsgUpdateParams{}
)

const (
	TypeMsgConvertCoin  = "convert_coin"
	TypeMsgConvertERC20 = "convert_erc20"
)

func (m *MsgConvertCoin) ValidateBasic() error {
	// TODO implement me
	panic("implement me")
}

func (m *MsgConvertCoin) GetSigners() []sdk.AccAddress {
	// TODO implement me
	panic("implement me")
}

func (m *MsgConvertERC20) ValidateBasic() error {
	// TODO implement me
	panic("implement me")
}

func (m *MsgConvertERC20) GetSigners() []sdk.AccAddress {
	// TODO implement me
	panic("implement me")
}

func (m *MsgUpdateParams) ValidateBasic() error {
	// TODO implement me
	panic("implement me")
}

func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	// TODO implement me
	panic("implement me")
}
