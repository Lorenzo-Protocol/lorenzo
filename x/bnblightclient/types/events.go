package types

import (
	_ "embed" //nolint: golint
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

//go:embed stake_plan_hub_abi.json
var stakePlanHubContractABIJSON []byte //nolint: golint

var stakePlanHubContractABI abi.ABI

func init() {
	// unmarshal the StakePlanHubContractABI
	err := json.Unmarshal(stakePlanHubContractABIJSON, &stakePlanHubContractABI)
	if err != nil {
		panic(err)
	}
}

// ABIstakePlanHub is the compiled StakePlanHub contract abi
func ABIstakePlanHub() *abi.ABI {
	return &stakePlanHubContractABI
}

// CrossChainEvent is a struct that contains the sender, plan id, BTC contract address, stake amount, and stBTC amount.
type CrossChainEvent struct {
	ChainID            uint32         `json:"chain_id"`
	Contract           common.Address `json:"contract"`
	Identifier         uint64         `json:"identifier"`
	Sender             common.Address `json:"sender"`
	PlanID             uint64         `json:"plan_id"`
	BTCcontractAddress common.Address `json:"btc_contract_address"`
	StakeAmount        big.Int        `json:"stake_amount"`
	StBTCAmount        big.Int        `json:"st_btc_amount"`
}