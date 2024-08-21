package types

import (
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
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

var (
	_ legacytx.LegacyMsg = (*MsgUpdateParams)(nil)
	_ legacytx.LegacyMsg = (*MsgUpgradePlan)(nil)
	_ legacytx.LegacyMsg = (*MsgCreatePlan)(nil)
	_ legacytx.LegacyMsg = (*MsgClaims)(nil)
	_ legacytx.LegacyMsg = (*MsgCreateYAT)(nil)
	_ legacytx.LegacyMsg = (*MsgUpdatePlanStatus)(nil)
	_ legacytx.LegacyMsg = (*MsgSetMinter)(nil)
	_ legacytx.LegacyMsg = (*MsgRemoveMinter)(nil)
	_ legacytx.LegacyMsg = (*MsgSetMerkleRoot)(nil)
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

func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateParams) Route() string {
	return ""
}

func (m *MsgUpdateParams) Type() string {
	return "lorenzo/plan/MsgUpdateParams"
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

func (m *MsgUpgradePlan) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUpgradePlan) Route() string {
	return ""
}

func (m *MsgUpgradePlan) Type() string {
	return "lorenzo/plan/MsgUpgradePlan"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgCreatePlan) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.YatContractAddress) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid yat contract address")
	}
	if len(strings.TrimSpace(m.Name)) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "plan name cannot be empty")
	}
	if m.AgentId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "agent id cannot be zero")
	}
	if m.PlanStartTime == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "plan start time cannot be zero")
	}
	if m.PeriodTime == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "period time cannot be zero")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgCreatePlan) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgCreatePlan) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgCreatePlan) Route() string {
	return ""
}

func (m *MsgCreatePlan) Type() string {
	return "lorenzo/plan/MsgCreatePlan"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgClaims) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.Receiver) {
		return errorsmod.Wrap(ErrReceiver, "invalid receiver address")
	}

	merkleProofs := strings.Split(m.MerkleProof, ",")
	if len(merkleProofs) == 0 {
		return fmt.Errorf("invalid merkle proof")
	}

	// check merkle proof is hex hash
	for _, merkleProof := range merkleProofs {
		merkleProof = strings.TrimSpace(merkleProof)
		if len(merkleProof) == 0 {
			return fmt.Errorf("invalid merkle proof")
		}
		merkleRoot := common.HexToHash(merkleProof)
		if merkleRoot.String() != merkleProof {
			return fmt.Errorf("invalid merkle proof")
		}
	}

	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgClaims) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgClaims) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgClaims) Route() string {
	return ""
}

func (m *MsgClaims) Type() string {
	return "lorenzo/plan/MsgClaims"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgCreateYAT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if len(strings.TrimSpace(m.Name)) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "yat name cannot be empty")
	}
	if len(strings.TrimSpace(m.Symbol)) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "yat symbol cannot be empty")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgCreateYAT) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgCreateYAT) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgCreateYAT) Route() string {
	return ""
}

func (m *MsgCreateYAT) Type() string {
	return "lorenzo/plan/MsgCreateYAT"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgUpdatePlanStatus) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if m.PlanId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "plan id cannot be zero")
	}
	if m.Status < 0 || m.Status > 1 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid status")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgUpdatePlanStatus) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdatePlanStatus) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgUpdatePlanStatus) Route() string {
	return ""
}

func (m *MsgUpdatePlanStatus) Type() string {
	return "lorenzo/plan/MsgUpdatePlanStatus"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgSetMinter) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.Minter) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid minter address")
	}
	if !common.IsHexAddress(m.ContractAddress) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid yat contract address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgSetMinter) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgSetMinter) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgSetMinter) Route() string {
	return ""
}

func (m *MsgSetMinter) Type() string {
	return "lorenzo/plan/MsgSetMinter"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgRemoveMinter) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(m.Minter) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid minter address")
	}
	if !common.IsHexAddress(m.ContractAddress) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid yat contract address")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgRemoveMinter) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgRemoveMinter) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgRemoveMinter) Route() string {
	return ""
}

func (m *MsgRemoveMinter) Type() string {
	return "lorenzo/plan/MsgRemoveMinter"
}

// ValidateBasic executes sanity validation on the provided data
func (m *MsgSetMerkleRoot) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(err, "invalid sender address")
	}

	if m.PlanId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "plan id cannot be zero")
	}

	merkleRoot := common.HexToHash(m.MerkleRoot)
	if merkleRoot.String() != m.MerkleRoot {
		return fmt.Errorf("invalid merkle root")
	}
	return nil
}

// GetSigners returns the expected signers for a MsgAddAgent message
func (m *MsgSetMerkleRoot) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m *MsgSetMerkleRoot) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(m))
}

func (m *MsgSetMerkleRoot) Route() string {
	return ""
}

func (m *MsgSetMerkleRoot) Type() string {
	return "lorenzo/plan/MsgSetMerkleRoot"
}
