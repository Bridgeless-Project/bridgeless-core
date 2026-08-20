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

	if _, found := m.Keeper.GetTokenInfo(ctx, msg.ChainId, msg.Address); !found {
		return nil, types.ErrTokenInfoNotFound
	}

	m.Keeper.SetTokenInfoMetadata(ctx, msg.ChainId, msg.Address, msg.Metadata)

	return &types.MsgSetTokenInfoMetadataResponse{}, nil
}
