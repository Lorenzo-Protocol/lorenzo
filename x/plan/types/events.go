package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Plan module event types
const (
	EventTypeCreatePlan = "create_plan"

	AttributeKeyCreatePlanId                 = "plan_id"
	AttributeKeyCreatePlanName               = "name"
	AttributeKeyCreatePlanDescUri            = "plan_desc_uri"
	AttributeKeyCreatePlanAgentId            = "agent_id"
	AttributeKeyCreatePlanPlanStartBlock     = "plan_start_block"
	AttributeKeyCreatePlanPeriodBlocks       = "period_blocks"
	AttributeKeyCreatePlanYatContractAddress = "yat_contract_address"
	AttributeKeyCreatePlanContractAddress    = "contract_address"

	EventTypeUpgradePlan                     = "upgrade_plan"
	AttributeKeyUpgradePlanOldImplementation = "old_implementation"
	AttributeKeyUpgradePlanNewImplementation = "new_implementation"

	EventClaims = "claims"

	AttributeKeyClaimsPlanId      = "plan_id"
	AttributeKeyClaimsReceiver    = "receiver"
	AttributeKeyClaimsRoundId     = "round_id"
	AttributeKeyClaimsAmount      = "amount"
	AttributeKeyClaimsMerkleProof = "merkle_proof"

	EventCreateYAT                       = "create_yat"
	AttributeKeyCreateYATContractAddress = "yat_contract_address"
	AttributeKeyCreateYATName            = "name"
	AttributeKeyCreateYATSymbol          = "symbol"

	// EventUpdatePlanStatus
	EventTypeUpdatePlanStatus             = "update_plan_status"
	AttributeKeyUpdatePlanStatusPlanId    = "plan_id"
	AttributeKeyUpdatePlanStatusOldStatus = "old_status"
	AttributeKeyUpdatePlanStatusNewStatus = "new_status"

	// SetMerkleRoot
	EventTypeSetMerkleRoot              = "set_merkle_root"
	AttributeKeySetMerkleRootMerkleRoot = "merkle_root"
	AttributeKeySetMerkleRootPlanId     = "plan_id"

	// Minter
	EventTypeSetMinter    = "set_minter"
	EventTypeRemoveMinter = "remove_minter"

	AttributeKeyMinter   = "minter"
	AttributeKeyContract = "contract"

	AttributeKeySender = sdk.AttributeKeySender

	EventSetParams         = "set_params"
	AttributeKeyBeaconAddr = "beacon_addr"
	AttributeKeyLogicAddr  = "logic_addr"
)

// NewCreatePlanEvent construct a new plan created sdk.Event
func NewCreatePlanEvent(sender sdk.AccAddress, plan Plan) sdk.Event {
	return sdk.NewEvent(
		EventTypeCreatePlan,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyCreatePlanId, fmt.Sprintf("%d", plan.Id)),
		sdk.NewAttribute(AttributeKeyCreatePlanName, plan.Name),
		sdk.NewAttribute(AttributeKeyCreatePlanDescUri, plan.PlanDescUri),
		sdk.NewAttribute(AttributeKeyCreatePlanAgentId, fmt.Sprintf("%d", plan.AgentId)),
		sdk.NewAttribute(AttributeKeyCreatePlanPlanStartBlock, fmt.Sprintf("%d", plan.PlanStartBlock)),
		sdk.NewAttribute(AttributeKeyCreatePlanPeriodBlocks, fmt.Sprintf("%d", plan.PeriodBlocks)),
		sdk.NewAttribute(AttributeKeyCreatePlanYatContractAddress, plan.YatContractAddress),
		sdk.NewAttribute(AttributeKeyCreatePlanContractAddress, plan.ContractAddress),
	)
}

// NewUpgradePlanEvent construct a new plan upgrade sdk.Event
func NewUpgradePlanEvent(sender sdk.AccAddress, oldImplementation, newImplementation string) sdk.Event {
	return sdk.NewEvent(
		EventTypeUpgradePlan,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyUpgradePlanOldImplementation, oldImplementation),
		sdk.NewAttribute(AttributeKeyUpgradePlanNewImplementation, newImplementation),
	)
}

// NewClaimsEvent construct a new yat created sdk.Event
func NewClaimsEvent(
	sender sdk.AccAddress,
	planId uint64,
	receiver string,
	roundId string,
	amount string,
	merkleProof string,
) sdk.Event {
	return sdk.NewEvent(
		EventClaims,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyClaimsPlanId, fmt.Sprintf("%d", planId)),
		sdk.NewAttribute(AttributeKeyClaimsReceiver, receiver),
		sdk.NewAttribute(AttributeKeyClaimsRoundId, roundId),
		sdk.NewAttribute(AttributeKeyClaimsAmount, amount),
		sdk.NewAttribute(AttributeKeyClaimsMerkleProof, merkleProof),
	)
}

// NewCreateYATEvent construct a new yat created sdk.Event
func NewCreateYATEvent(
	sender sdk.AccAddress,
	yatContractAddress string,
	name string,
	symbol string,
) sdk.Event {
	return sdk.NewEvent(
		EventCreateYAT,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyCreateYATContractAddress, yatContractAddress),
		sdk.NewAttribute(AttributeKeyCreateYATName, name),
		sdk.NewAttribute(AttributeKeyCreateYATSymbol, symbol),
	)
}

// NewSetMinterEvent construct a new set minter sdk.Event
func NewSetMinterEvent(
	sender sdk.AccAddress,
	minter string,
	contract string,
) sdk.Event {
	return sdk.NewEvent(
		EventTypeSetMinter,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyMinter, minter),
		sdk.NewAttribute(AttributeKeyContract, contract),
	)
}

// NewRemoveMinterEvent construct a new remove minter sdk.Event
func NewRemoveMinterEvent(
	sender sdk.AccAddress,
	minter string,
	contract string,
) sdk.Event {
	return sdk.NewEvent(
		EventTypeRemoveMinter,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyMinter, minter),
		sdk.NewAttribute(AttributeKeyContract, contract),
	)
}

// NewUpdatePlanStatusEvent construct a new update plan status sdk.Event
func NewUpdatePlanStatusEvent(
	sender sdk.AccAddress,
	planId uint64,
	oldStatus PlanStatus,
	newStatus PlanStatus,
) sdk.Event {
	return sdk.NewEvent(
		EventTypeUpdatePlanStatus,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyUpdatePlanStatusPlanId, fmt.Sprintf("%d", planId)),
		sdk.NewAttribute(AttributeKeyUpdatePlanStatusOldStatus, oldStatus.String()),
		sdk.NewAttribute(AttributeKeyUpdatePlanStatusNewStatus, newStatus.String()),
	)
}

// NewSetMerkleRootEvent construct a new set merkle root sdk.Event
func NewSetMerkleRootEvent(
	sender sdk.AccAddress,
	planId uint64,
	merkleRoot string,
) sdk.Event {
	return sdk.NewEvent(
		EventTypeSetMerkleRoot,
		sdk.NewAttribute(AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeySetMerkleRootPlanId, fmt.Sprintf("%d", planId)),
		sdk.NewAttribute(AttributeKeySetMerkleRootMerkleRoot, merkleRoot),
	)
}
