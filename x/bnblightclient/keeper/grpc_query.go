package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/bnblightclient keeper providing gRPC method
type Querier struct {
	k Keeper
}

// Header implements types.QueryServer.
func (q Querier) Header(goCtx context.Context, req *types.QueryHeaderRequest) (*types.QueryHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	header, found := q.k.GetHeader(ctx, req.Number)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHeaderNotFound, "header %d not found", req.Number)
	}
	return &types.QueryHeaderResponse{Header: header}, nil
}

// HeaderByHash implements types.QueryServer.
func (q Querier) HeaderByHash(goCtx context.Context, req *types.QueryHeaderByHashRequest) (*types.QueryHeaderByHashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	header, found := q.k.GetHeaderByHash(ctx, req.Hash)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHeaderNotFound, "header %x not found", req.Hash)
	}
	return &types.QueryHeaderByHashResponse{Header: header}, nil
}

// LatestedHeader implements types.QueryServer.
func (q Querier) LatestedHeader(goCtx context.Context, req *types.QueryLatestedHeaderRequest) (*types.QueryLatestedHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	header, found := q.k.GetLatestedHeader(ctx)
	if !found {
		return nil, errorsmod.Wrap(types.ErrHeaderNotFound, "latested header not found")
	}
	return &types.QueryLatestedHeaderResponse{Header: *header}, nil
}

// Params implements types.QueryServer.
func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
