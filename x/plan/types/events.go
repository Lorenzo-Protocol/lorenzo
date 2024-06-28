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
		sdk.NewAttribute(AttributeKeyCreatePlanPlanStartBlock, plan.PlanStartBlock.String()),
		sdk.NewAttribute(AttributeKeyCreatePlanPeriodBlocks, plan.PeriodBlocks.String()),
		sdk.NewAttribute(AttributeKeyCreatePlanYatContractAddress, plan.YatContractAddress),
		sdk.NewAttribute(AttributeKeyCreatePlanContractAddress, plan.ContractAddress),
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
