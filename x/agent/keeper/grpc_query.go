package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	*Keeper
}

func (q Querier) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Agents retrieves all agents from the Keeper.
//
// The function takes a Go context and a QueryAgentsRequest as parameters.
// It returns a QueryAgentsResponse containing a list of agents and an error.
func (q Querier) Agents(goCtx context.Context, req *types.QueryAgentsRequest) (*types.QueryAgentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	var agents []types.Agent
	var pageRes *query.PageResponse
	var err error
	agentStore := q.Keeper.GetAgentsPrefixStore(ctx)
	if pageRes, err = query.Paginate(agentStore, req.Pagination, func(_ []byte, value []byte) error {
		var plan types.Agent
		q.Keeper.cdc.MustUnmarshal(value, &plan)
		agents = append(agents, plan)
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.QueryAgentsResponse{
		Agents:     agents,
		Pagination: pageRes,
	}, nil
}

// Agent retrieves an agent from the Keeper based on the given ID.
//
// Parameters:
// - goctx: the Go context.
// - request: the QueryAgentRequest containing the ID of the agent.
//
// Returns:
// - *types.QueryAgentResponse: the QueryAgentResponse containing the agent with the given ID, or an error if the agent is not found.
func (q Querier) Agent(goCtx context.Context, request *types.QueryAgentRequest) (*types.QueryAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	agent, has := q.Keeper.GetAgent(ctx, request.Id)
	if !has {
		return nil, status.Errorf(codes.InvalidArgument, "not found agent: %d", request.Id)
	}
	return &types.QueryAgentResponse{Agent: agent}, nil
}

// NewQuerierImpl returns an implementation of the captains QueryServer interface.
func NewQuerierImpl(k *Keeper) types.QueryServer {
	return &Querier{k}
}
