# Changelog

## [Unreleased]

## v2.0.0

### New Features

- (ante) [\#58](https://github.com/Lorenzo-Protocol/lorenzo/pull/58) Unified fee settings for Cosmos and EVM transactions.
- (x/plan) [\#26](https://github.com/Lorenzo-Protocol/lorenzo/pull/26) Introduced the Plan module for managing staking plans.
- (x/agent) [\#23](https://github.com/Lorenzo-Protocol/lorenzo/pull/23) Introduced the Agent module for managing staking plan's agents.
- (x/token) [\#61](https://github.com/Lorenzo-Protocol/lorenzo/pull/61) Introduced the Token module for managing tokens between ERC20 contracts and SDK bank.
- (x/ibctransfer) [\#61](https://github.com/Lorenzo-Protocol/lorenzo/pull/61) Wrapped the IBC transfer module to support automatic conversion for IBC assets.
- (x/btcstaking) [\#40](https://github.com/Lorenzo-Protocol/lorenzo/pull/40) Implemented dynamic confirmation depth.
- (x/btcstaking) [\#42](https://github.com/Lorenzo-Protocol/lorenzo/pull/42) Added allow list for receivers with ETH addresses.
- (x/btcstaking) [\#67](https://github.com/Lorenzo-Protocol/lorenzo/pull/67) Added transaction output amount limit to filter out dust outputs.
- (x/btcstaking) [\#74](https://github.com/Lorenzo-Protocol/lorenzo/pull/74) Enabled minting stBTC to multiple chains.

### Bug Fixes

- (app) [\#63](https://github.com/Lorenzo-Protocol/lorenzo/pull/63) Fixed Tendermint client type registration.
- (client) [\#47](https://github.com/Lorenzo-Protocol/lorenzo/pull/47) Fixed error handling for BTC staking query commands.
- (x/btcstaking) [\#37](https://github.com/Lorenzo-Protocol/lorenzo/pull/37) Added BTC staking genesis validation.
- (x/btcstaking) [\#95](https://github.com/Lorenzo-Protocol/lorenzo/pull/95) Remove out-dated check
- (test) [\#89](https://github.com/Lorenzo-Protocol/lorenzo/pull/89) Fixed testnet cmd

### Improvements

- (client) [\#17](https://github.com/Lorenzo-Protocol/lorenzo/pull/17) BTC staking query now accepts transaction ID instead of transaction hash.
- (client) [\#53](https://github.com/Lorenzo-Protocol/lorenzo/pull/53) Renamed BTC staking query commands.
- (x/btcstaking) [\#66](https://github.com/Lorenzo-Protocol/lorenzo/pull/66) Accepted `OP_PUSHDATA2` and `OP_PUSHDATA` in `OP_RETURN` transaction outputs.
- (x/btcstaking) [\#98](https://github.com/Lorenzo-Protocol/lorenzo/pull/98) Refactor msg_server and fix some standard writing methods.


## 1.0.0

### Features

* (Lorenzo) [\#1](https://github.com/Lorenzo-Protocol/lorenzo/pull/1) Add btclightclient module from version v0.8.5 & v0.7.2 of third party
* (Lorenzo) [\#8](https://github.com/Lorenzo-Protocol/lorenzo/pull/8) support nonfee transaction
* (Lorenzo) [\#11](https://github.com/Lorenzo-Protocol/lorenzo/pull/11) set MsgInsertHeaders to be fee free
* (Lorenzo) [\#12](https://github.com/Lorenzo-Protocol/lorenzo/pull/12) add fee module
* (Lorenzo) [\#13](https://github.com/Lorenzo-Protocol/lorenzo/pull/13) added implementation for submitting bitcoin fee rate in module btclightclient

### Bug Fixes

* (Lorenzo) [\#5](https://github.com/Lorenzo-Protocol/lorenzo/pull/5) Fix btc lightclient
* (Lorenzo) [\#9](https://github.com/Lorenzo-Protocol/lorenzo/pull/9) Fix GetSigners method of MsgUpdateParams
* (Lorenzo) [\#14](https://github.com/Lorenzo-Protocol/lorenzo/pull/14) Fix btcstaking genesis init