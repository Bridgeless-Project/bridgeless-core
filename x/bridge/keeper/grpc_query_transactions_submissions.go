package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetTxsSubmissions(goCtx context.Context, req *types.QueryGetTxsSubmissions) (*types.QueryGetTxsSubmissionsResponse,
	error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	txsSubmissions, pages, err := k.GetPaginatedTransactionsSubmissions(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetTxsSubmissionsResponse{TxsSubmissions: txsSubmissions, Pagination: pages}, nil
}

func (k queryServer) GetTxSubmissionsByHash(goCtx context.Context, req *types.QueryGetTxSubmissionsByHash) (*types.QueryGetTxSubmissionsByHashResponse,
	error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	txSubmissions, found := k.GetTransactionSubmissions(ctx, req.TxHash)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTxSubmissionsByHashResponse{TxSubmissions: txSubmissions}, nil

}
