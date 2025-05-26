package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (k Keeper) SetTransactionSubmissions(sdkCtx sdk.Context, txSubmissions *types.TransactionSubmissions) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))

	tStore.Set([]byte(txSubmissions.TxHash), k.cdc.MustMarshal(txSubmissions))
}

func (k Keeper) GetPaginatedTransactionsSubmissions(sdkCtx sdk.Context, pagination *query.PageRequest) (
	[]types.TransactionSubmissions, *query.PageResponse, error) {

	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))

	var txsWithSubmissions []types.TransactionSubmissions
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var txSubmissions types.TransactionSubmissions
		if err := k.cdc.Unmarshal(value, &txSubmissions); err != nil {
			return err
		}
		txsWithSubmissions = append(txsWithSubmissions, txSubmissions)
		return nil
	})

	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to get paginated transactions submissions")
	}

	return txsWithSubmissions, pageRes, nil
}

func (k Keeper) GetTransactionSubmissions(sdkCtx sdk.Context, txHash string) (txSubmissions types.TransactionSubmissions, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))

	bz := tStore.Get([]byte(txHash))
	if bz == nil {
		return txSubmissions, false
	}

	k.cdc.MustUnmarshal(bz, &txSubmissions)

	return txSubmissions, true
}

func (k Keeper) TxHash(tx *types.Transaction) common.Hash {
	return crypto.Keccak256Hash(k.cdc.MustMarshal(tx))
}
