package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetEpochTransactions(goctx context.Context, req *types.QueryGetEpochTransactions) (*types.QueryGetEpochTransactionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k queryServer) GetEpochById(goctx context.Context, req *types.QueryGetEpoch) (*types.QueryGetEpochByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goctx)

	epoch, found := k.GetEpoch(ctx, req.EpochId)
	if !found {
		return nil, types.ErrEpochNotFound
	}

	return &types.QueryGetEpochByIdResponse{Epoch: epoch}, nil

}
