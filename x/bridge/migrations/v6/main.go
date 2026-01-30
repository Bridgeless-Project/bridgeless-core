package v6

import (
	"fmt"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.17-rc3 %s module migrations", types.ModuleName))

	txStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTransactionPrefix))
	iterator := sdk.KVStorePrefixIterator(txStore, []byte{})

	defer iterator.Close()

	// NOTE: comment out migration
	//for ; iterator.Valid(); iterator.Next() {
	//
	//	// get oldTx data with unnecessary commission_amount field
	//	var oldTx txSubmissions
	//	cdc.MustUnmarshal(iterator.Value(), &oldTx)
	//
	//	newTx := types.TransactionSubmissions{
	//		TxHash:     oldTx.TxHash,
	//		Submitters: oldTx.Submitters,
	//	}
	//
	//	// set new transaction instead of old one
	//	txStore.Set(
	//		[]byte(oldTx.TxHash),
	//		cdc.MustMarshal(&newTx),
	//	)
	//}

	return nil
}
