package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed compiled_contracts/YATProxy.json
	YATProxyJSON []byte //nolint: golint

	//go:embed compiled_contracts/UpgradeableBeacon.json
	BeaconJSON []byte //nolint: golint

	//go:embed compiled_contracts/YieldAccruingToken.json
	YieldAccruingTokenJSON []byte //nolint: golint

	// YieldAccruingTokenContract is the compiled yield accruing token contract
	YieldAccruingTokenContract evmtypes.CompiledContract

	// BeaconContract is the compiled beacon contract proxy
	BeaconContract evmtypes.CompiledContract

	// YATProxyContract is the compiled yat contract proxy
	YATProxyContract evmtypes.CompiledContract
)

func init() {
	// contract code
	// https://github.com/Lorenzo-Protocol/builtin-contracts/blob/main/contracts/YieldAccruingToken.sol
	err := json.Unmarshal(YieldAccruingTokenJSON, &YieldAccruingTokenContract)
	if err != nil {
		panic(err)
	}

	if len(YieldAccruingTokenContract.Bin) == 0 {
		panic("load YieldAccruingToken contract failed")
	}

	err = json.Unmarshal(BeaconJSON, &BeaconContract)
	if err != nil {
		panic(err)
	}

	if len(BeaconContract.Bin) == 0 {
		panic("load Beacon contract failed")
	}

	err = json.Unmarshal(YATProxyJSON, &YATProxyContract)
	if err != nil {
		panic(err)
	}

	if len(YATProxyContract.Bin) == 0 {
		panic("load YATProxy contract failed")
	}
}
