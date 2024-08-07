package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrConvertDisabled        = errorsmod.Register(ModuleName, 2, "token module is disabled")
	ErrInternalTokenPair      = errorsmod.Register(ModuleName, 3, "internal ethereum token mapping error")
	ErrTokenPairNotFound      = errorsmod.Register(ModuleName, 4, "token pair not found")
	ErrTokenPairAlreadyExists = errorsmod.Register(ModuleName, 5, "token pair already exists")
	ErrUndefinedOwner         = errorsmod.Register(ModuleName, 6, "undefined owner of contract pair")
	ErrBalanceInvariance      = errorsmod.Register(ModuleName, 7, "post transfer balance invariant failed")
	ErrUnexpectedEvent        = errorsmod.Register(ModuleName, 8, "unexpected event")
	ErrABIPack                = errorsmod.Register(ModuleName, 9, "contract ABI pack failed")
	ErrABIUnpack              = errorsmod.Register(ModuleName, 10, "contract ABI unpack failed")
	ErrEVMCall                = errorsmod.Register(ModuleName, 11, "EVM call unexpected error")
	ErrTokenPairDisabled      = errorsmod.Register(ModuleName, 12, "token pair is disabled")
	ErrInvalidToken           = errorsmod.Register(ModuleName, 13, "invalid token")
	ErrInvalidDenom           = errorsmod.Register(ModuleName, 14, "invalid denom")
)
