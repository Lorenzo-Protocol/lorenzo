package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = (*MsgUpdateParams)(nil)
	_ sdk.Msg = (*MsgUpgradePlan)(nil)
	_ sdk.Msg = (*MsgCreatePlan)(nil)
	_ sdk.Msg = (*MsgClaims)(nil)
	_ sdk.Msg = (*MsgCreateYAT)(nil)
	_ sdk.Msg = (*MsgUpdatePlanStatus)(nil)
	_ sdk.Msg = (*MsgSetMinter)(nil)
	_ sdk.Msg = (*MsgRemoveMinter)(nil)
	_ sdk.Msg = (*MsgSetMerkleRoot)(nil)
)

// ValidateBasic executes sanity validation on the provided data
func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	return m.Params.Validate()
}

// GetSigners returns the expected signers for a MsgUpdateParams message
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgUpgradePlan) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}
	if !common.IsHexAddress(m.Implementation) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "expecting a hex address, got %s", m.Implementation)
	}
	return nil
}

// GetSigners returns the expected signers for a MsgUpdateParams message
func (m *MsgUpgradePlan) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgCreatePlan) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.YatContractAddress) {
		return errorsmod.Wrap(ErrContractAddress, "invalid yat contract address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgCreatePlan) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgClaims) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.Receiver) {
		return errorsmod.Wrap(ErrReceiver, "invalid receiver address")
	}
	merkleProof := common.HexToHash(m.MerkleProof)
	if len(merkleProof.Bytes()) != 32 {
		return fmt.Errorf("invalid merkle proof")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgClaims) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgCreateYAT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgCreateYAT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgUpdatePlanStatus) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgUpdatePlanStatus) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgSetMinter) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.ContractAddress) {
		return errorsmod.Wrap(ErrContractAddress, "invalid yat contract address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgSetMinter) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgRemoveMinter) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.ContractAddress) {
		return errorsmod.Wrap(ErrContractAddress, "invalid yat contract address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgRemoveMinter) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgSetMerkleRoot) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	merkleRoot := common.HexToHash(m.MerkleRoot)
	if len(merkleRoot.Bytes()) != 32 {
		return fmt.Errorf("invalid merkle root")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgSetMerkleRoot) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}
