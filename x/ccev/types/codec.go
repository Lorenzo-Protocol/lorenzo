package types

import (
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

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