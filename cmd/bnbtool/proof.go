package main

import (
	"context"
	"encoding/base64"
	"errors"
	"math/big"
	"os"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
)

func (c client) genReceiptProof(ctx context.Context, height int64, txHash string, receiptFile, proofFile string) error {
	block, err := c.eth.BlockByNumber(context.Background(), big.NewInt(height))
	if err != nil {
		return err
	}

	blockNumber := rpc.BlockNumber(height)
	receipts, err := c.eth.BlockReceipts(ctx, rpc.BlockNumberOrHash{BlockNumber: &blockNumber})
	if err != nil {
		return err
	}

	txIndex := slices.IndexFunc(receipts,
		func(e *types.Receipt) bool {
			return e.TxHash == common.HexToHash(txHash)
		},
	)

	// 将交易回执写入文件
	bz, err := rlp.EncodeToBytes(receipts[txIndex])
	if err != nil {
		return err
	}
	if err := os.WriteFile(receiptFile, bz, 0644); err != nil {
		return err
	}

	proof, err := genReceiptProof(uint64(txIndex), block.ReceiptHash(), receipts)
	if err != nil {
		return err
	}

	// 将证明写入文件
	bz, err = rlp.EncodeToBytes(proof)
	if err != nil {
		return err
	}
	if err := os.WriteFile(proofFile, bz, 0644); err != nil {
		return err
	}
	return nil
}

func genReceiptProof(txIndex uint64, root common.Hash, receipts []*types.Receipt) (*Proof, error) {
	db := triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil)
	mpt := trie.NewEmpty(db)
	receiptHash := types.DeriveSha(types.Receipts(receipts), mpt)
	if receiptHash != root {
		return nil, errors.New("root hash mismatch")
	}

	var indexBuf []byte
	indexBuf = rlp.AppendUint64(indexBuf[:0], uint64(txIndex))
	valueTarget, err := mpt.Get(indexBuf)
	if err != nil {
		return nil, err
	}

	proof := NewProofPath()
	if err := mpt.Prove(indexBuf, &proof); err != nil {
		return nil, err
	}

	res := Proof{
		Index: indexBuf,
		Value: valueTarget,
		Path:  proof,
	}

	return &res, nil
}

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
