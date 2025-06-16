package v4

import (
	"fmt"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.16-rc1 %s module migrations", types.ModuleName))

	txStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTransactionPrefix))
	iterator := sdk.KVStorePrefixIterator(txStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {

		// get oldTx data without deposit and withdrawal amount fields
		var oldTx transaction
		cdc.MustUnmarshal(iterator.Value(), &oldTx)

		// set the values of new fields
		newTx := types.Transaction{
			DepositChainId:    oldTx.DepositChainId,
			DepositTxHash:     oldTx.DepositTxHash,
			DepositTxIndex:    oldTx.DepositTxIndex,
			DepositBlock:      oldTx.DepositBlock,
			DepositToken:      oldTx.DepositToken,
			DepositAmount:     oldTx.Amount,
			Depositor:         oldTx.Depositor,
			Receiver:          oldTx.Receiver,
			WithdrawalChainId: oldTx.WithdrawalChainId,
			WithdrawalTxHash:  oldTx.WithdrawalTxHash,
			WithdrawalToken:   oldTx.WithdrawalToken,
			Signature:         oldTx.Signature,
			IsWrapped:         oldTx.IsWrapped,
			WithdrawalAmount:  oldTx.Amount,
		}

		// set new transaction instead of old one
		txStore.Set(
			types.KeyTransaction(types.TransactionId(&newTx)),
			cdc.MustMarshal(&newTx),
		)
	}

	return nil
}
