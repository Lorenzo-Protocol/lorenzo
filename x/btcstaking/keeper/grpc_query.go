package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
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

// BTCBStakingRecord retrieves the BTCB staking record for the given contract and staking index.
//
// Parameters:
// - c: The context.Context object for the request.
// - req: The QueryBTCBStakingRecordRequest object containing the contract address and staking index.
//
// Returns:
// - *types.QueryBTCBStakingRecordResponse: The BTCB staking record response.
// - error: An error if the staking record retrieval fails.
func (k Keeper) BTCBStakingRecord(c context.Context, req *types.QueryBTCBStakingRecordRequest) (*types.QueryBTCBStakingRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	contract := common.HexToAddress(req.Contract)
	// get the staking record
	stakingRecord, err := k.getBTCBStakingRecord(ctx, req.ChainId, contract[:], req.StakingIdx)
	if err != nil {
		return nil, err
	}
	return &types.QueryBTCBStakingRecordResponse{Record: stakingRecord}, nil
}

// NewQuerierImpl returns an implementation of the captains QueryServer interface.
func NewQuerierImpl(k *Keeper) types.QueryServer {
	return &Querier{k}
}
