package v5

import (
	"fmt"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type txEntry struct {
	key []byte
	tx  types.Transaction
}

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.31 %s module migrations", types.ModuleName))

	txStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTransactionPrefix))
	txIterator := sdk.KVStorePrefixIterator(txStore, []byte{})
	defer txIterator.Close()

	var txEntries []txEntry
	for ; txIterator.Valid(); txIterator.Next() {
		var tx types.Transaction
		cdc.MustUnmarshal(txIterator.Value(), &tx)
		txEntries = append(txEntries, txEntry{
			key: append([]byte{}, txIterator.Key()...),
			tx:  tx,
		})
	}

	for _, e := range txEntries {
		txStore.Delete(e.key)
	}

	for _, e := range txEntries {
		newKey := []byte(types.TransactionId(&e.tx))
		if txStore.Has(newKey) {
			ctx.Logger().Info("skipping duplicate transaction", "key", string(newKey))
			continue
		}
		txStore.Set(newKey, cdc.MustMarshal(&e.tx))
	}

	stopTxstore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreStopListTransactionsPrefix))
	stTxIterator := sdk.KVStorePrefixIterator(stopTxstore, []byte{})
	defer stTxIterator.Close()

	var stopTxEntries []txEntry
	for ; stTxIterator.Valid(); stTxIterator.Next() {
		var tx types.Transaction
		cdc.MustUnmarshal(stTxIterator.Value(), &tx)
		stopTxEntries = append(stopTxEntries, txEntry{
			key: append([]byte{}, stTxIterator.Key()...),
			tx:  tx,
		})
	}

	for _, e := range stopTxEntries {
		stopTxstore.Delete(e.key)
	}

	for _, e := range stopTxEntries {
		newKey := []byte(types.TransactionId(&e.tx))
		if stopTxstore.Has(newKey) {
			ctx.Logger().Info("skipping duplicate stop list transaction", "key", string(newKey))
			continue
		}
		stopTxstore.Set(newKey, cdc.MustMarshal(&e.tx))
	}

	return nil
}
