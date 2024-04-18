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

func NewEventBurnCreated(signer sdk.AccAddress, btcTargetAddress btcutil.Address, amount, fee sdk.Coin) *EventBurnCreated {
	return &EventBurnCreated{
		Signer:           signer.String(),
		BtcTargetAddress: btcTargetAddress.String(),
		AmountDenom:      amount.Denom,
		AmountValue:      amount.Amount.Uint64(),
		FeeDenom:         fee.Denom,
		FeeValue:         fee.Amount.Uint64(),
	}
}
