package cli

import (
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group btcstaking queries under a subcommand
	btcstakingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the btcstaking module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	btcstakingQueryCmd.AddCommand(
		CmdGetParams(),
		CmdGetBTCStaingRecord(),
	)

	return btcstakingQueryCmd
}

func CmdGetParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-params",
		Short: "get btc staking params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetBTCStaingRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-btc-staking-record [btc_staking_tx_id]",
		Short: "get the btc staking record",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			txHashBytes, err := chainhash.NewHashFromStr(args[0])

			if err != nil {
				return err
			}
			res, err := queryClient.StakingRecord(cmd.Context(), &types.QueryStakingRecordRequest{TxHash: txHashBytes[:]})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
