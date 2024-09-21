package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// EventHandler defines an event processing interface
type EventHandler interface {
	// GetUniqueID takes the context, topics, and eventArgs of an event and returns the event id, which is a unique string
	// that can be used to identify the event. The event id is used to store the event in the events store.
	GetUniqueID(ctx sdk.Context, topics []common.Hash, eventArgs []any) string

	// Execute processes the event and executes the corresponding logic.
	// It takes the context, chain id, address of the contract that emitted the event, the topics of the event, and the event args.
	// It returns an error if the processing fails.
	Execute(ctx sdk.Context, chainID uint32, address common.Address, topics []common.Hash, eventArgs []any) error
}

// DecodeABI takes a byte slice that represents the abi of a contract and returns a pointer to the abi struct
// If the unmarshalling fails, it returns an error
func DecodeABI(abiBz []byte) (*abi.ABI, error) {
	abi := new(abi.ABI)
	// unmarshal the StakePlanHubContractABI
	if err := json.Unmarshal(abiBz, abi); err != nil {
		return nil, err
	}
	return abi, nil
}
