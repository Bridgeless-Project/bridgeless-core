package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) AllSwaps(goCtx context.Context, req *types.QueryAllSwaps) (*types.QueryAllSwapsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	swaps, page, err := k.GetSwapsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSwapsResponse{Swap: swaps, Pagination: page}, nil
}

func (k queryServer) GetSwapById(goCtx context.Context, req *types.QueryGetSwapById) (*types.QueryGetSwapByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	swap, found := k.GetSwap(ctx, req.TxHash, req.TxNonce, req.ChainId)
	if !found {
		return nil, status.Error(codes.NotFound, types.ErrSwapNotFound.Error())
	}

	return &types.QueryGetSwapByIdResponse{Swap: swap}, nil
}
