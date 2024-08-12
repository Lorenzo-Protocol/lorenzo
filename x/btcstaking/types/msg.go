package types

import (
	fmt "fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ensure that these message types implement the sdk.Msg interface
var (
	_ sdk.Msg = &MsgCreateBTCStaking{}
	_ sdk.Msg = &MsgBurnRequest{}
	_ sdk.Msg = &MsgRemoveReceiver{}
	_ sdk.Msg = &MsgAddReceiver{}
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgCreateBTCBStaking{}
)

func (m *MsgCreateBTCStaking) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Signer); err != nil {
		return errorsmod.Wrap(err, "invalid signer address")
	}

	if m.StakingTx == nil {
		return fmt.Errorf("empty staking tx info")
	}
	// staking tx should be correctly formatted
	if err := m.StakingTx.ValidateBasic(); err != nil {
		return err
	}
	if m.AgentId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "agent id cannot be zero")
	}

	return nil
}

func (m *MsgCreateBTCStaking) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}

	return []sdk.AccAddress{signer}
}

func (m *MsgBurnRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Signer); err != nil {
		return errorsmod.Wrap(err, "invalid signer address")
	}

	if len(strings.TrimSpace(m.BtcTargetAddress)) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "btc target address cannot be empty")
	}

	if m.Amount.IsZero() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount cannot be zero")
	}

	return nil
}

func (m *MsgBurnRequest) ValidateBtcAddress(btcNetworkParams *chaincfg.Params) error {
	_, err := btcutil.DecodeAddress(m.BtcTargetAddress, btcNetworkParams)
	if err != nil {
		return err
	}
	return nil
}

func (m *MsgBurnRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Signer)
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

func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	if err := m.Params.Validate(); err != nil {
		return err
	}
	return nil
}

// ValidateBasic implements sdk.Msg
func (m *MsgCreateBTCBStaking) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Signer); err != nil {
		return errorsmod.Wrap(err, "invalid signer address")
	}

	if len(m.Receipt) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "staking receipt cannot be empty")
	}

	if len(m.Proof) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "staking proof cannot be empty")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgCreateBTCBStaking) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Signer)
	return []sdk.AccAddress{addr}
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
