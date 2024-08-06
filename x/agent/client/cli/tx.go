package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/agent/types"
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

	cmd.AddCommand(CmdTxAddAgent())
	cmd.AddCommand(CmdTxEditAgent())
	cmd.AddCommand(CmdTxRemoveAgent())

	return cmd
}

// CmdTxAddAgent returns a cobra.Command that submits a transaction to add an agent.
//
// The command takes two arguments: name and btc_receiving_address. It returns an error if the command is not executed with exactly one argument.
// The command generates or broadcasts a transaction using the provided flags and the client transaction context.
// The command returns an error if there is an issue getting the client transaction context or creating the message.
// The command returns an error if there is an issue generating or broadcasting the transaction.
// The command returns nil if the transaction is successfully generated or broadcasted.
//
// Parameters:
// - cmd: the cobra.Command to be executed.
// - args: the arguments passed to the command.
//
// Returns:
// - error: an error if there is an issue executing the command.
func CmdTxAddAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name] [btc_receiving_address]",
		Short: "submit add agent tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ethAddr, err := cmd.Flags().GetString(FlagEthAddress)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			url, err := cmd.Flags().GetString(FlagURL)
			if err != nil {
				return err
			}
			msg := &types.MsgAddAgent{
				Name:                args[0],
				BtcReceivingAddress: args[1],
				EthAddr:             ethAddr,
				Description:         description,
				Url:                 url,
				Sender:              clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsAddAgent)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdTxEditAgent returns a cobra.Command that submits a transaction to edit an agent.
//
// The command takes one argument: agent-id. It returns an error if the command is not executed with exactly one argument.
// The command generates or broadcasts a transaction using the provided flags and the client transaction context.
// The command returns an error if there is an issue getting the client transaction context or creating the message.
// The command returns an error if there is an issue generating or broadcasting the transaction.
// The command returns nil if the transaction is successfully generated or broadcasted.
//
// Return type: error.
func CmdTxEditAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit [agent-id]",
		Short: "submit edit agent tx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			agentID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			url, err := cmd.Flags().GetString(FlagURL)
			if err != nil {
				return err
			}

			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}

			msg := types.MsgEditAgent{
				Name:        name,
				Description: description,
				Url:         url,
				Sender:      clientCtx.GetFromAddress().String(),
				Id:          agentID,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FsEditAgent)
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdTxRemoveAgent description of the Go function.
//
// CmdTxRemoveAgent submits a transaction to remove an agent.
// It takes no arguments and returns an error.
func CmdTxRemoveAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [agent-id]",
		Short: "submit remove agent tx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			agentID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			msg := types.MsgRemoveAgent{
				Sender: clientCtx.GetFromAddress().String(),
				Id:     agentID,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
