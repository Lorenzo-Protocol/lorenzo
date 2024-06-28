package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrABIPack           = errorsmod.Register(ModuleName, 2, "contract ABI pack failed")
	ErrABIUnpack         = errorsmod.Register(ModuleName, 3, "contract ABI unpack failed")
	ErrUnauthorized      = errorsmod.Register(ModuleName, 4, "unauthorized address")
	ErrReceiver          = errorsmod.Register(ModuleName, 5, "invalid receiver address")
	ErrContractNotFound  = errorsmod.Register(ModuleName, 6, "contract not found")
	ErrPlanNotFound      = errorsmod.Register(ModuleName, 7, "plan not found")
	ErrBeaconNotSet      = errorsmod.Register(ModuleName, 8, "beacon not set")
	ErrVMExecution       = errorsmod.Register(ModuleName, 9, "VM execution failed")
	ErrAgentNotFound     = errorsmod.Register(ModuleName, 10, "agent not found")
	ErrInvalidPlanStatus = errorsmod.Register(ModuleName, 11, "invalid plan status")
)
