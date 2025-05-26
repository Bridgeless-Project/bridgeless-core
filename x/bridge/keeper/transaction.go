package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (k Keeper) SetTransaction(sdkCtx sdk.Context, transaction types.Transaction) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))
	tStore.Set(types.KeyTransaction(types.TransactionId(&transaction)), k.cdc.MustMarshal(&transaction))

}

func (k Keeper) GetTransaction(sdkCtx sdk.Context, id string) (types.Transaction, bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))

	var transaction types.Transaction
	bz := tStore.Get(types.KeyTransaction(id))
	if bz == nil {
		return transaction, false
	}

	k.cdc.MustUnmarshal(bz, &transaction)
	return transaction, true
}

func (k Keeper) GetPaginatedTransactions(
	sdkCtx sdk.Context, pagination *query.PageRequest,
) (
	[]types.Transaction, *query.PageResponse, error,
) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))

	var transactions []types.Transaction
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var transaction types.Transaction
		if err := k.cdc.Unmarshal(value, &transaction); err != nil {
			return err
		}
		transactions = append(transactions, transaction)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return transactions, pageRes, nil
}

func (k Keeper) SubmitTx(ctx sdk.Context, transaction *types.Transaction, submitter string) error {
	// Check whether tx has enough submissions to be added to core
	txSubmissions, found := k.GetTransactionSubmissions(ctx, k.TxHash(transaction).String())

	threshold := k.GetParams(ctx).TssThreshold
	if !found {
		txSubmissions.TxHash = k.TxHash(transaction).String()
	}

	// If tx has been submitted before with the same address new submission is rejected
	if isSubmitter(txSubmissions.Submitters, submitter) {
		return errorsmod.Wrap(types.ErrTranscationAlreadySubmitted,
			"transaction has been already submitted by this address")
	}

	txSubmissions.Submitters = append(txSubmissions.Submitters, submitter)
	k.SetTransactionSubmissions(ctx, &txSubmissions)

	// If tx has not been submitted yet or has not enough submissions (less than tss threshold param)
	// it is not set to core
	if len(txSubmissions.Submitters) == int(threshold+1) {
		k.SetTransaction(ctx, *transaction)
		emitSubmitEvent(ctx, *transaction)
	}

	return nil
}

func isSubmitter(submitters []string, submitter string) bool {
	for _, s := range submitters {
		if submitter == s {
			return true
		}
	}

	return false
}
