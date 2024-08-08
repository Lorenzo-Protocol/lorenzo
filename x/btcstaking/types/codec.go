package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateBTCStaking{}, "btcstaking/MsgCreateBTCStaking", nil)
	cdc.RegisterConcrete(&MsgBurnRequest{}, "btcstaking/MsgBurnRequest", nil)
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
