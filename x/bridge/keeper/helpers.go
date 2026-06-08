package keeper

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func ConstructSystemTxHash(amount *big.Int, tokenAddress, receiver []byte) string {
	return hexutil.Encode(crypto.Keccak256(amount.Bytes(), tokenAddress, receiver))
}
