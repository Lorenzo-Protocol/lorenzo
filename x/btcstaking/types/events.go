package types

import (
	sdkmath "cosmossdk.io/math"
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

func NewEventMintStBTC(cosmosAddr, ethAddr string, amount sdkmath.Int) *EventMintStBTC {
	return &EventMintStBTC{
		CosmosAddress: cosmosAddr,
		EthAddress:    ethAddr,
		Amount:        amount,
	}
}
