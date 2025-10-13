package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) SetReferral(goCtx context.Context, msg *types.MsgSetReferral) (*types.MsgSetReferralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetReferral(ctx, msg.Referral.GetId())
	if found {
		return nil, errorsmod.Wrap(types.ErrAlreadyExists, "referral with this ID already exists")
	}

	m.AddReferral(ctx, msg.Referral)
	return &types.MsgSetReferralResponse{}, nil
}

func (m msgServer) RemoveReferral(goCtx context.Context, msg *types.MsgRemoveReferral) (*types.MsgRemoveReferralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetReferral(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(types.ErrReferralNotFound, "referral not found")
	}

	m.DeleteReferral(ctx, msg.Id)

	return &types.MsgRemoveReferralResponse{}, nil
}
