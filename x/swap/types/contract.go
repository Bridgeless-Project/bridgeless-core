package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SwapperWithdrawParams struct {
	Token      common.Address
	Amount     *big.Int
	TxHash     common.Hash
	TxNonce    *big.Int
	IsWrapped  bool
	Signatures [][]byte
}

type SwapperSwapParams struct {
	AmountIn                 *big.Int
	MinDestinationAmount     *big.Int
	SwapDeadline             *big.Int
	Path                     []common.Address
	IsDestinationTokenNative bool
}

type SwapperDepositParams struct {
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId *big.Int
}

type AmountsListResponse struct {
	Amounts []*big.Int `json:"amounts"`
}
