package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	amino = codec.NewLegacyAmino()
	// AminoCdc is the amino codec used for serialization.
	AminoCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
}

// RegisterLegacyAminoCodec registers the module's types for the given codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClient{}, "lorenzo/ccev/MsgCreateClient", nil)
	cdc.RegisterConcrete(&MsgUploadContract{}, "lorenzo/ccev/MsgUploadContract", nil)
	cdc.RegisterConcrete(&MsgUpdateHeader{}, "lorenzo/ccev/MsgUploadHeaders", nil)
	cdc.RegisterConcrete(&MsgUploadHeaders{}, "lorenzo/ccev/MsgUpdateHeader", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "lorenzo/ccev/MsgUpdateParams", nil)
}

// RegisterInterfaces registers implementations for sdk.Msg and MsgUpdateParams in the given InterfaceRegistry.
//
// Parameter:
// - registry: the InterfaceRegistry to register implementations to.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClient{},
		&MsgUploadContract{},
		&MsgUpdateParams{},
		&MsgUpdateHeader{},
		&MsgUploadHeaders{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// UnmarshalProof unmarshals a byte slice into a Proof struct.
//
// Parameters:
// - data: a byte slice containing the JSON representation of a Proof struct.
//
// Returns:
// - *Proof: a pointer to a Proof struct if the unmarshalling is successful.
// - error: an error if the unmarshalling fails.
func UnmarshalProof(data []byte) (*Proof, error) {
	proof := new(Proof)
	if err := rlp.DecodeBytes(data, proof); err != nil {
		return nil, ErrInvalidProof
	}
	return proof, nil
}

// UnmarshalReceipt unmarshals the given data into an evmtypes.Receipt object.
//
// Parameters:
// - data: a byte slice containing the JSON-encoded receipt data.
//
// Returns:
// - *evmtypes.Receipt: a pointer to the unmarshaled receipt object.
// - error: an error if the unmarshaling process fails.
func UnmarshalReceipt(data []byte) (*evmtypes.Receipt, error) {
	receipt := new(evmtypes.Receipt)
	if err := rlp.DecodeBytes(data, receipt); err != nil {
		return nil, err
	}
	return receipt, nil
}
