package cli

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
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
		CmdRecords(),
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

func ToRecordDisplay(record *types.BTCStakingRecord) *types.StakingRecordDisplay {
	resDisp := types.StakingRecordDisplay{}
	resDisp.TxId = (chainhash.Hash)(record.TxHash).String()
	resDisp.Amount = sdkmath.NewIntFromUint64(record.Amount).Mul(sdkmath.NewIntFromUint64(1e10)).String()
	resDisp.MintToAddress = common.BytesToAddress(record.MintToAddr).String()
	resDisp.BtcReceiverName = record.BtcReceiverName
	resDisp.BtcReceiverAddr = record.BtcReceiverAddr
	return &resDisp
}

func CmdGetBTCStakingRecord() *cobra.Command {
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
			if res.Record == nil {
				return fmt.Errorf("record not found")
			}
			if err != nil {
				return err
			}
			resDisp := ToRecordDisplay(res.Record)
			return clientCtx.PrintProto(resDisp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "records",
		Short: "retrieve staking records ",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Records(cmd.Context(), &types.QueryRecordsRequest{Pagination: pageReq})
			if err != nil {
				return err
			}
			recordsDisp := make([]*types.StakingRecordDisplay, 0)
			for _, record := range res.Records {
				recordsDisp = append(recordsDisp, ToRecordDisplay(record))
			}

			return clientCtx.PrintProto(&types.RecordsDisplay{
				Records:    recordsDisp,
				Pagination: res.Pagination,
			})
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "records")

	return cmd
}
