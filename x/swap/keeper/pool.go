package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetPool(ctx sdk.Context, pool types.SwapPool) {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StorePoolPrefix))
	pStore.Set(types.KeyPool(pool.TokenId), k.cdc.MustMarshal(&pool))
}

func (k Keeper) AddPool(ctx sdk.Context, pool types.SwapPool) error {
	_, found := k.GetPool(ctx, pool.TokenId)
	if found {
		return errorsmod.Wrap(sdkerrors.ErrConflict, "pool already exists")
	}

	k.SetPool(ctx, pool)
	return nil
}

func (k Keeper) GetPool(ctx sdk.Context, tokenID string) (pool types.SwapPool, found bool) {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StorePoolPrefix))
	bz := pStore.Get(types.KeyPool(tokenID))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &pool)
	found = true

	return
}

func (k Keeper) UpdatePool(ctx sdk.Context, pool types.SwapPool) error {
	_, found := k.GetPool(ctx, pool.TokenId)
	if !found {
		return errorsmod.Wrap(types.ErrPoolNotFound, "pool not found")
	}

	k.SetPool(ctx, pool)
	return nil
}

func (k Keeper) RemovePool(ctx sdk.Context, tokenID string) error {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StorePoolPrefix))

	_, found := k.GetPool(ctx, tokenID)
	if !found {
		return errorsmod.Wrap(types.ErrPoolNotFound, "pool not found")
	}

	pStore.Delete(types.KeyPool(tokenID))
	return nil
}

func (k Keeper) GetPoolsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.SwapPool, *query.PageResponse, error) {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StorePoolPrefix))
	var pools []types.SwapPool

	pageRes, err := query.Paginate(pStore, pagination, func(key []byte, value []byte) error {
		var pool types.SwapPool

		k.cdc.MustUnmarshal(value, &pool)
		pools = append(pools, pool)
		return nil
	})
	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return pools, pageRes, nil
}

func (k Keeper) GetAllPools(ctx sdk.Context) (pools []types.SwapPool) {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StorePoolPrefix))
	iterator := pStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var pool types.SwapPool
		k.cdc.MustUnmarshal(iterator.Value(), &pool)
		pools = append(pools, pool)
	}

	return
}
