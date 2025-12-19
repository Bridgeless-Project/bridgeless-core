package keeper

import (
	"time"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// update vesting state for each nft
func (k Keeper) EndBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	newEpochNeeded := k.GetNewEpochNeeded(ctx)

	if !newEpochNeeded {
		return
	}

	currentEpoch, ok := k.GetEpoch(ctx, k.GetParams(ctx).EpochSequence)
	if !ok {
		k.Logger(ctx).Error("current epoch not set")
		return
	}

	currentEpoch.EndBlock = ctx.BlockHeight()
	k.SetEpoch(ctx, currentEpoch)

	k.CreateEpoch(ctx, ctx.BlockHeight(), 0)
}
