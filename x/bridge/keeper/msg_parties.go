package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) SetParties(goCtx context.Context, msg *types.MsgSetParties) (*types.MsgSetPartiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.Keeper.GetParams(ctx)

	if msg.Creator != params.ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	params.Parties = msg.Parties
	m.SetParams(ctx, params)

	return &types.MsgSetPartiesResponse{}, nil
}
