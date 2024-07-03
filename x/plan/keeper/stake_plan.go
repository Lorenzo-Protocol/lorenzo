package keeper

import (
	"fmt"
	"math/big"

	contractsplan "github.com/Lorenzo-Protocol/lorenzo/contracts/plan"

	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeployStakePlanProxyContract deploys a new Stake Plan Proxy contract.
//
// Parameters:
// - ctx: the SDK context.
// - stakePlanName: the name of the stake plan.
// - planDescUri: the URI of the plan description.
// - planId: the plan ID.
// - agentId: the agent ID.
// - planStartBlock: the start block of the plan.
// - periodBlocks: the period blocks of the plan.
// - yatContractAddress: the address of the YAT contract.
//
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployStakePlanProxyContract(
	ctx sdk.Context,
	stakePlanName,
	planDescUri string,
	planId *big.Int,
	agentId *big.Int,
	planStartBlock *big.Int,
	periodBlocks *big.Int,
	yatContractAddress common.Address,
) (common.Address, error) {
	// pack contract arguments
	initArgs, err := contractsplan.StakePlanContract.ABI.Pack(
		types.StakePlanMethodInitialize,
		// args
		stakePlanName,
		planDescUri,
		planId,
		agentId,
		planStartBlock,
		periodBlocks,
		yatContractAddress,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(types.ErrABIPack, fmt.Sprintf("failed to pack contract arguments: %s", err))
	}

	params := k.GetParams(ctx)
	if params.Beacon == "" {
		return common.Address{}, errorsmod.Wrap(types.ErrBeaconNotSet, "beacon not set")
	}

	// pack proxy contract arguments
	contractArgs, err := contractsplan.StakePlanProxyContract.ABI.Pack(
		"",
		common.HexToAddress(params.Beacon),
		initArgs,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrap(types.ErrABIPack, fmt.Sprintf("failed to pack contract arguments: %s", err))
	}

	data := make([]byte, len(contractsplan.StakePlanProxyContract.Bin)+len(contractArgs))
	copy(data[:len(contractsplan.StakePlanProxyContract.Bin)], contractsplan.StakePlanProxyContract.Bin)
	copy(data[len(contractsplan.StakePlanProxyContract.Bin):], contractArgs)

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
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy contract for %s", stakePlanName)
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution,
			"failed to deploy contract for %s, reason: %s", stakePlanName, result.Revert())
	}

	return contractAddr, nil
}

// DeployStakePlanLogicContract deploys a new Stake Plan Logic contract.
//
// Parameters:
// - ctx: the SDK context.
// Returns:
// - common.Address: the address of the deployed contract.
// - error: an error if the deployment fails.
func (k Keeper) DeployStakePlanLogicContract(
	ctx sdk.Context,
) (common.Address, error) {
	// pack proxy contract arguments
	contractArgs, err := contractsplan.StakePlanContract.ABI.Pack(
		"",
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIPack, "contract arguments are invalid: %s", err.Error())
	}
	data := make([]byte, len(contractsplan.StakePlanContract.Bin)+len(contractArgs))
	copy(data[:len(contractsplan.StakePlanContract.Bin)], contractsplan.StakePlanContract.Bin)
	copy(data[len(contractsplan.StakePlanContract.Bin):], contractArgs)

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
			"failed to deploy contract for stake plan logic contract")
	}
	if result.Failed() {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrVMExecution,
			"failed to deploy contract for stake plan logic contract, reason: %s", result.Revert())
	}

	return contractAddr, nil
}

// ClaimYATToken claims YAT tokens.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - accountAddr: the address of the account to claim tokens for.
// - amount: the amount of tokens to claim.
// - merkleProof: the Merkle proof of the claim.
//
// Returns:
// - error: an error if the claim fails.
func (k Keeper) ClaimYATToken(
	ctx sdk.Context,
	contractAddress common.Address,
	account common.Address,
	roundId *big.Int,
	amount *big.Int,
	merkleProof string,
) error {
	merkleProofBytes := common.HexToHash(merkleProof)

	contractABI := contractsplan.StakePlanContract.ABI
	_, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.StakePlanMethodClaimYATToken,
		// args
		account,
		roundId,
		amount,
		merkleProofBytes,
	)
	if err != nil {
		return err
	}
	return nil
}

