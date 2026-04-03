package keeper_test

import (
	"testing"

	keepertest "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupPoolMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, sdk.Context, string, string) {
	k, ctx := keepertest.SwapKeeper(t)
	admin := sdk.AccAddress(make([]byte, 20)).String()
	nonAdmin := sdk.AccAddress(bytesOf(20, 1)).String()
	k.SetParams(ctx, types.NewParams(admin))

	return keeper.NewMsgServerImpl(*k), *k, ctx, admin, nonAdmin
}

func bytesOf(size int, value byte) []byte {
	b := make([]byte, size)
	for i := range b {
		b[i] = value
	}
	return b
}

func TestMsgAddPool(t *testing.T) {
	ms, k, ctx, admin, _ := setupPoolMsgServer(t)

	_, err := ms.AddPool(sdk.WrapSDKContext(ctx), &types.MsgAddPool{
		Creator: admin,
		Pool:    &types.SwapPool{Address: "0xpool1", TokenId: "token-1"},
	})
	require.NoError(t, err)

	_, found := k.GetPool(ctx, "token-1")
	require.True(t, found)
}

func TestMsgAddPoolNonAdmin(t *testing.T) {
	ms, _, ctx, _, nonAdmin := setupPoolMsgServer(t)

	_, err := ms.AddPool(sdk.WrapSDKContext(ctx), &types.MsgAddPool{
		Creator: nonAdmin,
		Pool:    &types.SwapPool{Address: "0xpool1", TokenId: "token-1"},
	})
	require.Error(t, err)
}

func TestMsgUpdatePool(t *testing.T) {
	ms, k, ctx, admin, _ := setupPoolMsgServer(t)
	k.SetPool(ctx, types.SwapPool{Address: "0xpool1", TokenId: "token-1"})

	_, err := ms.UpdatePool(sdk.WrapSDKContext(ctx), &types.MsgUpdatePool{
		Creator: admin,
		Pool:    &types.SwapPool{Address: "0xpool2", TokenId: "token-1"},
	})
	require.NoError(t, err)

	pool, found := k.GetPool(ctx, "token-1")
	require.True(t, found)
	require.Equal(t, "0xpool2", pool.Address)
}

func TestMsgUpdatePoolNonAdmin(t *testing.T) {
	ms, k, ctx, admin, nonAdmin := setupPoolMsgServer(t)
	k.SetPool(ctx, types.SwapPool{Address: "0xpool1", TokenId: "token-1"})

	_, err := ms.UpdatePool(sdk.WrapSDKContext(ctx), &types.MsgUpdatePool{
		Creator: nonAdmin,
		Pool:    &types.SwapPool{Address: "0xpool2", TokenId: "token-1"},
	})
	require.Error(t, err)

	pool, found := k.GetPool(ctx, "token-1")
	require.True(t, found)
	require.Equal(t, "0xpool1", pool.Address)
	_ = admin
}

func TestMsgUpdatePoolNotFound(t *testing.T) {
	ms, _, ctx, admin, _ := setupPoolMsgServer(t)

	_, err := ms.UpdatePool(sdk.WrapSDKContext(ctx), &types.MsgUpdatePool{
		Creator: admin,
		Pool:    &types.SwapPool{Address: "0xpool2", TokenId: "missing"},
	})
	require.Error(t, err)
}

func TestMsgRemovePool(t *testing.T) {
	ms, k, ctx, admin, _ := setupPoolMsgServer(t)
	k.SetPool(ctx, types.SwapPool{Address: "0xpool1", TokenId: "token-1"})

	_, err := ms.RemovePool(sdk.WrapSDKContext(ctx), &types.MsgRemovePool{
		Creator: admin,
		TokenId: "token-1",
	})
	require.NoError(t, err)

	_, found := k.GetPool(ctx, "token-1")
	require.False(t, found)
}

func TestMsgRemovePoolNonAdmin(t *testing.T) {
	ms, k, ctx, admin, nonAdmin := setupPoolMsgServer(t)
	k.SetPool(ctx, types.SwapPool{Address: "0xpool1", TokenId: "token-1"})

	_, err := ms.RemovePool(sdk.WrapSDKContext(ctx), &types.MsgRemovePool{
		Creator: nonAdmin,
		TokenId: "token-1",
	})
	require.Error(t, err)

	_, found := k.GetPool(ctx, "token-1")
	require.True(t, found)
	_ = admin
}

func TestMsgRemovePoolNotFound(t *testing.T) {
	ms, _, ctx, admin, _ := setupPoolMsgServer(t)

	_, err := ms.RemovePool(sdk.WrapSDKContext(ctx), &types.MsgRemovePool{
		Creator: admin,
		TokenId: "missing",
	})
	require.Error(t, err)
}

func TestMsgPoolNilMessage(t *testing.T) {
	ms, _, ctx, _, _ := setupPoolMsgServer(t)

	_, err := ms.AddPool(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)

	_, err = ms.UpdatePool(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)

	_, err = ms.RemovePool(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
}
