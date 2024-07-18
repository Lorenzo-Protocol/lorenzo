package keeper

import (
	"bytes"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec

	// the address capable of executing a MsgUpdateParams message. Typically, this should be the x/gov module account.
	authority string
}

// VerifyReceiptProof verifies the proof of a receipt in the Ethereum blockchain.
//
// Parameters:
// - ctx: the SDK context.
// - receipt: the Ethereum receipt to verify.
// - proof: the proof of the receipt.
//
// Returns:
// - error: an error if the proof verification fails.
func (k Keeper) VerifyReceiptProof(ctx sdk.Context, receipt *evmtypes.Receipt, proof types.Proof) error {
	db := trie.NewDatabase(rawdb.NewMemoryDatabase())
	mpt := trie.NewEmpty(db)
	_ = evmtypes.DeriveSha(evmtypes.Receipts{receipt}, mpt)

	var indexBuf []byte
	indexBuf = rlp.AppendUint64(indexBuf[:0], uint64(0))
	txValue := mpt.Get(indexBuf)

	header, exist := k.GetHeader(ctx, receipt.BlockNumber.Uint64())
	if !exist {
		return errorsmod.Wrapf(types.ErrHeaderNotFound, "header %d not found", header.Number)
	}

	val, err := trie.VerifyProof(common.Hash(header.ReceiptRoot), proof.Index, &proof.Path)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidProof, "invalid receipt proof")
	}

	if bytes.Equal(val, proof.Value) && bytes.Equal(val, txValue) {
		return errorsmod.Wrapf(types.ErrInvalidProof, "invalid receipt proof")
	}

	return nil
}
