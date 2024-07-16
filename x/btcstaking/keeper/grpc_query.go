package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	*Keeper
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// get the params
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) StakingRecord(c context.Context, req *types.QueryStakingRecordRequest) (*types.QueryStakingRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	if len(req.TxHash) != chainhash.HashSize {
		return nil, fmt.Errorf("invalid hash length of %v, want %v", len(req.TxHash), chainhash.HashSize)
	}

	var txHash chainhash.Hash
	if err := txHash.SetBytes(req.TxHash); err != nil {
		return nil, err
	}

	// get the staking record
	stakingRecord := k.getBTCStakingRecord(ctx, txHash)

	return &types.QueryStakingRecordResponse{Record: stakingRecord}, nil
}

func (k Keeper) Records(ctx context.Context, req *types.QueryRecordsRequest) (*types.QueryRecordsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Ensure that the pagination key corresponds to hash bytes
	if req.Pagination.Key != nil && len(req.Pagination.Key) != chainhash.HashSize && len(req.Pagination.Key) != 0 {
		return nil, fmt.Errorf("invalid hash length of %v, want %v or empty", len(req.Pagination.Key), chainhash.HashSize)
	}

	store := k.btcStakingRecordStore(sdkCtx)
	records := make([]*types.BTCStakingRecord, 0)
	pageRes, err := query.FilteredPaginate(store, req.Pagination, func(key []byte, recordBytes []byte, accumulate bool) (bool, error) {
		if accumulate {
			var record types.BTCStakingRecord
			k.cdc.MustUnmarshal(recordBytes, &record)
			records = append(records, &record)
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryRecordsResponse{Records: records, Pagination: pageRes}, nil
}

// NewQuerierImpl returns an implementation of the captains QueryServer interface.
func NewQuerierImpl(k *Keeper) types.QueryServer {
	return &Querier{k}
}
