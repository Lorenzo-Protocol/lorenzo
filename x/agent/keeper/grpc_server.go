package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

var _ types.QueryServer = Keeper{}

// Agents retrieves all agents from the Keeper.
//
// The function takes a Go context and a QueryAgentsRequest as parameters.
// It returns a QueryAgentsResponse containing a list of agents and an error.
func (k Keeper) Agents(goctx context.Context, request *types.QueryAgentsRequest) (*types.QueryAgentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	return &types.QueryAgentsResponse{Agents: k.getAgents(ctx)}, nil
}

// Agent retrieves an agent from the Keeper based on the given ID.
//
// Parameters:
// - goctx: the Go context.
// - request: the QueryAgentRequest containing the ID of the agent.
//
// Returns:
// - *types.QueryAgentResponse: the QueryAgentResponse containing the agent with the given ID, or an error if the agent is not found.
func (k Keeper) Agent(goctx context.Context, request *types.QueryAgentRequest) (*types.QueryAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	agent, has := k.GetAgent(ctx, request.Id)
	if !has {
		return nil, status.Errorf(codes.InvalidArgument, "not found agent:%d", request.Id)
	}
	return &types.QueryAgentResponse{Agent: agent}, nil
}

func (k Keeper) getAgents(ctx sdk.Context) (agents []types.Agent) {
	store := ctx.KVStore(k.storeKey)

	it := sdk.KVStorePrefixIterator(store, types.AgentKey)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var agent types.Agent
		k.cdc.MustUnmarshal(it.Value(), &agent)
		agents = append(agents, agent)
	}
	return agents
}
