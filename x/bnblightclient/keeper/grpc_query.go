package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/bnblightclient keeper providing gRPC method
type Querier struct {
	k Keeper
}

// Header implements types.QueryServer.
func (q Querier) Header(context.Context, *types.QueryHeaderRequest) (*types.QueryHeaderResponse, error) {
	panic("unimplemented")
}

// LatestedHeader implements types.QueryServer.
func (q Querier) LatestedHeader(context.Context, *types.QueryLatestedHeaderRequest) (*types.QueryLatestedHeaderResponse, error) {
	panic("unimplemented")
}

// Params implements types.QueryServer.
func (q Querier) Params(context.Context, *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	panic("unimplemented")
}
