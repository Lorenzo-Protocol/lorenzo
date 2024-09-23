package keeper

import (
	"bytes"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

// VerifyAndCallback verifies the receipt of a transaction using the provided proof, and
// if the verification is successful, calls the provided callback function for each
// cross-chain event extracted from the receipt.
//
// Parameters:
// - ctx: The SDK context.
// - chainID: The chain ID.
// - number: The block number.
// - receipt: The receipt to be verified.
// - proof: The proof to verify the receipt.
// - callback: A callback function to be called for each cross-chain event.
//
// Returns:
// - error: An error if the verification fails.
func (k Keeper) VerifyAndCallback(
	ctx sdk.Context,
	chainID uint32,
	number uint64,
	receiptBz []byte,
	proofBz []byte,
	handler types.EventHandler,
) error {
	proof, err := types.UnmarshalProof(proofBz)
	if err != nil {
		return err
	}

	receipt, err := types.UnmarshalReceipt(receiptBz)
	if err != nil {
		return err
	}

	if err := k.VerifyReceipt(ctx, chainID, number, receipt, proof); err != nil {
		return err
	}
	return k.handleReceipt(ctx, chainID, receipt, handler)
}

// VerifyReceipt verifies the receipt of a transaction using the provided proof.
//
// Parameters:
// - ctx: The SDK context.
// - number: The block number.
// - receipt: The receipt to be verified.
// - proof: The proof to verify the receipt.
//
// Returns:
// - error: An error if the verification fails.
func (k Keeper) VerifyReceipt(
	ctx sdk.Context,
	chainID uint32,
	number uint64,
	receipt *evmtypes.Receipt,
	proof *types.Proof,
) error {
	if receipt.Status != evmtypes.ReceiptStatusSuccessful {
		return errorsmod.Wrapf(types.ErrInvalidTransaction, "cannot verify failed transactions")
	}

	db := trie.NewDatabase(rawdb.NewMemoryDatabase())
	mpt := trie.NewEmpty(db)
	_ = evmtypes.DeriveSha(evmtypes.Receipts{receipt}, mpt)

	var indexBuf []byte
	indexBuf = rlp.AppendUint64(indexBuf[:0], uint64(0))
	txValue := mpt.Get(indexBuf)

	header, exist := k.GetHeader(ctx, chainID, number)
	if !exist {
		return errorsmod.Wrapf(types.ErrHeaderNotFound, "header %d not found", number)
	}

	val, err := trie.VerifyProof(common.HexToHash(header.ReceiptRoot), proof.Index, &proof.Path)
	if err != nil {
		return errorsmod.Wrapf(types.ErrInvalidProof, "invalid receipt proof")
	}

	if !bytes.Equal(val, proof.Value) || !bytes.Equal(val, txValue) {
		return errorsmod.Wrapf(types.ErrInvalidProof, "invalid receipt proof")
	}

	return nil
}

// handleReceipt parses the given Ethereum receipt and extracts the cross-chain events.
//
// Parameters:
// - ctx: The SDK context.
// - receipt: The Ethereum receipt to be parsed.
// - callback: The callback function that processes the extracted events.
//
// Returns:
// - error: An error if the parsing fails.
func (k Keeper) handleReceipt(
	ctx sdk.Context,
	chainID uint32,
	receipt *evmtypes.Receipt,
	handler types.EventHandler,
) error {
	if len(receipt.Logs) == 0 {
		return errorsmod.Wrapf(types.ErrInvalidEvent, "no event log found")
	}

	var events []*types.Event
	for _, log := range receipt.Logs {
		contract := k.getContract(ctx, chainID, log.Address)
		if contract == nil {
			continue
		}

		if len(log.Topics) == 0 {
			return errorsmod.Wrapf(types.ErrInvalidEvent, "event has wrong number of topics")
		}

		eventID := log.Topics[0]
		abi, err := types.DecodeABI(contract.Address, contract.Abi)
		if err != nil {
			continue
		}

		event, err := abi.EventByID(eventID)
		if err != nil {
			continue
		}

		if event.Name != contract.EventName {
			continue
		}

		eventArgs, err := abi.Unpack(event.Name, log.Data)
		if err != nil {
			return errorsmod.Wrapf(types.ErrInvalidEvent, "failed to unpack %s event", event.Name)
		}

		eventInfo := &types.Event{
			Address: log.Address,
			Topics: log.Topics,
			Args:   eventArgs,
		}
		identify,err := handler.GetUniqueID(ctx, eventInfo)
		if err != nil {
			return errorsmod.Wrapf(types.ErrInvalidEvent, "failed to get unique id: %s", err.Error())
		}

		if k.hasEvent(ctx, chainID, contract.Address, identify) {
			return errorsmod.Wrapf(types.ErrInvalidEvent, "repeated events")
		}
		k.setEvent(ctx, chainID, contract.Address, identify)

		events = append(events, eventInfo)
	}
	return handler.Execute(ctx, chainID, events)
}
