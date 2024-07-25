package keeper

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/ibctransfer/types"
	tokentypes "github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

var _ ibctransfertypes.MsgServer = Keeper{}

// Transfer overrides the ics-20 transfer method
func (k Keeper) Transfer(goCtx context.Context, msg *ibctransfertypes.MsgTransfer) (*ibctransfertypes.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// NOTE: avoid extra costs for extended logic
	kvGasCfg := ctx.KVGasConfig()
	transientKVGasCfg := ctx.TransientKVGasConfig()
	ctx = ctx.WithKVGasConfig(sdk.GasConfig{}).WithTransientKVGasConfig(sdk.GasConfig{})

	defer func() {
		ctx = ctx.WithKVGasConfig(kvGasCfg).WithTransientKVGasConfig(transientKVGasCfg)
	}()

	// 1. if not erc20 denom or not registered
	tokenArg := types.GetTokenFromDenom(msg.Token.Denom)
	id := k.tokenKeeper.GetTokenPairId(ctx, tokenArg)
	pair, found := k.tokenKeeper.GetTokenPair(ctx, id)
	if !found {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	// 2. if pair disabled
	if !pair.Enabled {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	// 3. if token module disabled
	if !k.tokenKeeper.IsConvertEnabled(ctx) {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	// 4. if balance is enough
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	balance := k.bankKeeper.GetBalance(ctx, sender, pair.Denom)
	if balance.Amount.GTE(msg.Token.Amount) {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	// if balance is not enough, try to convert erc20 before transfer
	difference := msg.Token.Amount.Sub(balance.Amount)
	if _, err := k.tokenKeeper.ConvertERC20(sdk.WrapSDKContext(ctx),
		&tokentypes.MsgConvertERC20{
			ContractAddress: pair.GetContractAddress(),
			Amount:          difference,
			Sender:          common.BytesToAddress(sender.Bytes()).String(),
			Receiver:        sender.String(),
		}); err != nil {
		return nil, err
	}

	return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
}
