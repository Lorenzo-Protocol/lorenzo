package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	// This should be a reference to Keeper
	k *Keeper
}

func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {

	if m.k.authority != req.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", m.k.authority, req.Authority)
	}
	if err := req.Params.Validate(); err != nil {
		return nil, govtypes.ErrInvalidProposalMsg.Wrapf("invalid parameter: %v", err)
	}
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	if err := m.k.SetParams(sdkCtx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (m msgServer) CreatePlan(goCtx context.Context, req *types.MsgCreatePlan) (*types.MsgCreatePlanResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.AccAddress(req.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}
	plan := types.Plan{
		Name:                  req.Name,
		Symbol:                req.Symbol,
		PlanDescUri:           req.PlanDescUri,
		AgentId:               req.AgentId,
		SubscriptionStartTime: req.SubscriptionStartTime,
		SubscriptionEndTime:   req.SubscriptionEndTime,
		EndTime:               req.EndTime,
	}
	planResult, err := m.k.AddPlan(ctx, plan)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		types.NewCreatePlanEvent(sender, planResult),
	)
	return &types.MsgCreatePlanResponse{Id: planResult.Id}, nil
}

func (m msgServer) Claims(goCtx context.Context, req *types.MsgClaims) (*types.MsgClaimsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender := sdk.AccAddress(req.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "unauthorized")
	}

	if err := req.ValidateBasic(); err != nil {
		return nil, err
	}

	return &types.MsgClaimsResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper}
}
