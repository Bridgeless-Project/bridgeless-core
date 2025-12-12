package v7

import (
	"fmt"

	oldTypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/migrations/v8/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.25-rc1 %s module migrations", types.ModuleName))

	tokenStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenPrefix))
	iterator := sdk.KVStorePrefixIterator(tokenStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oldToken oldTypes.Token
		cdc.MustUnmarshal(iterator.Value(), &oldToken)

		newToken := types.Token{
			Id:       oldToken.Id,
			Metadata: types.TokenMetadata(oldToken.Metadata),
			Info:     getUpdatedTokenInfo(ctx, oldToken, storeKey, cdc),
		}

		// set new token instead of old one
		tokenStore.Set(
			types.KeyToken(newToken.Id),
			cdc.MustMarshal(&newToken),
		)
	}

	return nil
}

func setTokenPairs(sdkCtx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, current types.TokenInfo, pairs ...oldTypes.TokenInfo) {
	pStore := prefix.NewStore(sdkCtx.KVStore(storeKey), types.Prefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.TokenPairPrefix(current.ChainId, current.Address))

	for _, pair := range pairs {
		if pair.ChainId == current.ChainId {
			continue
		}

		// convert old token info to new one
		newPair := types.TokenInfo{
			Address:             pair.Address,
			ChainId:             pair.ChainId,
			CommissionRate:      current.CommissionRate,
			MinWithdrawalAmount: pair.MinWithdrawalAmount,
			Decimals:            pair.Decimals,
			IsWrapped:           pair.IsWrapped,
			TokenId:             pair.TokenId,
		}
		srcBranchStore.Set(types.KeyTokenPair(pair.ChainId), cdc.MustMarshal(&newPair))
	}
}

// getUpdatedTokenInfo iterate over old token info and convert it to new one. Once converted, it sets new token info to new
// store and returns array of new token info. Also this function sets updated token info to token pairs store.
func getUpdatedTokenInfo(ctx sdk.Context, token oldTypes.Token, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) []types.TokenInfo {
	tokenInfoStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenInfoPrefix))
	tokenInfos := make([]types.TokenInfo, len(token.Info))

	// convert old token info to new one and set to new store and tokenInfos arrays
	for i, info := range token.Info {
		tokenInfos[i] = types.TokenInfo{
			Address:             info.Address,
			ChainId:             info.ChainId,
			CommissionRate:      token.CommissionRate,
			MinWithdrawalAmount: info.MinWithdrawalAmount,
			Decimals:            info.Decimals,
			IsWrapped:           info.IsWrapped,
			TokenId:             info.TokenId,
		}

		// set new token info instead of old one
		tokenInfoStore.Set(
			types.KeyTokenInfo(tokenInfos[i].ChainId, tokenInfos[i].Address),
			cdc.MustMarshal(&tokenInfos[i]),
		)

		// set token pairs
		setTokenPairs(ctx, storeKey, cdc, tokenInfos[i], token.Info...)
	}

	return tokenInfos
}
