package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetChain(sdkCtx sdk.Context, chain types.Chain) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreChainPrefix))
	cStore.Set(types.KeyChain(chain.Id), k.cdc.MustMarshal(&chain))
}

func (k Keeper) GetChain(sdkCtx sdk.Context, id string) (chain types.Chain, found bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreChainPrefix))
	bz := cStore.Get(types.KeyChain(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &chain)
	found = true

	return
}

func (k Keeper) GetChainsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Chain, *query.PageResponse, error) {
	cStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreChainPrefix))
	var chains []types.Chain

	pageRes, err := query.Paginate(cStore, pagination, func(key []byte, value []byte) error {
		var chain types.Chain

		k.cdc.MustUnmarshal(value, &chain)

		chains = append(chains, chain)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return chains, pageRes, nil
}

func (k Keeper) GetAllChains(sdkCtx sdk.Context) (chains []types.Chain) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreChainPrefix))
	iterator := cStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var chain types.Chain
		k.cdc.MustUnmarshal(iterator.Value(), &chain)
		chains = append(chains, chain)
	}

	return
}

func (k Keeper) RemoveChain(sdkCtx sdk.Context, id string) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreChainPrefix))
	cStore.Delete(types.KeyChain(id))
}
