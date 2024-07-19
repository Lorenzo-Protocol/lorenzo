package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// StakePlanHubContractABI is the ABI for the StakePlanHub contract. TODO
var StakePlanHubContractABI abi.ABI

// BNBCrossChainEvent is a struct that contains the sender, plan id, BTC contract address, stake amount, and stBTC amount.
type BNBCrossChainEvent struct {
	EventIndex         uint64         `json:"event_index"`
	Sender             common.Address `json:"sender"`
	PlanID             uint64         `json:"plan_id"`
	BTCcontractAddress common.Address `json:"btc_contract_address"`
	StakeAmount        *big.Int       `json:"stake_amount"`
	StBTCAmount        *big.Int       `json:"st_btc_amount"`
}

// Key returns the unique key of the EventRecord struct
func(e EventRecord) Key() []byte {
	return KeyEventRecord(e.BlockNumber, e.Contract, e.Index)
}
