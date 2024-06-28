package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) AddPlan(ctx sdk.Context, plan types.Plan) (types.Plan, error) {
	// check if the agent not exists
	_, agentFound := k.agentKeeper.GetAgent(ctx, plan.AgentId)
	if !agentFound {
		return types.Plan{}, types.ErrAgentNotFound
	}

	// generate the next plan ID
	planId := k.GetNextNumber(ctx)
	plan.Id = planId
	plan.Enabled = types.PlanStatus_Enabled
	// Deploy the plan contract for plan
	planIdBigint := sdk.NewIntFromUint64(planId)
	agentIdBigint := sdk.NewIntFromUint64(plan.AgentId)
	yatContractAddr := common.HexToAddress(plan.YatContractAddress)
	contractAddress, err := k.DeployStakePlanProxyContract(
		ctx,
		plan.Name,
		plan.PlanDescUri,
		planIdBigint.BigInt(),
		agentIdBigint.BigInt(),
		plan.PlanStartBlock.BigInt(),
		plan.PeriodBlocks.BigInt(),
		yatContractAddr,
	)
	if err != nil {
		return types.Plan{}, err
	}
	plan.ContractAddress = contractAddress.Hex()

	// set the plan
	k.setPlan(ctx, plan)
	// increment the next plan ID
	k.setNextNumber(ctx, planId+1)
	return plan, nil
}

func (k Keeper) UpdatePlanStatus(ctx sdk.Context, planId uint64, status types.PlanStatus) error {
	plan, found := k.GetPlan(ctx, planId)
	if !found {
		return types.ErrPlanNotFound
	}

	planAddress := common.HexToAddress(plan.ContractAddress)
	if status == types.PlanStatus_Enabled {
		if err := k.AdminPauseBridge(ctx, planAddress); err != nil {
			return err
		}
	} else if status == types.PlanStatus_Disabled {
		if err := k.AdminUnpauseBridge(ctx, planAddress); err != nil {
			return err
		}
	}
	// update the plan status
	plan.Enabled = status
	k.setPlan(ctx, plan)
	return nil
}

// GetPlan retrieves a plan by the plan ID from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
// - planId: the plan ID.
//
// Returns:
// - types.Plan: the plan, or an empty plan if it is not found in the store.
func (k Keeper) GetPlan(ctx sdk.Context, planId uint64) (types.Plan, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPlan(planId)
	bz := store.Get(key)
	if bz == nil {
		return types.Plan{}, false
	}

	var plan types.Plan
	k.cdc.MustUnmarshal(bz, &plan)
	return plan, true
}

// GetPlans retrieves all plans from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
//
// Returns:
// - []types.Plan: the plans.
func (k Keeper) GetPlans(ctx sdk.Context) []types.Plan {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixPlan)
	defer iterator.Close() // nolint: errcheck

	var plans []types.Plan
	for ; iterator.Valid(); iterator.Next() {
		var plan types.Plan
		k.cdc.MustUnmarshal(iterator.Value(), &plan)
		plans = append(plans, plan)
	}
	return plans
}

// GetNextNumber retrieves the next number from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
//
// Returns:
// - int32: the next number, or 1 if it is not found in the store.
func (k Keeper) GetNextNumber(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyNextNumber())
	if bz == nil {
		return 1
	}

	return sdk.BigEndianToUint64(bz)
}

// GetPlanIdByContractAddr retrieves the plan ID by the contract address from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
// - contractAddr: the contract address.
//
// Returns:
// - uint64: the plan ID, or 0 if it is not found in the store.
func (k Keeper) GetPlanIdByContractAddr(ctx sdk.Context, contractAddr string) uint64 {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPlanContract(contractAddr)
	bz := store.Get(key)
	if bz != nil {
		var plan types.Plan
		k.cdc.MustUnmarshal(bz, &plan)
		return plan.Id
	}
	return 0
}

// GetContractAddrByPlanId retrieves the contract address by the plan ID from the Keeper's store.
//
// Parameters:
// - ctx: the SDK context.
// - planId: the plan ID.
//
// Returns:
// - string: the contract address, or an empty string if it is not found in the store.
func (k Keeper) GetContractAddrByPlanId(ctx sdk.Context, planId uint64) string {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPlan(planId)
	bz := store.Get(key)
	if bz != nil {
		var plan types.Plan
		k.cdc.MustUnmarshal(bz, &plan)
		return plan.ContractAddress
	}
	return ""
}

// GetNodesPrefixStore returns the store for the plans
func (k Keeper) GetNodesPrefixStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.KeyPrefixPlan)
}

// setPlan sets a plan in the Keeper's store.
//
// ctx - the SDK context.
// plan - the plan to be set.
func (k Keeper) setPlan(ctx sdk.Context, plan types.Plan) {
	// contractAddress --> plan
	store := ctx.KVStore(k.storeKey)
	key := types.KeyPlanContract(plan.ContractAddress)
	planBz := k.cdc.MustMarshal(&plan)
	store.Set(key, planBz)
	// planId --> plan
	key = types.KeyPlan(plan.Id)
	store.Set(key, planBz)
}

// setNextNumber sets the next number in the Keeper's store.
//
// ctx - the SDK context.
// number - the number to be set.
func (k Keeper) setNextNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyNextNumber(), sdk.Uint64ToBigEndian(number))
}
