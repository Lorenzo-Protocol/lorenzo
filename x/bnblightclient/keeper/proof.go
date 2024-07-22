package keeper

import (
	"bytes"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

// VerifyReceiptProof verifies the receipt proof for a BNB cross-chain event.
//
// ctx - context in which the verification is done
// receipt - the EVM transaction receipt to verify
// proof - the proof object containing the necessary data for verification
// Returns an array of BNBCrossChainEvent and an error if the verification fails.
func (k Keeper) VerifyReceiptProof(ctx sdk.Context, receipt *evmtypes.Receipt, proof types.Proof) ([]types.CrossChainEvent, error) {
	if err := k.verifyProof(ctx, receipt, proof); err != nil {
		return nil, err
	}

	events, err := k.parseEvents(ctx, receipt)
	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return nil, errorsmod.Wrapf(types.ErrInvalidEvent, "invalid receipt, no events found")
	}
	return events, nil
}

// GetAllEvmEvents retrieves all event records stored in the context.
//
// ctx - The context object.
// []*types.EvmEvent - A slice of event records.
func (k Keeper) GetAllEvmEvents(ctx sdk.Context) (events []*types.EvmEvent) {
	store := ctx.KVStore(k.storeKey)

	it := sdk.KVStorePrefixIterator(store, types.KeyPrefixEventRecord)
	defer it.Close() //nolint:errcheck

	for ; it.Valid(); it.Next() {
		var event types.EvmEvent
		k.cdc.MustUnmarshal(it.Value(), &event)
		events = append(events, &event)
	}
	return events
}

func (k Keeper) verifyProof(ctx sdk.Context, receipt *evmtypes.Receipt, proof types.Proof) error {
	if receipt.Status != evmtypes.ReceiptStatusSuccessful {
		return errorsmod.Wrapf(types.ErrInvalidTransaction, "cannot verify failed transactions")
	}

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

	if !bytes.Equal(val, proof.Value) || !bytes.Equal(val, txValue) {
		return errorsmod.Wrapf(types.ErrInvalidProof, "invalid receipt proof")
	}

	return nil
}

func (k Keeper) parseEvents(ctx sdk.Context, receipt *evmtypes.Receipt) ([]types.CrossChainEvent, error) {
	if len(receipt.Logs) == 0 {
		return nil, errorsmod.Wrapf(types.ErrInvalidEvent, "no event log found")
	}

	params := k.GetParams(ctx)
	contractAddr := common.HexToAddress(params.StakePlanHubAddress).Bytes()

	events := make([]types.CrossChainEvent, 0, len(receipt.Logs))
	for _, log := range receipt.Logs {
		if !bytes.Equal(contractAddr, log.Address.Bytes()) {
			continue
		}

		if len(log.Topics) != 4 {
			return nil, errorsmod.Wrapf(
				types.ErrInvalidEvent,
				"event has wrong number of topics, expected 4, actual: %d",
				len(log.Topics),
			)
		}

		eventID := log.Topics[0]
		event, err := types.ABIstakePlanHub().EventByID(eventID)
		if err != nil {
			continue
		}

		if event.Name != params.EventName {
			continue
		}

		// stakeIndex
		identifier := new(big.Int).SetBytes(log.Topics[1].Bytes())
		// planId
		planID := new(big.Int).SetBytes(log.Topics[2].Bytes())
		// sender
		sender := common.BytesToAddress(log.Topics[3].Bytes())
		
		eventArgs, err := types.ABIstakePlanHub().Unpack(event.Name, log.Data)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrInvalidEvent, "failed to unpack %s event", event.Name)
		}

		if len(eventArgs) != 3 {
			return nil, errorsmod.Wrapf(
				types.ErrInvalidEvent,
				"event has wrong number of parameters, expected 2, actual: %d",
				len(eventArgs),
			)
		}
		
		record := &types.EvmEvent{
			BlockNumber: receipt.BlockNumber.Uint64(),
			Identifier:  identifier.Uint64(),
			Contract:    log.Address.Bytes(),
		}
		if k.hasEvmEvent(ctx, record) {
			return nil, errorsmod.Wrapf(types.ErrInvalidEvent, "event identifier %d already exists", identifier.Uint64())
		}
		k.setEvmEvent(ctx, record)

		// btcContractAddress
		btcContractAddress, ok := eventArgs[0].(common.Address)
		if !ok {
			return nil, errorsmod.Wrap(
				types.ErrInvalidEvent,
				"event `btcContractAddress` parameters is invalid, expected `common.Address`",
			)
		}

		// stakeAmount
		stakeAmount, ok := eventArgs[1].(*big.Int)
		if !ok {
			return nil, errorsmod.Wrap(
				types.ErrInvalidEvent,
				"event `stakeAmount` parameters is invalid, expected `*big.Int`",
			)
		}

		// stBTCAmount
		stBTCAmount, ok := eventArgs[2].(*big.Int)
		if !ok {
			return nil, errorsmod.Wrap(
				types.ErrInvalidEvent,
				"event `stBTCAmount` parameters is invalid, expected `*big.Int`",
			)
		}

		bnbEvent := types.CrossChainEvent{
			Identifier:         identifier.Uint64(),
			Sender:             sender,
			PlanID:             planID.Uint64(),
			BTCcontractAddress: btcContractAddress,
			StakeAmount:        stakeAmount,
			StBTCAmount:        stBTCAmount,
		}
		events = append(events, bnbEvent)
	}
	return events, nil
}

func (k Keeper) hasEvmEvent(ctx sdk.Context, record *types.EvmEvent) bool {
	store := ctx.KVStore(k.storeKey)
	key := record.Key()
	return store.Has(key)
}

func (k Keeper) setEvmEvent(ctx sdk.Context, record *types.EvmEvent) {
	store := ctx.KVStore(k.storeKey)
	key := record.Key()

	bz := k.cdc.MustMarshal(record)
	store.Set(key[:], bz)
}
