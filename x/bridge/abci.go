package bridge

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	params := k.GetParams(ctx)
	nextEpoch, found := k.GetEpoch(ctx, params.Epoch+1)
	if !found || nextEpoch.Status != types.EpochStatus_FINALIZING {
		return
	}

	nextEpoch.Status = types.EpochStatus_COMPLETE
	k.SetEpoch(ctx, &nextEpoch)

	params = types.Params{
		ModuleAdmin:     params.ModuleAdmin,
		Epoch:           nextEpoch.Id,
		TssThreshold:    nextEpoch.TssThreshold,
		Parties:         nextEpoch.Parties,
		RelayerAccounts: params.RelayerAccounts,
	}
	k.SetParams(ctx, params)
}
