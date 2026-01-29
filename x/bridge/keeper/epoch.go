package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) SetEpoch(sdkCtx sdk.Context, epoch *types.Epoch) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochPrefix))
	tStore.Set(types.KeyEpoch(epoch.Id), k.cdc.MustMarshal(epoch))
}

func (k Keeper) GetPaginatedEpochs(sdkCtx sdk.Context, pagination *query.PageRequest) ([]types.Epoch, *query.PageResponse, error) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochPrefix))

	var epochs []types.Epoch
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var epoch types.Epoch
		if err := k.cdc.Unmarshal(value, &epoch); err != nil {
			return err
		}
		epochs = append(epochs, epoch)
		return nil
	})
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to get paginated epoch")
	}

	return epochs, pageRes, nil
}

func (k Keeper) GetEpoch(sdkCtx sdk.Context, epochId uint32) (epoch types.Epoch, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochPrefix))

	bz := tStore.Get(types.KeyEpoch(epochId))
	if bz == nil {
		return epoch, false
	}

	k.cdc.MustUnmarshal(bz, &epoch)
	return epoch, true
}

func (k Keeper) RemoveEpoch(sdkCtx sdk.Context, epochId uint32) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochPrefix))
	tStore.Delete(types.KeyEpoch(epochId))
}

// ------------------- Epoch Signature ------------------
func (k Keeper) SetEpochChainSignature(sdkCtx sdk.Context, epochSig *types.EpochChainSignatures) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochChainSignaturePrefix))
	tStore.Set(types.KeyEpochChainSignature(epochSig.ChainId, epochSig.EpochId), k.cdc.MustMarshal(epochSig))
}

func (k Keeper) GetEpochChainSignature(sdkCtx sdk.Context, epochId uint32, chainId string) (epochChainSignatures types.EpochChainSignatures, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochChainSignaturePrefix))
	bz := tStore.Get(types.KeyEpochChainSignature(chainId, epochId))
	if bz == nil {
		return epochChainSignatures, false
	}

	k.cdc.MustUnmarshal(bz, &epochChainSignatures)
	return epochChainSignatures, true
}

// --------------- Transactions ---------------------
func (k Keeper) SetEpochTransaction(sdkCtx sdk.Context, chainId uint32, epochTransaction types.TransactionIdentifier) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochTransactionPrefix))
	tStore.Set(types.KeyEpochTransaction(chainId, epochTransaction.DepositTxIndex, epochTransaction.DepositTxHash, epochTransaction.DepositTxHash), k.cdc.MustMarshal(&epochTransaction))
}

func (k Keeper) RemoveEpochTransaction(sdkCtx sdk.Context, chainId uint32, epochTransaction types.TransactionIdentifier) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochTransactionPrefix))
	tStore.Delete(types.KeyEpochTransaction(chainId, epochTransaction.DepositTxIndex, epochTransaction.DepositTxHash, epochTransaction.DepositTxHash))
}

func (k Keeper) GetPaginatedEpochTransactions(sdkCtx sdk.Context, pagination *query.PageRequest) ([]types.TransactionIdentifier, *query.PageResponse, error) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreEpochTransactionPrefix))

	var transactionIds []types.TransactionIdentifier
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var tx types.TransactionIdentifier
		if err := k.cdc.Unmarshal(value, &tx); err != nil {
			return err
		}
		transactionIds = append(transactionIds, tx)
		return nil
	})
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "failed to get paginated tx ids for epoch")
	}

	return transactionIds, pageRes, nil
}
