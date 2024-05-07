package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the name of the module
	ModuleName = "plan"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey = ModuleName
)

const (
	YATMethodMint                      = "mint"
	YATMethodClaimRewardAndWithDrawBTC = "claimRewardAndWithDrawBTC"
	YATMethodOnlyClaimReward           = "onlyClaimReward"
	YATMethodBurnWithstBTCBurn         = "burnWithstBTCBurn"
	YATMethodSetRewardTokenAddress     = "setRewardTokenAddress"

	// query method
	YATMethodPlanId                = "planId"
	YATMethodAgentId               = "agentId"
	YATMethodSubscriptionStartTime = "subscriptionStartTime"
	YATMethodSubscriptionEndTime   = "subscriptionEndTime"
	YATMethodEndTime               = "endTime"
	YATMethodPlanDesc              = "planDesc"
	YATMethodRewardTokenAddress    = "rewardTokenAddress"
)

// ModuleAddress is the native module address for the module
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

var (
	// PlanKey is the key to store the plan in the store
	PlanKey       = []byte{0x01}
	NextNumberKey = []byte{0x02}
)

func KeyPlan(id uint64) []byte {
	bz := sdk.Uint64ToBigEndian(id)
	return append(PlanKey, bz...)
}

func KeyNextNumber() []byte {
	return NextNumberKey
}
