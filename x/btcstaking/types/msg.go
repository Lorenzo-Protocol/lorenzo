package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure that these message types implement the sdk.Msg interface
var (
	_ sdk.Msg = &MsgCreateBTCStaking{}
	_ sdk.Msg = &MsgBurnRequest{}
)

func (m *MsgCreateBTCStaking) ValidateBasic() error {
	if m.StakingTx == nil {
		return fmt.Errorf("empty staking tx info")
	}
	// staking tx should be correctly formatted
	if err := m.StakingTx.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

func (msg *MsgCreateBTCStaking) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}

	return []sdk.AccAddress{signer}
}

func (m *MsgBurnRequest) ValidateBasic() error {
	return nil
}

func (msg *MsgBurnRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}

	return []sdk.AccAddress{signer}
}

func NewMsgBurnRequest(signer, btcTargetAddress string, amount uint64) MsgBurnRequest {
	return MsgBurnRequest{
		Signer:           signer,
		BtcTargetAddress: btcTargetAddress,
		Amount:           amount,
	}
}
