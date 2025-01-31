package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetParties(goCtx context.Context, msg *types.MsgSetParties) (*types.MsgSetPartiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.Keeper.GetParams(ctx)

	if msg.Creator != params.ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	params.Parties = msg.Parties
	m.SetParams(ctx, params)

	return &types.MsgSetPartiesResponse{}, nil
}
