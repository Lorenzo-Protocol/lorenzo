package keeper

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

const (
	EthAddrLen        = 20
	ChainIDLen        = 4
	SatoshiToStBTCMul = 1e10
	PlanIDLen         = 8
)

const (
	Dep0Amount = 4e5
	Dep1Amount = 2e6
	Dep2Amount = 1e7
	Dep3Amount = 5e7
)

func NewBTCTxFromBytes(txBytes []byte) (*wire.MsgTx, error) {
	var msgTx wire.MsgTx
	rbuf := bytes.NewReader(txBytes)
	if err := msgTx.Deserialize(rbuf); err != nil {
		return nil, err
	}

	return &msgTx, nil
}

func ExtractPaymentTo(tx *wire.MsgTx, addr btcutil.Address) (uint64, error) {
	payToAddrScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return 0, fmt.Errorf("invalid address")
	}
	var amt uint64 = 0
	for _, out := range tx.TxOut {
		if bytes.Equal(out.PkScript, payToAddrScript) {
			amt += uint64(out.Value)
		}
	}
	return amt, nil
}

const maxOpReturnPkScriptSize = 83

func ExtractPaymentToWithOpReturnIdAndDust(tx *wire.MsgTx, addr btcutil.Address, dustAmount int64) (uint64, []byte, error) {
	payToAddrScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return 0, nil, fmt.Errorf("invalid address")
	}
	var amt uint64 = 0
	foundOpReturnId := false
	var opReturnId []byte
	for _, out := range tx.TxOut {
		if bytes.Equal(out.PkScript, payToAddrScript) && out.Value >= dustAmount {
			amt += uint64(out.Value)
		} else {
			pkScript := out.PkScript
			pkScriptLen := len(pkScript)
			// valid op return script will have at least 2 bytes
			// - fisrt byte should be OP_RETURN marker
			// - second byte should indicate how many bytes there are in opreturn script
			if pkScriptLen > 1 &&
				pkScriptLen <= maxOpReturnPkScriptSize &&
				pkScript[0] == txscript.OP_RETURN {

				// if this is OP_PUSHDATA1, we need to drop first 3 bytes as those are related
				// to script iteslf i.e OP_RETURN + OP_PUSHDATA1 + len of bytes
				if pkScript[1] == txscript.OP_PUSHDATA1 {
					opReturnId = pkScript[3:]
				} else if pkScript[1] == txscript.OP_PUSHDATA2 {
					opReturnId = pkScript[4:]
				} else if pkScript[1] == txscript.OP_PUSHDATA4 {
					opReturnId = pkScript[6:]
				} else {
					// this should be one of OP_DATAXX opcodes we drop first 2 bytes
					opReturnId = pkScript[2:]
				}
				foundOpReturnId = true
			}
		}
	}
	if !foundOpReturnId {
		return 0, nil, fmt.Errorf("expected op_return_id not found")
	}
	return amt, opReturnId, nil
}

func checkBTCTxDepth(stakingTxDepth uint64, btcAmount uint64) error {
	if btcAmount < Dep0Amount { // no depth check required
	} else if btcAmount < Dep1Amount { // at least 1 depth
		if stakingTxDepth < 1 {
			return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=1; depth=%d", stakingTxDepth)
		}
	} else if btcAmount < Dep2Amount {
		if stakingTxDepth < 2 {
			return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=2; depth=%d", stakingTxDepth)
		}
	} else if btcAmount < Dep3Amount {
		if stakingTxDepth < 3 {
			return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=3; depth=%d", stakingTxDepth)
		}
	} else if stakingTxDepth < 4 {
		return types.ErrBlkHdrNotConfirmed.Wrapf("not k-deep: k=4; depth=%d", stakingTxDepth)
	}
	return nil
}

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

func opReturnMsgLenCheck(opReturnMsg []byte) bool {
	return len(opReturnMsg) == EthAddrLen || len(opReturnMsg) == EthAddrLen+ChainIDLen || len(opReturnMsg) == EthAddrLen+ChainIDLen+PlanIDLen
}

func opReturnMsgContainsChainId(opReturnMsg []byte) bool {
	return len(opReturnMsg) >= EthAddrLen+ChainIDLen
}

func opReturnMsgGetChainId(opReturnMsg []byte) uint32 {
	base := EthAddrLen
	return binary.BigEndian.Uint32(opReturnMsg[base : base+ChainIDLen])
}

func opReturnMsgGetEthAddr(opReturnMsg []byte) []byte {
	return opReturnMsg[:EthAddrLen]
}

func opReturnMsgContainsPlanId(opReturnMsg []byte) bool {
	return len(opReturnMsg) >= EthAddrLen+ChainIDLen+PlanIDLen
}

func opReturnMsgGetPlanId(opReturnMsg []byte) uint64 {
	base := EthAddrLen + ChainIDLen
	return binary.BigEndian.Uint64(opReturnMsg[base : base+PlanIDLen])
}
