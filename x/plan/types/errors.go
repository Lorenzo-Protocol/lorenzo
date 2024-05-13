package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrABIPack          = errorsmod.Register(ModuleName, 2, "contract ABI pack failed")
	ErrABIUnpack        = errorsmod.Register(ModuleName, 3, "contract ABI unpack failed")
	ErrUnauthorized     = errorsmod.Register(ModuleName, 4, "unauthorized address")
	ErrReceiver         = errorsmod.Register(ModuleName, 5, "invalid receiver address")
	ErrContractNotFound = errorsmod.Register(ModuleName, 6, "contract not found")
	ErrPlanNotFound     = errorsmod.Register(ModuleName, 7, "plan not found")
)
