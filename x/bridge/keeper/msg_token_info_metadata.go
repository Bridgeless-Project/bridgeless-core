package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) SetTokenInfoMetadata(goCtx context.Context, msg *types.MsgSetTokenInfoMetadata) (*types.MsgSetTokenInfoMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if _, found := m.Keeper.GetTokenInfo(ctx, msg.Metadata.ChainId, msg.Metadata.Address); !found {
		return nil, types.ErrTokenInfoNotFound
	}

	m.Keeper.SetTokenInfoMetadata(ctx, msg.Metadata)

	return &types.MsgSetTokenInfoMetadataResponse{}, nil
}

func (m msgServer) RemoveTokenInfoMetadata(goCtx context.Context, msg *types.MsgRemoveTokenInfoMetadata) (*types.MsgRemoveTokenInfoMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if _, found := m.Keeper.GetTokenInfoMetadata(ctx, msg.ChainId, msg.Address); !found {
		return nil, types.ErrTokenInfoMetadataNotFound
	}

	m.Keeper.RemoveTokenInfoMetadata(ctx, msg.ChainId, msg.Address)

	return &types.MsgRemoveTokenInfoMetadataResponse{}, nil
}
