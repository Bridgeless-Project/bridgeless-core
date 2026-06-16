package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) DistributeFees(goCtx context.Context, msg *types.MsgDistributeFees) (*types.MsgDistributeFeesResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidDataType, "message cannot be nil")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := m.Keeper.GetParams(ctx)
	if params.ModuleAdmin != msg.Creator {
		return nil, errorsmod.Wrapf(types.ErrPermissionDenied, "only module admin can start a fee distribution")
	}

	if _, found := m.GetEpoch(ctx, msg.EpochId); !found {
		return nil, errorsmod.Wrapf(types.ErrInvalidEpochID, "epoch %d not found", msg.EpochId)
	}
	emitDistributeFees(ctx, msg.EpochId)

	return &types.MsgDistributeFeesResponse{}, nil
}

func (m msgServer) ProcessSystemWithdrawal(goCtx context.Context, msg *types.MsgProcessSystemWithdrawal) (*types.MsgProcessSystemWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsParty(ctx, msg.Creator) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	for _, tx := range msg.Withdrawal {
		if err := m.SystemWithdrawal(ctx, &tx, msg.Creator); err != nil {
			return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
		}
	}

	return &types.MsgProcessSystemWithdrawalResponse{}, nil

}
