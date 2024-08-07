package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	PlanNotFount    = "Plan not found"
	AgentIdNotMatch = "AgentId not match"
	Success         = "ok"
)

func (k Keeper) Mint(
	ctx sdk.Context,
	btcStakingRecord *types.BTCStakingRecord,
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

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, coins); err != nil {
		return errorsmod.Wrapf(types.ErrTransferToAddr, "failed to send coins from module to account: %v", err)
	}

	if err := k.addBTCStakingRecord(ctx, btcStakingRecord); err != nil {
		return errorsmod.Wrapf(types.ErrRecordStaking, "failed to record staking: %v", err)
	}

	// mint yat can is error
	plan, found := k.planKeeper.GetPlan(ctx, planId)
	if !found {
		btcStakingRecord.MintYatResult = PlanNotFount
	} else if plan.AgentId != agentId {
		btcStakingRecord.MintYatResult = AgentIdNotMatch
	} else {
		// mint yat
		yatMintErr := k.planKeeper.Mint(ctx, planId, common.BytesToAddress(receiverAddr), toMintAmount.BigInt())
		if yatMintErr != nil {
			btcStakingRecord.MintYatResult = yatMintErr.Error()
		} else {
			btcStakingRecord.MintYatResult = Success
		}
	}

	return nil
}

func (k Keeper) Burn(ctx sdk.Context, sender sdk.AccAddress, amount sdk.Coin) error {
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
