package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetChainById(goctx context.Context, req *types.QueryGetChainById) (*types.QueryGetChainByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	token, found := k.GetChain(ctx, req.Id)
	if !found {
		return nil, types.ErrChainNotFound
	}

	return &types.QueryGetChainByIdResponse{Chain: token}, nil
}

func (k queryServer) GetChains(goctx context.Context, req *types.QueryGetChains) (*types.QueryGetChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	chains, page, err := k.GetChainsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, types.ErrChainNotFound
	}

	return &types.QueryGetChainsResponse{Chains: chains, Pagination: page}, nil
}
