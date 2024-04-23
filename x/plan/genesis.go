package plan

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new fee genesis
func InitGenesis(ctx sdk.Context, keeper *keeper.Keeper, data types.GenesisState) {

}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper *keeper.Keeper) *types.GenesisState {
	return nil
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data types.GenesisState) error {
	return nil
}
