package types

import (
	errorsmod "cosmossdk.io/errors"
	"errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func TransactionId(tx *Transaction) string {
	return fmt.Sprintf("%s/%v/%s", tx.DepositTxHash, tx.DepositTxIndex, tx.DepositChainId)
}

func validateTransaction(tx *Transaction) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}

	if _, set := big.NewInt(0).SetString(tx.DepositAmount, 10); !set {
		return errors.New(fmt.Sprintf("invalid deposit amount: %s", tx.DepositAmount))
	}

	if _, set := big.NewInt(0).SetString(tx.WithdrawalAmount, 10); !set {
		return errors.New(fmt.Sprintf("invalid withdrawal amount: %s", tx.WithdrawalAmount))
	}

	if _, set := big.NewFloat(0).SetString(tx.CommissionAmount); !set {
		return errors.New(fmt.Sprintf("invalid commission amount: %s", tx.CommissionAmount))
	}

	if len(tx.DepositChainId) == 0 || len(tx.WithdrawalChainId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deposit chain id and withdrawal id cannot be empty")
	}
	if tx.WithdrawalChainId == tx.DepositChainId {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "withdrawal chain id cannot be the same as deposit chain id ")
	}

	if tx.Receiver == "" {
		return errors.New("receiver cannot be empty")
	}

	if tx.DepositTxHash == "" {
		return errors.New("deposit tx hash cannot be empty")
	}

	return nil
}

func validateTransactionSubmissions(txSubmissions *TransactionSubmissions) error {
	if _, err := hexutil.Decode(txSubmissions.TxHash); err != nil {
		return errorsmod.Wrap(ErrInvalidTxHash, err.Error())
	}

	for _, submitter := range txSubmissions.Submitters {
		if !common.IsHexAddress(submitter) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid submitter: %s", submitter))
		}
	}

	return nil
}

// ValidateChainTransaction is used to validate transaction details depending on certain chain type
func ValidateChainTransaction(tx *Transaction, chainType ChainType) error {
	if tx == nil {
		return errors.New("transaction is nil")
	}
	if len(chainType.String()) == 0 {
		return errors.New("chain type is empty")
	}

	// Only ZANO chain can have empty depositor address
	if len(tx.Depositor) == 0 && chainType != ChainType_ZANO {
		return errors.New("depositor address can`t be empty, unless it`s ZANO chain")
	}

	// Place for custom validation logic
	switch chainType {
	case ChainType_EVM:
	case ChainType_ZANO:
	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_OTHER:
	}

	return nil
}
