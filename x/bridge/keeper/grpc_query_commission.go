package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetCommissionByToken(goctx context.Context, req *types.QueryGetCommissionByToken) (*types.QueryGetCommissionByTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goctx)
	commission, found := k.GetCommission(ctx, req.TokenId)
	if !found {
		return nil, types.ErrCommissionNotFound
	}

	return &types.QueryGetCommissionByTokenResponse{Commission: commission}, nil
}

func (k queryServer) GetCommissions(goctx context.Context, req *types.QueryGetCommissions) (*types.QueryGetCommissionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	commissions, pages, err := k.GetCommissionsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetCommissionsResponse{Commissions: commissions, Pagination: pages}, nil
}
