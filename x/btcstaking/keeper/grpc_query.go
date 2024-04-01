package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.

func (k Keeper) BTCReceivingAddr(c context.Context, req *types.QueryBTCReceivingAddrRequest) (*types.QueryBTCReceivingAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// get the receiving address
	receivingAddr := k.GetBTCReceivingAddr(ctx)

	return &types.QueryBTCReceivingAddrResponse{Addr: receivingAddr}, nil
}

func (k Keeper) StakingRecord(c context.Context, req *types.QueryStakingRecordRequest) (*types.QueryStakingRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if len(req.TxHash) != chainhash.HashSize {
		return nil, fmt.Errorf("invalid hash length of %v, want %v", len(req.TxHash), chainhash.HashSize)
	}

	var txHash chainhash.Hash
	txHash.SetBytes(req.TxHash)

	// get the staking record
	stakingRecord := k.getBTCStakingRecord(ctx, txHash)

	return &types.QueryStakingRecordResponse{Record: stakingRecord}, nil
}
