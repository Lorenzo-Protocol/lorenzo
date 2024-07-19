package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	// ErrHeaderNotFound is returned when header is not found
	ErrHeaderNotFound = errorsmod.Register(ModuleName, 2, "header not found")

	// ErrInvalidHeader is returned when header is not valid
	ErrInvalidHeader = errorsmod.Register(ModuleName, 3, "header not found")
	// ErrUnauthorized is returned when tx is not authorized
	ErrUnauthorized = errorsmod.Register(ModuleName, 4, "tx unauthorized")
	// ErrInvalidProof is returned when proof is not valid
	ErrInvalidProof = errorsmod.Register(ModuleName, 5, "invalid proof")
	// ErrInvalidTransaction is returned when transaction is not valid
	ErrInvalidTransaction = errorsmod.Register(ModuleName, 6, "invalid transaction")
	// ErrInvalidEvent is returned when event is not valid
	ErrInvalidEvent = errorsmod.Register(ModuleName, 7, "invalid event")
)
