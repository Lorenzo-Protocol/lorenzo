package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/contracts"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeployYATContract creates and deploys a YAT contract on the EVM with the
// plan module account as owner.
func (k Keeper) DeployYATContract(
	ctx sdk.Context,
	deployer common.Address,
	name,
	symbol,
	planDescUri string,
	planId uint64,
	agentId uint64,
	subscriptionStartTime uint64,
	subscriptionEndTime uint64,
	endTime uint64,
) (common.Address, error) {
	contractArgs, err := contracts.YieldAccruingTokenContract.ABI.Pack(
		"",
		name,
		symbol,
		planDescUri,
		planId,
		agentId,
		subscriptionStartTime,
		subscriptionEndTime,
		endTime,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "failed to pack contract arguments: %s", err)
	}
	data := make([]byte, len(contracts.YieldAccruingTokenContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.YieldAccruingTokenContract.Bin)], contracts.YieldAccruingTokenContract.Bin)
	copy(data[len(contracts.YieldAccruingTokenContract.Bin):], contractArgs)
	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	// generate contract address
	contractAddr := crypto.CreateAddress(deployer, nonce)
	_, err = k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", name)
	}
	return contractAddr, nil
}

// Mint creates a new plan token and mints it to the specified address.
func (k Keeper) Mint(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
	amount uint64,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		true,
		types.YATMethodMint,
		// args
		accountAddr,
		amount,
	)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ClaimRewardAndWithDrawBTC(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		true,
		types.YATMethodClaimRewardAndWithDrawBTC,
		// args
		accountAddr,
	)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) onlyClaimReward(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		true,
		types.YATMethodOnlyClaimReward,
		// args
		accountAddr,
	)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) burnWithstBTCBurn(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
	amount uint64,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		true,
		types.YATMethodBurnWithstBTCBurn,
		// args
		accountAddr,
		amount,
	)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) setRewardTokenAddress(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	rewardTokenAddress common.Address,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		true,
		types.YATMethodSetRewardTokenAddress,
		// args
		rewardTokenAddress,
	)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) PlanId(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodPlanId,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodPlanId, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, err
	}
	planId, ok := unpacked[0].(uint64)
	if !ok {
		return 0, err
	}

	return planId, nil
}

func (k Keeper) AgentId(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodAgentId,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodAgentId, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, err
	}
	agentId, ok := unpacked[0].(uint64)
	if !ok {
		return 0, err
	}

	return agentId, nil
}

func (k Keeper) SubscriptionStartTime(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodSubscriptionStartTime,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodSubscriptionStartTime, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, err
	}
	subscriptionStartTime, ok := unpacked[0].(uint64)
	if !ok {
		return 0, err
	}

	return subscriptionStartTime, nil
}

func (k Keeper) SubscriptionEndTime(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodSubscriptionEndTime,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodSubscriptionEndTime, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, err
	}
	subscriptionEndTime, ok := unpacked[0].(uint64)
	if !ok {
		return 0, err
	}

	return subscriptionEndTime, nil
}

func (k Keeper) EndTime(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodEndTime,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodEndTime, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, err
	}
	endTime, ok := unpacked[0].(uint64)
	if !ok {
		return 0, err
	}

	return endTime, nil
}

func (k Keeper) PlanDesc(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (string, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodPlanDesc,
	)
	if err != nil {
		return "", err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodPlanDesc, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", err
	}
	planDesc, ok := unpacked[0].(string)
	if !ok {
		return "", err
	}

	return planDesc, nil
}

func (k Keeper) RewardTokenAddress(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (common.Address, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		types.ModuleAddress,
		contractAddress,
		false,
		types.YATMethodRewardTokenAddress,
	)
	if err != nil {
		return common.Address{}, err
	}
	unpacked, err := contractABI.Unpack(types.YATMethodRewardTokenAddress, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return common.Address{}, err
	}
	rewardTokenAddress, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, err
	}

	return rewardTokenAddress, nil
}
