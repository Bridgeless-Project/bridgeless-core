package v6

import (
	"fmt"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const defaultMaxWithdrawalAmount = "0"

type entry struct {
	key   []byte
	value []byte
}

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.32 %s module migrations", types.ModuleName))

	// token info is stored in three places: the token itself, the token info store and the token pairs store
	migrateTokens(ctx, storeKey, cdc)
	migrateTokenInfos(ctx, storeKey, cdc, types.StoreTokenInfoPrefix)
	migrateTokenInfos(ctx, storeKey, cdc, types.StoreTokenPairsPrefix)

	return nil
}

func migrateTokens(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	var entries []entry
	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		cdc.MustUnmarshal(iterator.Value(), &token)

		for i := range token.Info {
			setDefaultMaxWithdrawalAmount(&token.Info[i])
		}

		entries = append(entries, entry{
			key:   append([]byte{}, iterator.Key()...),
			value: cdc.MustMarshal(&token),
		})
	}

	for _, e := range entries {
		store.Set(e.key, e.value)
	}
}

// migrateTokenInfos migrates the store under the given prefix which values are types.TokenInfo
func migrateTokenInfos(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, storePrefix string) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(storePrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	var entries []entry
	for ; iterator.Valid(); iterator.Next() {
		var tokenInfo types.TokenInfo
		cdc.MustUnmarshal(iterator.Value(), &tokenInfo)

		setDefaultMaxWithdrawalAmount(&tokenInfo)

		entries = append(entries, entry{
			key:   append([]byte{}, iterator.Key()...),
			value: cdc.MustMarshal(&tokenInfo),
		})
	}

	for _, e := range entries {
		store.Set(e.key, e.value)
	}
}

func setDefaultMaxWithdrawalAmount(info *types.TokenInfo) {
	if info.MaxWithdrawalAmount == "" {
		info.MaxWithdrawalAmount = defaultMaxWithdrawalAmount
	}
}
