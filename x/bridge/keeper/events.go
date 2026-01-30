package keeper

import (
	"math/big"
	"strconv"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
		sdk.NewAttribute(types.AttributeKeyCommissionAmount, transaction.CommissionAmount),
		sdk.NewAttribute(types.AttributeEpochId, big.NewInt(int64(transaction.EpochId)).String()),
		sdk.NewAttribute(types.AttributeKeyMerkleProof, transaction.MerkleProof)))
}

func emitRemoveTransactionEvent(sdkCtx sdk.Context, transaction types.Transaction) {
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(types.EventType_TRANSACTION_DELETED.String(),
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
		sdk.NewAttribute(types.AttributeKeyCommissionAmount, transaction.CommissionAmount),
		sdk.NewAttribute(types.AttributeKeyMerkleProof, transaction.MerkleProof),
		sdk.NewAttribute(types.AttributeEpochId, big.NewInt(int64(transaction.EpochId)).String())))
}

func emitSubmitEvent(sdkCtx sdk.Context, transaction types.Transaction) {
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(types.EventType_DEPOSIT_SUBMITTED.String(),
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
		sdk.NewAttribute(types.AttributeKeyCommissionAmount, transaction.CommissionAmount),
		sdk.NewAttribute(types.AttributeKeyMerkleProof, transaction.MerkleProof),
		sdk.NewAttribute(types.AttributeEpochId, big.NewInt(int64(transaction.EpochId)).String()),
	))
}

func emitStartEpochEvent(sdkCtx sdk.Context, epochId uint32, info string) {
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventType_STARTED_NEW_EPOCH.String(),
			sdk.NewAttribute(types.AttributeEpochId, big.NewInt(int64(epochId)).String()),
			sdk.NewAttribute(types.AttributeTssInfo, info),
		),
	)

}
