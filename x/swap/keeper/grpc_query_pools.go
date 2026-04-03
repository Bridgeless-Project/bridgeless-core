package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) AllPool(goCtx context.Context, req *types.QueryAllPools) (*types.QueryAllPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pools, page, err := k.GetPoolsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolsResponse{Pool: pools, Pagination: page}, nil
}

func (k queryServer) GetPoolByTokenId(goCtx context.Context, req *types.QueryGetPoolByTokenId) (*types.QueryGetPoolByTokenIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pool, found := k.GetPool(ctx, req.TokenId)
	if !found {
		return nil, status.Error(codes.NotFound, types.ErrPoolNotFound.Error())
	}

	return &types.QueryGetPoolByTokenIdResponse{Pool: pool}, nil
}
