package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func validateChain(chain *Chain) error {
	if chain == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "nil chain")
	}
	if len(chain.BridgeAddress) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "empty bridge address")
	}

	switch chain.Type {
	case ChainType_EVM:
		if _, set := big.NewInt(0).SetString(chain.Id, 10); !set {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid chain id: %s", chain.Id)
		}
		if !common.IsHexAddress(chain.BridgeAddress) {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid bridge address: %s", chain.BridgeAddress)
		}
		if !common.IsHexAddress(chain.Operator) {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid operator address: %s", chain.Operator)
		}
	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_OTHER:
	case ChainType_ZANO:
	default:
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid chain type: %s", chain.Type)
	}

	return nil
}
