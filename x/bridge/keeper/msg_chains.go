package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (m msgServer) InsertChain(goCtx context.Context, msg *types.MsgInsertChain) (*types.MsgInsertChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetChain(ctx, msg.Chain.Id)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "chain already exists")
	}

	m.SetChain(ctx, msg.Chain)

	return &types.MsgInsertChainResponse{}, nil
}

func (m msgServer) DeleteChain(goCtx context.Context, msg *types.MsgDeleteChain) (*types.MsgDeleteChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetChain(ctx, msg.ChainId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "chain not found")
	}

	m.RemoveChain(ctx, msg.ChainId)

	return &types.MsgDeleteChainResponse{}, nil
}
