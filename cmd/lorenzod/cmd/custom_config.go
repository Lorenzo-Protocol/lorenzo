package cmd

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"

	bbn "github.com/Lorenzo-Protocol/lorenzo/v2/types"
)

type BtcConfig struct {
	Network string `mapstructure:"network"`
}

func defaultLorenzoBtcConfig() BtcConfig {
	return BtcConfig{
		Network: string(bbn.BtcMainnet),
	}
}

type LorenzoAppConfig struct {
	serverconfig.Config `mapstructure:",squash"`
	BtcConfig           BtcConfig `mapstructure:"btc-config"`
}

func DefaultLorenzoConfig() *LorenzoAppConfig {
	return &LorenzoAppConfig{
		Config:    *serverconfig.DefaultConfig(),
		BtcConfig: defaultLorenzoBtcConfig(),
	}
}

func DefaultLorenzoTemplate() string {
	return serverconfig.DefaultConfigTemplate + `
[json-rpc]
# FeeHistoryCap sets the global cap for total number of blocks that can be fetched
feehistory-cap = 100

[btc-config]

# Configures which bitcoin network should be used for checkpointing
# valid values are: [mainnet, testnet, simnet, regtest]
network = "{{ .BtcConfig.Network }}"
`
}
