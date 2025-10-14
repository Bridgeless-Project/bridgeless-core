package keeper

import (
	"context"
	"math/big"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) UpdateTransaction(goCtx context.Context, msg *types.MsgUpdateTransaction) (*types.MsgUpdateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsParty(ctx, msg.Submitter) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	if err := m.UpdateTx(ctx, &msg.Transaction, msg.Submitter); err != nil {
		return nil, errorsmod.Wrap(types.InvalidTransaction, err.Error())
	}

	return &types.MsgUpdateTransactionResponse{}, nil
}

func emitUpdateTransactionEvent(sdkCtx sdk.Context, transaction types.Transaction) {
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(types.EventType_TRANSACTION_UPDATED.String(),
		sdk.NewAttribute(types.AttributeKeyDepositTxHash, transaction.DepositTxHash),
		sdk.NewAttribute(types.AttributeKeyDepositNonce, big.NewInt(int64(transaction.DepositTxIndex)).String()),
		sdk.NewAttribute(types.AttributeKeyDepositChainId, transaction.DepositChainId),
		sdk.NewAttribute(types.AttributeKeyDepositAmount, transaction.DepositAmount),
		sdk.NewAttribute(types.AttributeKeyDepositBlock, big.NewInt(int64(transaction.DepositBlock)).String()),
		sdk.NewAttribute(types.AttributeKeyDepositToken, transaction.DepositToken),
		sdk.NewAttribute(types.AttributeKeyWithdrawalAmount, transaction.WithdrawalAmount),
		sdk.NewAttribute(types.AttributeKeyDepositor, transaction.Depositor),
		sdk.NewAttribute(types.AttributeKeyReceiver, transaction.Receiver),
		sdk.NewAttribute(types.AttributeKeyWithdrawalChainID, transaction.WithdrawalChainId),
		sdk.NewAttribute(types.AttributeKeyWithdrawalTxHash, transaction.WithdrawalTxHash),
		sdk.NewAttribute(types.AttributeKeyWithdrawalToken, transaction.WithdrawalToken),
		sdk.NewAttribute(types.AttributeKeySignature, transaction.Signature),
		sdk.NewAttribute(types.AttributeKeyIsWrapped, strconv.FormatBool(transaction.IsWrapped)),
		sdk.NewAttribute(types.AttributeKeyCommissionAmount, transaction.CommissionAmount)))
}
