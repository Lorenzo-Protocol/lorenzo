package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

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
func (m msgServer) UpdateHeader(goCtx context.Context, msg *types.MsgUpdateHeader) (*types.MsgUpdateHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.Allow(ctx, msg.Sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "address %s is not in allowlist", msg.Sender)
	}

	if err := m.Keeper.UpdateHeader(ctx, msg.ChainId, &msg.Header); err != nil {
		return nil, err
	}
	return &types.MsgUpdateHeaderResponse{}, nil
}

// UpdateParams implements types.MsgServer.
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if m.authority != msg.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", m.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.SetParams(ctx, &msg.Params); err != nil {
		return nil, err
	}
	return &types.MsgUpdateParamsResponse{}, nil
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
func (m msgServer) UploadHeaders(goCtx context.Context, msg *types.MsgUploadHeaders) (*types.MsgUploadHeadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.Allow(ctx, msg.Sender) {
		return nil, errorsmod.Wrapf(types.ErrUnauthorized, "address %s is not in allowlist", msg.Sender)
	}

	if err := m.Keeper.UploadHeaders(ctx, msg.ChainId, msg.Headers); err != nil {
		return nil, err
	}
	return &types.MsgUploadHeadersResponse{}, nil
}
