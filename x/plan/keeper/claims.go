package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/contracts"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) Withdraw(ctx sdk.Context, claimsType types.ClaimsType, planId uint64, receiver string) error {
	receiverEvmAddress := common.HexToAddress(receiver)
	// get contract address
	contractAddress := k.GetContractAddrByPlanId(ctx, planId)
	if len(contractAddress) == 0 {
		return types.ErrContractNotFound
	}
	contractAddressHex := common.HexToAddress(contractAddress)
	yatABI := contracts.YieldAccruingTokenContract.ABI
	// call the evm module to withdraw the reward
	switch claimsType {
	case types.ClaimsType_ONLY_CLAIM_REWARD:
		return k.OnlyClaimReward(
			ctx,
			contractAddressHex,
			yatABI,
			receiverEvmAddress,
		)
	default:
		return k.ClaimRewardAndWithDrawBTC(
			ctx,
			contractAddressHex,
			yatABI,
			receiverEvmAddress,
		)
	}

}
