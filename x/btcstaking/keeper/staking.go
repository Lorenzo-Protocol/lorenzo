package keeper

import (
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	bnblightclienttypes "github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
)

// DepositBTCB deposits BTCB to the staking module.
//
// ctx: the context of the current blockchain.
// depositor: the address of the depositor.
// number: the block height of the receipt.
// receiptBz: the byte array representation of the receipt.
// proofBz: the byte array representation of the proof.
// error: returns an error if there was a problem with the deposit.
func (k Keeper) DepositBTCB(
	ctx sdk.Context,
	depositor sdk.AccAddress,
	number uint64,
	receiptBz,
	proofBz []byte,
) error {
	proof, err := bnblightclienttypes.UnmarshalProof(proofBz)
	if err != nil {
		return err
	}

	receipt, err := bnblightclienttypes.UnmarshalReceipt(receiptBz)
	if err != nil {
		return err
	}

	events, err := k.bnblcKeeper.VerifyReceiptProof(ctx, number, receipt, proof)
	if err != nil {
		return err
	}

	totalStBTCAmt := new(big.Int)
	for i := range events {
		event := events[i]
		if k.hasBTCBStakingRecord(ctx, event.ChainID, event.Contract.Bytes(), event.Identifier) {
			return types.ErrDuplicateStakingEvent.Wrapf("duplicate event,planID %d,stakingIdx %d,contract %s",
				event.PlanID,
				event.Identifier,
				event.Contract.String(),
			)
		}
		amount := new(big.Int).SetBytes(event.StBTCAmount.Bytes())
		result := ""

		// TODO: Mint YAT yet to be implemented
		// result := types.MintYatSuccess

		//// mint yat to the sender
		//if err := k.planKeeper.Mint(ctx, event.PlanID, event.Sender, amount); err != nil {
		//	k.Logger(ctx).Error("mint yat error",
		//		"planID", event.PlanID,
		//		"stakingIdx", event.Identifier,
		//		"contract", event.Contract.String(),
		//		"sender", event.Sender.String(),
		//		"amount", amount.String(),
		//		"error", err,
		//	)
		//	result = types.MintYatFailed
		//}

		totalStBTCAmt = totalStBTCAmt.Add(totalStBTCAmt, amount)

		btcbStakingRecord := &types.BTCBStakingRecord{
			StakingIdx:    event.Identifier,
			Contract:      event.Contract.Bytes(),
			ReceiverAddr:  event.Sender.String(),
			Amount:        math.NewIntFromBigInt(amount),
			ChainId:       event.ChainID,
			MintYatResult: result,
			PlanId:        event.PlanID,
		}

		k.addBTCBStakingRecord(ctx, btcbStakingRecord)

		// emit an event
		ctx.EventManager().EmitTypedEvent(types.NewEventBTCBStakingCreated(btcbStakingRecord)) //nolint:errcheck,gosec

	}

	// mint stBTC to the bridgeAddr
	totalStBTC := sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, sdk.NewIntFromBigInt(totalStBTCAmt)))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, totalStBTC); err != nil {
		return err
	}

	bridgeAddr := sdk.AccAddress(common.HexToAddress(k.GetParams(ctx).BridgeAddr).Bytes())
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bridgeAddr, totalStBTC); err != nil {
		return err
	}
	return nil
}

func (k Keeper) addBTCBStakingRecord(ctx sdk.Context, record *types.BTCBStakingRecord) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(record)
	store.Set(types.KeyBTCBStakingRecord(record.ChainId, record.Contract, record.StakingIdx), bz)
}

func (k Keeper) hasBTCBStakingRecord(ctx sdk.Context, chainID uint32, contract []byte, stakingIdx uint64) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyBTCBStakingRecord(chainID, contract, stakingIdx)
	return store.Has(key)
}

func (k Keeper) getBTCBStakingRecord(ctx sdk.Context, chainID uint32, contract []byte, stakingIdx uint64) (*types.BTCBStakingRecord, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyBTCBStakingRecord(chainID, contract, stakingIdx)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, types.ErrStakingRecordNotFound
	}
	var record types.BTCBStakingRecord
	k.cdc.MustUnmarshal(bz, &record)
	return &record, nil
}
