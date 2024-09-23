package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Lorenzo-Protocol/lorenzo/v3/x/ccev/types"
)

const (
	// FlagHeaderHash is the flag for the header hash of the block
	FlagHeaderHash = "header-hash"
	// FlagHeaderNumber is the flag for the header number of the block
	FlagHeaderNumber = "header-number"
	// FlagReceiptRoot is the flag for the receipt root of the block
	FlagReceiptRoot     = "receipt-root"
	// FlagContractAddress is the flag for the contract address
	FlagContractAddress = "contract-address"
	// FlagContractABI is the flag for the contract ABI
	FlagContractABI     = "contract-abi"
	// FlagEventName is the flag for the event name
	FlagEventName       = "event-name"
)

// FsCreateClient defines the flags for creating a new client
var FsCreateClient = pflag.NewFlagSet("", pflag.ContinueOnError)

// FsUploadContract defines the flags for uploading a contract
var FsUploadContract = pflag.NewFlagSet("", pflag.ContinueOnError)
// FsUploadHeaders defines the flags for uploading headers
var FsUploadHeaders = pflag.NewFlagSet("", pflag.ContinueOnError)
var FsUpdateHeaders = pflag.NewFlagSet("", pflag.ContinueOnError)

func init() {
	FsCreateClient.String(FlagHeaderHash, "", "The hash of the header")
	FsCreateClient.String(FlagHeaderNumber, "", "The number of the header")
	FsCreateClient.String(FlagReceiptRoot, "", "The receipt root of the header")

	FsUploadContract.String(FlagContractAddress, "", "The contract address")
	FsUploadContract.String(FlagContractABI, "", "The contract abi")
	FsUploadContract.String(FlagEventName, "", "The event name")

	FsUploadHeaders.String(FlagHeaderHash, "", "The hash of the header")
	FsUploadHeaders.String(FlagHeaderNumber, "", "The number of the header")
	FsUploadHeaders.String(FlagReceiptRoot, "", "The receipt root of the header")

	FsUpdateHeaders.String(FlagHeaderHash, "", "The hash of the header")
	FsUpdateHeaders.String(FlagHeaderNumber, "", "The number of the header")
	FsUpdateHeaders.String(FlagReceiptRoot, "", "The receipt root of the header")
}

func parseHeader(cmd *cobra.Command) (*types.TinyHeader, error) {
	headerHash, err := cmd.Flags().GetString(FlagHeaderHash)
	if err != nil {
		return nil, err
	}
	headerNumber, err := cmd.Flags().GetUint64(FlagHeaderNumber)
	if err != nil {
		return nil, err
	}
	receiptRoot, err := cmd.Flags().GetString(FlagReceiptRoot)
	if err != nil {
		return nil, err
	}
	return &types.TinyHeader{Hash: headerHash, Number: headerNumber, ReceiptRoot: receiptRoot}, nil
}
