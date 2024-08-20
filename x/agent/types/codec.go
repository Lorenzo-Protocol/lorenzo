package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
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
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lorenzo/agent/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgAddAgent{}, "lorenzo/agent/MsgAddAgent", nil)
	cdc.RegisterConcrete(&MsgEditAgent{}, "lorenzo/agent/MsgEditAgent", nil)
	cdc.RegisterConcrete(&MsgRemoveAgent{}, "lorenzo/agent/MsgRemoveAgent", nil)
}

// RegisterInterfaces registers implementations for sdk.Msg and MsgUpdateParams in the given InterfaceRegistry.
//
// Parameter:
// - registry: the InterfaceRegistry to register implementations to.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAgent{},
		&MsgEditAgent{},
		&MsgRemoveAgent{},
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
