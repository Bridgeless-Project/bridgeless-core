package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetStopListTxsById(goCtx context.Context, req *types.QueryGetStopListTxById) (*types.QueryGetStopListTxByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	transaction, ok := k.GetTxFromStopList(sdkCtx, types.TransactionId(&types.Transaction{
		DepositTxHash:  req.TxHash,
		DepositChainId: req.ChainId,
		DepositTxIndex: req.TxNonce,
	}))

	if !ok {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	return &types.QueryGetStopListTxByIdResponse{Transaction: transaction}, nil
}

func (k queryServer) GetStopListTxs(goCtx context.Context, req *types.QueryGetStopListTxs) (*types.QueryGetStopListTxsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	transactions, pageRes, err := k.GetTxsFromStopListWithPagination(sdkCtx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetStopListTxsResponse{Pagination: pageRes, Transactions: transactions}, nil
}
