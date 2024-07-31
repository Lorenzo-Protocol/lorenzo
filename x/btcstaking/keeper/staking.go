package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	bnblightclienttypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
)

// DepositBTCB deposits BTCB to the staking module.
//
// ctx: the context of the current blockchain.
// depositor: the address of the depositor.
// receiptBz: the byte array representation of the receipt.
// proofBz: the byte array representation of the proof.
// error: returns an error if there was a problem with the deposit.
func (k Keeper) DepositBTCB(ctx sdk.Context, depositor sdk.AccAddress, receiptBz, proofBz []byte) error {
	proof, err := bnblightclienttypes.UnmarshalProof(proofBz)
	if err != nil {
		return err
	}

	receipt, err := bnblightclienttypes.UnmarshalReceipt(receiptBz)
	if err != nil {
		return err
	}

	event, err := k.bnblcKeeper.VerifyReceiptProof(ctx, receipt, proof)
	if err != nil {
		return err
	}

	bridgeAddr := sdk.MustAccAddressFromBech32(k.GetParams(ctx).BridgeAddr)
	totalStBTCAmt := new(big.Int)
	for _, e := range event {
		// mint yat to the sender
		if err := k.planKeeper.Mint(ctx, e.PlanID, e.Sender, &e.StBTCAmount); err != nil {
			return err
		}

		totalStBTCAmt = totalStBTCAmt.Add(totalStBTCAmt, &e.StBTCAmount)
	}

	// mint stBTC to the bridgeAddr
	totalStBTC := sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, sdk.NewIntFromBigInt(totalStBTCAmt)))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, totalStBTC); err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bridgeAddr, totalStBTC)
	if err != nil {
		return err
	}

	// TODO add BTCBStakingRecord
	// TODO emit BTCBStakingRecord event
	return nil
}
