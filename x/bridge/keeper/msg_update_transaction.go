package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) UpdateTransaction(goCtx context.Context, msg *types.MsgUpdateTransaction) (*types.MsgUpdateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.isRelayerAccount(msg.Submitter, m.GetParams(ctx).RelayerAccounts) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	if err := m.UpdateTx(ctx, &msg.Transaction); err != nil {
		return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
	}

	return &types.MsgUpdateTransactionResponse{}, nil
}

func (m msgServer) isRelayerAccount(account string, relayerAccounts []string) bool {
	for _, acc := range relayerAccounts {
		if account == acc {
			return true
		}
	}

	return false
}
