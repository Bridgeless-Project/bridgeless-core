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

func TestPoolsQuery(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	k.SetPool(ctx, types.SwapPool{Address: "0xpool1", TokenId: "token-1"})
	k.SetPool(ctx, types.SwapPool{Address: "0xpool2", TokenId: "token-2"})

	res, err := qs.AllPool(wctx, &types.QueryAllPools{Pagination: &query.PageRequest{Limit: 10, CountTotal: true}})
	require.NoError(t, err)
	require.Len(t, res.Pool, 2)
	require.NotNil(t, res.Pagination)
	require.Equal(t, uint64(2), res.Pagination.Total)
}

func TestPoolByTokenIDQuery(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	pool := types.SwapPool{Address: "0xpool1", TokenId: "token-1"}
	k.SetPool(ctx, pool)

	res, err := qs.GetPoolByTokenId(wctx, &types.QueryGetPoolByTokenId{TokenId: pool.TokenId})
	require.NoError(t, err)
	require.Equal(t, pool, res.Pool)
}

func TestPoolByTokenIDQueryNotFound(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := qs.GetPoolByTokenId(wctx, &types.QueryGetPoolByTokenId{TokenId: "missing"})
	require.Error(t, err)
	require.Equal(t, codes.NotFound, status.Code(err))
}

func TestPoolsQueryInvalidRequest(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	qs := keeper.NewQueryServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)

	_, err := qs.AllPool(wctx, nil)
	require.Error(t, err)
	require.Equal(t, codes.InvalidArgument, status.Code(err))
}
