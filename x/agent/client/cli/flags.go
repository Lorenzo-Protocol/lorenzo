package cli

import (
	flag "github.com/spf13/pflag"

	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

const (
	// FlagName is the flag for the name of the agent
	FlagName = "name"
	// FlagEthAddress is the flag for the eth address of the agent
	FlagEthAddress = "eth-address"
	// FlagDescription is the flag for the description of the agent
	FlagDescription = "description"
	// FlagURL is the flag for the url of the agent
	FlagURL = "url"
)

var (
	// FsAddAgent defines the flags for adding an agent
	FsAddAgent = flag.NewFlagSet("", flag.ContinueOnError)
	// FsEditAgent defines the flags for editing an agent
	FsEditAgent = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsAddAgent.String(FlagEthAddress, "", "The agent eth address")
	FsAddAgent.String(FlagDescription, "", "The agent description")
	FsAddAgent.String(FlagURL, "", "The agent url")

	FsEditAgent.String(FlagName, types.DoNotModifyDesc, "The agent name")
	FsEditAgent.String(FlagDescription, types.DoNotModifyDesc, "The agent description")
	FsEditAgent.String(FlagURL, types.DoNotModifyDesc, "The agent url")
}
