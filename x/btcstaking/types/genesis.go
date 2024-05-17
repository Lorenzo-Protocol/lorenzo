package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: &Params{
			BtcConfirmationsDepth: 6,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if gs.Params == nil {
		return fmt.Errorf("params cannot be nil")
	}
	receivers := map[string]bool{}
	for _, receiver := range gs.Params.Receivers {
		if _, receiverExists := receivers[receiver.Name]; receiverExists {
			return fmt.Errorf("duplicate receiver name: %s", receiver)
		}
		if len(receiver.Name) == 0 {
			return fmt.Errorf("receiver name cannot be empty")
		}
		if len(receiver.Addr) == 0 {
			return fmt.Errorf("receiver addr cannot be empty")
		}
		receivers[receiver.Name] = true
	}
	if gs.Params.BtcConfirmationsDepth == 0 {
		return fmt.Errorf("btc confirmations depth cannot be 0")
	}
	return nil
}
