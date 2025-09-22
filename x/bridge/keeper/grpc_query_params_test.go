package keeper_test

import (
	"testing"

	testkeeper "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	kp, ctx := testkeeper.BridgeKeeper(t)
	qs := keeper.NewQueryServerImpl(*kp)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	kp.SetParams(ctx, params)

	response, err := qs.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
