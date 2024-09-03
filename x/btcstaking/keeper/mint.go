package keeper

import (
	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) MintStBTC(ctx sdk.Context, toAddr sdk.AccAddress, amount sdkmath.Int) error {
	coins := sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, amount))

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, coins); err != nil {
		return err
	}

	ethAddr := common.BytesToAddress(toAddr.Bytes())

	// emit the mint event
	ctx.EventManager().EmitTypedEvent(types.NewEventMintStBTC(toAddr.String(), ethAddr.Hex(), amount)) //nolint:errcheck,gosec
	return nil
}
