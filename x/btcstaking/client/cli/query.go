package cli

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/common"
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
		CmdGetBTCStakingRecord(),
	)

	return btcstakingQueryCmd
}

func CmdGetParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
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

func CmdGetBTCStakingRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "btc-staking-record [btc_staking_tx_id]",
		Aliases: []string{"record"},
		Short:   "get the btc staking record",
		Args:    cobra.ExactArgs(1),
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
			if res.Record == nil {
				return fmt.Errorf("record not found")
			}
			resDisp := types.StakingRecordDisplay{
				TxId:            (chainhash.Hash)(res.Record.TxHash).String(),
				Amount:          sdkmath.NewIntFromUint64(res.Record.Amount).Mul(sdkmath.NewIntFromUint64(1e10)).String(),
				ReceiverAddress: common.BytesToAddress(res.Record.ReceiverAddr).String(),
				AgentName:       res.Record.AgentName,
				AgentBtcAddr:    res.Record.AgentBtcAddr,
				ChainId:         res.Record.ChainId,
			}
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(&resDisp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
