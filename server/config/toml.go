package config

const DefaultConfigTemplate = `

###############################################################################
###                             BTC Configuration                           ###
###############################################################################

[btc-config]

# Configures which bitcoin network should be used for checkpointing
# valid values are: [mainnet, testnet, simnet, regtest]
network = "{{ .BtcConfig.Network }}"
`
