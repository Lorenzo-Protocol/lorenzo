package keeper

import (
	"math/big"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) Withdraw(
	ctx sdk.Context,
	planId uint64,
	receiver string,
	roundId *big.Int,
	amount *big.Int,
	merkleProof string,
) error {
	receiverEvmAddress := common.HexToAddress(receiver)
	// get contract address
	contractAddress := k.GetContractAddrByPlanId(ctx, planId)
	if len(contractAddress) == 0 {
		return types.ErrContractNotFound
	}
	contractAddressHex := common.HexToAddress(contractAddress)
	return k.ClaimYATToken(
		ctx,
		contractAddressHex,
		receiverEvmAddress,
		roundId,
		amount,
		merkleProof,
	)
}
