package keeper

import (
	"bytes"
	"math"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

func hashConcat(a []byte, b []byte) chainhash.Hash {
	c := []byte{}
	c = append(c, a...)
	c = append(c, b...)
	return chainhash.DoubleHashH(c)
}

func ParseMerkleBlock(proof []byte) (*wire.MsgMerkleBlock, error) {
	var msgMerkleBlk wire.MsgMerkleBlock
	if err := msgMerkleBlk.BtcDecode(bytes.NewReader(proof), wire.BIP0037Version, wire.WitnessEncoding); err != nil {
		return nil, err
	}
	return &msgMerkleBlk, nil
}

func getFlag(bitfield []byte, pos int) bool {
	return (bitfield[pos/8]>>(pos%8))&1 > 0
}

func calcWidth(nTxes, hei uint32) uint32 {
	return (nTxes + (1 << hei) - 1) >> hei
}

func calcHeight(nTxes int) uint32 {
	return uint32(math.Ceil(math.Log2(float64(nTxes))))
}

func traverseMerkleBlock(msg *wire.MsgMerkleBlock, hei uint32, pos uint32, bit_used *int, hash_used *int, proof *[]byte, txIndex *uint32) ([]byte, bool) {
	parentOfMatch := getFlag(msg.Flags, *bit_used)
	*bit_used += 1
	if hei == 0 || !parentOfMatch {
		hash := msg.Hashes[*hash_used]
		*hash_used += 1
		if hei == 0 && parentOfMatch {
			*txIndex = pos
		}
		return hash[:], parentOfMatch
	} else {
		left, lis := traverseMerkleBlock(msg, hei-1, pos*2, bit_used, hash_used, proof, txIndex)
		var right []byte
		if pos*2+1 < calcWidth(msg.Transactions, hei-1) {
			right, _ = traverseMerkleBlock(msg, hei-1, pos*2+1, bit_used, hash_used, proof, txIndex)
		} else {
			right = left
		}
		var hash chainhash.Hash
		if parentOfMatch {
			if lis {
				*proof = append(*proof, right...)
			} else {
				*proof = append(*proof, left...)
			}
			hash = hashConcat(left, right)
		} else {
			hash = hashConcat(right, left)
		}
		return hash[:], parentOfMatch
	}
}

// XXX: missing some checks, not a safe function to use on chain.
func ParseBTCProof(msgMerkleBlk *wire.MsgMerkleBlock) (uint32, []byte, error) {
	hei := calcHeight(int(msgMerkleBlk.Transactions))
	bit_used, hash_used := 0, 0

	proof := []byte{}
	txIndex := uint32(0)
	traverseMerkleBlock(msgMerkleBlk, hei, 0, &bit_used, &hash_used, &proof, &txIndex)
	return txIndex, proof, nil
}
