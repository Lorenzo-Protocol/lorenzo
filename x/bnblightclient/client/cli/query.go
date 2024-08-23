package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/bnblightclient/types"
)

// GetQueryCmd returns the cli query commands for bnblightclient module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryHeader())
	cmd.AddCommand(CmdQueryHeaderHash())
	cmd.AddCommand(CmdQueryLatestHeader())
	return cmd
}

// CmdQueryParams returns a new Cobra command for showing the parameters of the bnblightclient module.
//
// No parameters.
// Returns *cobra.Command.
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "query the parameters of the fee module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryHeader returns a new Cobra command for querying the header of the bnblightclient module by number.
//
// It takes no parameters and returns a pointer to a Cobra command.
func CmdQueryHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header",
		Short: "query the header of the bnblightclient module by number",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			number, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid number: %s, error: %s", args[0], err.Error())
			}

			res, err := queryClient.Header(context.Background(), &types.QueryHeaderRequest{Number: number})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryHeaderHash queries the header of the bnblightclient module by hash.
//
// It takes exact 1 argument.
// Returns an error.
func CmdQueryHeaderHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header-hash",
		Short: "query the header of the bnblightclient module by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.HeaderByHash(
				context.Background(),
				&types.QueryHeaderByHashRequest{Hash: common.FromHex(args[0])},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryLatestHeader returns a new Cobra command for querying the latest header of the bnblightclient module.
//
// It takes exact 1 argument.
// Returns an error.
func CmdQueryLatestHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-header",
		Short: "query the latest header of the bnblightclient module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.LatestHeader(
				context.Background(),
				&types.QueryLatestHeaderRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
