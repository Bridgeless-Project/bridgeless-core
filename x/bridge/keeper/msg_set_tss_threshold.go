package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetTssThreshold(goCtx context.Context, msg *types.MsgSetTssThreshold) (*types.MsgSetTssThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := m.Keeper.GetParams(ctx)

	if msg.Creator != params.ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	params.TssThreshold = msg.Threshold
	m.Keeper.SetParams(ctx, params)

	return &types.MsgSetTssThresholdResponse{}, nil
}
