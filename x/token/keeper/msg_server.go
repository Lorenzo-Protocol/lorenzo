package keeper

import (
	"context"

	v1 "github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

var _ v1.MsgServer = msgServer{}

type msgServer struct {
	*Keeper
}

func (m msgServer) ConvertCoin(ctx context.Context, coin *v1.MsgConvertCoin) (*v1.MsgConvertCoinResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (m msgServer) ConvertERC20(ctx context.Context, erc20 *v1.MsgConvertERC20) (*v1.MsgConvertERC20Response, error) {
	// TODO implement me
	panic("implement me")
}

func (m msgServer) UpdateParams(ctx context.Context, params *v1.MsgUpdateParams) (*v1.MsgUpdateParamsResponse, error) {
	// TODO implement me
	panic("implement me")
}
