package keeper

import (
	"context"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m msgServer) SubmitSwapTx(goCtx context.Context, msg *types.MsgSubmitSwapTx) (*types.MsgSubmitSwapTxResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.bridge.IsParty(ctx, msg.Creator) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "creator is not an authorized bridge party")
	}

	commission, err := m.computeCommission(ctx, msg.Tx)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to compute commission")
	}

	if _, found := m.GetSwap(ctx, msg.Tx.Tx.DepositTxHash, msg.Tx.Tx.DepositTxIndex, msg.Tx.Tx.DepositChainId); found {
		return nil, errorsmod.Wrap(types.ErrAlreadyProcessed, "swap was already executed")
	}

	requestHash := m.SwapHash(msg).Hex()
	submissions, found := m.GetSwapSubmissions(ctx, requestHash)
	if !found {
		submissions = bridgetypes.Submissions{Hash: requestHash}
	}

	if hasSubmitter(submissions.Submitters, msg.Creator) {
		return nil, errorsmod.Wrap(types.ErrAlreadySubmitted, "swap has already been submitted by this creator")
	}

	submissions.Submitters = append(submissions.Submitters, msg.Creator)
	m.SetSwapSubmissions(ctx, &submissions)

	threshold := m.bridge.GetParams(ctx).TssThreshold
	if len(submissions.Submitters) != int(threshold+1) {
		return &types.MsgSubmitSwapTxResponse{}, nil
	}

	swap, err := m.executeSwap(ctx, msg)
	if err != nil {
		return nil, err
	}

	m.SetSwap(ctx, *swap)
	if msg.Tx.IsFeeDistribution {
		if commission == nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFee, "commission is nil")
		}

		m.bridge.SetCommission(ctx, msg.Tx.Tx.EpochId, *commission)
	}

	return &types.MsgSubmitSwapTxResponse{}, nil
}

func (m msgServer) computeCommission(ctx sdk.Context, tx *types.SwapTransaction) (*bridgetypes.Commission, error) {
	if !tx.IsFeeDistribution {
		return nil, nil
	}
	depositTokenInfo, found := m.bridge.GetTokenInfo(ctx, tx.Tx.DepositToken, tx.Tx.DepositChainId)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrTokenInfoNotFound, "token info not found for %s on chain %s", tx.Tx.WithdrawalToken, tx.Tx.WithdrawalChainId)
	}

	commission, found := m.bridge.GetCommission(ctx, tx.Tx.EpochId, depositTokenInfo.TokenId)
	if !found {
		return nil, errorsmod.Wrapf(bridgetypes.ErrCommissionNotFound, "commission not found for token %s", depositTokenInfo.TokenId)
	}

	commissionAmount, ok := new(big.Int).SetString(commission.Amount, 10)
	if !ok {
		return nil, errorsmod.Wrapf(bridgetypes.ErrInvalidCommission, "invalid commission amount: %s", commission.Amount)
	}

	withdrawalAmount, ok := new(big.Int).SetString(commission.Amount, 10)
	if !ok {
		return nil, errorsmod.Wrapf(bridgetypes.ErrInvalidAmount, "invalid withdrawal amount: %s", tx.Tx.WithdrawalAmount)
	}

	commissionAmount.Sub(commissionAmount, withdrawalAmount)
	if commissionAmount.Sign() < 0 {
		return nil, errorsmod.Wrapf(bridgetypes.ErrInvalidCommission, "withdrawal amount %s exceeds commission amount %s", withdrawalAmount.String(), commission.Amount)
	}

	commission.Amount = commissionAmount.String()
	return &commission, nil
}
