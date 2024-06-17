package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Plan module event types
const (
	EventTypeCreatePlan = "create_plan"

	AttributeKeyPlanId                    = "plan_id"
	AttributeKeyPlanName                  = "name"
	AttributeKeyPlanSymbol                = "symbol"
	AttributeKeyPlanDescUri               = "plan_desc_uri"
	AttributeKeyPlanAgentId               = "agent_id"
	AttributeKeyPlanSubscriptionStartTime = "subscription_start_time"
	AttributeKeyPlanSubscriptionEndTime   = "subscription_end_time"
	AttributeKeyPlanEndTime               = "end_time"
	AttributeKeyPlanContractAddress       = "contract_address"

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
		sdk.NewAttribute(AttributeKeyPlanId, fmt.Sprintf("%d", plan.Id)),
		sdk.NewAttribute(AttributeKeyPlanName, plan.Name),
		sdk.NewAttribute(AttributeKeyPlanSymbol, plan.Symbol),
		sdk.NewAttribute(AttributeKeyPlanDescUri, plan.PlanDescUri),
		sdk.NewAttribute(AttributeKeyPlanAgentId, fmt.Sprintf("%d", plan.AgentId)),
		sdk.NewAttribute(AttributeKeyPlanSubscriptionStartTime, fmt.Sprintf("%d", plan.SubscriptionStartTime)),
		sdk.NewAttribute(AttributeKeyPlanSubscriptionEndTime, fmt.Sprintf("%d", plan.SubscriptionEndTime)),
		sdk.NewAttribute(AttributeKeyPlanEndTime, fmt.Sprintf("%d", plan.EndTime)),
		sdk.NewAttribute(AttributeKeyPlanContractAddress, plan.ContractAddress),
	)
}
