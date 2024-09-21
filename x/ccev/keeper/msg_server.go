package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	*Keeper
}

// CreateClient implements types.MsgServer.
func (m msgServer) CreateClient(goCtx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.Allow(ctx, msg.Sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "address %s is not in allowlist", msg.Sender)
	}

	if err := m.Keeper.CreateClient(ctx, &msg.Client); err != nil {
		return nil, err
	}
	return &types.MsgCreateClientResponse{}, nil
}

// UpdateHeader implements types.MsgServer.
func (m msgServer) UpdateHeader(context.Context, *types.MsgUpdateHeader) (*types.MsgUpdateHeaderResponse, error) {
	panic("unimplemented")
}

// UpdateParams implements types.MsgServer.
func (m msgServer) UpdateParams(context.Context, *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	panic("unimplemented")
}

// UploadCrossChainContract implements types.MsgServer.
func (m msgServer) UploadCrossChainContract(goCtx context.Context, msg *types.MsgUploadCrossChainContract) (*types.MsgUploadCrossChainContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.Allow(ctx, msg.Address) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "address %s is not in allowlist", msg.Address)
	}

	m.UploadContract(ctx, msg.ChainId, msg.Address, msg.EventName, msg.Abi)
	return &types.MsgUploadCrossChainContractResponse{}, nil
}

// UploadHeaders implements types.MsgServer.
func (m msgServer) UploadHeaders(context.Context, *types.MsgUploadHeaders) (*types.MsgUploadHeadersResponse, error) {
	panic("unimplemented")
}
