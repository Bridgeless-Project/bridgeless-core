package v3

import (
	"fmt"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.12-rc1 %s module migrations", types.ModuleName))

	// reading old structures of Transaction (without `signature` and `is_wrapped`)
	txStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTransactionPrefix))
	tStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenInfoPrefix))
	iterator := sdk.KVStorePrefixIterator(txStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Transaction
		cdc.MustUnmarshal(iterator.Value(), &val)

		// set empty string as default signature
		val.Signature = ""

		bz := tStore.Get(types.KeyTokenInfo(val.WithdrawalChainId, val.WithdrawalToken))
		if bz == nil {
			// set false as default `IsWrapped`
			val.IsWrapped = false
			txStore.Set(types.KeyTransaction(types.TransactionId(&val)), cdc.MustMarshal(&val))
			continue
		}

		tokenInfo := new(types.TokenInfo)
		cdc.MustUnmarshal(bz, tokenInfo)

		// set new value
		val.IsWrapped = tokenInfo.IsWrapped
		txStore.Set(types.KeyTransaction(
			types.TransactionId(&val)),
			cdc.MustMarshal(&val),
		)
	}

	return nil
}
