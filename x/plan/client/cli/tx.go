package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/Lorenzo-Protocol/lorenzo/x/plan/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(GetCmdCreatePlan())
	cmd.AddCommand(GetClaimsCmd())
	return cmd
}

func GetCmdCreatePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-plan [path/to/plan.json]",
		Short: "Creates a new plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			contents, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var msgCreatePlan *types.MsgCreatePlan
			err = json.Unmarshal(contents, msgCreatePlan)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCreatePlan)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetClaimsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claims [plan-id] [receiver] [claims_type]",
		Short: "Handle claims for a plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			contents, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var msgCreatePlan = &types.MsgCreatePlan{}
			err = json.Unmarshal(contents, msgCreatePlan)
			if err != nil {
				return err
			}
			if msgCreatePlan.ValidateBasic() != nil {
				return err

			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCreatePlan)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
