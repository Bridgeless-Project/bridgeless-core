package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/bridge module sentinel errors
var (
	ErrSample                       = errorsmod.Register(ModuleName, 1100, "sample error")
	ErrSourceChainNotSupported      = errorsmod.Register(ModuleName, 1101, "source chain not supported")
	ErrDestinationChainNotSupported = errorsmod.Register(ModuleName, 1102, "destination chain not supported")
	ErrTranscationAlreadySubmitted  = errorsmod.Register(ModuleName, 1103, "transaction already submitted")
	ErrOperationNotAllowed          = errorsmod.Register(ModuleName, 1104, "operation not allowed")
	ErrTokenNotFound                = errorsmod.Register(ModuleName, 1105, "token not found")
	ErrTokenPairNotFound            = errorsmod.Register(ModuleName, 1106, "token pair not found")
	ErrTokenInfoNotFound            = errorsmod.Register(ModuleName, 1109, "token info not found")
	ErrChainNotFound                = errorsmod.Register(ModuleName, 1107, "chain not found")
	ErrPermissionDenied             = errorsmod.Register(ModuleName, 1108, "permission denied")
	InvalidTransaction              = errorsmod.Register(ModuleName, 1110, "invalid transaction")
	ErrInvalidPartiesList           = errorsmod.Register(ModuleName, 1111, "invalid parties list")
	ErrInvalidCommissionRate        = errorsmod.Register(ModuleName, 1112, "invalid commission rate")
	ErrInvalidTssThreshold          = errorsmod.Register(ModuleName, 1113, "invalid tss threshold")
	ErrInvalidTxHash                = errorsmod.Register(ModuleName, 1114, "invalid tx hash")
)
