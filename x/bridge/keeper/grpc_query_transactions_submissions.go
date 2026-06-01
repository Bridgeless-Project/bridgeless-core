package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k queryServer) GetSystemTxsSubmissions(goCtx context.Context, req *types.QueryGetSystemTxsSubmissions) (
	*types.QueryGetSystemTxsSubmissionsResponse,
	error,
) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	txsSubmissions, pages, err := k.GetPaginatedSystemTransactionsSubmissions(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetSystemTxsSubmissionsResponse{TxsSubmissions: txsSubmissions, Pagination: pages}, nil
}

func (k queryServer) GetSystemTxSubmissionsByHash(goCtx context.Context, req *types.QueryGetSystemTxSubmissionsByHash) (
	*types.QueryGetSystemTxSubmissionsByHashResponse,
	error,
) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	txSubmissions, found := k.GetSystemTransactionSubmissions(ctx, req.TxHash)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSystemTxSubmissionsByHashResponse{TxSubmissions: txSubmissions}, nil

}
