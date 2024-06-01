package types

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ensure that these message types implement the sdk.Msg interface
var (
	_ sdk.Msg = &MsgCreateBTCStaking{}
	_ sdk.Msg = &MsgBurnRequest{}
	_ sdk.Msg = &MsgRemoveReceiver{}
	_ sdk.Msg = &MsgAddReceiver{}
	_ sdk.Msg = &MsgUpdateAllowList{}
)

func (m *MsgCreateBTCStaking) ValidateBasic() error {
	if m.StakingTx == nil {
		return fmt.Errorf("empty staking tx info")
	}
	if len(m.Receiver) == 0 {
		return fmt.Errorf("receiver name cannot be empty")
	}
	// staking tx should be correctly formatted
	if err := m.StakingTx.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

func (m *MsgAddReceiver) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgAddReceiver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	if err := m.Receiver.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgRemoveReceiver) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgRemoveReceiver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	if len(m.Receiver) == 0 {
		return fmt.Errorf("receiver name cannot be empty")
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

func (m *MsgBurnRequest) ValidateBtcAddress(btcNetworkParams *chaincfg.Params) error {
	_, err := btcutil.DecodeAddress(m.BtcTargetAddress, btcNetworkParams)
	if err != nil {
		return err
	}
	return nil
}

func (msg *MsgBurnRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}

	return []sdk.AccAddress{signer}
}

func NewMsgBurnRequest(signer, btcTargetAddress string, amount math.Int) MsgBurnRequest {
	return MsgBurnRequest{
		Signer:           signer,
		BtcTargetAddress: btcTargetAddress,
		Amount:           amount,
	}
}

func (m *MsgUpdateAllowList) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdateAllowList) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	if err := ValidateAddressList(m.MinterAllowList); err != nil {
		return err
	}
	return nil
}
