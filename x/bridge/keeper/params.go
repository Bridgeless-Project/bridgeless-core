package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) IsParty(ctx sdk.Context, sender string) bool {
	parties := k.GetParams(ctx).Parties
	for _, party := range parties {
		if party.Address == sender {
			return true
		}
	}
	return false
}

func (k Keeper) IsEpochParty(ctx sdk.Context, sender string, epochId uint32) bool {
	epoch, found := k.GetEpoch(ctx, epochId)
	if !found {
		return false
	}

	for _, party := range epoch.Parties {
		if party.Address == sender {
			return true
		}
	}
	return false

}
