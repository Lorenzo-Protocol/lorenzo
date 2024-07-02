package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	*Keeper
}

func (q Querier) TokenPairs(ctx context.Context, request *types.QueryTokenPairsRequest) (*types.QueryTokenPairsResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (q Querier) TokenPair(ctx context.Context, request *types.QueryTokenPairRequest) (*types.QueryTokenPairResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (q Querier) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	// TODO implement me
	panic("implement me")
}
