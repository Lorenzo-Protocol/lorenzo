package cli

import (
	"encoding/hex"
	"fmt"

	sdkmath "cosmossdk.io/math"
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
			resDisp := types.StakingRecordDisplay{}
			resDisp.TxId = (chainhash.Hash)(res.Record.TxHash).String()
			resDisp.Amount = sdkmath.NewIntFromUint64(res.Record.Amount).Mul(sdkmath.NewIntFromUint64(1e10)).String()
			resDisp.MintToAddress = "0x" + hex.EncodeToString(res.Record.MintToAddr)
			resDisp.BtcReceiverName = res.Record.BtcReceiverName
			resDisp.BtcReceiverAddr = res.Record.BtcReceiverAddr
			return clientCtx.PrintProto(&resDisp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
