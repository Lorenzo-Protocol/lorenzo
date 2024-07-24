package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/x/bnblightclient/types"
)

// GetTxCmd returns the transaction commands for bnblightclient module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(CmdTxUploadHeaders())
	cmd.AddCommand(CmdTxUpdateHeader())
	return cmd
}

// CmdTxUploadHeaders creates a cobra.Command to upload BNB light client headers.
//
// It takes a single argument, which is the path to a JSON file containing the headers.
// The function reads the content of the file, unmarshals it into a slice of types.Header,
// and creates a MsgUploadHeaders message with the signer and the headers.
//
// The function returns an error if there is any issue reading the file or unmarshaling
// the content.
//
// The function also returns an error if there is any issue generating or broadcasting
// the transaction.
//
// The function returns a pointer to a cobra.Command.
func CmdTxUploadHeaders() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-headers [headers-file]",
		Short: "upload bnb light client header",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var headers []*types.Header
			if err := json.Unmarshal(content, &headers); err != nil {
				return err
			}

			msg := types.MsgUploadHeaders{
				Signer:  clientCtx.GetFromAddress().String(),
				Headers: headers,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdTxUpdateHeader creates a cobra.Command to update the BNB light client header.
//
// It takes a single argument, which is the path to a JSON file containing the header.
// The function reads the content of the file, unmarshals it into a types.Header struct,
// and creates a MsgUpdateHeader message with the signer and the header.
//
// The function returns an error if there is any issue reading the file or unmarshaling
// the content.
//
// The function also returns an error if there is any issue generating or broadcasting
// the transaction.
//
// The function returns a pointer to a cobra.Command.
func CmdTxUpdateHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-header [header-file]",
		Short: "update bnb light client header",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var header types.Header
			if err := json.Unmarshal(content, &header); err != nil {
				return err
			}

			msg := types.MsgUpdateHeader{
				Signer: clientCtx.GetFromAddress().String(),
				Header: &header,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
