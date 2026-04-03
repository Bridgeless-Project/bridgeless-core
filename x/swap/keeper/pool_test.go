package keeper_test

import (
	"testing"

	testkeeper "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestSetGetPool(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pool := types.SwapPool{
		Address: "0xpool1",
		TokenId: "token-1",
	}

	k.SetPool(ctx, pool)

	got, found := k.GetPool(ctx, pool.TokenId)
	require.True(t, found)
	require.Equal(t, pool, got)
}

func TestGetPoolNotFound(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)

	_, found := k.GetPool(ctx, "missing")
	require.False(t, found)
}

func TestGetAllPools(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pools := []types.SwapPool{
		{Address: "0xpool1", TokenId: "token-1"},
		{Address: "0xpool2", TokenId: "token-2"},
	}

	for _, pool := range pools {
		k.SetPool(ctx, pool)
	}

	require.ElementsMatch(t, pools, k.GetAllPools(ctx))
}

func TestGetPoolsWithPagination(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pools := []types.SwapPool{
		{Address: "0xpool1", TokenId: "token-1"},
		{Address: "0xpool2", TokenId: "token-2"},
	}

	for _, pool := range pools {
		k.SetPool(ctx, pool)
	}

	got, page, err := k.GetPoolsWithPagination(ctx, &query.PageRequest{Limit: 1, CountTotal: true})
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.NotNil(t, page)
	require.Equal(t, uint64(len(pools)), page.Total)
}

func TestAddPool(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pool := types.SwapPool{
		Address: "0xpool1",
		TokenId: "token-1",
	}

	err := k.AddPool(ctx, pool)
	require.NoError(t, err)

	got, found := k.GetPool(ctx, pool.TokenId)
	require.True(t, found)
	require.Equal(t, pool, got)
}

func TestAddPoolConflict(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pool := types.SwapPool{
		Address: "0xpool1",
		TokenId: "token-1",
	}

	k.SetPool(ctx, pool)

	err := k.AddPool(ctx, pool)
	require.Error(t, err)
}

func TestUpdatePool(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pool := types.SwapPool{
		Address: "0xpool1",
		TokenId: "token-1",
	}
	updatedPool := types.SwapPool{
		Address: "0xpool2",
		TokenId: "token-1",
	}

	k.SetPool(ctx, pool)

	err := k.UpdatePool(ctx, updatedPool)
	require.NoError(t, err)

	got, found := k.GetPool(ctx, updatedPool.TokenId)
	require.True(t, found)
	require.Equal(t, updatedPool, got)
}

func TestUpdatePoolNotFound(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)

	err := k.UpdatePool(ctx, types.SwapPool{Address: "0xpool1", TokenId: "missing"})
	require.Error(t, err)
}

func TestRemovePool(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	pool := types.SwapPool{
		Address: "0xpool1",
		TokenId: "token-1",
	}

	k.SetPool(ctx, pool)

	err := k.RemovePool(ctx, pool.TokenId)
	require.NoError(t, err)

	_, found := k.GetPool(ctx, pool.TokenId)
	require.False(t, found)
}

func TestRemovePoolNotFound(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)

	err := k.RemovePool(ctx, "missing")
	require.Error(t, err)
}
