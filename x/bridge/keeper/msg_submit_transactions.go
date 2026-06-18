package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) SubmitTransactions(goCtx context.Context, msg *types.MsgSubmitTransactions) (*types.MsgSubmitTransactionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsParty(ctx, msg.Submitter) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	for _, tx := range msg.Transactions {
		chain, found := m.GetChain(ctx, tx.DepositChainId)
		if !found {
			return nil, types.ErrSourceChainNotSupported
		}
		if _, found = m.GetChain(ctx, tx.WithdrawalChainId); !found {
			return nil, types.ErrDestinationChainNotSupported
		}

		// Custom validation of transaction for certain chain type
		err := types.ValidateChainTransaction(&tx, chain.Type)
		if err != nil {
			return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
		}

		if err = m.SubmitTx(ctx, &tx, msg.Submitter); err != nil {
			return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
		}
	}

	return &types.MsgSubmitTransactionsResponse{}, nil
}
