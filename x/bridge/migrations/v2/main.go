package v2

import (
	"fmt"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.10-rc1 %s module migrations", types.ModuleName))

	// reading old structures of token info (without field "is_wrapped" into new one)
	store := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		cdc.MustUnmarshal(iterator.Value(), &token)

		// just resetting correct structures where "is_wrapped" exists
		store.Set(types.KeyToken(token.Id), cdc.MustMarshal(&token))

		for _, tokenInfo := range token.Info {
			// adding new store

			tiStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenInfoPrefix))
			tiStore.Set(types.KeyTokenInfo(tokenInfo.ChainId, tokenInfo.Address), cdc.MustMarshal(&tokenInfo))

			pStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenPairsPrefix))
			srcBranchStore := prefix.NewStore(pStore, types.TokenPairPrefix(tokenInfo.ChainId, tokenInfo.Address))

			for _, pair := range token.Info {
				if pair.ChainId == tokenInfo.ChainId {
					continue
				}

				srcBranchStore.Set(types.KeyTokenPair(pair.ChainId), cdc.MustMarshal(&pair))
			}
		}
	}

	return nil
}
