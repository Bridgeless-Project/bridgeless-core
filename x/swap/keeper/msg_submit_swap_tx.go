package keeper

import (
	"context"

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
		return nil, errorsmod.Wrap(err, "failed to execute swap")
	}

	m.SetSwap(ctx, *swap)
	if msg.Tx.IsFeeDistribution {
		if commission == nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInsufficientFee, "commission is nil")
		}

		m.bridge.SetCommission(ctx, msg.Tx.Tx.EpochId, *commission)

		amount, ok := sdk.NewIntFromString(commission.Amount)
		if !ok {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid commission amount: %s", commission.Amount)
		}

		err = m.bridge.PartiesDistributeFee(ctx, msg.Tx.Tx.EpochId, sdk.NewCoin(m.staking.BondDenom(ctx), amount))
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to distribute fee among parties")
		}
	}

	return &types.MsgSubmitSwapTxResponse{}, nil
}
