package keeper

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/contracts"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeployYATProxyContract deploys a new Yield Accruing Token (YAT) contract.
//
// Parameters:
// - ctx: the SDK context.
// - deployer: the address of the account deploying the contract.
// - name: the name of the YAT contract.
// - symbol: the symbol of the YAT contract.
// - planDescUri: the URI of the plan description.
// - planId: the ID of the plan.
// - agentId: the ID of the agent.
// - subscriptionStartTime: the start time of the subscription.
// - subscriptionEndTime: the end time of the subscription.
// - endTime: the end time of the plan.
// - merkleRoot: the Merkle root of the plan.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployYATProxyContract(
	ctx sdk.Context,
	name,
	symbol,
	planDescUri string,
	planId uint64,
	agentId uint64,
	subscriptionStartTime uint64,
	subscriptionEndTime uint64,
	endTime uint64,
	merkleRoot string,
) (common.Address, error) {
	merkleRootBytes, _ := hexutil.Decode(merkleRoot)

	// pack contract arguments
	initArgs, err := contracts.YieldAccruingTokenContract.ABI.Pack(
		types.YATMethodInitialize,
		name,
		symbol,
		planDescUri,
		planId,
		agentId,
		subscriptionStartTime,
		subscriptionEndTime,
		endTime,
		merkleRootBytes,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(types.ErrABIPack, fmt.Sprintf("failed to pack contract arguments: %s", err))
	}

	params := k.GetParams(ctx)
	if params.Beacon == "" {
		return common.Address{}, errorsmod.Wrap(types.ErrBeaconNotSet, "beacon not set")
	}

	// pack proxy contract arguments
	contractArgs, err := contracts.YATProxyContract.ABI.Pack(
		"",
		common.HexToAddress(params.Beacon),
		initArgs,
	)

	data := make([]byte, len(contracts.YieldAccruingTokenContract.Bin)+len(contractArgs))
	copy(data[:len(contracts.YieldAccruingTokenContract.Bin)], contracts.YieldAccruingTokenContract.Bin)
	copy(data[len(contracts.YieldAccruingTokenContract.Bin):], contractArgs)

	// deployer is the module address
	deployer := k.getModuleEthAddress(ctx)
	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	// generate contract address
	contractAddr := crypto.CreateAddress(deployer, nonce)
	result, err := k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", name)
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(types.ErrVMExecution, "failed to deploy contract for %s, reason: %s", name, result.Revert())
	}

	return contractAddr, nil
}

// DeployYATLogicContract deploys a new Yield Accruing Token (YAT) contract.
//
// Parameters:
// - ctx: the SDK context.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployYATLogicContract(
	ctx sdk.Context,
) (common.Address, error) {
	data := make([]byte, len(contracts.YieldAccruingTokenContract.Bin))
	copy(data[:len(contracts.YieldAccruingTokenContract.Bin)], contracts.YieldAccruingTokenContract.Bin)

	deployer := k.getModuleEthAddress(ctx)
	nonce, err := k.accountKeeper.GetSequence(ctx, deployer.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	// generate contract address
	contractAddr := crypto.CreateAddress(deployer, nonce)
	result, err := k.CallEVMWithData(ctx, deployer, nil, data, true)
	if err != nil {

		return common.Address{}, errorsmod.Wrapf(
			err,
			"failed to deploy contract for yat logic contract")
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution,
			"failed to deploy contract for yat logic contract, reason: %s", result.Revert())
	}

	return contractAddr, nil
}

// Mint mints YAT tokens to an account.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - contractABI: the ABI of the YAT contract.
// - accountAddr: the address of the account to mint tokens to.
// - amount: the amount of tokens to mint.
//
// Returns:
// - error: an error if the minting fails.
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
		k.getModuleEthAddress(ctx),
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

// ClaimYATToken claims YAT tokens.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - contractABI: the ABI of the YAT contract.
// - accountAddr: the address of the account to claim tokens for.
// - amount: the amount of tokens to claim.
// - merkleProof: the Merkle proof of the claim.
//
// Returns:
// - error: an error if the claim fails.
func (k Keeper) ClaimYATToken(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
	amount uint64,
	merkleProof string,
) error {
	merkleProofBytes, _ := hexutil.Decode(merkleProof)
	_, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodClaimYATToken,
		// args
		accountAddr,
		amount,
		merkleProofBytes,
	)
	if err != nil {
		return err
	}
	return nil
}

// ClaimRewardAndWithDrawBTC claims rewards and withdraws BTC.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - contractABI: the ABI of the YAT contract.
// - accountAddr: the address of the account to claim rewards and withdraw BTC for.
// - amount: the amount of BTC to withdraw.
//
// Returns:
// - error: an error if the claim and withdrawal fails.
func (k Keeper) ClaimRewardAndWithDrawBTC(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
	amount uint64,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.YATMethodClaimRewardAndWithDrawBTC,
		// args
		accountAddr,
		amount,
	)
	if err != nil {
		return err
	}
	return nil
}

// OnlyClaimReward claims rewards.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - contractABI: the ABI of the YAT contract.
// - accountAddr: the address of the account to claim rewards for.
//
// Returns:
// - error: an error if the claim fails.
func (k Keeper) OnlyClaimReward(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	accountAddr common.Address,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// BurnWithstBTCBurn burns stBTC and mints YAT tokens.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - contractABI: the ABI of the YAT contract.
// - accountAddr: the address of the account to burn stBTC and mint tokens for.
// - amount: the amount of stBTC to burn.
//
// Returns:
// - error: an error if the burn and mint fails.
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
		k.getModuleEthAddress(ctx),
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

// SetRewardTokenAddress sets the reward token address for the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - contractABI: the ABI of the YAT contract.
// - rewardTokenAddress: the address of the reward token.
//
// Returns:
// - error: an error if the setting fails.
func (k Keeper) SetRewardTokenAddress(
	ctx sdk.Context,
	contractAddress common.Address,
	contractABI abi.ABI,
	rewardTokenAddress common.Address,
) error {
	_, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// PlanId gets the plan id from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - uint64: the plan ID.
// - error: an error if the getting fails.
func (k Keeper) PlanId(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// AgentId gets the agent ID from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - uint64: the agent ID.
// - error: an error if the getting fails.
func (k Keeper) AgentId(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// SubscriptionStartTime gets the subscription start time from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - uint64: the subscription start time.
// - error: an error if the getting fails.
func (k Keeper) SubscriptionStartTime(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// SubscriptionEndTime gets the subscription end time from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - uint64: the subscription end time.
// - error: an error if the getting fails.
func (k Keeper) SubscriptionEndTime(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// EndTime gets the end time from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - uint64: the end time.
// - error: an error if the getting fails.
func (k Keeper) EndTime(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (uint64, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// PlanDesc gets the plan description from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - string: the plan description.
// - error: an error if the getting fails.
func (k Keeper) PlanDesc(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (string, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

// RewardTokenAddress gets the reward token address from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractABI: the address of the YAT contract.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - common.Address: the reward token address.
// - error: an error if the getting fails.
func (k Keeper) RewardTokenAddress(
	ctx sdk.Context,
	contractABI abi.ABI,
	contractAddress common.Address,
) (common.Address, error) {
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
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

func (k Keeper) getModuleEthAddress(ctx sdk.Context) common.Address {
	moduleAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	return common.BytesToAddress(moduleAccount.GetAddress().Bytes())
}
