package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/bnblightclient/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

// UpdateHeader implements types.MsgServer.
func (m msgServer) UpdateHeader(goCtx context.Context, req *types.MsgUpdateHeader) (*types.MsgUpdateHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.k.GetParams(ctx)
	if !slices.Contains(params.AllowList, req.Signer) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "signer %s is not in allowlist", req.Signer)
	}

	if err := m.k.UpdateHeader(ctx, req.Header); err != nil {
		return nil, err
	}
	return &types.MsgUpdateHeaderResponse{}, nil
}

// UpdateParams implements types.MsgServer.
func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.k.authority != req.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", m.k.authority, req.Authority)
	}

	if err := req.Params.Validate(); err != nil {
		return nil, govtypes.ErrInvalidProposalMsg.Wrapf("invalid parameter: %v", err)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.k.SetParams(ctx, &req.Params); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
}

// UploadHeaders implements types.MsgServer.
func (m msgServer) UploadHeaders(goCtx context.Context, req *types.MsgUploadHeaders) (*types.MsgUploadHeadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.k.GetParams(ctx)
	if !slices.Contains(params.AllowList, req.Signer) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "signer %s is not in allowlist", req.Signer)
	}

	if err := m.k.UploadHeaders(ctx, req.Headers); err != nil {
		return nil, err
	}
	return &types.MsgUploadHeadersResponse{}, nil
}
