package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Lorenzo-Protocol/lorenzo/x/token/types"
)

// NewTxCmd returns a root CLI command handler for token transaction commands
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "token module subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConvertCoinCmd(),
		NewConvertERC20Cmd(),
	)

	return txCmd
}

// NewConvertCoinCmd returns a CLI command handler for converting coin to erc20 token.
func NewConvertCoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-coin [coin] [receiver_hex_address]",
		Short: "Convert sdk coin to erc20 token and send it to the receiver hex address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			if !common.IsHexAddress(args[1]) {
				return fmt.Errorf("invalid hex address %s", args[1])
			}

			sender := cliCtx.GetFromAddress().String()

			msg := &types.MsgConvertCoin{
				Coin:     coin,
				Receiver: args[1],
				Sender:   sender,
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewConvertERC20Cmd returns a CLI command handler for converting erc20 token to sdk coin.
func NewConvertERC20Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-erc20 [contract_address] [amount] [receiver_bech32_address]",
		Short: "Convert erc20 token to sdk coin and send it to the receiver bech32 addr",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if !common.IsHexAddress(args[0]) {
				return fmt.Errorf("invalid contract address %s", args[0])
			}

			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount %s", args[1])
			}

			sender := common.BytesToAddress(cliCtx.GetFromAddress().Bytes())

			msg := types.MsgConvertERC20{
				ContractAddress: args[0],
				Amount:          amount,
				Receiver:        args[2],
				Sender:          sender.Hex(),
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
