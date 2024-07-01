package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	//go:embed compiled_contracts/StakePlanProxy.json
	StakePlanProxyJSON []byte //nolint: golint
	//go:embed compiled_contracts/StakePlan.json
	StakePlanJSON []byte //nolint: golint

	//go:embed compiled_contracts/UpgradeableBeacon.json
	BeaconJSON []byte //nolint: golint

	//go:embed compiled_contracts/YieldAccruingToken.json
	YieldAccruingTokenJSON []byte //nolint: golint

	// YieldAccruingTokenContract is the compiled yield accruing token contract
	YieldAccruingTokenContract evmtypes.CompiledContract

	// StakePlanContract is the compiled StakePlan contract
	StakePlanContract evmtypes.CompiledContract

	// BeaconContract is the compiled beacon contract proxy
	BeaconContract evmtypes.CompiledContract

	// StakePlanProxyContract is the compiled StakePlan contract proxy
	StakePlanProxyContract evmtypes.CompiledContract
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

	// contract code
	// https://github.com/Lorenzo-Protocol/builtin-contracts/blob/main/contracts/StakePlan.sol
	// unmarshal the compiled StakePlanContract contract
	err = json.Unmarshal(StakePlanJSON, &StakePlanContract)
	if err != nil {
		panic(err)
	}

	if len(StakePlanContract.Bin) == 0 {
		panic("load StakePlan contract failed")
	}

	// unmarshal the compiled BeaconContract contract
	err = json.Unmarshal(BeaconJSON, &BeaconContract)
	if err != nil {
		panic(err)
	}

	if len(BeaconContract.Bin) == 0 {
		panic("load Beacon contract failed")
	}

	// unmarshal the compiled StakePlanProxyContract contract
	err = json.Unmarshal(StakePlanProxyJSON, &StakePlanProxyContract)
	if err != nil {
		panic(err)
	}

	if len(StakePlanProxyContract.Bin) == 0 {
		panic("load StakePlanProxy contract failed")
	}
}
