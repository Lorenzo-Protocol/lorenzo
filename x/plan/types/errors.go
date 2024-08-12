package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrABIPack                      = errorsmod.Register(ModuleName, 2, "contract ABI pack failed")
	ErrABIUnpack                    = errorsmod.Register(ModuleName, 3, "contract ABI unpack failed")
	ErrUnauthorized                 = errorsmod.Register(ModuleName, 4, "unauthorized address")
	ErrReceiver                     = errorsmod.Register(ModuleName, 5, "invalid receiver address")
	ErrContractNotFound             = errorsmod.Register(ModuleName, 6, "contract not found")
	ErrPlanNotFound                 = errorsmod.Register(ModuleName, 7, "plan not found")
	ErrBeaconNotSet                 = errorsmod.Register(ModuleName, 8, "beacon not set")
	ErrVMExecution                  = errorsmod.Register(ModuleName, 9, "VM execution failed")
	ErrAgentNotFound                = errorsmod.Register(ModuleName, 10, "agent not found")
	ErrInvalidPlanStatus            = errorsmod.Register(ModuleName, 11, "invalid plan status")
	ErrContractAddress              = errorsmod.Register(ModuleName, 12, "invalid contract address")
	ErrEthAddress                   = errorsmod.Register(ModuleName, 13, "invalid Ethereum address")
	ErrInvalidUpdateMinterType      = errorsmod.Register(ModuleName, 14, "invalid update minter type")
	ErrYatContractNotFound          = errorsmod.Register(ModuleName, 15, "yat contract not found")
	ErrYatContractNotContract       = errorsmod.Register(ModuleName, 16, "yat contract is not a contract")
	ErrStakePlanContractNotFound    = errorsmod.Register(ModuleName, 17, "stake plan contract not found")
	ErrStakePlanContractNotContract = errorsmod.Register(ModuleName, 18, "stake plan contract is not a contract")
	ErrInvalidPlanStartTime         = errorsmod.Register(ModuleName, 19, "invalid plan start time")
	ErrPlanPaused                   = errorsmod.Register(ModuleName, 20, "plan is paused")
	ErrMerkelRootIsInvalid          = errorsmod.Register(ModuleName, 21, "merkle root is invalid")
	ErrMerkleProofIsInvalid         = errorsmod.Register(ModuleName, 22, "merkle proof is invalid")
)
