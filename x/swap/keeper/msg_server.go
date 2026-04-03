package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

func (m msgServer) AddPool(goCtx context.Context, msg *types.MsgAddPool) (*types.MsgAddPoolResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	if msg.Pool == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "pool cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if err := m.Keeper.AddPool(ctx, *msg.Pool); err != nil {
		return nil, err
	}

	return &types.MsgAddPoolResponse{}, nil
}

func (m msgServer) RemovePool(goCtx context.Context, msg *types.MsgRemovePool) (*types.MsgRemovePoolResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if err := m.Keeper.RemovePool(ctx, msg.TokenId); err != nil {
		return nil, err
	}

	return &types.MsgRemovePoolResponse{}, nil
}

func (m msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	if msg.Pool == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "pool cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if err := m.Keeper.UpdatePool(ctx, *msg.Pool); err != nil {
		return nil, err
	}

	return &types.MsgUpdatePoolResponse{}, nil
}

func (m msgServer) SubmitSwapTx(ctx context.Context, tx *types.MsgSubmitSwapTx) (*types.MsgSubmitSwapTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
