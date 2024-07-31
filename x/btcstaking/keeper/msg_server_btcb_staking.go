package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
)

// CreateBTCStakingFromBNB implements types.MsgServer.
func (ms msgServer) CreateBTCBStaking(goctx context.Context, req *types.MsgCreateBTCBStaking) (*types.MsgCreateBTCBStakingResponse, error) {
	depositor,err := sdk.AccAddressFromBech32(req.Signer)
	if err != nil {
		return nil, err
	}
	
	ctx := sdk.UnwrapSDKContext(goctx)
	if err = ms.DepositBTCB(ctx, depositor, req.Receipt, req.Proof); err != nil {
		return nil, err
	}
	return &types.MsgCreateBTCBStakingResponse{}, nil
}