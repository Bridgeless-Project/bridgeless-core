package keeper

import (
	"math/big"
	"testing"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func Test_GetCommissionAmount(t *testing.T) {

	type tc struct {
		withdrawalAmount *big.Int
		commissionRate   string
		expected         *big.Int
	}

	testCases := map[string]tc{
		"should get commission amount for integer rate": {
			withdrawalAmount: big.NewInt(1000),
			commissionRate:   "1",
			expected:         big.NewInt(1000),
		},

		"should get commission amount for float rate": {
			withdrawalAmount: big.NewInt(1000_000_000),
			commissionRate:   "0.5",
			expected:         big.NewInt(500_000_000),
		},
		"should get commission amount for float rate with many decimals": {
			withdrawalAmount: big.NewInt(1000_000_000),
			commissionRate:   "0.532256666",
			expected:         big.NewInt(532256666),
		},
		"should make zero commission as precision is too small": {
			withdrawalAmount: big.NewInt(100),
			commissionRate:   "0.0000000000004",
			expected:         big.NewInt(0),
		},
		"should make 50% commission": {
			withdrawalAmount: big.NewInt(100),
			commissionRate:   "0.5",
			expected:         big.NewInt(50),
		},
		"should make 100% commission": {
			withdrawalAmount: big.NewInt(100),
			commissionRate:   "1",
			expected:         big.NewInt(100),
		},
		"should make minimal commission with precision 5": {
			withdrawalAmount: big.NewInt(100_000_000),
			commissionRate:   "0.00001",
			expected:         big.NewInt(1000),
		},
		"should make zero commission amount": {
			withdrawalAmount: big.NewInt(100),
			commissionRate:   "0",
			expected:         big.NewInt(0),
		},
	}

	for name, tCase := range testCases {
		t.Run(name, func(t *testing.T) {
			result, _ := types.GetCommissionAmount(tCase.withdrawalAmount, tCase.commissionRate)
			require.Equal(t, tCase.expected, result)
		})
	}

}
