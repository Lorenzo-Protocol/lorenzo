package types

import (
	"bytes"
	"errors"
	"fmt"
	big "math/big"
	"time"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v3/types"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func (ti *TransactionInfo) ValidateBasic() error {
	if ti.Key == nil {
		return fmt.Errorf("key in TransactionInfo is nil")
	}
	if ti.Transaction == nil {
		return fmt.Errorf("transaction in TransactionInfo is nil")
	}
	if ti.Proof == nil {
		return fmt.Errorf("proof in TransactionInfo is nil")
	}
	return nil
}

func ParseTransaction(bytes []byte) (*btcutil.Tx, error) {
	tx, e := btcutil.NewTxFromBytes(bytes)

	if e != nil {
		return nil, e
	}

	e = blockchain.CheckTransactionSanity(tx)

	if e != nil {
		return nil, e
	}

	return tx, nil
}

func ValidateBTCHeader(header *wire.BlockHeader, powLimit *big.Int) error {
	msgBlock := &wire.MsgBlock{Header: *header}
	block := btcutil.NewBlock(msgBlock)

	// The upper limit for the power to be spent
	// Use the one maintained by btcd
	err := blockchain.CheckProofOfWork(block, powLimit)
	if err != nil {
		return err
	}

	if !header.Timestamp.Equal(time.Unix(header.Timestamp.Unix(), 0)) {
		str := fmt.Sprintf("block timestamp of %v has a higher "+
			"precision than one second", header.Timestamp)
		return errors.New(str)
	}

	return nil
}

func hashConcat(a []byte, b []byte) chainhash.Hash {
	c := []byte{}
	c = append(c, a...)
	c = append(c, b...)
	return chainhash.DoubleHashH(c)
}

func verify(tx *btcutil.Tx, merkleRoot *chainhash.Hash, intermediateNodes []byte, index uint32) bool {
	txHash := tx.Hash()

	// Shortcut the empty-block case
	if txHash.IsEqual(merkleRoot) && index == 0 && len(intermediateNodes) == 0 {
		return true
	}

	proof := []byte{}
	proof = append(proof, txHash[:]...)
	proof = append(proof, intermediateNodes...)
	proof = append(proof, merkleRoot[:]...)

	var current chainhash.Hash

	idx := index

	proofLength := len(proof)

	if proofLength%32 != 0 {
		return false
	}

	if proofLength == 64 {
		return false
	}

	root := proof[proofLength-32:]

	cur := proof[:32:32]
	copy(current[:], cur)

	numSteps := (proofLength / 32) - 1

	for i := 1; i < numSteps; i++ {
		start := i * 32
		end := i*32 + 32
		next := proof[start:end:end]
		if idx%2 == 1 {
			current = hashConcat(next, current[:])
		} else {
			current = hashConcat(current[:], next)
		}
		idx >>= 1
	}

	return bytes.Equal(current[:], root)
}

func (ti *TransactionInfo) VerifyInclusion(btcHeader *lrz.BTCHeaderBytes, powLimit *big.Int) error {
	if err := ti.ValidateBasic(); err != nil {
		return err
	}
	if !ti.Key.Hash.Eq(btcHeader.Hash()) {
		return fmt.Errorf("the given btcHeader is different from that in TransactionInfo")
	}

	tx, err := ParseTransaction(ti.Transaction)
	if err != nil {
		return err
	}

	header := btcHeader.ToBlockHeader()
	if err := ValidateBTCHeader(header, powLimit); err != nil {
		return err
	}

	if !verify(tx, &header.MerkleRoot, ti.Proof, ti.Key.Index) {
		return fmt.Errorf("header failed validation due to failed proof")
	}

	return nil
}
