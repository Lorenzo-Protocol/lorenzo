package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
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
	cdc.RegisterConcrete(&MsgInsertHeaders{}, "lorenzo/btclightclient/MsgInsertHeaders", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lorenzo/btclightclient/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgUpdateFeeRate{}, "lorenzo/btclightclient/MsgUpdateFeeRate", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// Register messages
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertHeaders{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
