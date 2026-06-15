package keeper

import (
	"errors"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
	contractstypes "github.com/Bridgeless-Project/bridgeless-core/v12/contracts/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/utils"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	contractEventWithdrawn = "WithdrawnERC20"
)

const transferMethod = "transfer"

// PostTxProcessing listens for configured bridge contract withdrawal events and
// distributes the corresponding stored system withdrawal fees.
func (k Keeper) PostTxProcessing(ctx sdk.Context, _ core.Message, receipt *ethtypes.Receipt) error {
	if receipt == nil || len(receipt.Logs) == 0 {
		k.Logger(ctx).Error("PostTxProcessing receipt is nil or empty")
		return nil
	}

	// Getting this address from params preventing panic during EVM flow
	params := k.GetParams(ctx)
	contractAddress := common.HexToAddress(params.BridgeAddress)
	for _, evmLog := range receipt.Logs {
		if evmLog == nil || evmLog.Address != contractAddress || len(evmLog.Topics) == 0 {
			k.Logger(ctx).Debug("skipping event with empty topics")
			continue
		}

		event, err := contracts.BridgeContract.ABI.EventByID(evmLog.Topics[0])
		if err != nil {
			k.Logger(ctx).Error(errorsmod.Wrap(err, "failed to resolve bridge contract event").Error())
			continue
		}

		if event.Name != contractEventWithdrawn {
			k.Logger(ctx).Info("skipping event:", "got", event.Name, "want", contractEventWithdrawn)
			continue
		}

		eventBody := contractstypes.BridgeWithdrawnERC20{}
		if err = utils.UnpackLog(contracts.BridgeContract.ABI, &eventBody, event.Name, evmLog); err != nil {
			k.Logger(ctx).Error("failed to unpack event body")
			continue
		}

		withdrawal, found := k.GetSystemTransaction(ctx, ConstructSystemTxHash(eventBody.Amount, eventBody.Token.Bytes(), eventBody.Receiver.Bytes()))
		if !found {
			k.Logger(ctx).Info(
				"system withdrawal for EVM withdrawal log not found",
				"tx_hash", evmLog.TxHash.Hex(),
				"tx_index", evmLog.TxIndex,
				"log_index", evmLog.Index,
			)
			continue
		}

		if len(withdrawal.Result) != 0 {
			k.Logger(ctx).Info("already have result for EVM withdrawal log")
			continue
		}

		if err = k.FeeDistribute(ctx, withdrawal, eventBody.Token); err != nil {
			return errorsmod.Wrap(err, "failed to distribute system withdrawal fees")
		}
	}

	return nil
}

func (k Keeper) FeeDistribute(ctx sdk.Context, withdrawal types.SystemWithdrawal, tokenAddress common.Address) error {
	k.Logger(ctx).Info("start fee distribution", "remaining", withdrawal.Amount)
	remaining, ok := new(big.Int).SetString(withdrawal.Amount, 10)
	if !ok {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "amount %s", withdrawal.Amount)
	}
	if remaining.Sign() < 0 {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "negative remaining:  %s", remaining.String())
	}

	initialAmount := new(big.Int).Set(remaining)
	tokenInfo, found := k.GetTokenInfo(ctx, utils.GetChainId(ctx), tokenAddress.Hex())
	if !found {
		return errorsmod.Wrap(types.ErrTokenInfoNotFound, "fee distribution token not found")
	}

	epoch, found := k.GetEpoch(ctx, withdrawal.EpochId)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, "epoch %d not found", withdrawal.EpochId)
	}
	if len(epoch.Parties) == 0 {
		return errorsmod.Wrap(types.ErrInvalidPartiesList, "epoch has no parties")
	}

	// TODO: integrate referral withdrawal
	//for _, referralRewards := range withdrawal.ReferralRewards {
	//	// returns the referralsRewardAmount with native decimals
	//	_, txhash, address, err := k.distributeReferralReward(ctx, referralRewards, tokenAddress)
	//	if err != nil {
	//		return errorsmod.Wrap(err, "failed to distribute referral rewards")
	//	}
	//
	//	results = append(results, types.TxResult{
	//		TxHash:     txhash,
	//		Address:    address,
	//		ReferralId: referralRewards.ReferralId,
	//	})
	// remaining.Sub(remaining, referralsRewardAmount) // total amount - referral rewards
	//}
	//if remaining.Sign() < 0 {
	//	return errorsmod.Wrapf(types.ErrInvalidAmount, "referral rewards are negative: %s", remaining.String())
	//}

	if remaining.Sign() == 0 {
		k.Logger(ctx).Info("skipping fee distribute: remaining is zero")
		return nil
	}

	share := new(big.Int).Div(remaining, big.NewInt(int64(len(epoch.Parties))))
	if share.Sign() == 0 {
		k.Logger(ctx).Info("skipping fee distribute: share is zero")
		return nil
	}

	// if one of txs  will fail - all transaction will fail
	results, err := k.distributeTokensBetweenParties(ctx, share, tokenAddress, epoch, remaining)
	if err != nil {
		k.Logger(ctx).Error("filed to distribute fees")
		return errorsmod.Wrap(err, "failed to distribute fees")
	}

	withdrawal.Result = results
	distributedAmount := new(big.Int).Sub(initialAmount, remaining)
	if err := k.SubtractCommissionNative(ctx, withdrawal.EpochId, tokenInfo, distributedAmount); err != nil {
		return errorsmod.Wrap(err, "failed to subtract distributed commission")
	}

	k.Logger(ctx).Info("saving results", "results", results)
	k.SetSystemTransaction(
		ctx,
		withdrawal,
	)

	return nil
}

