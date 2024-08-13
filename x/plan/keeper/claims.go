package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const EmptyMerkleRoot = "0x0000000000000000000000000000000000000000000000000000000000000000"

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
	// check if merkle root not set
	merkelRoot, err := k.MerkleRoot(ctx, contractAddressHex, roundId)
	if err != nil || EmptyMerkleRoot == merkelRoot {
		return errorsmod.Wrapf(types.ErrMerkelRootIsInvalid, "merkelRoot: %s", merkelRoot)
	}
	return k.ClaimYATToken(
		ctx,
		contractAddressHex,
		receiverEvmAddress,
		roundId,
		amount,
		merkleProof,
	)
}
