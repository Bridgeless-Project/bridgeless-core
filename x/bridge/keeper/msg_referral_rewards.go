package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) SetReferralRewards(goCtx context.Context, msg *types.MsgSetReferralRewards) (*types.MsgSetReferralRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetReferralRewards(ctx, msg.Rewards.ReferralId, msg.Rewards.TokenId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrAlreadyExists, "Referral rewards with this Referral ID and Token ID already exists")
	}

	m.AddReferralRewards(ctx, msg.Rewards.ReferralId, msg.Rewards.TokenId, msg.Rewards)

	return &types.MsgSetReferralRewardsResponse{}, nil
}

func (m msgServer) RemoveReferralRewards(goCtx context.Context, msg *types.MsgRemoveReferralRewards) (*types.MsgRemoveReferralRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetReferralRewards(ctx, msg.ReferrerId, msg.TokenId)
	if found {
		return nil, errorsmod.Wrap(types.ErrAlreadyExists, "Referral rewards with this Referral ID and Token ID already exists")
	}

	m.DeleteReferralRewards(ctx, msg.ReferrerId, msg.TokenId)

	return &types.MsgRemoveReferralRewardsResponse{}, nil
}
