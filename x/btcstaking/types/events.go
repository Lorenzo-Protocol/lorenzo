package types

import (
	"github.com/btcsuite/btcd/btcutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewEventBTCStakingCreated(record *BTCStakingRecord) *EventBTCStakingCreated {
	return &EventBTCStakingCreated{
		Record: record,
	}
}

func NewEventBTCBStakingCreated(record *BTCBStakingRecord) *EventBTCBStakingCreated {
	return &EventBTCBStakingCreated{
		Record: record,
	}
}

func NewEventBurnCreated(signer sdk.AccAddress, btcTargetAddress btcutil.Address, amount sdk.Coin) *EventBurnCreated {
	return &EventBurnCreated{
		Signer:           signer.String(),
		BtcTargetAddress: btcTargetAddress.String(),
		Amount:           amount,
	}
}
