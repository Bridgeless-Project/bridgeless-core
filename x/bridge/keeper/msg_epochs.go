package keeper

import (
	"context"
	"encoding/json"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) StartEpoch(goCtx context.Context, msg *types.MsgStartEpoch) (*types.MsgStartEpochResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidDataType, "message cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.Keeper.GetParams(ctx)
	if params.Epoch+1 != msg.EpochId {
		return nil, errorsmod.Wrapf(types.ErrInvalidEpochID, "expected epoch ID %d, got %d", params.Epoch+1, msg.EpochId)
	}

	if params.ModuleAdmin != msg.Creator {
		return nil, errorsmod.Wrapf(types.ErrPermissionDenied, "only module admin can start a new epoch")
	}

	_, found := m.Keeper.GetEpoch(ctx, msg.EpochId+1)
	if found {
		return nil, errorsmod.Wrapf(types.ErrInvalidEpochID, "epoch %d already started", msg.EpochId)
	}

	parties, err := determineEpochSigners(params.Parties, msg.Info)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidPartiesList, "failed to determine epoch signers: %v", err)
	}

	epoch := &types.Epoch{
		Id:           msg.EpochId,
		Status:       types.EpochStatus_INITIATED,
		TssInfo:      msg.Info,
		Parties:      parties,
		StartBlock:   uint64(ctx.BlockHeight()),
		TssThreshold: msg.TssThreshold,
	}

	m.Keeper.SetEpoch(ctx, epoch)

	// broadcast event
	tssInfo, err := json.Marshal(msg.Info)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrPackEvent, "failed to marshal TSS info: %v", err)
	}

	emitStartEpochEvent(ctx, msg.EpochId, string(tssInfo))

	return &types.MsgStartEpochResponse{}, nil
}

func (m msgServer) SetEpochSignature(goCtx context.Context, msg *types.MsgSetEpochSignature) (*types.MsgSetEpochSignatureResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidDataType, "message cannot be nil")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if !m.IsParty(ctx, msg.Creator) {
		return nil, errorsmod.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	params := m.Keeper.GetParams(ctx)
	isReadyToMigration := true
	for _, sig := range msg.EpochChainSignatures {
		_, found := m.Keeper.GetEpoch(ctx, sig.EpochId)
		if !found {
			return nil, errorsmod.Wrap(types.ErrInvalidEpochID, "epoch not found")
		}

		submissions, found := m.Keeper.GetEpochChainSignatureSubmission(ctx, sig.EpochId, sig.ChainType, m.Keeper.EpochSignatureHash(&sig).String())
		if !found {
			submissions.Hash = m.Keeper.EpochSignatureHash(&sig).String()
		}

		// If tx has been submitted before with the same address new submission is rejected
		if isSubmitter(submissions.Submitters, msg.Creator) {
			return nil, errorsmod.Wrap(types.ErrTranscationAlreadySubmitted,
				"transaction has been already submitted by this address")
		}

		submissions.Submitters = append(submissions.Submitters, msg.Creator)
		m.Keeper.SetEpochChainSignatureSubmission(ctx, sig.EpochId, sig.ChainType, submissions)

		if len(submissions.Submitters) < int(params.TssThreshold+1) {
			isReadyToMigration = false
			continue
		}
		if sig.Address != "" {
			chains := m.GetAllChainsByType(ctx, sig.ChainType)
			for _, chain := range chains {
				chain.BridgeAddress = sig.Address
				m.SetChain(ctx, chain)
				m.SetChainByType(ctx, chain)
			}
		}
		m.Keeper.SetEpochChainSignature(ctx, &sig)
	}

	if isReadyToMigration {
		epoch, _ := m.Keeper.GetEpoch(ctx, msg.EpochChainSignatures[0].EpochId)
		epoch.Status = types.EpochStatus_FINALIZING
		m.Keeper.SetEpoch(ctx, &epoch)
	}

	return &types.MsgSetEpochSignatureResponse{}, nil
}

func determineEpochSigners(tssParties []*types.Party, tssInfo []types.TSSInfo) ([]*types.Party, error) {
	var err error
	validate := func(tssParties []*types.Party, info types.TSSInfo) ([]*types.Party, error) {
		for i, party := range tssParties {
			if party.Address == info.Address {
				if info.Active {
					return nil, errorsmod.Wrap(types.ErrInvalidPartiesList, "duplicate active party found")
				}
				tssParties = slices.Delete(tssParties, i, i+1)
				return tssParties, nil
			}
		}
		return append(tssParties, &types.Party{
			Address: info.Address,
		}), nil
	}

	for _, info := range tssInfo {
		tssParties, err = validate(tssParties, info)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrInvalidPartiesList, "failed to validate TSS parties list")
		}
	}

	return tssParties, nil
}
