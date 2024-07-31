package types

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// Proof is a struct that contains the index, value, and proof.
type Proof struct {
	Index []byte    `json:"index"`
	Value []byte    `json:"value"`
	Path  ProofPath `json:"path"`
}

// ProofPath is a struct that contains the keys and values of the proof.
type ProofPath struct {
	Keys   []string `json:"keys"`
	Values [][]byte `json:"values"`
}

// NewProofPath creates a new ProofPath.
func NewProofPath() ProofPath {
	return ProofPath{
		Keys:   make([]string, 0),
		Values: make([][]byte, 0),
	}
}

// Put adds a new key-value pair to the ProofPath.
func (pm *ProofPath) Put(key []byte, value []byte) error {
	if pm.Keys == nil || pm.Values == nil {
		return errors.New("proofMap db is nil")
	}

	vIdx := pm.valueIdx(key)
	if vIdx != -1 {
		pm.Values[vIdx] = value
		return nil
	}

	pm.Keys = append(pm.Keys, pm.encodeKey(key))
	pm.Values = append(pm.Values, value)
	return nil
}

// Delete removes a key-value pair from the ProofPath.
func (pm *ProofPath) Delete(key []byte) error {
	if pm.Keys == nil || pm.Values == nil {
		return errors.New("proofMap db is nil")
	}

	vIdx := pm.valueIdx(key)
	if vIdx == -1 {
		return errors.New("key not found")
	}

	pm.Keys = append(pm.Keys[:vIdx], pm.Keys[vIdx+1:]...)
	pm.Values = append(pm.Values[:vIdx], pm.Values[vIdx+1:]...)

	return nil
}

// Has checks if a key exists in the ProofPath.
func (pm *ProofPath) Has(key []byte) (bool, error) {
	if pm.Keys == nil || pm.Values == nil {
		return false, errors.New("proofMap db is nil")
	}
	vIdx := pm.valueIdx(key)
	if vIdx == -1 {
		return false, nil
	}
	return true, nil
}

// Get returns the value of a key in the ProofPath.
func (pm *ProofPath) Get(key []byte) ([]byte, error) {
	if pm.Keys == nil || pm.Values == nil {
		return nil, errors.New("proofMap db is nil")
	}

	vIdx := pm.valueIdx(key)
	if vIdx == -1 {
		return nil, errors.New("value not found")
	}
	return pm.Values[vIdx], nil
}

func (pm *ProofPath) encodeKey(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

func (pm *ProofPath) valueIdx(key []byte) int {
	encodedKey := pm.encodeKey(key)
	for i, k := range pm.Keys {
		if k == encodedKey {
			return i
		}
	}
	return -1
}

// GenReceiptProof generates a Merkle Patricia Trie (MPT) proof for a specific transaction index in a list of receipts.
//
// Parameters:
// - txIndex: the index of the transaction for which the proof is generated.
// - root: the root hash of the MPT.
// - receipts: a list of receipts.
//
// Returns:
// - *mptproof.MPTProof: a pointer to an MPTProof struct containing the index, value, and proof of the transaction.
// - error: an error if the proof generation fails.
func GenReceiptProof(txIndex uint64, root common.Hash, receipts []*evmtypes.Receipt) (*Proof, error) {
	db := trie.NewDatabase(rawdb.NewMemoryDatabase())
	mpt := trie.NewEmpty(db)
	receiptHash := evmtypes.DeriveSha(evmtypes.Receipts(receipts), mpt)
	if receiptHash != root {
		return nil, errors.New("root hash mismatch")
	}

	var indexBuf []byte
	indexBuf = rlp.AppendUint64(indexBuf[:0], txIndex)
	valueTarget := mpt.Get(indexBuf)

	proof := NewProofPath()
	err := mpt.Prove(indexBuf, 0, &proof)
	if err != nil {
		return nil, err
	}

	res := Proof{
		Index: indexBuf,
		Value: valueTarget,
		Path:  proof,
	}

	return &res, nil
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
	var proof Proof
	if err := json.Unmarshal(data, &proof); err != nil {
		return nil, ErrInvalidProof
	}
	return &proof, nil
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
	receipt := &evmtypes.Receipt{}
	if err := receipt.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	return receipt, nil
}
