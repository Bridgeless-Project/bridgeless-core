package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetReferralRewards(sdkCtx sdk.Context, referralId uint32, tokenId uint64, ReferralRewards types.ReferralRewards) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	cStore.Set(types.KeyReferralRewards(referralId, tokenId), k.cdc.MustMarshal(&ReferralRewards))
}

func (k Keeper) GetReferralRewards(sdkCtx sdk.Context, referralId uint32, tokenId uint64) (ReferralRewards types.ReferralRewards, found bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	bz := cStore.Get(types.KeyReferralRewards(referralId, tokenId))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &ReferralRewards)
	found = true

	return
}

func (k Keeper) GetReferralRewardssWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.ReferralRewards, *query.PageResponse, error) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	var ReferralRewardss []types.ReferralRewards

	pageRes, err := query.Paginate(cStore, pagination, func(key []byte, value []byte) error {
		var ReferralRewards types.ReferralRewards

		k.cdc.MustUnmarshal(value, &ReferralRewards)

		ReferralRewardss = append(ReferralRewardss, ReferralRewards)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return ReferralRewardss, pageRes, nil
}

func (k Keeper) GetAllReferralRewardss(sdkCtx sdk.Context) (ReferralRewardss []types.ReferralRewards) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	iterator := cStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ReferralRewards types.ReferralRewards
		k.cdc.MustUnmarshal(iterator.Value(), &ReferralRewards)
		ReferralRewardss = append(ReferralRewardss, ReferralRewards)
	}

	return
}

func (k Keeper) RemoveReferralRewards(sdkCtx sdk.Context, referralId uint32, tokenId uint64) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	cStore.Delete(types.KeyReferralRewards(referralId, tokenId))
}
