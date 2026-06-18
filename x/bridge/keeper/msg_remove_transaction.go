package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) RemoveTransaction(goCtx context.Context, msg *types.MsgRemoveTransaction) (*types.MsgRemoveTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParams(ctx)
	if params.ModuleAdmin != msg.Creator {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "creator isn`t the module admin")
	}

	if err := m.DeleteTx(ctx, msg.DepositTxHash, msg.DepositTxIndex, msg.DepositChainId); err != nil {
		return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
	}

	return &types.MsgRemoveTransactionResponse{}, nil
}
