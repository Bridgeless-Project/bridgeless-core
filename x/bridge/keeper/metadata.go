package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetTokenInfoMetadata(sdkCtx sdk.Context, metadata types.TokenInfoMetadata) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoMetadataPrefix))
	tStore.Set(types.KeyTokenInfoMetadata(metadata.ChainId, metadata.Address), k.cdc.MustMarshal(&metadata))
}

func (k Keeper) GetTokenInfoMetadata(sdkCtx sdk.Context, chain, address string) (metadata types.TokenInfoMetadata, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoMetadataPrefix))
	bz := tStore.Get(types.KeyTokenInfoMetadata(chain, address))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &metadata)
	found = true

	return
}

func (k Keeper) GetAllTokenInfoMetadata(sdkCtx sdk.Context) (metadata []types.TokenInfoMetadata) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoMetadataPrefix))
	iterator := tStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var entry types.TokenInfoMetadata
		k.cdc.MustUnmarshal(iterator.Value(), &entry)
		metadata = append(metadata, entry)
	}

	return
}

func (k Keeper) GetTokensInfoMetadataWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.TokenInfoMetadata, *query.PageResponse, error) {
	tStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoMetadataPrefix))
	var metadata []types.TokenInfoMetadata

	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var entry types.TokenInfoMetadata
		k.cdc.MustUnmarshal(value, &entry)
		metadata = append(metadata, entry)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return metadata, pageRes, nil
}

func (k Keeper) RemoveTokenInfoMetadata(sdkCtx sdk.Context, chain, addr string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoMetadataPrefix))
	tStore.Delete(types.KeyTokenInfoMetadata(chain, addr))
}
