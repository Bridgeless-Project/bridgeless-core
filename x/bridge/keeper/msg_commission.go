package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) SetCommission(goCtx context.Context, msg *types.MsgSetCommission) (*types.MsgSetCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	m.Keeper.SetCommission(ctx, types.Commission{
		TokenId: msg.TokenId,
		Amount:  msg.Amount,
	})

	return &types.MsgSetCommissionResponse{}, nil
}

func (m msgServer) UpdateCommission(goCtx context.Context, msg *types.MsgUpdateCommission) (*types.MsgUpdateCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	commission, found := m.Keeper.GetCommission(ctx, msg.TokenId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrCommissionNotFound, "commission with this TokenID is not found")
	}

	commission.Amount = msg.Amount
	m.Keeper.SetCommission(ctx, commission)

	return &types.MsgUpdateCommissionResponse{}, nil
}

func (m msgServer) RemoveCommission(goCtx context.Context, msg *types.MsgRemoveCommission) (*types.MsgRemoveCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	commission, found := m.Keeper.GetCommission(ctx, msg.TokenId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrCommissionNotFound, "commission with this TokenID is not found")
	}

	m.Keeper.RemoveCommission(ctx, commission)

	return &types.MsgRemoveCommissionResponse{}, nil
}
