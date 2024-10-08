package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	PlanNotFound    = "Plan not found"
	AgentIdNotMatch = "AgentId not match"
	Success         = "Ok"
)

func (k Keeper) Delegate(
	ctx sdk.Context,
	btcStakingRecord *types.BTCStakingRecord,
	mintToAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	btcAmount uint64,
	planId uint64,
	agentId uint64,
) error {
	toMintAmount := sdkmath.NewIntFromUint64(btcAmount).Mul(sdkmath.NewIntFromUint64(SatoshiToStBTCMul))
	coins := []sdk.Coin{
		{
			Denom:  types.NativeTokenDenom,
			Amount: toMintAmount,
		},
	}

	// mint stBTC to module account
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return errorsmod.Wrapf(types.ErrMintToModule, "failed to mint coins: %v", err)
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, mintToAddr, coins); err != nil {
		return errorsmod.Wrapf(types.ErrTransferToAddr, "failed to send coins from module to account: %v", err)
	}

	if planId != 0 {
		// TODO: Mint YAT yet to be implemented
		//// Mint Yat can be wrong if plan not found or agentId not match
		//plan, found := k.planKeeper.GetPlan(ctx, planId)
		//if !found {
		//	btcStakingRecord.MintYatResult = PlanNotFound
		//} else if plan.AgentId != agentId {
		//	btcStakingRecord.MintYatResult = AgentIdNotMatch
		//} else {
		//	// mint yat
		//	yatMintErr := k.planKeeper.Mint(ctx, planId, common.BytesToAddress(receiverAddr), toMintAmount.BigInt())
		//	if yatMintErr != nil {
		//		btcStakingRecord.MintYatResult = yatMintErr.Error()
		//	} else {
		//		btcStakingRecord.MintYatResult = Success
		//	}
		//}
		// Only record planId
		btcStakingRecord.PlanId = planId
	}
	if err := k.AddBTCStakingRecord(ctx, btcStakingRecord); err != nil {
		return errorsmod.Wrapf(types.ErrRecordStaking, "failed to record staking: %v", err)
	}

	return nil
}

func (k Keeper) Undelegate(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coin) error {
	balance := k.bankKeeper.GetBalance(ctx, sender, types.NativeTokenDenom)
	if balance.IsLT(amount) {
		return types.ErrBurnInsufficientBalance
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, []sdk.Coin{amount})
	if err != nil {
		return errorsmod.Wrapf(types.ErrBurn, "failed to send coins from account to module: %v", err)
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, []sdk.Coin{amount})
	if err != nil {
		return errorsmod.Wrapf(types.ErrBurn, "failed to burn coins: %v", err)
	}

	return nil
}
