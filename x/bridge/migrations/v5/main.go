package v5

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.17-rc1 %s module migrations", types.ModuleName))

	if err := migrateTxs(ctx, storeKey, cdc); err != nil {
		return errorsmod.Wrap(err, "failed to migrate transactions for v12.1.17-rc1 update")
	}

	if err := migrateTokens(ctx, storeKey, cdc); err != nil {
		return errorsmod.Wrap(err, "failed to migrate tokens for v12.1.17-rc1 update")
	}

	return nil
}

func migrateTxs(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	txStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTransactionPrefix))
	iterator := sdk.KVStorePrefixIterator(txStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {

		// get oldToken data without tx_data field
		var oldToken transaction1
		cdc.MustUnmarshal(iterator.Value(), &oldToken)

		// set the values of new fields
		newTx := types.Transaction{
			DepositChainId:    oldToken.DepositChainId,
			DepositTxHash:     oldToken.DepositTxHash,
			DepositTxIndex:    oldToken.DepositTxIndex,
			DepositBlock:      oldToken.DepositBlock,
			DepositToken:      oldToken.DepositToken,
			DepositAmount:     oldToken.DepositAmount,
			Depositor:         oldToken.Depositor,
			Receiver:          oldToken.Receiver,
			WithdrawalChainId: oldToken.WithdrawalChainId,
			WithdrawalTxHash:  oldToken.WithdrawalTxHash,
			WithdrawalToken:   oldToken.WithdrawalToken,
			Signature:         oldToken.Signature,
			IsWrapped:         oldToken.IsWrapped,
			WithdrawalAmount:  oldToken.DepositAmount,
			CommissionAmount:  "",
			TxData:            "",
		}

		// set new transaction instead of old one
		txStore.Set(
			types.KeyTransaction(types.TransactionId(&newTx)),
			cdc.MustMarshal(&newTx),
		)
	}

	return nil
}

func migrateTokens(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	tokenStore := prefix.NewStore(ctx.KVStore(storeKey), types.Prefix(types.StoreTokenPrefix))
	iterator := sdk.KVStorePrefixIterator(tokenStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {

		// get oldToken data without commission_rate field
		var oldToken token
		cdc.MustUnmarshal(iterator.Value(), &oldToken)

		// set the values of new fields
		newToken := types.Token{
			Id: oldToken.Id,
			Metadata: types.TokenMetadata{
				Name:   oldToken.Metadata.Name,
				Symbol: oldToken.Metadata.Symbol,
				Uri:    oldToken.Metadata.Uri,
			},
			Info: transformTokenInfo(oldToken.Info),
			//CommissionRate: "0.01",
		}

		// set new token instead of old one
		tokenStore.Set(
			types.KeyToken(newToken.Id),
			cdc.MustMarshal(&newToken),
		)
	}
	return nil
}

func transformTokenInfo(oldInfo []tokenInfo) (newInfo []types.TokenInfo) {
	for _, info := range oldInfo {
		newInfo = append(newInfo, types.TokenInfo{
			Address:   info.Address,
			Decimals:  info.Decimals,
			ChainId:   info.ChainId,
			TokenId:   info.TokenId,
			IsWrapped: info.IsWrapped,
		})
	}
	return
}
