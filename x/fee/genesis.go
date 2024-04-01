package fee

import (
	"fmt"

	"github.com/Lorenzo-Protocol/lorenzo/x/fee/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/fee/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new fee genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	if err := ValidateGenesis(data); err != nil {
		panic(fmt.Errorf("failed to initialize fee genesis state: %s", err.Error()))
	}
	if err := keeper.SetParams(ctx, data.Params); err != nil {
		panic(fmt.Errorf("failed to set fee genesis state: %s", err.Error()))
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	params := keeper.GetParams(ctx)
	return types.NewGenesisState(params)
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data types.GenesisState) error {
	return data.Params.Validate()
}


