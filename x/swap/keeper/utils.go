package keeper

import (
	"math/big"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func parseUintString(value string) (*big.Int, error) {
	parsed, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid big int: %s", value)
	}

	if parsed.Sign() < 0 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "big int cannot be negative: %s", value)
	}

	return parsed, nil
}

func isZeroAddress(address string) bool {
	return common.HexToAddress(address) == (common.Address{})
}

func getChainId(ctx sdk.Context) string {
	// the chian-id returns something like cosmos_1234-1
	prefixAndChain := strings.Split(ctx.ChainID(), "_") // split to [cosmos, 1234-1]
	if len(prefixAndChain) != 2 {
		return ctx.ChainID()
	}

	evmChainIDWithSuffix := strings.Split(prefixAndChain[1], "-") // split to [1234, 1]
	if len(evmChainIDWithSuffix) != 2 {
		return prefixAndChain[1]
	}

	return evmChainIDWithSuffix[0]
}

func txHashToBytes32(txHash string) [32]byte {
	var res [32]byte
	hashBytes, err := hexutil.Decode(txHash)
	if err != nil || len(hashBytes) != 32 {
		bytes := crypto.Keccak256(([]byte)(txHash))
		copy(res[:], bytes)
		return res
	}

	copy(res[:], hashBytes)
	return res
}
