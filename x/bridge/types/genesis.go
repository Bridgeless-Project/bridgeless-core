package types

import (
	errorsmod "cosmossdk.io/errors"
	bridgeTypes "github.com/hyle-team/bridgeless-core/v12/types"
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

	chains := make(map[string]struct{})
	for _, chain := range gs.Chains {
		if _, ok := chains[chain.Id]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate chain id: %s", chain.Id)
		} else {
			chains[chain.Id] = struct{}{}
		}

		if err := validateChain(&chain); err != nil {
			return errorsmod.Wrapf(err, "invalid chain %s", chain.Id)
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
		if _, ok := txsSubmissions[txSubmissions.TxHash]; ok {
			return errorsmod.Wrapf(bridgeTypes.ErrDuplicatedValue, "duplicate tx hash: %v", txSubmissions.TxHash)
		}

		if err := validateTransactionSubmissions(&txSubmissions); err != nil {
			return errorsmod.Wrapf(err, "invalid tx submissions %v", txSubmissions.TxHash)
		}
	}

	return nil
}
