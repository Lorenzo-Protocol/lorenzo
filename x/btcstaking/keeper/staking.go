package keeper

import (
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	bnblightclienttypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
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
		amount := new(big.Int).SetBytes(event.StBTCAmount.Bytes())
		result := types.MintYatSuccess

		// mint yat to the sender
		if err := k.planKeeper.Mint(ctx, event.PlanID, event.Sender, amount); err != nil {
			result = types.MintYatFailed
		}

		totalStBTCAmt = totalStBTCAmt.Add(totalStBTCAmt, amount)
		k.addBTCBStakingRecord(ctx, &types.BTCBStakingRecord{
			EventIdx:      event.Identifier,
			ReceiverAddr:  event.Sender.String(),
			Amount:        math.NewIntFromBigInt(amount),
			ChainId:       event.ChainID,
			MintYatResult: result,
		})
	}

	// mint stBTC to the bridgeAddr
	totalStBTC := sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, sdk.NewIntFromBigInt(totalStBTCAmt)))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, totalStBTC); err != nil {
		return err
	}

	bridgeAddr := sdk.MustAccAddressFromBech32(k.GetParams(ctx).BridgeAddr)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bridgeAddr, totalStBTC); err != nil {
		return err
	}
	return nil
}

func (k Keeper) addBTCBStakingRecord(ctx sdk.Context, record *types.BTCBStakingRecord) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(record)
	store.Set(types.KeyBTCBStakingRecord(record.EventIdx), bz)
}
