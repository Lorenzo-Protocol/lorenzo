package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/token/types"
)

// GetQueryCmd returns the parent command for token module query commands
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the token module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetQueryTokenPairsCmd(),
		GetQueryTokenPairCmd(),
		GetQueryParamsCmd(),
		GetQueryBalanceCmd(),
	)
	return cmd
}

// GetQueryTokenPairsCmd queries all registered token pairs
func GetQueryTokenPairsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-pairs",
		Short: "Query token pairs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryTokenPairsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TokenPairs(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetQueryTokenPairCmd returns the command handler for token pair querying
func GetQueryTokenPairCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-pair [token]",
		Short: "Query a token pair",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryTokenPairRequest{
				Token: args[0],
			}

			res, err := queryClient.TokenPair(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetQueryParamsCmd returns the command handler for token module params querying
func GetQueryParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Get token module params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

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

// GetQueryBalanceCmd returns the command handler for token module balance querying
func GetQueryBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance [account_address] [token]",
		Short: "Get balance of an address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			addrString := args[0]
			if !common.IsHexAddress(addrString) {
				_, err := sdk.AccAddressFromBech32(addrString)
				if err != nil {
					return fmt.Errorf("invalid account address: %s, should be either hex addr or bech32 addr", addrString)
				}
			}

			res, err := queryClient.Balance(context.Background(), &types.QueryBalanceRequest{
				AccountAddress: args[0],
				Token:          args[1],
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
