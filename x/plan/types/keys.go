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
	YATMethodInitialize                = "initialize"
	YATMethodMint                      = "mint"
	YATMethodClaimYATToken             = "claimYATToken"
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

const (
	BeaconMethodUpgradeTo = "upgradeTo"
)

// ModuleAddress is the native module address for the module
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

var (
	ParamsKey             = []byte{0x01}
	KeyPrefixNextNumber   = []byte{0x02}
	KeyPrefixPlan         = []byte{0x03}
	KeyPrefixPlanContract = []byte{0x04}

	Delimiter = []byte{0x00}
)

func KeyPlan(id uint64) []byte {
	bz := sdk.Uint64ToBigEndian(id)
	return append(KeyPrefixPlan, bz...)
}

func KeyNextNumber() []byte {
	return KeyPrefixNextNumber
}

func KeyPlanContract(contractAddr string) []byte {
	key := make([]byte, len(contractAddr))
	copy(key, KeyPrefixPlanContract)
	copy(key[len(KeyPrefixPlanContract):], contractAddr)
	return key
}
