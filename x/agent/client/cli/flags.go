package cli

import (
	flag "github.com/spf13/pflag"

	"github.com/Lorenzo-Protocol/lorenzo/x/agent/types"
)

const (
	FlagName                = "name"
	FlagBtcReceivingAddress = "btc-receiving-address"
	FlagDescription         = "description"
	FlagUrl                 = "url"
)

var (
	FsAddAgent  = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditAgent = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsAddAgent.String(FlagDescription, "", "The agent description")
	FsAddAgent.String(FlagUrl, "", "The agent url")

	FsEditAgent.String(FlagName, types.DoNotModifyDesc, "The agent name")
	FsEditAgent.String(FlagBtcReceivingAddress, types.DoNotModifyDesc, "the Bitcoin receiving address of the agent")
	FsEditAgent.String(FlagDescription, types.DoNotModifyDesc, "The agent description")
	FsEditAgent.String(FlagUrl, types.DoNotModifyDesc, "The agent url")
}
