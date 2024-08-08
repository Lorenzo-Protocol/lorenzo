package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlock is a method of the Keeper struct that is responsible for pruning headers at the end of each block.
//
// It takes a context parameter of type sdk.Context, which represents the current state of the blockchain.
// The method does not return any value.
func (k Keeper) EndBlock(ctx sdk.Context) {
	k.prune(ctx)
}
