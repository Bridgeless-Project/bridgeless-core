package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) SetCommission(ctx sdk.Context, commission types.Commission) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	cStore.Set(types.KeyCommission(commission.TokenId), k.cdc.MustMarshal(&commission))
}

func (k Keeper) GetCommission(ctx sdk.Context, tokenId uint64) (types.Commission, bool) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))

	var commission types.Commission
	bz := cStore.Get(types.KeyCommission(tokenId))
	if bz == nil {
		return commission, false
	}
	k.cdc.MustUnmarshal(bz, &commission)
	return commission, true
}

func (k Keeper) RemoveCommission(ctx sdk.Context, commission types.Commission) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))
	cStore.Delete(types.KeyCommission(commission.TokenId))
}

func (k Keeper) GetCommissionWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Commission, *query.PageResponse, error) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreCommissionPrefix))

	var commissions []types.Commission
	pageRes, err := query.Paginate(cStore, pagination, func(key []byte, value []byte) error {
		var commission types.Commission
		if err := k.cdc.Unmarshal(value, &commission); err != nil {
			return err
		}
		commissions = append(commissions, commission)
		return nil
	})
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to get paginated commission")
	}

	return commissions, pageRes, nil
}
