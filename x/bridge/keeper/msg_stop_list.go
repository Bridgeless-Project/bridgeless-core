package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) AddTxToStopList(goCtx context.Context, msg *types.MsgAddTxToStopList) (*types.MsgAddTxToStopListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.Keeper.GetParams(ctx)

	if msg.Creator != params.ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if _, ok := m.GetTxFromStopList(ctx, types.TransactionId(msg.Transaction)); ok {
		return nil, errorsmod.Wrap(types.ErrAlreadyExists, "transaction already exists")
	}

	m.SetTxToStopList(ctx, *msg.Transaction)

	return &types.MsgAddTxToStopListResponse{}, nil
}

func (m msgServer) RemoveTxFromStopList(goCtx context.Context, msg *types.MsgRemoveTxFromStopList) (*types.MsgRemoveTxFromStopListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.Keeper.GetParams(ctx)
	tx := &types.Transaction{
		DepositTxHash:  msg.TxHash,
		DepositChainId: msg.ChainId,
		DepositTxIndex: msg.TxNonce,
	}

	if msg.Creator != params.ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if _, ok := m.GetTxFromStopList(ctx, types.TransactionId(tx)); !ok {
		return nil, types.ErrTransactionNotFound
	}

	m.DeleteTxFromStopList(ctx, types.TransactionId(tx))

	return &types.MsgRemoveTxFromStopListResponse{}, nil

}
