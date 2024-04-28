package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func (k Keeper) DeployYATContract(
	ctx sdk.Context,
	deployer common.Address,
	name string,
	owner common.Address,
) (common.Address, error) {
	panic("to be implemented")
}

// Mint mint a new YAT token
func (k Keeper) Mint() error {
	panic("to be implemented")
}
