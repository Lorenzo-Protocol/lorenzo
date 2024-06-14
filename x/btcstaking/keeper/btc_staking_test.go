package keeper

import (
	"bytes"
	"encoding/hex"
	"testing"

	lrz "github.com/Lorenzo-Protocol/lorenzo/types"
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/stretchr/testify/assert"
)

func TestBTCTx(t *testing.T) {
	assert := assert.New(t)
	txBytes, _ := hex.DecodeString("0200000000010274171503dbc24539663844c6f8e6c290947cdf021ddadcd3fdecc3ee049ea0ee0000000000fdffffff4e9c069460c3386eaac1c88d0c07cf89ee1399bcf1af12c9b39439ba9f44313b0200000000fdffffff034a01000000000000225120b3312a2c36383e101f2fa446eb16f00f1d2edef0eb7b839889e13f4592d80ecf00000000000000004a6a4800000a580000000000000000000000001c779b46ca5ffcf144f53aa09c17a5c3372de33e41494e4e000000000000000000000000000000000000000000000001158e460913d00000a7860000000000002251205aca89ca6c96635892be129cdf0ceed6e7c27c2713ddea812b678343455161b201406305bf2c1a3df628c581642f32e85dffa07b3dac3da4bca6b735c4b6c9dcca1e79959a589df01260464d41f8a7cd451a4f932b3fa50a76396503c0e87e380f530140e9ab6a00fe4a00f01cb61f013331efddb441dc807904f0a46c0f6d8da9eab950d5729db413d5fda9c3f3b0dcec33f8fdef8f6ad6776eb1baf9a08e9d6d0c826f00000000")
	proof, _ := hex.DecodeString("00000020657d06debfee161b0de46492a1cc776d6b56ee063c862ada1200000000000000dd80a6d8798a78c6671aa8af826404d2d8b796eef3784a99808d8b85f90ba8ecbcaa0366ffff001d04055dacf90100000a01b7ec9c4c909a9e65859935cc8d542f43b50e215473df6342b0b6373fe943ab932ffcb2edc6484d093a3a14513913a5e7ac4f741488fb9a1021361357c929b1070b29d5f13fe5ae24f115eece7c7f73d8e1a5eb6700781794435c2ef66861b21a27f8dda8f1f4597775c64dadfe30f9e03f65b21c1c483ebefc629b345a59d04e9c069460c3386eaac1c88d0c07cf89ee1399bcf1af12c9b39439ba9f44313b45e627c5780e608bdd72d2614088cb02d72e129889ec9faeb902f9ebae95adf9ef9c0d15192bd3e47cc08efef4f2366b14d3ec8075fcb842d541fbf4418ce0d1b1196114246b0cab7cbcf5633bb0ec3ab1d8a02cadcf1aecc853a8641bf3332fcf1e609853bd5c868240f7535ab6e988c53d3b16ea424960c0a453b9dced7256f1db8093992557847d26df562bf630ea2d6606aab2d44446fa40a7cace6d1246035d5b00")
	tx, _ := types.ParseTransaction(txBytes)
	merkleBlk, _ := ParseMerkleBlock(proof)
	txIndex, proofBytes, _ := ParseBTCProof(merkleBlk)
	blkHdr := &merkleBlk.Header

	var blkHdrHashBytes lrz.BTCHeaderHashBytes
	tmp := blkHdr.BlockHash()
	blkHdrHashBytes.FromChainhash(&tmp)
	txInfo := types.TransactionInfo{
		Key: &types.TransactionKey{
			Index: txIndex,
			Hash:  &blkHdrHashBytes,
		},
		Transaction: txBytes,
		Proof:       proofBytes,
	}
	var blkHdrBytesbuf bytes.Buffer
	err := blkHdr.Serialize(&blkHdrBytesbuf)
	assert.Equal(nil, err, "serialize should work")
	tmp2 := blkHdrBytesbuf.Bytes()
	assert.Equal(nil, txInfo.VerifyInclusion((*lrz.BTCHeaderBytes)(&tmp2), chaincfg.TestNet3Params.PowLimit), "inclusion should work")

	addr := "tb1ptt9gnjnvje343y47z2wd7r8w6mnuylp8z0w74qftv7p5x323vxeq9jrn6f"
	btcAddr, _ := btcutil.DecodeAddress(addr, &chaincfg.TestNet3Params)
	amt, opReturnId, e := ExtractPaymentToWithOpReturnId(tx.MsgTx(), btcAddr)
	assert.Equal(nil, e, "fail to extract")
	assert.Equal(uint64(34471), amt, "wrong amount")
	assert.Equal([]byte{0, 0, 10, 88, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 28, 119, 155, 70, 202, 95, 252, 241, 68, 245, 58, 160, 156, 23, 165, 195, 55, 45, 227, 62, 65, 73, 78, 78, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 21, 142, 70, 9, 19, 208, 0, 0}, opReturnId, "unexpected op_return_id")
}
