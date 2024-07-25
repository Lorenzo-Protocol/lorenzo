package keeper

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) getModuleEthAddress(ctx sdk.Context) common.Address {
	moduleAccount := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	return common.BytesToAddress(moduleAccount.GetAddress().Bytes())
}
