package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k queryServer) GetAllEpoches(ctx context.Context, query *types.QueryAllEpoches) (*types.QueryAllEpochesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	epoches, page, err := k.GetEpochesWithPagination(sdkCtx, query.Pagination)
	if err != nil {
		return nil, types.ErrEpochesNotFound
	}

	return &types.QueryAllEpochesResponse{
		Epoches:    epoches,
		Pagination: page,
	}, nil
}

func (k queryServer) GetEpochById(ctx context.Context, id *types.QueryEpochById) (*types.QueryEpochByIdResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	epoch, err := k.GetEpochById(sdkCtx, id)
	if err != nil {
		return nil, types.ErrEpochNotFound
	}

	return epoch, nil
}
