package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetReferralRewardsByToken(goctx context.Context, req *types.QueryGetReferralRewardByToken) (*types.QueryGetReferralRewardByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goctx)
	rewards, found := k.GetReferralRewards(ctx, req.ReferralId, req.TokenId)
	if !found {
		return nil, types.ErrReferralRewardsNotFound
	}

	return &types.QueryGetReferralRewardByIdResponse{Rewards: rewards}, nil
}

func (k queryServer) GetReferralsRewardsById(goctx context.Context, req *types.QueryGetReferralRewardsById) (*types.QueryGetReferralRewardsByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	rewards, page, err := k.GetReferralRewardsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetReferralRewardsByIdResponse{Rewards: rewards, Pagination: page}, nil
}
