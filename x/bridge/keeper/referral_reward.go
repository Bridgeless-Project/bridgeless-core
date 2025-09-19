package keeper

import (
	"math/big"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InsertReferralRewards(sdkCtx sdk.Context, referralId uint32, tokenId uint64, ReferralRewards types.ReferralRewards) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	cStore.Set(types.KeyReferralRewards(referralId, tokenId), k.cdc.MustMarshal(&ReferralRewards))
}

func (k Keeper) AddReferralRewards(sdkCtx sdk.Context, referralId uint32, tokenId uint64, newReferralRewards types.ReferralRewards) error {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	bz := cStore.Get(types.KeyReferralRewards(referralId, tokenId))
	var referralRewards types.ReferralRewards
	if bz == nil {
		cStore.Set(types.KeyReferralRewards(referralId, tokenId), k.cdc.MustMarshal(&newReferralRewards))
		return nil
	}
	k.cdc.MustUnmarshal(bz, &referralRewards)

	if referralRewards.ToClaim == "" || referralRewards.TotalCollectedAmount == "" {
		cStore.Set(types.KeyReferralRewards(referralId, tokenId), k.cdc.MustMarshal(&newReferralRewards))
		return nil
	}

	// Sum the rewards
	toClaim, ok := big.NewInt(0).SetString(referralRewards.ToClaim, 10)
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid to-claim amount in existing referral rewards")
	}
	newToClaim, ok := big.NewInt(0).SetString(newReferralRewards.ToClaim, 10)
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid to-claim amount in new referral rewards")
	}
	totalCollectedAmount, ok := big.NewInt(0).SetString(referralRewards.TotalCollectedAmount, 10)
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid total collected amount in existing referral rewards")
	}
	newTotalCollectedAmount, ok := big.NewInt(0).SetString(newReferralRewards.TotalCollectedAmount, 10)
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid total collected amount in new referral rewards")
	}

	referralRewards.ToClaim = toClaim.Add(toClaim, newToClaim).String()
	referralRewards.TotalCollectedAmount = totalCollectedAmount.Add(totalCollectedAmount, newTotalCollectedAmount).String()

	// Update the store
	cStore.Set(types.KeyReferralRewards(referralId, tokenId), k.cdc.MustMarshal(&referralRewards))
	return nil
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

func (k Keeper) GetReferralRewardsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.ReferralRewards, *query.PageResponse, error) {
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

func (k Keeper) GetAllReferralRewards(sdkCtx sdk.Context) (ReferralRewardss []types.ReferralRewards) {
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

func (k Keeper) DeleteReferralRewards(sdkCtx sdk.Context, referralId uint32, tokenId uint64) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralRewardsPrefix))
	cStore.Delete(types.KeyReferralRewards(referralId, tokenId))
}
