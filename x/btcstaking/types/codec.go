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
	cdc.RegisterConcrete(&MsgCreateBTCStaking{}, "lorenzo/btcstaking/MsgCreateBTCStaking", nil)
	cdc.RegisterConcrete(&MsgCreateBTCBStaking{}, "lorenzo/btcstaking/MsgCreateBTCBStaking", nil)
	cdc.RegisterConcrete(&MsgBurnRequest{}, "lorenzo/btcstaking/MsgBurnRequest", nil)
	cdc.RegisterConcrete(&MsgRemoveReceiver{}, "lorenzo/btcstaking/MsgAddReceiver", nil)
	cdc.RegisterConcrete(&MsgAddReceiver{}, "lorenzo/btcstaking/MsgRemoveReceiver", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lorenzo/btcstaking/MsgUpdateParams", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// Register messages
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateBTCStaking{},
		&MsgCreateBTCBStaking{},
		&MsgBurnRequest{},
		&MsgRemoveReceiver{},
		&MsgAddReceiver{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
