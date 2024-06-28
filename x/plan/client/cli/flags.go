package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTo = "to"
)

var (
	FsClaims = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsClaims.String(FlagTo, "", "Address to which the claims is to be received")
}
