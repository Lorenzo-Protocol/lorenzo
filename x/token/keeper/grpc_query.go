package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/Lorenzo-Protocol/lorenzo/v2/contracts/erc20"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
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

func (q Querier) Balance(goCtx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check token validity
	if !common.IsHexAddress(req.Token) {
		if err := sdk.ValidateDenom(req.Token); err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"invalid token format: %s, should be either hex addr or coin denom", req.Token,
			)
		}
	}

	var (
		accAddress       sdk.AccAddress
		hexAddress       common.Address
		erc20TokenAmount string
	)

	// check account validity and assign address
	if !common.IsHexAddress(req.AccountAddress) {
		addr, err := sdk.AccAddressFromBech32(req.AccountAddress)
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"invalid account address: %s, should be either hex addr or bech32 addr", req.AccountAddress,
			)
		}
		accAddress = addr
		hexAddress = common.BytesToAddress(accAddress.Bytes())
	} else {
		hexAddress = common.HexToAddress(req.AccountAddress)
		accAddress = hexAddress.Bytes()
	}

	// query if token exists
	id := q.GetTokenPairId(ctx, req.Token)
	pair, found := q.GetTokenPair(ctx, id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "token pair for token '%s' not found", req.Token)
	}

	// query erc20 token amount
	erc20ABI := erc20.ERC20MinterBurnerDecimalsContract.ABI // nolint: staticcheck
	erc20Addr := pair.GetERC20ContractAddress()
	erc20Balance := q.ERC20BalanceOf(ctx, erc20ABI, erc20Addr, hexAddress)
	if erc20Balance != nil {
		erc20TokenAmount = erc20Balance.String()
	}

	// query coin balance
	coin := q.bankKeeper.GetBalance(ctx, accAddress, pair.Denom)

	return &types.QueryBalanceResponse{
		Coin:             coin,
		Erc20Address:     erc20Addr.String(),
		Erc20TokenAmount: erc20TokenAmount,
	}, nil
}
