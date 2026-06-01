package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) Transactions(goCtx context.Context, req *types.QueryTransactionsRequest) (*types.QueryTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	txs, pages, err := k.GetPaginatedTransactions(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTransactionsResponse{Transactions: txs, Pagination: pages}, nil
}

func (k queryServer) TransactionById(goCtx context.Context, req *types.QueryTransactionByIdRequest) (*types.QueryTransactionByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	tx, found := k.GetTransaction(ctx, types.TransactionId(&types.Transaction{
		DepositChainId: req.ChainId,
		DepositTxHash:  req.TxHash,
		DepositTxIndex: req.TxNonce,
	}))
	if !found {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	return &types.QueryTransactionByIdResponse{Transaction: tx}, nil
}

func (k queryServer) SystemWithdrawals(goCtx context.Context, req *types.QuerySystemWithdrawalsRequest) (*types.QuerySystemWithdrawalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	withdrawals, pages, err := k.GetPaginatedSystemTransactions(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySystemWithdrawalsResponse{Withdrawals: withdrawals, Pagination: pages}, nil
}

func (k queryServer) SystemWithdrawalById(goCtx context.Context, req *types.QuerySystemWithdrawalByIdRequest) (*types.QuerySystemWithdrawalByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	withdrawal, found := k.GetSystemTransaction(ctx, types.SystemTransactionId(&types.SystemWithdrawal{
		TxHash:  req.TxHash,
		TxIndex: req.TxIndex,
	}))
	if !found {
		return nil, status.Error(codes.NotFound, "system withdrawal not found")
	}

	return &types.QuerySystemWithdrawalByIdResponse{Withdrawal: withdrawal}, nil
}
