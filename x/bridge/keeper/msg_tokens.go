package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) InsertToken(goCtx context.Context, msg *types.MsgInsertToken) (*types.MsgInsertTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetToken(ctx, msg.Token.Id)
	if found {
		return nil, errorsmod.Wrap(sdkerrors.ErrConflict, "token already exists")
	}

	m.SetToken(ctx, msg.Token)
	for _, info := range msg.Token.Info {
		m.SetTokenInfo(ctx, info)
		m.SetTokenPairs(ctx, info, msg.Token.Info...)
	}

	return &types.MsgInsertTokenResponse{}, nil
}

func (m msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	token.Metadata = msg.Metadata
	m.SetToken(ctx, token)

	return &types.MsgUpdateTokenResponse{}, nil
}

func (m msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	m.RemoveToken(ctx, msg.TokenId)
	for _, info := range token.Info {
		m.Keeper.RemoveTokenInfo(ctx, info.ChainId, info.Address)
		m.RemoveTokenPairs(ctx, info, token.Info...)
	}

	return &types.MsgDeleteTokenResponse{}, nil
}
