package v200

import (
	"github.com/Lorenzo-Protocol/lorenzo/v2/app/upgrades"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func migrateAgentFromBTCStakingToAgent(
	ctx sdk.Context,
	app upgrades.AppKeepers,
) error {
	btcStakingParams := app.BTCStakingKeeper.GetParams(ctx)
	if btcStakingParams == nil {
		return nil
	}
	for _, receiver := range btcStakingParams.Receivers {
		app.AgentKeeper.AddAgent(
			ctx,
			receiver.Name,
			receiver.Addr, receiver.EthAddr,
			"",
			"",
		)
	}
	// TODO: Is params.Receiver of btcstaking module removed?
	btcStakingParams.Receivers = nil
	if err := app.BTCStakingKeeper.SetParams(ctx, btcStakingParams); err != nil {
		return err
	}
	return nil
}
