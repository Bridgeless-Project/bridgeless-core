package keeper_test

import (
	"testing"

	testkeeper "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func sampleSwap(chainID, txHash string, txNonce uint64) types.SwapTransaction {
	return types.SwapTransaction{
		Tx: bridgetypes.Transaction{
			DepositChainId: chainID,
			DepositTxHash:  txHash,
			DepositTxIndex: txNonce,
		},
		FinalReceiver:      "receiver",
		FinalAmount:        "100",
		FinalDepositTxHash: "0xfinal",
	}
}

func TestSetGetSwap(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	swap := sampleSwap("chain-a", "0xhash1", 1)

	k.SetSwap(ctx, swap)

	got, found := k.GetSwap(ctx, swap.Tx.DepositTxHash, swap.Tx.DepositTxIndex, swap.Tx.DepositChainId)
	require.True(t, found)
	require.Equal(t, swap, got)
}

func TestGetSwapNotFound(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)

	_, found := k.GetSwap(ctx, "missing", 1, "chain-a")
	require.False(t, found)
}

func TestGetAllSwaps(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	swaps := []types.SwapTransaction{
		sampleSwap("chain-a", "0xhash1", 1),
		sampleSwap("chain-b", "0xhash2", 2),
	}

	for _, swap := range swaps {
		k.SetSwap(ctx, swap)
	}

	require.ElementsMatch(t, swaps, k.GetAllSwaps(ctx))
}

func TestGetSwapsWithPagination(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	swaps := []types.SwapTransaction{
		sampleSwap("chain-a", "0xhash1", 1),
		sampleSwap("chain-b", "0xhash2", 2),
	}

	for _, swap := range swaps {
		k.SetSwap(ctx, swap)
	}

	got, page, err := k.GetSwapsWithPagination(ctx, &query.PageRequest{Limit: 1, CountTotal: true})
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.NotNil(t, page)
	require.Equal(t, uint64(len(swaps)), page.Total)
}
