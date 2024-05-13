package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddPlan(ctx sdk.Context, plan types.Plan) error {
	// generate the next plan ID
	planId := k.GetNextNumber(ctx)
	plan.Id = planId
	// set the plan
	k.setPlan(ctx, planId, plan)
	return nil
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
	//todo: implement
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
	//todo: implement
	return ""
}

// setPlan sets a plan in the Keeper's store.
//
// ctx - the SDK context.
// plan - the plan to be set.
func (k Keeper) setPlan(ctx sdk.Context, planId uint64, plan types.Plan) {
	//todo: implement
}

// setNextNumber sets the next number in the Keeper's store.
//
// ctx - the SDK context.
// number - the number to be set.
func (k Keeper) setNextNumber(ctx sdk.Context, number uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyNextNumber(), sdk.Uint64ToBigEndian(number))
}
