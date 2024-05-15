package types

import (
	fmt "fmt"

	"cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	satBaseUint      = 1e10
	btcDustThreshold = 546 * satBaseUint
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
	if m.Amount.ModRaw(satBaseUint) != math.ZeroInt() {
		return fmt.Errorf("amount must be a multiple of %v", satBaseUint)
	}

	if m.Amount.LTE(math.NewInt(btcDustThreshold)) {
		return fmt.Errorf("amount must be greater than %v", btcDustThreshold)
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
