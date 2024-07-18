package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// UpdateHeader implements types.MsgServer.
func (m msgServer) UpdateHeader(context.Context, *types.MsgUpdateHeader) (*types.MsgUpdateHeaderResponse, error) {
	panic("unimplemented")
}

// UpdateParams implements types.MsgServer.
func (m msgServer) UpdateParams(context.Context, *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	panic("unimplemented")
}

// UploadHeaders implements types.MsgServer.
func (m msgServer) UploadHeaders(context.Context, *types.MsgUploadHeaders) (*types.MsgUploadHeadersResponse, error) {
	panic("unimplemented")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}


