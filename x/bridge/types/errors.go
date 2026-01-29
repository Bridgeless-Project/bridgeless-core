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
	ErrInvalidConfirmationsNumber   = errorsmod.Register(ModuleName, 1115, "invalid confirmations number")
	ErrInvalidChainName             = errorsmod.Register(ModuleName, 1116, "invalid chain name")
	ErrTransactionNotFound          = errorsmod.Register(ModuleName, 1117, "transaction not found")
	ErrReferralNotFound             = errorsmod.Register(ModuleName, 1118, "referral not found")
	ErrReferralRewardsNotFound      = errorsmod.Register(ModuleName, 1119, "referral rewards not found")
	ErrAlreadyExists                = errorsmod.Register(ModuleName, 1120, "entity already exists")
	ErrReferralIdMustBePositive     = errorsmod.Register(ModuleName, 1121, "referral id must be positive")
	ErrTokenIdMustBePositive        = errorsmod.Register(ModuleName, 1122, "token id must be positive")
	ErrInvalidDataType              = errorsmod.Register(ModuleName, 1123, "invalid data type")
	ErrInvalidEpochID               = errorsmod.Register(ModuleName, 1124, "invalid epoch ID")
	ErrPackEvent                    = errorsmod.Register(ModuleName, 1125, "failed to pack event")
	ErrEpochNotFound                = errorsmod.Register(ModuleName, 1126, "epoch not found")
)
