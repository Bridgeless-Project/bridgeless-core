package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k queryServer) GetStopListTxsById(goCtx context.Context, msg *types.QueryGetStopListTxById) (*types.QueryGetStopListTxByIdResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	transaction, ok := k.GetTxFromStopList(sdkCtx, types.TransactionId(&types.Transaction{
		DepositTxHash:  msg.TxHash,
		DepositChainId: msg.ChainId,
		DepositTxIndex: msg.TxNonce,
	}))

	if !ok {
		return nil, types.ErrTransactionNotFound
	}

	return &types.QueryGetStopListTxByIdResponse{Transaction: transaction}, nil
}

func (k queryServer) GetStopListTxs(goCtx context.Context, msg *types.QueryGetStopListTxs) (*types.QueryGetStopListTxsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)

	transactions, pageRes, err := k.GetTxsFromStopListWithPagination(sdkCtx, msg.Pagination)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to query stop list")
	}

	return &types.QueryGetStopListTxsResponse{Pagination: pageRes, Transactions: transactions}, nil
}
