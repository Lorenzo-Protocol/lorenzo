package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
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
	cmd.AddCommand(CmdQueryClient())
	cmd.AddCommand(CmdQueryClients())
	cmd.AddCommand(CmdQueryHeader())
	cmd.AddCommand(CmdQueryHeaderHash())
	cmd.AddCommand(CmdQueryLatestHeader())
	cmd.AddCommand(CmdQueryContract())
	cmd.AddCommand(CmdQueryParams())
	return cmd
}

// CmdQueryClient returns a new Cobra command for querying a client of the ccev module
// by the specified chainId.
//
// It takes exact 1 argument.
// Returns an error.
func CmdQueryClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: " query the client by the specified chainId",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Client(context.Background(), &types.QueryClientRequest{
				ChainId: chainID,
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

// CmdQueryClients returns a new Cobra command for querying all the clients of the ccev module.
//
// No parameters.
// Returns an error.
func CmdQueryClients() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clients",
		Short: "query all the clients",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Clients(context.Background(), &types.QueryClientsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryParams returns a new Cobra command for showing the parameters of the ccev module.
//
// No parameters.
// Returns *cobra.Command.
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "query the parameters of the ccev module",
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

// CmdQueryHeader returns a new Cobra command for querying the header of the specified chain-id by number.
//
// It takes no parameters and returns a pointer to a Cobra command.
func CmdQueryHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header",
		Short: "query the header of the chain by number",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			number, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid number: %s, error: %s", args[0], err.Error())
			}

			res, err := queryClient.Header(context.Background(), &types.QueryHeaderRequest{
				ChainId: chainID,
				Number:  number,
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

// CmdQueryHeaderHash queries the header of the specified chain_id by hash.
//
// It takes exact 1 argument.
// Returns an error.
func CmdQueryHeaderHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header-hash",
		Short: "query the header of the specified chain by hash",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.HeaderByHash(
				context.Background(),
				&types.QueryHeaderByHashRequest{
					ChainId: chainID,
					Hash:    args[1],
				},
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

// CmdQueryLatestHeader returns a new Cobra command for querying the latest header of the specified chain-id.
//
// It takes exact 1 argument.
// Returns an error.
func CmdQueryLatestHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-header",
		Short: "query the latest header of the specified chain-id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.LatestHeader(
				context.Background(),
				&types.QueryLatestHeaderRequest{
					ChainId: chainID,
				},
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

// CmdQueryContract returns a new Cobra command for querying the contract by the specified chainId and address.
//
// It takes exact 2 arguments.
// Returns an error.
func CmdQueryContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract",
		Short: "query the contract by the specified chainId and address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Contract(context.Background(), &types.QueryContractRequest{
				ChainId: chainID,
				Address: args[1],
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

func parseChainID(clientID string) (uint32, error) {
	chainID, err := strconv.ParseUint(clientID, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(chainID), nil
}
