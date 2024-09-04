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

		if err := k.MintStBTC(ctx, receiver, receiverInfo.Amount); err != nil {
			return errorsmod.Wrapf(types.ErrMintStBTC, "failed to mint stBTC: %v", err)
		}
	}
	return nil
}
