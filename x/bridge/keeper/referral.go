package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddReferral(sdkCtx sdk.Context, Referral types.Referral) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralPrefix))
	cStore.Set(types.KeyReferral(Referral.Id), k.cdc.MustMarshal(&Referral))
}

func (k Keeper) GetReferral(sdkCtx sdk.Context, id uint32) (Referral types.Referral, found bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralPrefix))
	bz := cStore.Get(types.KeyReferral(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &Referral)
	found = true

	return
}

func (k Keeper) GetReferralsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Referral, *query.PageResponse, error) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreReferralPrefix))
	var Referrals []types.Referral

	pageRes, err := query.Paginate(cStore, pagination, func(key []byte, value []byte) error {
		var Referral types.Referral

		k.cdc.MustUnmarshal(value, &Referral)

		Referrals = append(Referrals, Referral)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return Referrals, pageRes, nil
}

func (k Keeper) GetAllReferrals(sdkCtx sdk.Context) (Referrals []types.Referral) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralPrefix))
	iterator := cStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var Referral types.Referral
		k.cdc.MustUnmarshal(iterator.Value(), &Referral)
		Referrals = append(Referrals, Referral)
	}

	return
}

func (k Keeper) DeleteReferral(sdkCtx sdk.Context, id uint32) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreReferralPrefix))
	cStore.Delete(types.KeyReferral(id))
}
