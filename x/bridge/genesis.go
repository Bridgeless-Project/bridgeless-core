package bridge

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/pkg/errors"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, chain := range genState.Chains {
		k.SetChain(ctx, chain)
		k.SetChainByType(ctx, chain)
	}
	for _, token := range genState.Tokens {
		k.SetToken(ctx, token)
		for _, info := range token.Info {
			k.SetTokenInfo(ctx, info)
			k.SetTokenPairs(ctx, info, token.Info...)
		}
	}
	for _, tx := range genState.Transactions {
		k.SetTransaction(ctx, tx)
		if tx.EpochId != 0 {
			k.SetEpochTransaction(ctx, tx.EpochId, types.TransactionIdentifier{
				DepositTxHash:  tx.DepositTxHash,
				DepositTxIndex: tx.DepositTxIndex,
				DepositChainId: tx.DepositChainId,
			})
		}
	}

	for _, txSubmissions := range genState.TransactionsSubmissions {
		k.SetTransactionSubmissions(ctx, &txSubmissions)
	}

	for _, referral := range genState.Referrals {
		k.AddReferral(ctx, referral)
	}

	for _, referralRewards := range genState.ReferralsRewards {
		k.InsertReferralRewards(ctx, referralRewards.ReferralId, referralRewards.TokenId, referralRewards)
	}

	if err := genState.Validate(); err != nil {
		panic(errors.Wrap(err, "invalid genesis state"))
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	txs, _, err := k.GetPaginatedTransactions(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis transactions"))
	}

	tokens, _, err := k.GetTokensWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis tokens"))
	}

	chains, _, err := k.GetChainsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis chains"))
	}

	txsWithSubmissions, _, err := k.GetPaginatedTransactionsSubmissions(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to export genesis transaction submissions"))
	}

	referrals := k.GetAllReferrals(ctx)
	referralsRewards := k.GetAllReferralRewards(ctx)

	return &types.GenesisState{
		Params:                  k.GetParams(ctx),
		Chains:                  chains,
		Tokens:                  tokens,
		Transactions:            txs,
		TransactionsSubmissions: txsWithSubmissions,
		Referrals:               referrals,
		ReferralsRewards:        referralsRewards,
	}
}
