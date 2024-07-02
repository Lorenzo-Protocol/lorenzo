package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	*Keeper
}

func NewMsgServer(k *Keeper) types.MsgServer {
	return &msgServer{k}
}

func (m msgServer) ConvertCoin(ctx context.Context, coin *types.MsgConvertCoin) (*types.MsgConvertCoinResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (m msgServer) ConvertERC20(ctx context.Context, erc20 *types.MsgConvertERC20) (*types.MsgConvertERC20Response, error) {
	// TODO implement me
	panic("implement me")
}

func (m msgServer) UpdateParams(ctx context.Context, params *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	// TODO implement me
	panic("implement me")
}
