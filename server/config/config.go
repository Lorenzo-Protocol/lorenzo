package config

import (
	bbn "github.com/Lorenzo-Protocol/lorenzo/v3/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/evmos/ethermint/server/config"
)

type BtcConfig struct {
	Network string `mapstructure:"network"`
}

type Config struct {
	config.Config

	BtcConfig BtcConfig `mapstructure:"btc-config"`
}

// ValidateBasic returns an error any of the application configuration fields are invalid
func (c Config) ValidateBasic() error {
	return c.Config.ValidateBasic()
}

func defaultLorenzoBtcConfig() BtcConfig {
	return BtcConfig{
		Network: string(bbn.BtcMainnet),
	}
}

// AppConfig helps to override default appConfig template and configs.
// return "", nil if no custom configuration is required for the application.
func AppConfig(denom string) (string, interface{}) {
	// Optionally allow the chain developer to overwrite the SDK's default
	// server config.
	srvCfg := config.DefaultConfig()

	// The SDK's default minimum gas price is set to "" (empty value) inside
	// app.toml. If left empty by validators, the node will halt on startup.
	// However, the chain developer can set a default app.toml value for their
	// validators here.
	//
	// In summary:
	// - if you leave srvCfg.MinGasPrices = "", all validators MUST tweak their
	//   own app.toml config,
	// - if you set srvCfg.MinGasPrices non-empty, validators CAN tweak their
	//   own app.toml to override, or use this default value.
	//
	// In ethermint, we set the min gas prices to 0.
	if denom != "" {
		srvCfg.MinGasPrices = "0" + denom
	}

	customAppConfig := Config{
		Config:    *srvCfg,
		BtcConfig: defaultLorenzoBtcConfig(),
	}

	customAppTemplate := serverconfig.DefaultConfigTemplate + config.DefaultConfigTemplate + DefaultConfigTemplate

	return customAppTemplate, customAppConfig
}

func DefaultLorenzoTemplate() string {
	return serverconfig.DefaultConfigTemplate + config.DefaultConfigTemplate + DefaultConfigTemplate
}

func DefaultConfig() *Config {
	return &Config{
		Config:    *config.DefaultConfig(),
		BtcConfig: defaultLorenzoBtcConfig(),
	}
}
