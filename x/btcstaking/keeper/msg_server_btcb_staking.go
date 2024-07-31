package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
)

// CreateBTCStakingFromBNB implements types.MsgServer.
func (ms msgServer) CreateBTCBStaking(goctx context.Context, req *types.MsgCreateBTCBStaking) (*types.MsgCreateBTCBStakingResponse, error) {	
	return &types.MsgCreateBTCBStakingResponse{}, nil
}