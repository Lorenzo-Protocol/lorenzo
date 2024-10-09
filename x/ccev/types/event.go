package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var abiCache = map[string]*abi.ABI{}

// EventHandler defines an event processing interface
type EventHandler interface {
	// Processed returns true if the event has been processed
	PreProcessed(ctx sdk.Context, chainID uint32, eventInfo *Event) error

	// Process processes the event and executes the corresponding logic.
	// It takes the context, chain id, address of the contract that emitted the event, the topics of the event, and the event args.
	// It returns an error if the processing fails.
	Process(ctx sdk.Context, chainID uint32, events []*Event) error
}

// Event is a struct that contains the topics and args of an event
type Event struct {
	Address common.Address
	Topics  []common.Hash
	Args    []any
}

// DecodeABI takes a byte slice that represents the abi of a contract and returns a pointer to the abi struct
// If the unmarshalling fails, it returns an error
func DecodeABI(address string, abiBz []byte) (*abi.ABI, error) {
	cached, ok := abiCache[address]
	if ok {
		return cached, nil
	}

	abi := new(abi.ABI)
	if err := json.Unmarshal(abiBz, abi); err != nil {
		return nil, err
	}
	abiCache[address] = abi
	return abi, nil
}
