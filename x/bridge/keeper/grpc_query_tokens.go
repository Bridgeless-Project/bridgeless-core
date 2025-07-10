package keeper

import (
	"context"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetTokenById(goctx context.Context, req *types.QueryGetTokenById) (*types.QueryGetTokenByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	token, found := k.GetToken(ctx, req.Id)
	if !found {
		return nil, types.ErrTokenNotFound
	}

	return &types.QueryGetTokenByIdResponse{Token: token}, nil
}

func (k queryServer) GetTokens(goctx context.Context, req *types.QueryGetTokens) (*types.QueryGetTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	tokens, page, err := k.GetTokensWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetTokensResponse{Tokens: tokens, Pagination: page}, nil
}

func (k queryServer) GetTokenInfo(goctx context.Context, req *types.QueryGetTokenInfo) (*types.QueryGetTokenInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	tokenInfo, found := k.Keeper.GetTokenInfo(ctx, req.Chain, req.Address)
	if !found {
		return nil, types.ErrTokenInfoNotFound
	}

	return &types.QueryGetTokenInfoResponse{Info: tokenInfo}, nil
}
