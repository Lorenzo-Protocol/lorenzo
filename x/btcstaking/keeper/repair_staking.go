package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Compensate(ctx sdk.Context, receiverInfos []*types.ReceiverInfo) error {
	for _, receiverInfo := range receiverInfos {
		receiver, err := sdk.AccAddressFromBech32(receiverInfo.Address)
		if err != nil {
			return err
		}

		coins := sdk.NewCoins(sdk.NewCoin(types.NativeTokenDenom, receiverInfo.Amount))

		// mint stBTC to module account
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
			return errorsmod.Wrapf(types.ErrMintToModule, "failed to mint coins: %v", err)
		}

		// send coins to receiver
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, coins); err != nil {
			return errorsmod.Wrapf(types.ErrTransferToAddr, "failed to send coins from module to account: %v", err)
		}
	}
	return nil
}
