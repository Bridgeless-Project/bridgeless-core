package bridge

import (
	"cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	if err := finishMigrationEpoch(ctx, k, params); err != nil {
		k.Logger(ctx).Error("failed to finish migration epoch", "error", err)
		return
	}
	if err := finishEpochSupport(ctx, k, params); err != nil {
		k.Logger(ctx).Error("failed to finish epoch support", "error", err)
		return
	}
}

func finishMigrationEpoch(ctx sdk.Context, k keeper.Keeper, params types.Params) error {
	nextEpoch, found := k.GetEpoch(ctx, params.Epoch+1)
	if !found || nextEpoch.Status != types.EpochStatus_MIGRATION_FINALIZING {
		return nil // any other status means no update is needed
	}

	nextEpoch.Status = types.EpochStatus_RUNNING
	k.SetEpoch(ctx, &nextEpoch)

	epoch, _ := k.GetEpoch(ctx, params.Epoch)
	epoch.Status = types.EpochStatus_SHUTDOWN
	epoch.EndBlock = uint64(ctx.BlockHeight()) + params.SupportingTime
	k.SetEpoch(ctx, &epoch)

	params = types.Params{
		ModuleAdmin:     params.ModuleAdmin,
		Epoch:           nextEpoch.Id,
		TssThreshold:    nextEpoch.TssThreshold,
		Parties:         nextEpoch.Parties,
		RelayerAccounts: params.RelayerAccounts,
		SupportingTime:  params.SupportingTime,
	}
	k.SetParams(ctx, params)

	if err := broadcastEpochUpdatedEvent(ctx, k, nextEpoch.Id, true); err != nil {
		return errors.Wrap(err, "broadcast epoch updated event")
	}

	return nil
}

func finishEpochSupport(ctx sdk.Context, k keeper.Keeper, params types.Params) error {
	if params.Epoch == 0 {
		return nil
	}

	prevEpoch, _ := k.GetEpoch(ctx, params.Epoch-1)
	// ol=nly if the epoch status
	if prevEpoch.Status == types.EpochStatus_SHUTDOWN && prevEpoch.EndBlock <= uint64(ctx.BlockHeight()) {
		prevEpoch.Status = types.EpochStatus_UNSUPPORTED
		k.SetEpoch(ctx, &prevEpoch)

		if err := broadcastEpochUpdatedEvent(ctx, k, prevEpoch.Id, false); err != nil {
			return errors.Wrap(err, "broadcast epoch updated event")
		}
	}
	return nil
}

func broadcastEpochUpdatedEvent(ctx sdk.Context, k keeper.Keeper, epochId uint32, isAdding bool) error {
	chainTypes := []types.ChainType{
		types.ChainType_EVM,
		types.ChainType_SOLANA,
		types.ChainType_TON,
	}

	for _, chainType := range chainTypes {
		signature, found := k.GetEpochChainSignature(ctx, epochId, chainType)
		if !found {
			return errors.Wrap(types.ErrEpochSignatureNotFound, chainType.String())
		}
		var epochChainSignatures types.EpochSignature

		if isAdding {
			epochChainSignatures = *signature.AddedSignature
		} else {
			epochChainSignatures = *signature.RemovedSignature
		}

		for _, chain := range k.GetAllChainsByType(ctx, chainType) {
			// Broadcast the epoch updated event for each chain in the epoch
			k.EmitEpochUpdatedEvent(ctx, epochId, chain.Id, epochChainSignatures, isAdding)
		}
	}

	return nil
}
