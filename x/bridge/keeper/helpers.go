package keeper

import (
	"math/big"

	"github.com/Zano-Execution-Layer/go-ethereum/common/hexutil"
	"github.com/Zano-Execution-Layer/go-ethereum/crypto"
)

func ConstructSystemTxHash(amount *big.Int, tokenAddress, receiver []byte) string {
	return hexutil.Encode(crypto.Keccak256(amount.Bytes(), tokenAddress, receiver))
}

func TransformAmount(amount *big.Int, currentDecimals uint64, targetDecimals uint64) *big.Int {
	result, _ := new(big.Int).SetString(amount.String(), 10)

	if currentDecimals == targetDecimals {
		return result
	}

	if currentDecimals < targetDecimals {
		for i := uint64(0); i < targetDecimals-currentDecimals; i++ {
			result.Mul(result, new(big.Int).SetInt64(10))
		}
	} else {
		for i := uint64(0); i < currentDecimals-targetDecimals; i++ {
			result.Div(result, new(big.Int).SetInt64(10))
		}
	}

	return result
}
