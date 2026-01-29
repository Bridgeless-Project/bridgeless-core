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

	if params.Epoch+1 != msg.EpochId {
		return nil, errorsmod.Wrapf(types.ErrInvalidEpochID, "invalid epoch ID: expected %d, got %d", params.Epoch+1, msg.EpochId)
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
	// we do not need to set up a relayer_addresses here

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
	_, found := m.Keeper.GetEpoch(ctx, msg.EpochChainSignatures.EpochId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidEpochID, "epoch not found")
	}

	// TODO: add submitters
	m.Keeper.SetEpochChainSignature(ctx, &msg.EpochChainSignatures)
	return &types.MsgSetEpochSignatureResponse{}, nil
}

func (m msgServer) FinishEpochMigration(goCtx context.Context, msg *types.MsgFinishEpochMigration) (*types.MsgFinishEpochMigrationResponse, error) {
	if msg == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidDataType, "message cannot be nil")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.Keeper.GetParams(ctx)

	if params.Epoch+1 != msg.EpochId {
		return nil, errorsmod.Wrapf(types.ErrInvalidEpochID, "expected epoch ID %d, got %d", params.Epoch+1, msg.EpochId)
	}

	// validate that for all chains the signature set
	epoch, found := m.Keeper.GetEpoch(ctx, msg.EpochId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidEpochID, "epoch not found")
	}

	epoch.Status = types.EpochStatus_FINALIZING
	m.Keeper.SetEpoch(ctx, &epoch)

	return &types.MsgFinishEpochMigrationResponse{}, nil
}

func determineEpochSigners(tssParties []*types.Party, tssInfo []types.TSSInfo) (result []*types.Party, err error) {
	validate := func(tssParties []*types.Party, info types.TSSInfo) ([]*types.Party, error) {
		for i, party := range tssParties {
			if party.Address == info.Address {
				if info.Active {
					return nil, errorsmod.Wrap(types.ErrInvalidPartiesList, "duplicate active party found")
				}
				tssParties = slices.Delete(tssParties, i, 1)
				return tssParties, nil
			}
		}
		return append(tssParties, &types.Party{
			Address: info.Address,
		}), nil
	}

	for _, info := range tssInfo {
		result, err = validate(result, info)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrInvalidPartiesList, "failed to validate TSS parties list")
		}
	}

	return result, nil
}
