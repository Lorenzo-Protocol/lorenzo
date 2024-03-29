package cli

import (
	//"encoding/hex"
	"fmt"
	//"strings"

	//sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	/*"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"*/
	"github.com/spf13/cobra"

	//lrz "github.com/Lorenzo-Protocol/lorenzo/types"
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
)

const (
	FlagMoniker         = "moniker"
	FlagIdentity        = "identity"
	FlagWebsite         = "website"
	FlagSecurityContact = "security-contact"
	FlagDetails         = "details"
	FlagCommissionRate  = "commission-rate"
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

	//FIXME: add command
	cmd.AddCommand()

	return cmd
}
