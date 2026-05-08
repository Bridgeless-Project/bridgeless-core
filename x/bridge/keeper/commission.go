package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) SetCommission(sdkCtx sdk.Context, epochId uint32, commission types.Commission) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	eStore.Set(types.KeyEpochCommission(epochId, commission.TokenId), k.cdc.MustMarshal(&commission))
}

func (k Keeper) GetCommission(sdkCtx sdk.Context, epochId uint32, tokenId uint64) (types.Commission, bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	var commission types.Commission
	bz := eStore.Get(types.KeyCommission(tokenId))
	if bz == nil {
		return commission, false
	}
	k.cdc.MustUnmarshal(bz, &commission)
	return commission, true
}

func (k Keeper) RemoveCommission(sdkCtx sdk.Context, epochId uint32, tokenId uint64) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	eStore.Delete(types.KeyCommission(tokenId))
}

func (k Keeper) GetCommissionsWithPagination(sdkCtx sdk.Context, epochId uint32, pagination *query.PageRequest) ([]types.Commission, *query.PageResponse, error) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	var commissions []types.Commission
	pageRes, err := query.Paginate(eStore, pagination, func(key []byte, value []byte) error {
		var commission types.Commission
		if err := k.cdc.Unmarshal(value, &commission); err != nil {
			return err
		}
		commissions = append(commissions, commission)
		return nil
	})
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to get paginated commissions")
	}

	return commissions, pageRes, nil
}
