package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/version"

	"github.com/ethereum/go-ethereum/common"

	sdkmath "cosmossdk.io/math"

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
	cmd.AddCommand(GetCmdUpgradeYAT())
	cmd.AddCommand(GetCmdCreatePlan())
	cmd.AddCommand(GetClaimsCmd())
	cmd.AddCommand(GetCreateYATCmd())
	cmd.AddCommand(GetUpdatePlanStatusCmd())
	return cmd
}

func GetCmdUpgradeYAT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade-plan [implementation]",
		Short: "Upgrade a plan",
		Example: fmt.Sprintf(
			"$ %s tx plan upgrade-plan [implementation]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			implementation := args[0]
			if !common.IsHexAddress(implementation) {
				return fmt.Errorf("invalid implementation address")
			}

			msgUpgradeYAT := &types.MsgUpgradeYAT{
				Implementation: implementation,
				Authority:      from.String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgUpgradeYAT)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdCreatePlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-plan [path/to/plan.json]",
		Short: "Creates a new plan",
		Example: fmt.Sprintf(
			"$ %s tx plan create-plan [path/to/plan.json]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
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
		Use:   "claims [plan-id] [round-id] [amount] [proof]",
		Short: "Handle claims for a plan",
		Example: fmt.Sprintf(
			"$ %s tx plan claims [plan-id][round-id] [amount] [proof]"+
				"--to=\"0x0eeb8ec40c6705b669469346ff8f9ce5cad57ed5\" "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			planId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("plan-id must be an integer")
			}

			roundId, ok := sdkmath.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("round-id must be an integer")
			}

			amount, ok := sdkmath.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("amount must be an integer")
			}
			hexHash := common.HexToHash(args[3])
			if len(hexHash.Bytes()) != 32 {
				return fmt.Errorf("invalid merkle proof")
			}

			receiver, err := cmd.Flags().GetString(FlagTo)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()
			if len(receiver) <= 0 {
				// set default receiver
				receiver = common.BytesToAddress(from.Bytes()).Hex()
			}

			msgClaims := &types.MsgClaims{
				PlanId:      planId,
				Receiver:    receiver,
				RoundId:     roundId,
				Amount:      amount,
				MerkleProof: args[3],
				Sender:      from.String(),
			}

			if msgClaims.ValidateBasic() != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClaims)
		},
	}
	cmd.Flags().AddFlagSet(FsClaims)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCreateYATCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "creat-yat [name] [symbol]",
		Short: "Creates a new YAT contract",
		Example: fmt.Sprintf(
			"$ %s tx plan creat-yat [name] [symbol]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			msgCreateYAT := &types.MsgCreateYAT{
				Name:   args[0],
				Symbol: args[1],
				Sender: from.String(),
			}

			if msgCreateYAT.ValidateBasic() != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgCreateYAT)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetUpdatePlanStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-plan-status [name] [symbol]",
		Short: "Update plan status",
		Example: fmt.Sprintf(
			"$ %s tx plan update-plan-status [name] [symbol]"+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			planId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("plan-id must be an integer")
			}

			planStatusUint32, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("plan-status must be an integer")
			}

			// Performing upper limit check
			if planStatusUint32 > 3 {
				return fmt.Errorf("plan-status must be less than 3")
			}

			planStatus := types.PlanStatus(planStatusUint32)

			from := clientCtx.GetFromAddress()

			msgUpdatePlanStatus := &types.MsgUpdatePlanStatus{
				PlanId: planId,
				Status: planStatus,
				Sender: from.String(),
			}

			if msgUpdatePlanStatus.ValidateBasic() != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgUpdatePlanStatus)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
