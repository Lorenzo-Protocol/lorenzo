package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
)

// Params queries the parameters of the module.
func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Plans queries all plans.
func (k Keeper) Plans(goCtx context.Context, req *types.PlansRequest) (*types.PlansResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Plan queries a plan by id.
func (k Keeper) Plan(goCtx context.Context, req *types.PlanRequest) (*types.PlanResponse, error) {
	//TODO implement me
	panic("implement me")
}
