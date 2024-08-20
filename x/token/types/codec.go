package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino    = codec.NewLegacyAmino()
	AminoCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterCoin{}, "lorenzo/token/MsgRegisterCoin", nil)
	cdc.RegisterConcrete(&MsgRegisterERC20{}, "lorenzo/token/MsgRegisterERC20", nil)
	cdc.RegisterConcrete(&MsgToggleConversion{}, "lorenzo/token/MsgToggleConversion", nil)
	cdc.RegisterConcrete(&MsgConvertCoin{}, "lorenzo/token/MsgConvertCoin", nil)
	cdc.RegisterConcrete(&MsgConvertERC20{}, "lorenzo/token/MsgConvertERC20", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lorenzo/token/MsgUpdateParams", nil)
}

// RegisterInterfaces register implementations
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterCoin{},
		&MsgRegisterERC20{},
		&MsgToggleConversion{},
		&MsgConvertCoin{},
		&MsgConvertERC20{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
