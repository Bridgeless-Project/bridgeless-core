package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SwapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
