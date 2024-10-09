package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/ccev keeper providing gRPC method
type Querier struct {
	Keeper
}

// NewQuerier creates a new instance of Querier.
//
// Args:
// - k is the keeper of the module.
//
// Returns:
// - a new instance of Querier.
func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

// Client implements types.QueryServer.
func (q Querier) Client(goCtx context.Context, req *types.QueryClientRequest) (*types.QueryClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	client := q.getClient(ctx, req.ChainId)
	if client == nil {
		return &types.QueryClientResponse{}, errorsmod.Wrapf(types.ErrNotFoundClient, "client %d not found", req.ChainId)
	}
	return &types.QueryClientResponse{Client: client}, nil
}

// Clients implements types.QueryServer.
func (q Querier) Clients(goCtx context.Context, req *types.QueryClientsRequest) (*types.QueryClientsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	clients := q.getAllClients(ctx)
	return &types.QueryClientsResponse{Clients: clients}, nil
}

// Contract implements types.QueryServer.
func (q Querier) Contract(goCtx context.Context, req *types.QueryContractRequest) (*types.QueryContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !q.hasClient(ctx, req.ChainId) {
		return nil, errorsmod.Wrapf(types.ErrNotFoundClient, "client %d not found", req.ChainId)
	}

	contract := q.getContract(ctx, req.ChainId, common.HexToAddress(req.Address))
	if contract == nil {
		return &types.QueryContractResponse{}, errorsmod.Wrapf(types.ErrNotFoundContract, "contract %d not found, cannot update", req.ChainId)
	}
	return &types.QueryContractResponse{Contract: contract}, nil
}

// Header implements types.QueryServer.
func (q Querier) Header(goCtx context.Context, req *types.QueryHeaderRequest) (*types.QueryHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !q.hasClient(ctx, req.ChainId) {
		return nil, errorsmod.Wrapf(types.ErrNotFoundClient, "client %d not found", req.ChainId)
	}

	header, found := q.GetHeader(ctx, req.ChainId, req.Number)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHeaderNotFound, "header %d not found", req.Number)
	}
	return &types.QueryHeaderResponse{Header: header}, nil
}

// HeaderByHash implements types.QueryServer.
func (q Querier) HeaderByHash(goCtx context.Context, req *types.QueryHeaderByHashRequest) (*types.QueryHeaderByHashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !q.hasClient(ctx, req.ChainId) {
		return nil, errorsmod.Wrapf(types.ErrNotFoundClient, "client %d not found", req.ChainId)
	}

	header, found := q.GetHeaderByHash(ctx, req.ChainId, req.Hash)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrHeaderNotFound, "header %x not found", req.Hash)
	}
	return &types.QueryHeaderByHashResponse{Header: header}, nil
}

// LatestHeader implements types.QueryServer.
func (q Querier) LatestHeader(goCtx context.Context, req *types.QueryLatestHeaderRequest) (*types.QueryLatestHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !q.hasClient(ctx, req.ChainId) {
		return nil, errorsmod.Wrapf(types.ErrNotFoundClient, "client %d not found", req.ChainId)
	}

	header, found := q.GetLatestHeader(ctx, req.ChainId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrHeaderNotFound, "latested header not found")
	}
	return &types.QueryLatestHeaderResponse{Header: header}, nil
}

// Params implements types.QueryServer.
func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := q.GetParams(ctx)
	return &types.QueryParamsResponse{Params: *params}, nil
}
