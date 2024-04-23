package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	// This should be a reference to Keeper
	k *Keeper
}

func (m msgServer) UpdateParams(ctx context.Context, params *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) CreatePlan(ctx context.Context, request *types.CreatePlanRequest) (*types.CreatePlanResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) Claims(ctx context.Context, request *types.ClaimsRequest) (*types.ClaimsResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{keeper}
}
