package keeper

import (
	"math/big"
	"reflect"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
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

func (k Keeper) RemoveTransaction(sdkCtx sdk.Context, id string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))
	tStore.Delete(types.KeyTransaction(id))
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
	threshold := k.GetParams(ctx).TssThreshold
	txSubmissions, found := k.GetTransactionSubmissions(ctx, k.TxHash(transaction).String())
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
	if len(txSubmissions.Submitters) != int(threshold+1) {
		return nil
	}

	k.SetTransaction(ctx, *transaction)
	emitSubmitEvent(ctx, *transaction)

	if types.IsDefaultReferralId(transaction.ReferralId) {
		return nil
	}

	referral, ok := k.GetReferral(ctx, transaction.ReferralId)
	if !ok {
		return errorsmod.Wrap(types.ErrReferralNotFound, "referral ID not found")
	}

	token, ok := k.GetTokenInfo(ctx, transaction.DepositChainId, transaction.DepositToken)
	if !ok {
		return errorsmod.Wrap(types.ErrTokenInfoNotFound, "token info not found for deposit token")
	}

	// Rewards for referral are taken from CommissionAmount
	commissionAmount, ok := big.NewInt(0).SetString(transaction.CommissionAmount, 10)
	if !ok {
		return errorsmod.Wrap(types.ErrInvalidDataType, "invalid withdrawal amount")
	}

	rewards, err := types.GetCommissionAmount(commissionAmount, referral.CommissionRate)
	if err != nil {
		return errorsmod.Wrap(err, "failed to calculate referral rewards")
	}

	referralRewards := types.ReferralRewards{
		ReferralId:         transaction.ReferralId,
		TokenId:            token.TokenId,
		ToClaim:            sdk.NewIntFromBigInt(rewards).String(),
		TotalClaimedAmount: sdk.NewInt(0).String(), // not used when adding referral rewards and should be 0
	}

	err = k.AddReferralRewards(ctx, transaction.ReferralId, token.TokenId, referralRewards)
	if err != nil {
		return errorsmod.Wrap(err, "failed to add referral rewards")
	}

	return nil
}

func (k Keeper) DeleteTx(ctx sdk.Context, depositTxHash string, depositTxIndex uint64, depositChainId string) error {
	txId := types.TransactionId(&types.Transaction{DepositTxHash: depositTxHash, DepositTxIndex: depositTxIndex, DepositChainId: depositChainId})
	transaction, ok := k.GetTransaction(ctx, txId)
	if !ok {
		return errorsmod.Wrap(types.ErrTransactionNotFound, "failed to get transaction")
	}

	// Delete tx from core
	k.RemoveTransaction(ctx, txId)

	// Delete tx submissions
	txSubmissions, found := k.GetTransactionSubmissions(ctx, k.TxHash(&transaction).String())
	if found {
		k.RemoveTransactionSubmissions(ctx, txSubmissions.TxHash)
	}

	// Minus referral rewards
	// if referral ID is default no need to minus rewards, just emit event and return
	if types.IsDefaultReferralId(transaction.ReferralId) {
		emitRemoveTransactionEvent(ctx, transaction)
		return nil
	}

	// If referral ID is not default minus rewards
	referral, ok := k.GetReferral(ctx, transaction.ReferralId)
	if !ok {
		return errorsmod.Wrap(types.ErrReferralNotFound, "referral ID not found")
	}

	token, ok := k.GetTokenInfo(ctx, transaction.DepositChainId, transaction.DepositToken)
	if !ok {
		return errorsmod.Wrap(types.ErrTokenInfoNotFound, "token info not found for deposit token")
	}

	// Rewards for referral are taken from CommissionAmount
	commissionAmount, ok := big.NewInt(0).SetString(transaction.CommissionAmount, 10)
	if !ok {
		return errorsmod.Wrap(types.ErrInvalidDataType, "invalid withdrawal amount")
	}

	rewards, err := types.GetCommissionAmount(commissionAmount, referral.CommissionRate)
	if err != nil {
		return errorsmod.Wrap(err, "failed to calculate referral rewards")
	}

	// convert rewards to negative value to minus it
	referralRewards := types.ReferralRewards{
		ReferralId:         transaction.ReferralId,
		TokenId:            token.TokenId,
		ToClaim:            sdk.NewIntFromBigInt(rewards).Neg().String(),
		TotalClaimedAmount: sdk.NewInt(0).String(), // not used when adding referral rewards and should be 0
	}

	err = k.AddReferralRewards(ctx, transaction.ReferralId, token.TokenId, referralRewards)
	if err != nil {
		return errorsmod.Wrap(err, "failed to minus referral rewards")
	}

	emitRemoveTransactionEvent(ctx, transaction)
	return nil
}

func (k Keeper) UpdateTx(ctx sdk.Context, transaction *types.Transaction) error {
	// validate that tx already exists in the store
	oldtx, ok := k.GetTransaction(ctx, types.TransactionId(transaction))
	if !ok {
		return errorsmod.Wrap(types.ErrTransactionNotFound, "failed to get transaction")
	}

	// validate that txs are the same except WithdrawalTxHash
	if err := compareTxs(oldtx, *transaction); err != nil {
		return errorsmod.Wrap(err, "failed to compare transactions")
	}

	k.SetTransaction(ctx, *transaction)
	emitUpdateTransactionEvent(ctx, *transaction)

	return nil
}

func compareTxs(tx, tx2 types.Transaction) error {
	txValue := reflect.ValueOf(tx)
	tx2Value := reflect.ValueOf(tx2)

	txTypes := reflect.TypeOf(tx)

	if txValue.NumField() != tx2Value.NumField() {
		return errorsmod.Wrap(types.ErrInvalidDataType, "transactions have different number of fields")
	}

	for i := 0; i < txValue.NumField(); i++ {
		if txTypes.Field(i).Name == "WithdrawalTxHash" {
			continue
		}

		if txValue.Field(i).Interface() != tx2Value.Field(i).Interface() {
			return errorsmod.Wrapf(types.ErrInvalidDataType, "field %s is different: %v != %v", txTypes.Field(i).Name, txValue.Field(i).Interface(), tx2Value.Field(i).Interface())
		}
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
