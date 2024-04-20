package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryAgents())
	cmd.AddCommand(CmdQueryAgent())

	return cmd
}

// CmdQueryAgents returns a cobra.Command object for querying agents.
//
// The function takes no parameters and returns a pointer to a cobra.Command object.
func CmdQueryAgents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "retrieve the hashes maintained by this module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Agents(context.Background(), &types.QueryAgentsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryAgent returns a cobra.Command object for querying a specific agent.
//
// It takes no parameters and returns a pointer to a cobra.Command object.
func CmdQueryAgent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent [agent-id]",
		Short: "retrieve the hashes maintained by this module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.Agent(context.Background(), &types.QueryAgentRequest{
				Id: id,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
