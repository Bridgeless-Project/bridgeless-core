package keeper

import (
	"context"
	"errors"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) ProcessSystemWithdrawal(goCtx context.Context, msg *types.MsgProcessSystemWithdrawal) (*types.MsgProcessSystemWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsEpochParty(ctx, msg.Creator, msg.EpochId) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	for _, tx := range msg.Withdrawal {
		if tx.EpochId != msg.EpochId {
			return nil, errors.New("epoch mismatch")
		}
		if err := m.SystemWithdrawal(ctx, &tx, msg.Creator); err != nil {
			return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
		}
	}

	return &types.MsgProcessSystemWithdrawalResponse{}, nil

}
