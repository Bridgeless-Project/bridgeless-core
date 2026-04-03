package keeper_test

import (
	"testing"

	testkeeper "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSwapsQuery(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	swap1 := sampleSwap("chain-a", "0xhash1", 1)
	swap2 := sampleSwap("chain-b", "0xhash2", 2)
	k.SetSwap(ctx, swap1)
	k.SetSwap(ctx, swap2)

	res, err := qs.AllSwaps(wctx, &types.QueryAllSwaps{Pagination: &query.PageRequest{Limit: 10, CountTotal: true}})
	require.NoError(t, err)
	require.Len(t, res.Swap, 2)
	require.NotNil(t, res.Pagination)
	require.Equal(t, uint64(2), res.Pagination.Total)
}

func TestSwapByIDQuery(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	swap := sampleSwap("chain-a", "0xhash1", 1)
	k.SetSwap(ctx, swap)

	res, err := qs.GetSwapById(wctx, &types.QueryGetSwapById{
		ChainId: swap.Tx.DepositChainId,
		TxHash:  swap.Tx.DepositTxHash,
		TxNonce: swap.Tx.DepositTxIndex,
	})
	require.NoError(t, err)
	require.Equal(t, swap, res.Swap)
}

func TestSwapByIDQueryNotFound(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := qs.GetSwapById(wctx, &types.QueryGetSwapById{ChainId: "chain-a", TxHash: "missing", TxNonce: 1})
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestSwapsQueryInvalidRequest(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := qs.AllSwaps(wctx, nil)
	require.Error(t, err)
	require.Equal(t, codes.InvalidArgument, status.Code(err))
}