func (k Keeper) distributeReferralReward(ctx sdk.Context, rewards types.ReferralRewards, tokenAddress common.Address) (*big.Int, string, string, error) {
	if types.IsDefaultReferralId(rewards.ReferralId) || rewards.ToClaim == "" {
		return big.NewInt(0), "", "", nil
	}

	if err := k.validateToken(ctx, rewards.TokenId, tokenAddress); err != nil {
		return big.NewInt(0), "", "", errorsmod.Wrap(err, "failed to validate token")
	}

	referral, found := k.GetReferral(ctx, rewards.ReferralId)
	if !found {
		return nil, "", "", errorsmod.Wrapf(types.ErrReferralNotFound, "referral %d not found", rewards.ReferralId)
	}

	rewardsFromStore, ok := k.GetReferralRewards(ctx, rewards.ReferralId, rewards.TokenId)
	if !ok {
		return nil, "", "", errors.New("referral reward not found")
	}

	toClaimFromStore, ok := new(big.Int).SetString(rewardsFromStore.ToClaim, 10)
	if !ok {
		return nil, "", "", errors.New("invalid referral reward claim from store")
	}

	claimedRewardsFromStore, ok := new(big.Int).SetString(rewardsFromStore.TotalClaimedAmount, 10)
	if !ok {
		return nil, "", "", errors.New("invalid referral reward claim from store")
	}

	toClaim, ok := new(big.Int).SetString(rewards.ToClaim, 10)
	if !ok {
		return nil, "", "", errors.New("invalid referral claimed rewards")
	}

	token, err := k.tokenOnBridgeless(ctx, rewards.TokenId)
	if err != nil {
		return nil, "", "", errorsmod.Wrap(err, "failed to get token on bridgeless")
	}

	// convert toClaim to 18 decimals to successfully sub from toClaimFromStore
	toClaimDecimals18 := TransformAmount(toClaim, token.Decimals, types.DefaultChainDecimals)

	if toClaimFromStore.Cmp(toClaimDecimals18) == -1 {
		return nil, "", "", errors.New("not enough referral reward to claim")
	}

	withdrawalAddress, err := sdk.AccAddressFromBech32(referral.WithdrawalAddress)
	if err != nil {
		return nil, "", "", errorsmod.Wrapf(err, "invalid referral withdrawal address %s", referral.WithdrawalAddress)
	}

	evmAddress, err := utils.CosmosToEVM(withdrawalAddress)
	if err != nil {
		return nil, "", "", errorsmod.Wrap(err, "failed to convert address")
	}

	txhash, err := k.sendTokens(ctx, toClaim, tokenAddress, evmAddress)
	if err != nil {
		return nil, "", "", errorsmod.Wrap(err, "failed to send tokens")
	}

	change := toClaimFromStore.Sub(toClaimFromStore, toClaimDecimals18)

	//both amounts here with decimals 18
	rewards.ToClaim = change.String()
	rewards.TotalClaimedAmount = claimedRewardsFromStore.Add(claimedRewardsFromStore, toClaimDecimals18).String()

	k.InsertReferralRewards(ctx, referral.Id, rewards.TokenId, rewards)

	// toClaim here has native decimals
	return toClaim, txhash, evmAddress.String(), nil
}

func (k Keeper) sendTokens(ctx sdk.Context, amount *big.Int, token common.Address, receiver common.Address) (string, error) {
	tx, err := k.erc20.CallEVMAsTx(
		ctx,
		contracts.ERC20BurnableContract.ABI,
		types.ModuleAddress,
		token,
		true,
		transferMethod,
		receiver,
		amount,
	)
	if err != nil {
		return "", errorsmod.Wrapf(err, "failed to call erc20 burn")
	}
	return tx.Hash, nil
}

func (k Keeper) validateToken(ctx sdk.Context, tokenId uint64, tokenAddress common.Address) error {
	rewardTokenInfo, err := k.tokenOnBridgeless(ctx, tokenId)
	if err != nil {
		return errorsmod.Wrap(err, "failed to resolve token")
	}
	if !common.IsHexAddress(rewardTokenInfo.Address) {
		return errorsmod.Wrapf(types.ErrTokenInfoNotFound,
			"invalid bridgeless token address %s for token ID %d", rewardTokenInfo.Address, tokenId)
	}
	rewardTokenAddress := common.HexToAddress(rewardTokenInfo.Address)
	if rewardTokenAddress != tokenAddress {
		return errorsmod.Wrapf(types.ErrInvalidDataType,
			"token ID %d resolves to %s, but withdrawal token is %s",
			tokenId, rewardTokenAddress.Hex(), tokenAddress.Hex())
	}

	return nil
}

func (k Keeper) distributeTokensBetweenParties(
	ctx sdk.Context,
	share *big.Int,
	tokenAddress common.Address,
	epoch types.Epoch,
	remaining *big.Int,
) ([]types.TxResult, error) {
	results := make([]types.TxResult, 0)

	for _, party := range epoch.Parties {
		partyAddress, err := sdk.AccAddressFromBech32(party.Address)
		if err != nil {
			return nil, errorsmod.Wrapf(err, "invalid party address %s", party.Address)
		}

		evmAddress, err := utils.CosmosToEVM(partyAddress)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to convert address")
		}

		txhash, err := k.sendTokens(ctx, share, tokenAddress, evmAddress)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to send tokens")
		}

		//  we should compute `remaining - share` to avoid double distribution
		remaining.Sub(remaining, share)
		results = append(results, types.TxResult{
			TxHash:     txhash,
			Address:    evmAddress.String(),
			ReferralId: 0,
		})

		k.Logger(ctx).Info("fee distribute", "share", share.String(), "address", evmAddress.String())
	}

	return results, nil
}
