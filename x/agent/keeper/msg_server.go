package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns a new instance of msgServer with the provided Keeper.
//
// Parameter(s):
// - k: Keeper
// Return type(s): msgServer
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return msgServer{
		k: k,
	}
}

type msgServer struct {
	// This should be a reference to Keeper
	k Keeper
}

// AddAgent description of the Go function.
//
// AddAgent adds an agent to the msgServer.
// It takes the following parameter(s):
// - gctx: the context.Context object representing the context of the function.
// - msg: a pointer to the types.MsgAddAgent object representing the message to be added.
//
// It returns a pointer to the types.MsgAddAgentResponse object and an error.
func (m msgServer) AddAgent(goctx context.Context, msg *types.MsgAddAgent) (*types.MsgAddAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnAuthorized, "invalid sender :%s, not authorized", msg.Sender)
	}

	agentID := m.k.AddAgent(ctx, msg.Name, msg.BtcReceivingAddress, msg.Description, msg.Url)
	return &types.MsgAddAgentResponse{
		Id: agentID,
	}, nil
}

// EditAgent description of the Go function.
//
// EditAgent edits an existing agent in the msgServer.
// It takes the following parameter(s):
// - goctx: the context.Context object representing the context of the function.
// - msg: a pointer to the types.MsgEditAgent object representing the agent to be edited.
//
// It returns a pointer to the types.MsgEditAgentResponse object and an error.
func (m msgServer) EditAgent(goctx context.Context, msg *types.MsgEditAgent) (*types.MsgEditAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnAuthorized, "invalid sender :%s, not authorized", msg.Sender)
	}

	agent, has := m.k.GetAgent(ctx, msg.Id)
	if !has {
		return nil, errorsmod.Wrapf(types.ErrAgentNotFound, "not found agent:%d", msg.Id)
	}

	if msg.BtcReceivingAddress != types.DoNotModifyDesc {
		agent.BtcReceivingAddress = msg.BtcReceivingAddress
	}
	if msg.Description != types.DoNotModifyDesc {
		agent.Description = msg.Description
	}
	if msg.Url != types.DoNotModifyDesc {
		agent.Url = msg.Url
	}
	m.k.setAgent(ctx, agent)
	return &types.MsgEditAgentResponse{}, nil
}

// RemoveAgent removes an agent from the msgServer.
//
// It takes the following parameter(s):
// - gctx: the context.Context object representing the context of the function.
// - msg: a pointer to the types.MsgRemoveAgent object representing the agent to be removed.
//
// It returns a pointer to the types.MsgRemoveAgentResponse object and an error.
func (m msgServer) RemoveAgent(goctx context.Context, msg *types.MsgRemoveAgent) (*types.MsgRemoveAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	if !m.k.Authorized(ctx, sender) {
		return nil, errorsmod.Wrapf(types.ErrUnAuthorized, "invalid sender :%s, not authorized", msg.Sender)
	}

	_, has := m.k.GetAgent(ctx, msg.Id)
	if !has {
		return nil, errorsmod.Wrapf(types.ErrAgentNotFound, "not found agent:%d", msg.Id)
	}

	m.k.removeAgent(ctx, msg.Id)
	return &types.MsgRemoveAgentResponse{}, nil
}
