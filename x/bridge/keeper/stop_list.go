package keeper

import (
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetTxToStopList(sdkCtx sdk.Context, tx types.Transaction) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreBlacklistTransactionsPrefix))
	tStore.Set(types.KeyTransaction(types.TransactionId(&tx)), k.cdc.MustMarshal(&tx))
}

func (k Keeper) DeleteTxFromStopList(sdkCtx sdk.Context, txKey string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreBlacklistTransactionsPrefix))
	tStore.Delete(types.KeyTransaction(txKey))
}

func (k Keeper) GetTxFromStopList(sdkCtx sdk.Context, txKey string) (types.Transaction, bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreBlacklistTransactionsPrefix))

	var transaction types.Transaction
	txBytes := tStore.Get(types.KeyTransaction(txKey))
	if txBytes == nil {
		return transaction, false
	}
	k.cdc.MustUnmarshal(txBytes, &transaction)

	return transaction, true
}

func (k Keeper) GetTxsFromStopListWithPagination(sdkCtx sdk.Context, pagination *query.PageRequest) ([]types.Transaction, *query.PageResponse, error) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTokenPrefix))
	var txs []types.Transaction

	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var tx types.Transaction
		k.cdc.MustUnmarshal(value, &tx)
		txs = append(txs, tx)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return txs, pageRes, nil
}
