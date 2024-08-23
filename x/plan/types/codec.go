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
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lorenzo/plan/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgUpgradePlan{}, "lorenzo/plan/MsgUpgradePlan", nil)
	cdc.RegisterConcrete(&MsgCreatePlan{}, "lorenzo/plan/MsgCreatePlan", nil)
	cdc.RegisterConcrete(&MsgSetMerkleRoot{}, "lorenzo/plan/MsgSetMerkleRoot", nil)
	cdc.RegisterConcrete(&MsgClaims{}, "lorenzo/plan/MsgClaims", nil)
	cdc.RegisterConcrete(&MsgUpdatePlanStatus{}, "lorenzo/plan/MsgUpdatePlanStatus", nil)
	cdc.RegisterConcrete(&MsgCreateYAT{}, "lorenzo/plan/MsgCreateYAT", nil)
	cdc.RegisterConcrete(&MsgSetMinter{}, "lorenzo/plan/MsgSetMinter", nil)
	cdc.RegisterConcrete(&MsgRemoveMinter{}, "lorenzo/plan/MsgRemoveMinter", nil)
}

// RegisterInterfaces registers implementations for sdk.Msg and MsgUpdateParams in the given InterfaceRegistry.
//
// Parameter:
// - registry: the InterfaceRegistry to register implementations to.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgCreatePlan{},
		&MsgClaims{},
		&MsgUpgradePlan{},
		&MsgCreateYAT{},
		&MsgUpdatePlanStatus{},
		&MsgSetMinter{},
		&MsgRemoveMinter{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
