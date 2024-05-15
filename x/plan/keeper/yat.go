package keeper

import (
	"fmt"
	"math/big"

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
		return common.Address{}, errorsmod.Wrap(types.ErrABIPack, fmt.Sprintf("failed to pack contract arguments: %s", err))
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

func (k Keeper) OnlyClaimReward(
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

func (k Keeper) BurnWithstBTCBurn(
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

func (k Keeper) SetRewardTokenAddress(
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
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack plan ID from contract %s", contractAddress.Hex(),
		)
	}
	planId, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert plan ID to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return planId.Uint64(), nil
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
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack agent ID from contract %s", contractAddress.Hex(),
		)
	}
	agentId, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert agent ID to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return agentId.Uint64(), nil
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
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack subscription start time from contract %s", contractAddress.Hex(),
		)
	}
	subscriptionStartTime, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert subscription start time to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return subscriptionStartTime.Uint64(), nil
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
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack subscription end time from contract %s", contractAddress.Hex(),
		)
	}
	subscriptionEndTime, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert subscription end time to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return subscriptionEndTime.Uint64(), nil
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
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack end time from contract %s", contractAddress.Hex(),
		)
	}
	endTime, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert end time to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return endTime.Uint64(), nil
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
		return "", errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack plan description from contract %s", contractAddress.Hex(),
		)
	}
	planDesc, ok := unpacked[0].(string)
	if !ok {
		return "", errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert plan description to string from contract %s", contractAddress.Hex(),
		)
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
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack reward token address from contract %s", contractAddress.Hex(),
		)
	}
	rewardTokenAddress, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert reward token address to common.Address from contract %s", contractAddress.Hex(),
		)
	}

	return rewardTokenAddress, nil
}
