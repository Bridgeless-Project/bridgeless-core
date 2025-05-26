package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

func validateToken(token *Token) error {
	if token == nil {
		return errors.New("token is nil")
	}

	rate, err := sdkmath.LegacyNewDecFromStr(token.CommissionRate)

	if err != nil {
		return errorsmod.Wrap(ErrInvalidCommissionRate, err.Error())
	}

	if rate.IsNegative() {
		return errorsmod.Wrap(ErrInvalidCommissionRate, "commission rate must be positive")
	}

	if rate.GT(sdkmath.LegacyNewDec(1)) {
		return errorsmod.Wrap(ErrInvalidCommissionRate, "commission rate must be <= 100%")
	}

	return validateTokenMetadata(&token.Metadata)
}

func validateTokenMetadata(metadata *TokenMetadata) error {
	if metadata == nil {
		return errors.New("metadata is nil")
	}

	if metadata.Name == "" {
		return errors.New("name cannot be empty")
	}
	if metadata.Symbol == "" {
		return errors.New("symbol cannot be empty")
	}
	if metadata.Uri == "" {
		return errors.New("uri cannot be empty")
	}

	return nil
}

func validateTokenInfo(info *TokenInfo, chainType *ChainType) error {
	if info == nil {
		return errors.New("info is nil")
	}
	if info.ChainId == "" {
		return errors.New("chain id cannot be empty")

	}
	if info.Address == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token address is empty")
	}
	if chainType == nil {
		return nil
	}

	switch *chainType {
	case ChainType_EVM:
		if !common.IsHexAddress(info.Address) {
			return errors.New(fmt.Sprintf("invalid token address: %s", info.Address))
		}
		if info.Decimals == 0 {
			return errors.New(fmt.Sprintf("invalid token decimals: %v", info.Decimals))
		}
	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_ZANO:
	case ChainType_OTHER:
	default:
		return errors.New(fmt.Sprintf("invalid chain type: %v", *chainType))
	}

	return nil
}
