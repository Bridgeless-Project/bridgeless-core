package v4

import (
	"fmt"
	"math/big"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.30 %s module migrations", types.ModuleName))

	skipCounter := 0
	var epochId uint32 = 0
	commissions := getCommissions(ctx, storeKey, cdc, epochId)

	trStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTransactionPrefix))
	iterator := sdk.KVStorePrefixIterator(trStore, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var transaction types.Transaction
		cdc.MustUnmarshal(iterator.Value(), &transaction)

		token, found := getTokenInfo(ctx, storeKey, cdc, transaction.WithdrawalChainId, transaction.WithdrawalToken)
		if !found {
			fmt.Println(fmt.Sprintf("Failed transaction %s %s %s", transaction.DepositTxHash, transaction.DepositChainId, transaction.DepositToken))
			fmt.Println("Deposit block: ", transaction.DepositBlock)
			skipCounter++
			continue
		}

		commsission, ok := commissions[token.TokenId]
		if !ok {
			commsission = types.Commission{
				TokenId: token.TokenId,
				Amount:  "0",
			}
		}

		commissionAmount, ok := new(big.Int).SetString(commsission.Amount, 10)
		if !ok {
			skipCounter++
			continue
		}

		trComAmount, ok := new(big.Int).SetString(transaction.CommissionAmount, 10)
		if !ok {
			skipCounter++
			continue
		}

		// convert native decimals to 18
		trComAmount = TransformAmount(trComAmount, token.Decimals, types.DefaultChainDecimals)

		commissionAmount = commissionAmount.Add(commissionAmount, trComAmount)
		commsission.Amount = commissionAmount.String()

		commissions[token.TokenId] = commsission
	}

	for _, commission := range commissions {
		setCommission(ctx, storeKey, cdc, epochId, commission)
	}

	fmt.Println("counter: ", skipCounter)
	return nil
}

func getCommissions(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, epochId uint32) map[uint64]types.Commission {
	commissions := make(map[uint64]types.Commission)
	cStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	iterator := eStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var commission types.Commission
		cdc.MustUnmarshal(iterator.Value(), &commission)

		commissions[commission.TokenId] = commission
	}

	return commissions
}

func getTokenInfo(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, chain, address string) (tokenInfo types.TokenInfo, found bool) {
	tStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenInfoPrefix))
	bz := tStore.Get(types.KeyTokenInfo(chain, address))
	if bz == nil {
		return
	}

	cdc.MustUnmarshal(bz, &tokenInfo)
	found = true

	return
}

func setCommission(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, epochId uint32, commission types.Commission) {
	cStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreCommissionPrefix))
	eStore := prefix.NewStore(cStore, types.KeyEpoch(epochId))

	eStore.Set(types.KeyEpochCommission(epochId, commission.TokenId), cdc.MustMarshal(&commission))
}

func TransformAmount(amount *big.Int, currentDecimals uint64, targetDecimals uint64) *big.Int {
	result, _ := new(big.Int).SetString(amount.String(), 10)

	if currentDecimals == targetDecimals {
		return result
	}

	if currentDecimals < targetDecimals {
		for i := uint64(0); i < targetDecimals-currentDecimals; i++ {
			result.Mul(result, new(big.Int).SetInt64(10))
		}
	} else {
		for i := uint64(0); i < currentDecimals-targetDecimals; i++ {
			result.Div(result, new(big.Int).SetInt64(10))
		}
	}

	return result
}
