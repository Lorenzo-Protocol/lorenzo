package cli

import (
	"fmt"
	"strconv"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btclightclient/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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

	cmd.AddCommand(CmdTxInsertHeader())
	cmd.AddCommand(CmdTxUpdateFeeRate())

	return cmd
}

func CmdTxInsertHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert-headers [headers-bytes]",
		Short: "submit BTC headers bytes",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg, err := types.NewMsgInsertHeaders(clientCtx.GetFromAddress(), args[0])
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdTxUpdateFeeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-fee-rate fee-rate",
		Short: "submit bitcoin fee rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			feeRate, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			msg := types.MsgUpdateFeeRate{Signer: clientCtx.GetFromAddress().String(), FeeRate: feeRate}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
