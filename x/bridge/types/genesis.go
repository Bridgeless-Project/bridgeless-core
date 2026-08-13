package types

import (
	"fmt"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	bridgeTypes "github.com/Bridgeless-Project/bridgeless-core/v12/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:       DefaultParams(),
		Chains:       []Chain{},
		Tokens:       []Token{},
		Transactions: []Transaction{},
		Epochs:       []Epoch{},
		Commissions:  []GenesisCommission{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if err := gs.Params.Validate(); err != nil {
		return errorsmod.Wrap(err, "invalid params")
	}

	for _, tx := range gs.Transactions {
		if err := validateTransaction(&tx); err != nil {
			return errorsmod.Wrapf(err, "invalid transaction %s", TransactionId(&tx))
		}
	}

	chains := make(map[uint32]struct{})
	for _, chain := range gs.Chains {
		if _, ok := chains[chain.Id]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate chain id: %d", chain.Id)
		} else {
			chains[chain.Id] = struct{}{}
		}

		if err := validateChain(&chain); err != nil {
			return errorsmod.Wrapf(err, "invalid chain %d", chain.Id)
		}
	}

	uniqueTokens := make(map[uint64]struct{})
	for _, token := range gs.Tokens {
		if _, ok := uniqueTokens[token.Id]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate token id: %v", token.Id)
		} else {
			uniqueTokens[token.Id] = struct{}{}
		}

		if err := validateToken(&token); err != nil {
			return errorsmod.Wrapf(err, "invalid token %v", token.Id)
		}

		for _, info := range token.Info {
			if err := validateTokenInfo(&info, nil); err != nil {
				return errorsmod.Wrapf(err, "invalid token info for token %v", token.Id)
			}
		}
	}

	txsSubmissions := make(map[string]struct{})
	for _, txSubmissions := range gs.TransactionsSubmissions {
		if _, ok := txsSubmissions[txSubmissions.Hash]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate tx hash: %v", txSubmissions.Hash)
		}
		txsSubmissions[txSubmissions.Hash] = struct{}{}

		if err := validateTransactionSubmissions(&txSubmissions); err != nil {
			return errorsmod.Wrapf(err, "invalid tx submissions %v", txSubmissions.Hash)
		}
	}

	epochs := make(map[uint32]struct{})
	for _, epoch := range gs.Epochs {
		if _, ok := epochs[epoch.Id]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate epoch id: %d", epoch.Id)
		} else {
			epochs[epoch.Id] = struct{}{}
		}

		if err := validateEpoch(&epoch); err != nil {
			return errorsmod.Wrapf(err, "invalid epoch %d", epoch.Id)
		}
	}

	commissions := make(map[string]struct{})
	for _, entry := range gs.Commissions {
		key := fmt.Sprintf("%d/%d", entry.EpochId, entry.Commission.TokenId)
		if _, ok := commissions[key]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate commission: %s", key)
		}
		commissions[key] = struct{}{}

		amount, ok := new(big.Int).SetString(entry.Commission.Amount, 10)
		if !ok || amount.Sign() < 0 {
			return errorsmod.Wrapf(ErrInvalidCommission, "invalid commission amount %q for %s", entry.Commission.Amount, key)
		}
	}

	return nil
}
