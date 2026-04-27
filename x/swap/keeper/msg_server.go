package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

func (m msgServer) SubmitSwapTx(goCtx context.Context, msg *types.MsgSubmitSwapTx) (*types.MsgSubmitSwapTxResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	if msg.Tx == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "swap transaction cannot be nil")
	}
	if m.bridge == nil || m.erc20 == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidConfig, "swap keeper dependencies are not configured")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.bridge.IsParty(ctx, msg.Creator) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "creator is not an authorized bridge party")
	}

	if _, found := m.GetSwap(ctx, msg.Tx.Tx.DepositTxHash, msg.Tx.Tx.DepositTxIndex, msg.Tx.Tx.DepositChainId); found {
		return nil, errorsmod.Wrap(types.ErrAlreadyProcessed, "swap was already executed")
	}

	requestHash := m.SwapRequestHash(msg).Hex()
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
	return &types.MsgSubmitSwapTxResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
