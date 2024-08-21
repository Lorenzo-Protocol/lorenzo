package types

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

var (
	_ sdk.Msg = &MsgRegisterCoin{}
	_ sdk.Msg = &MsgRegisterERC20{}
	_ sdk.Msg = &MsgToggleConversion{}
	_ sdk.Msg = &MsgConvertCoin{}
	_ sdk.Msg = &MsgConvertERC20{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// ValidateBasic implements sdk.Msg. It checks:
// - metadata follows Metadata validation
// - base denom follows native denom and ibc denom validation
// - base denom must not be erc20/hex format.
func (m *MsgRegisterCoin) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	if len(m.Metadata) == 0 {
		return errorsmod.Wrap(errors.New("metadata cannot be empty"), "should at least have one metadata")
	}

	// TODO: enforce ibc and erc20 denom validation on metadata unit denom as well?
	seenDenom := make(map[string]bool)
	for _, metadata := range m.Metadata {
		if err := metadata.Validate(); err != nil {
			return err
		}

		// validate denom follows native denom and ibc denom spec.
		if err := ibctransfertypes.ValidateIBCDenom(metadata.Base); err != nil {
			return err
		}

		// NOTE: doesn't expect denom be erc20 denom.
		if err := ValidateERC20Denom(metadata.Base); err == nil {
			return errorsmod.Wrap(errors.New("unexpected denom"), "should not be erc20 denom")
		}

		if seenDenom[metadata.Base] {
			return errorsmod.Wrapf(errortypes.ErrInvalidCoins, "duplicate base denom %s", metadata.Base)
		}
		seenDenom[metadata.Base] = true
	}
	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgRegisterCoin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgRegisterCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgRegisterCoin) Route() string {
	return ""
}

func (m *MsgRegisterCoin) Type() string {
	return "lorenzo/token/MsgRegisterCoin"
}

// ValidateBasic implements sdk.Msg
func (m *MsgRegisterERC20) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	if len(m.ContractAddresses) == 0 {
		return errorsmod.Wrap(errors.New("contract addresses cannot be empty"), "should at least have one contract address")
	}

	seenAddr := make(map[string]bool)
	for _, addr := range m.ContractAddresses {
		if !common.IsHexAddress(addr) {
			return errorsmod.Wrapf(
				errortypes.ErrInvalidAddress, "address %s is not a valid ethereum hex address", addr)
		}

		if seenAddr[addr] {
			return errorsmod.Wrapf(
				errortypes.ErrInvalidAddress, "duplicate contract address %s", addr)
		}

		seenAddr[addr] = true
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgRegisterERC20) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgRegisterERC20) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgRegisterERC20) Route() string {
	return ""
}

func (m *MsgRegisterERC20) Type() string {
	return "lorenzo/token/MsgRegisterERC20"
}

// ValidateBasic implements sdk.Msg
func (m *MsgToggleConversion) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	if err := sdk.ValidateDenom(m.Token); err != nil {
		if !common.IsHexAddress(m.Token) {
			return errorsmod.Wrapf(ErrInvalidToken, "%s is neither valid sdk denom nor evm hex address", m.Token)
		}
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgToggleConversion) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgToggleConversion) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgToggleConversion) Route() string {
	return ""
}

func (m *MsgToggleConversion) Type() string {
	return "lorenzo/token/MsgToggleConversion"
}

// ValidateBasic implements sdk.Msg
func (m *MsgConvertCoin) ValidateBasic() error {
	if err := ValidateERC20Denom(m.Coin.Denom); err != nil {
		if err := ibctransfertypes.ValidateIBCDenom(m.Coin.Denom); err != nil {
			return errorsmod.Wrapf(ErrInvalidDenom,
				"%s is neither valid erc20, nor native denom, nor ibc denom", m.Coin.Denom)
		}
	}

	if !m.Coin.Amount.IsPositive() {
		return errorsmod.Wrapf(errortypes.ErrInvalidCoins, "non-positive amount")
	}

	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}

	if !common.IsHexAddress(m.Receiver) {
		return errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid receiver hex address %s", m.Receiver)
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgConvertCoin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgConvertCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgConvertCoin) Route() string {
	return ""
}

func (m *MsgConvertCoin) Type() string {
	return "lorenzo/token/MsgConvertCoin"
}

// ValidateBasic implements sdk.Msg
func (m *MsgConvertERC20) ValidateBasic() error {
	if !common.IsHexAddress(m.ContractAddress) {
		return errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid contract hex address %s", m.ContractAddress)
	}

	if !m.Amount.IsPositive() {
		return errorsmod.Wrapf(errortypes.ErrInvalidCoins, "non-positive amount")
	}

	if !common.IsHexAddress(m.Sender) {
		return errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid sender hex address %s", m.Sender)
	}

	_, err := sdk.AccAddressFromBech32(m.Receiver)
	if err != nil {
		return errorsmod.Wrap(err, "invalid receiver address")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgConvertERC20) GetSigners() []sdk.AccAddress {
	addr := common.HexToAddress(m.Sender)
	return []sdk.AccAddress{addr.Bytes()}
}

func (m *MsgConvertERC20) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgConvertERC20) Route() string {
	return ""
}

func (m *MsgConvertERC20) Type() string {
	return "lorenzo/token/MsgConvertERC20"
}

// ValidateBasic implements sdk.Msg
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements sdk.Msg
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateParams) Route() string {
	return ""
}

func (m *MsgUpdateParams) Type() string {
	return "lorenzo/token/MsgUpdateParams"
}
