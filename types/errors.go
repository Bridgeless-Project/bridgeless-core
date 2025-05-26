// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "bridgeless"

var (

	// ErrInvalidChainID returns an error resulting from an invalid chain ID.
	ErrInvalidChainID = errorsmod.Register(RootCodespace, 188, "invalid chain ID")

	// ErrInvalidGasCap returns an error if a the gas cap value is negative or invalid
	ErrInvalidGasCap = errorsmod.Register(RootCodespace, 189, "invalid gas cap")

	// ErrInvalidBaseFee returns an error if a the base fee cap value is invalid
	ErrInvalidBaseFee = errorsmod.Register(RootCodespace, 190, "invalid base fee")

	// ErrInvalidGasPrice returns an error if an invalid gas price is provided to the tx.
	ErrInvalidGasPrice = errorsmod.Register(RootCodespace, 191, "invalid gas price")

	// ErrDuplicatedValue returns an error if a duplicated value was used
	ErrDuplicatedValue = errorsmod.Register(RootCodespace, 192, "duplicated value")

	// ErrInvalidTxHash returns an error if used tx hash is invalid
	ErrInvalidTxHash = errorsmod.Register(RootCodespace, 193, "invalid tx hash")

	// ErrInvalidAmount returns an error if invalid amount is used
	ErrInvalidAmount = errorsmod.Register(RootCodespace, 194, "invalid amount")

	// ErrInvalidActionLength returns an error if invalid action length is used
	ErrInvalidActionLength = errorsmod.Register(RootCodespace, 195, "invalid action length")

	// ErrInvalidDenom returns an error if an invalid denomination is used
	ErrInvalidDenom = errorsmod.Register(RootCodespace, 196, "invalid denom")

	// ErrInvalidBlock returns an error if given block is invalid
	ErrInvalidBlock = errorsmod.Register(RootCodespace, 197, "invalid block")

	// ErrInvalidGasMultiplier returns an error if provided gas multiplier is invalid
	ErrInvalidGasMultiplier = errorsmod.Register(RootCodespace, 198, "invalid gas multiplier")

	// ErrInvalidDenominator returns an error if denominator is invalid
	ErrInvalidDenominator = errorsmod.Register(RootCodespace, 199, "invalid denominator")

	// ErrFailedToGetMsgs returns an error if core failed to get msgs
	ErrFailedToGetMsgs = errorsmod.Register(RootCodespace, 200, "failed to get msgs")

	// ErrInvalidTimeout returns an error if timeout is invalid
	ErrInvalidTimeout = errorsmod.Register(RootCodespace, 201, "invalid timeout")

	// ErrInvalidMetadata returns an error if metadata is invalid
	ErrInvalidMetadata = errorsmod.Register(RootCodespace, 202, "invalid metadata")
)
