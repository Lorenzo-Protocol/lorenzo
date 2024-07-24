package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	*Keeper
}

func NewQuerier(k *Keeper) Querier {
	return Querier{k}
}

// TokenPairs implements the Query/TokenPairs gRPC method
func (q Querier) TokenPairs(goCtx context.Context, req *types.QueryTokenPairsRequest) (*types.QueryTokenPairsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pairs []types.TokenPair
	prefixStore := prefix.NewStore(ctx.KVStore(q.storeKey), types.KeyPrefixTokenPair)

	pageRes, err := query.Paginate(prefixStore, req.Pagination, func(_, value []byte) error {
		var pair types.TokenPair
		if err := q.cdc.Unmarshal(value, &pair); err != nil {
			return err
		}
		pairs = append(pairs, pair)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTokenPairsResponse{
		TokenPairs: pairs,
		Pagination: pageRes,
	}, nil
}

// TokenPair implements the Query/TokenPair gRPC method
func (q Querier) TokenPair(goCtx context.Context, req *types.QueryTokenPairRequest) (*types.QueryTokenPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !common.IsHexAddress(req.Token) {
		if err := sdk.ValidateDenom(req.Token); err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"invalid token format: %s, should be either hex addr or coin denom", req.Token,
			)
		}
	}

	id := q.GetTokenPairId(ctx, req.Token)
	pair, found := q.GetTokenPair(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "token pair for token '%s' not found", req.Token)
	}

	return &types.QueryTokenPairResponse{TokenPair: pair}, nil
}

// Params implements the Query/Params gRPC method
func (q Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
