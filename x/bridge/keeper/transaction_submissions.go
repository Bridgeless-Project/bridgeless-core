package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k Keeper) SetTransactionSubmissions(sdkCtx sdk.Context, txSubmissions *types.Submissions) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))
	tStore.Set(types.KeyTransactionSubmissions(txSubmissions.Hash), k.cdc.MustMarshal(txSubmissions))
}

func (k Keeper) GetPaginatedTransactionsSubmissions(sdkCtx sdk.Context, pagination *query.PageRequest) (
	[]types.Submissions, *query.PageResponse, error) {

	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))

	var txsWithSubmissions []types.Submissions
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var txSubmissions types.Submissions
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

func (k Keeper) GetTransactionSubmissions(sdkCtx sdk.Context, txHash string) (txSubmissions types.Submissions, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))

	bz := tStore.Get(types.KeyTransactionSubmissions(txHash))
	if bz == nil {
		return txSubmissions, false
	}

	k.cdc.MustUnmarshal(bz, &txSubmissions)

	return txSubmissions, true
}

func (k Keeper) RemoveTransactionSubmissions(sdkCtx sdk.Context, txHash string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))
	tStore.Delete(types.KeyTransactionSubmissions(txHash))
}

func (k Keeper) TxHash(tx codec.ProtoMarshaler) common.Hash {
	return crypto.Keccak256Hash(k.cdc.MustMarshal(tx))
}

func (k Keeper) EpochSignaturesHash(epochSigs []types.EpochChainSignatures) common.Hash {
	bytes := make([]byte, 0)
	for _, sig := range epochSigs {
		bytes = append(bytes, k.cdc.MustMarshal(&sig)...)
	}

	return crypto.Keccak256Hash(bytes)
}

//-------------------SYSTEM TRANSACTIONS---------------------------------

func (k Keeper) SetSystemTransactionSubmissions(sdkCtx sdk.Context, txSubmissions *types.Submissions) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreSystemTransactionSubmissionsPrefix))
	tStore.Set(types.KeyTransactionSubmissions(txSubmissions.Hash), k.cdc.MustMarshal(txSubmissions))
}

func (k Keeper) GetPaginatedSystemTransactionsSubmissions(sdkCtx sdk.Context, pagination *query.PageRequest) (
	[]types.Submissions, *query.PageResponse, error) {

	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreSystemTransactionSubmissionsPrefix))

	var txsWithSubmissions []types.Submissions
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var txSubmissions types.Submissions
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

func (k Keeper) GetSystemTransactionSubmissions(sdkCtx sdk.Context, txHash string) (txSubmissions types.Submissions, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreSystemTransactionSubmissionsPrefix))

	bz := tStore.Get(types.KeyTransactionSubmissions(txHash))
	if bz == nil {
		return txSubmissions, false
	}

	k.cdc.MustUnmarshal(bz, &txSubmissions)

	return txSubmissions, true
}

func (k Keeper) RemoveSystemTransactionSubmissions(sdkCtx sdk.Context, txHash string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreSystemTransactionSubmissionsPrefix))
	tStore.Delete(types.KeyTransactionSubmissions(txHash))
}
