package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidID                = errorsmod.Register(ModuleName, 2, "invalid agent id")
	ErrNameEmpty                = errorsmod.Register(ModuleName, 3, "name cannot be empty")
	ErrBtcReceivingAddressEmpty = errorsmod.Register(ModuleName, 4, "btcReceivingAddress cannot be empty")
	ErrAgentNotFound            = errorsmod.Register(ModuleName, 5, "agent not found")
	ErrUnAuthorized             = errorsmod.Register(ModuleName, 6, "unauthorized")
	ErrInvalidBtcAddress        = errorsmod.Register(ModuleName, 7, "invalid btc address")
	ErrInvalidEthAddress        = errorsmod.Register(ModuleName, 8, "invalid eth address")
	ErrAdminExists              = errorsmod.Register(ModuleName, 9, "admin already exists")
)
