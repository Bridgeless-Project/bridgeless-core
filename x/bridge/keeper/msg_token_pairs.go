package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m msgServer) AddTokenInfo(goCtx context.Context, msg *types.MsgAddTokenInfo) (*types.MsgAddTokenInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.Info.TokenId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	for _, info := range token.Info {
		if info.ChainId == msg.Info.ChainId {
			return nil, errorsmod.Wrap(sdkerrors.ErrConflict, "token info already exists")
		}
	}

	token.Info = append(token.Info, msg.Info)
	m.SetToken(ctx, token)
	m.SetTokenInfo(ctx, msg.Info)
	m.SetTokenPairs(ctx, msg.Info, token.Info...) // mapping src chain -> dest chain
	for _, info := range token.Info {
		m.SetTokenPairs(ctx, info, msg.Info) // reverse mapping dest chain -> src chain
	}

	return &types.MsgAddTokenInfoResponse{}, nil
}

func (m msgServer) RemoveTokenInfo(goCtx context.Context, msg *types.MsgRemoveTokenInfo) (*types.MsgRemoveTokenInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	var idx = -1
	for i, info := range token.Info {
		if info.ChainId == msg.ChainId {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "token info not found")
	}

	old := token.Info[idx]
	m.Keeper.RemoveTokenInfo(ctx, old.ChainId, old.Address)
	m.RemoveTokenPairs(ctx, old, token.Info...) // mapping src chain -> dest chain
	for _, info := range token.Info {
		m.RemoveTokenPairs(ctx, info, old) // reverse mapping dest chain -> src chain
	}

	token.Info = append(token.Info[:idx], token.Info[idx+1:]...)
	m.SetToken(ctx, token)

	return &types.MsgRemoveTokenInfoResponse{}, nil
}