// MintFromStakePlan mints YAT tokens to an account.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the YAT contract.
// - to: the address of the account to mint tokens for.
// - amount: the amount of tokens to mint.
//
// Returns:
// - error: an error if the mint fails.
func (k Keeper) MintFromStakePlan(
	ctx sdk.Context,
	contractAddress, to common.Address,
	amount *big.Int,
) error {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.StakePlanMethodMint,
		// args
		to,
		amount,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to stake plan contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

// SetMerkleRoot sets the Merkle root of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
// - merkleProof: the Merkle proof to set.
//
// Returns:
// - error: an error if the setting fails.
func (k Keeper) SetMerkleRoot(
	ctx sdk.Context,
	contractAddress common.Address,
	merkleProof string,
) error {
	merkleProofBytes := common.HexToHash(merkleProof)

	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.StakePlanMethodClaimYATToken,
		// args
		merkleProofBytes,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to stake plan contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

// AdminPauseBridge pauses the bridge of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - error: an error if the pausing fails.
func (k Keeper) AdminPauseBridge(
	ctx sdk.Context,
	contractAddress common.Address,
) error {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.StakePlanMethodAdminPauseBridge,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to stake plan contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

// AdminUnpauseBridge unpauses the bridge of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - error: an error if the unpausing fails.
func (k Keeper) AdminUnpauseBridge(
	ctx sdk.Context,
	contractAddress common.Address,
) error {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.StakePlanMethodAdminUnpauseBridge,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to stake plan contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

// SetPlanDesc sets the description of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
// - planDesc: the description to set.
//
// Returns:
// - error: an error if the setting fails.
func (k Keeper) SetPlanDesc(
	ctx sdk.Context,
	contractAddress common.Address,
	planDesc string,
) error {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		true,
		types.StakePlanMethodSetPlanDesc,
		// args
		planDesc,
	)
	if err != nil {
		return err
	}
	if res.Failed() {
		return errorsmod.Wrapf(
			types.ErrVMExecution, "failed to stake plan contract: %s, reason: %s",
			contractAddress.String(),
			res.Revert(),
		)
	}
	return nil
}

// StakePlanName gets the name of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - string: the name of the StakePlan contract.
// - error: an error if the getting fails.
func (k Keeper) StakePlanName(
	ctx sdk.Context,
	contractAddress common.Address,
) (string, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodStakePlanName,
	)
	if err != nil {
		return "", err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodStakePlanName, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack stake plan name from contract %s", contractAddress.Hex(),
		)
	}
	stakePlanName, ok := unpacked[0].(string)
	if !ok {
		return "", errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert stake plan name to string from contract %s", contractAddress.Hex(),
		)
	}

	return stakePlanName, nil
}

// PlanId gets the plan id from the YAT contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the ABI of the YAT contract.
//
// Returns:
// - uint64: the plan ID.
// - error: an error if the getting fails.
func (k Keeper) PlanId(
	ctx sdk.Context,
	contractAddress common.Address,
) (uint64, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodPlanId,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodPlanId, res.Ret)
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
	contractAddress common.Address,
) (uint64, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodAgentId,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodAgentId, res.Ret)
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
	contractAddress common.Address,
) (string, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodPlanDesc,
	)
	if err != nil {
		return "", err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodPlanDesc, res.Ret)
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

// PlanStartBlock gets the start block of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - uint64: the start block.
// - error: an error if the getting fails.
func (k Keeper) PlanStartBlock(
	ctx sdk.Context,
	contractAddress common.Address,
) (uint64, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodPlanStartBlock,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodPlanStartBlock, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack Plan Start Block from contract %s", contractAddress.Hex(),
		)
	}
	planStartBlock, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert Plan Start Block to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return planStartBlock.Uint64(), nil
}

