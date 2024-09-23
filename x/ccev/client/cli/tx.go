package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

// GetTxCmd returns the transaction commands for ccev module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(CmdTxCreateClient())
	cmd.AddCommand(CmdTxUploadContract())
	cmd.AddCommand(CmdTxUploadHeaders())
	cmd.AddCommand(CmdTxUpdateHeader())
	return cmd
}

// CmdTxCreateClient creates a new client.
//
// The command takes two arguments, the ID of the chain and the name of the chain.
// The command generates or broadcasts a transaction using the provided flags and the client transaction context.
// The command returns an error if there is an issue getting the client transaction context or creating the message.
// The command returns an error if there is an issue generating or broadcasting the transaction.
//
// The command returns nil if the transaction is successfully generated or broadcasted.
//
// Example:
// lorenzo tx ccev create-client 1 BinanceSmartChain --header-hash 0x1234567890abcdef --header-number 1 --receipt-root 0x1234567890abcdef
func CmdTxCreateClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-client [chain-id] [chain-name]",
		Short: "create a new client",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			header, err := parseHeader(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgCreateClient{
				Sender: clientCtx.GetFromAddress().String(),
				Client: types.Client{
					ChainId:      chainID,
					ChainName:    args[1],
					InitialBlock: *header,
				},
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FsCreateClient)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagHeaderHash)
	_ = cmd.MarkFlagRequired(FlagHeaderNumber)
	_ = cmd.MarkFlagRequired(FlagReceiptRoot)
	return cmd
}

// CmdTxUploadContract creates a cobra.Command to upload a contract.
// It takes a single argument, which is the chain-id of the client.
// The function reads the content of the flags, unmarshals it into a MsgUploadContract message,
// and creates a transaction with the signer and the message.
// The function returns an error if there is any issue generating or broadcasting the transaction.
// The function returns a pointer to a cobra.Command.
func CmdTxUploadContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-contract [chain-id]",
		Short: "upload contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			contractAddress, err := cmd.Flags().GetString(FlagContractAddress)
			if err != nil {
				return err
			}

			contractABI, err := cmd.Flags().GetString(FlagContractABI)
			if err != nil {
				return err
			}

			eventName, err := cmd.Flags().GetString(FlagEventName)
			if err != nil {
				return err
			}

			msg := types.MsgUploadContract{
				ChainId:   chainID,
				Address:   contractAddress,
				EventName: eventName,
				Abi:       []byte(contractABI),
				Sender:    clientCtx.GetFromAddress().String(),
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FsUploadContract)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagContractAddress)
	_ = cmd.MarkFlagRequired(FlagContractABI)
	_ = cmd.MarkFlagRequired(FlagEventName)
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
		Use:   "upload-headers [chain-id]",
		Short: "upload light client header",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			header, err := parseHeader(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUploadHeaders{
				ChainId: chainID,
				Sender:  clientCtx.GetFromAddress().String(),
				Headers: []types.TinyHeader{*header},
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FsUploadHeaders)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagHeaderHash)
	_ = cmd.MarkFlagRequired(FlagHeaderNumber)
	_ = cmd.MarkFlagRequired(FlagReceiptRoot)
	return cmd
}



// CmdTxUpdateHeader creates a new Cobra command for updating a light client header.
//
// It takes a single argument, which is the path to a JSON file containing the header.
// The command reads the content of the file, unmarshals it into a types.TinyHeader,
// and creates a MsgUpdateHeader message with the signer and the header.
//
// The command returns an error if there is any issue reading the file or unmarshaling
// the content.
//
// The command also returns an error if there is any issue generating or broadcasting
// the transaction.
//
// The command returns a pointer to a Cobra command.
func CmdTxUpdateHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-header [chain-id]",
		Short: "update light client header",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID, err := parseChainID(args[0])
			if err != nil {
				return err
			}

			header, err := parseHeader(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateHeader{
				ChainId: chainID,
				Sender: clientCtx.GetFromAddress().String(),
				Header: *header,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().AddFlagSet(FsUpdateHeaders)
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagHeaderHash)
	_ = cmd.MarkFlagRequired(FlagHeaderNumber)
	_ = cmd.MarkFlagRequired(FlagReceiptRoot)
	return cmd
}
