package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
    "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/keeper"
    keepertest "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SwapKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
