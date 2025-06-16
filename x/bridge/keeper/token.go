package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetToken(sdkCtx sdk.Context, token types.Token) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenPrefix))
	tStore.Set(types.KeyToken(token.Id), k.cdc.MustMarshal(&token))
}

func (k Keeper) SetTokenInfo(sdkCtx sdk.Context, tokenInfo types.TokenInfo) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoPrefix))
	tStore.Set(types.KeyTokenInfo(tokenInfo.ChainId, tokenInfo.Address), k.cdc.MustMarshal(&tokenInfo))
}

func (k Keeper) GetToken(sdkCtx sdk.Context, id uint64) (token types.Token, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenPrefix))
	bz := tStore.Get(types.KeyToken(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &token)
	found = true

	return
}

func (k Keeper) GetTokenInfo(sdkCtx sdk.Context, chain, address string) (tokenInfo types.TokenInfo, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoPrefix))
	bz := tStore.Get(types.KeyTokenInfo(chain, address))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &tokenInfo)
	found = true

	return
}

func (k Keeper) GetTokensWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Token, *query.PageResponse, error) {
	tStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.Prefix(types.StoreTokenPrefix))
	var chains []types.Token

	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var chain types.Token
		k.cdc.MustUnmarshal(value, &chain)
		chains = append(chains, chain)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return chains, pageRes, nil
}

func (k Keeper) RemoveToken(sdkCtx sdk.Context, id uint64) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenPrefix))
	tStore.Delete(types.KeyToken(id))
}

func (k Keeper) RemoveTokenInfo(sdkCtx sdk.Context, chain, addr string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenInfoPrefix))
	tStore.Delete(types.KeyTokenInfo(chain, addr))
}
