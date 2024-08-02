package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// RegisterInterfaces registers implementations for sdk.Msg and MsgUpdateParams in the given InterfaceRegistry.
//
// Parameter:
// - registry: the InterfaceRegistry to register implementations to.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgUpdateHeader{},
		&MsgUploadHeaders{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// UnmarshalBNBHeader decodes the input data into a BNBHeader struct.
//
// It takes a byte slice data as input and returns a BNBHeader pointer and an error.
func UnmarshalBNBHeader(rawHeader []byte) (*BNBHeader, error) {
	bnbHeader := new(BNBHeader)
	if err := rlp.DecodeBytes(rawHeader, bnbHeader); err != nil {
		return nil, err
	}
	return bnbHeader, nil
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