// PeriodBlocks gets the period blocks of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - uint64: the period blocks.
// - error: an error if the getting fails.
func (k Keeper) PeriodBlocks(
	ctx sdk.Context,
	contractAddress common.Address,
) (uint64, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodPeriodBlocks,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodPeriodBlocks, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack period blocks from contract %s", contractAddress.Hex(),
		)
	}
	periodBlocks, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert period blocks  to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return periodBlocks.Uint64(), nil
}

// NextRewardReceiveBlock gets the next reward receive block of the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - uint64: the next reward receive block.
// - error: an error if the getting fails.
func (k Keeper) NextRewardReceiveBlock(
	ctx sdk.Context,
	contractAddress common.Address,
) (uint64, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodNextRewardReceiveBlock,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodNextRewardReceiveBlock, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack Next Reward Receive Block from contract %s", contractAddress.Hex(),
		)
	}
	nextRewardReceiveBlock, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert Next Reward Receive Block to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return nextRewardReceiveBlock.Uint64(), nil
}

// YatContractAddress gets the YAT contract address from the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - common.Address: the address of the YAT contract.
// - error: an error if the getting fails.
func (k Keeper) YatContractAddress(
	ctx sdk.Context,
	contractAddress common.Address,
) (common.Address, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodYatContractAddress,
	)
	if err != nil {
		return common.Address{}, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodYatContractAddress, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack YAT contract address from contract %s", contractAddress.Hex(),
		)
	}
	yatContractAddress, ok := unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert YAT contract address to address from contract %s", contractAddress.Hex(),
		)
	}

	return yatContractAddress, nil
}

// ClaimRoundId gets the claim round ID from the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - uint64: the claim round ID.
// - error: an error if the getting fails.
func (k Keeper) ClaimRoundId(
	ctx sdk.Context,
	contractAddress common.Address,
) (uint64, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodClaimRoundId,
	)
	if err != nil {
		return 0, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodClaimRoundId, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack Claim Round Id from contract %s", contractAddress.Hex(),
		)
	}
	claimRoundId, ok := unpacked[0].(*big.Int)
	if !ok {
		return 0, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert Claim Round Id to uint64 from contract %s", contractAddress.Hex(),
		)
	}

	return claimRoundId.Uint64(), nil
}

// MerkleRoot gets the Merkle root from the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
//
// Returns:
// - string: the Merkle root.
// - error: an error if the getting fails.
func (k Keeper) MerkleRoot(
	ctx sdk.Context,
	contractAddress common.Address,
	roundId *big.Int,
) (string, error) {
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodMerkleRoot,
		roundId,
	)
	if err != nil {
		return "", err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodMerkleRoot, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return "", errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack Merkle Root from contract %s", contractAddress.Hex(),
		)
	}

	// unpacked to bytes32
	merkleRoot, ok := unpacked[0].(string)
	if !ok {
		return "", errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to convert Merkle Root to string from contract %s", contractAddress.Hex(),
		)
	}

	return merkleRoot, nil
}

// ClaimLeafNodeFromPlan claims the leaf node from the StakePlan contract.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddress: the address of the StakePlan contract.
// - roundId: the round ID.
// - leafNode: the leaf node to claim.
//
// Returns:
// - bool: true if the claim is successful.
// - error: an error if the claiming fails.
func (k Keeper) ClaimLeafNodeFromPlan(
	ctx sdk.Context,
	contractAddress common.Address,
	roundId *big.Int,
	leafNode string,
) (bool, error) {
	leafNodeBytes := common.HexToHash(leafNode)
	contractABI := contractsplan.StakePlanContract.ABI
	res, err := k.CallEVM(
		ctx,
		contractABI,
		k.getModuleEthAddress(ctx),
		contractAddress,
		false,
		types.StakePlanMethodClaimLeafNode,
		roundId,
		leafNodeBytes,
	)
	if err != nil {
		return false, err
	}
	unpacked, err := contractABI.Unpack(types.StakePlanMethodMerkleRoot, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return false, errorsmod.Wrapf(
			types.ErrABIUnpack, "failed to unpack Merkle Root from contract %s", contractAddress.Hex(),
		)
	}

	return true, nil
}
