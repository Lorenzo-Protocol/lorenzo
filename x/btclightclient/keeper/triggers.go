package keeper

import (
	"context"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btclightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) triggerHeaderInserted(goCtx context.Context, headerInfo *types.BTCHeaderInfo) {
	// Trigger AfterBTCHeaderInserted hook
	k.AfterBTCHeaderInserted(goCtx, headerInfo)
	// Emit HeaderInserted event
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitTypedEvent(&types.EventBTCHeaderInserted{Header: headerInfo}) //nolint:errcheck,gosec
}

func (k Keeper) triggerRollBack(goCtx context.Context, headerInfo *types.BTCHeaderInfo) {
	// Trigger AfterBTCRollBack hook
	k.AfterBTCRollBack(goCtx, headerInfo)
	// Emit BTCRollBack event
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitTypedEvent(&types.EventBTCRollBack{Header: headerInfo}) //nolint:errcheck,gosec
}

func (k Keeper) triggerRollForward(goCtx context.Context, headerInfo *types.BTCHeaderInfo) {
	// Trigger AfterBTCRollForward hook
	k.AfterBTCRollForward(goCtx, headerInfo)
	// Emit BTCRollForward event
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitTypedEvent(&types.EventBTCRollForward{Header: headerInfo}) //nolint:errcheck,gosec
}

func (k Keeper) triggerFeeRateUpdated(goCtx context.Context, feeRate uint64) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx.EventManager().EmitTypedEvent(&types.EventBTCFeeRateUpdated{FeeRate: feeRate}) //nolint:errcheck,gosec
}
