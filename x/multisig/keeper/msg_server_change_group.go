package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/v12/x/multisig/types"
)

func (k msgServer) ChangeGroup(goCtx context.Context, msg *types.MsgChangeGroup) (*types.MsgChangeGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := ensureMsgAuthZ([]sdk.Msg{msg}, msg.Group)
	if err != nil {
		return nil, err
	}

	group, found := k.GetGroup(ctx, msg.Group)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "group (%s) not found", msg.Group)
	}

	group.Members = msg.Members
	group.Threshold = msg.Threshold

	k.SetGroup(ctx, group)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeGroup,
			sdk.NewAttribute(types.AttributeKeyGroup, group.Account),
		),
	)

	return &types.MsgChangeGroupResponse{}, nil
}
