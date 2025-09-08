package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetReferralById(goctx context.Context, req *types.QueryGetReferralById) (*types.QueryGetReferralByIdResponse, error) {
	if req == nil || req.ReferralId == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	referral, found := k.GetReferral(ctx, req.ReferralId)
	if !found {
		return nil, types.ErrReferralNotFound
	}

	return &types.QueryGetReferralByIdResponse{Referral: referral}, nil

}

func (k queryServer) GetQueryGetReferrals(goctx context.Context, req *types.QueryGetReferrals) (*types.QueryGetReferralsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goctx)
	referrals, page, err := k.GetReferralsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, types.ErrReferralNotFound
	}

	return &types.QueryGetReferralsResponse{Referrals: referrals, Pagination: page}, nil

}
