package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/query"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
)

type Querier struct {
	*Keeper
}

// Params queries the parameters of the module.
func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Plans queries all plans.
func (q Querier) Plans(goCtx context.Context, req *types.PlansRequest) (*types.PlansResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	var plans []types.Plan
	var pageRes *query.PageResponse
	var err error
	planStore := q.GetNodesPrefixStore(ctx)
	if pageRes, err = query.Paginate(planStore, req.Pagination, func(_ []byte, value []byte) error {
		var plan types.Plan
		q.cdc.MustUnmarshal(value, &plan)
		plans = append(plans, plan)
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.PlansResponse{
		Plans:      plans,
		Pagination: pageRes,
	}, nil
}

// Plan queries a plan by id.
func (q Querier) Plan(goCtx context.Context, req *types.PlanRequest) (*types.PlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	plan, found := q.GetPlan(ctx, req.Id)
	if !found {
		return nil, types.ErrPlanNotFound
	}

	return &types.PlanResponse{Plan: plan}, nil
}

// NewQuerierImpl returns an implementation of the captains QueryServer interface.
func NewQuerierImpl(k *Keeper) types.QueryServer {
	return &Querier{k}
}
