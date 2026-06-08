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
		return nil
	}
	bridgelessChain, ok := k.GetChain(ctx, utils.GetChainId(ctx))
	if !ok {
		return errors.New("chain not found")
	}

	contractAddress := common.HexToAddress(bridgelessChain.BridgeAddress)
	for _, evmLog := range receipt.Logs {
		if evmLog == nil || evmLog.Address != contractAddress || len(evmLog.Topics) == 0 {
			continue
		}

		event, err := contracts.BridgeContract.ABI.EventByID(evmLog.Topics[0])
		if err != nil {
			return errorsmod.Wrap(err, "failed to resolve bridge contract event")
		}

		if event.Name != contractEventWithdrawn {
			continue
		}

		eventBody := contractstypes.BridgeWithdrawnERC20{}
		if err = utils.UnpackLog(contracts.BridgeContract.ABI, &eventBody, event.Name, evmLog); err != nil {
			k.Logger(ctx).Info("failed to unpack event body")
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
			k.Logger(ctx).Debug("already have result for EVM withdrawal log")
			continue
		}

		if err = k.FeeDistribute(ctx, withdrawal, eventBody.Token); err != nil {
			return errorsmod.Wrap(err, "failed to distribute system withdrawal fees")
		}
	}

	return nil
}

func (k Keeper) FeeDistribute(ctx sdk.Context, withdrawal types.SystemWithdrawal, tokenAddress common.Address) error {
	remaining, ok := new(big.Int).SetString(withdrawal.Amount, 10)
	if !ok {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "amount %s", withdrawal.Amount)
	}

	// Update commission
	defer func() {
		tokenInfo, found := k.GetTokenInfo(ctx, utils.GetChainId(ctx), tokenAddress.Hex())
		if !found {
			k.Logger(ctx).Error("token info not found")
			return
		}

		k.SetCommission(ctx, withdrawal.EpochId, types.Commission{
			TokenId: tokenInfo.TokenId,
			Amount:  remaining.String(),
		})
	}()

	epoch, found := k.GetEpoch(ctx, withdrawal.EpochId)
	if !found {
		return errorsmod.Wrapf(types.ErrEpochNotFound, "epoch %d not found", withdrawal.EpochId)
	}
	if len(epoch.Parties) == 0 {
		return errorsmod.Wrap(types.ErrInvalidPartiesList, "epoch has no parties")
	}

	results := make([]types.TxResult, 0)
	for _, referralRewards := range withdrawal.ReferralRewards {
		referralsRewardAmount, txhash, address, err := k.distributeReferralReward(ctx, referralRewards, tokenAddress)
		if err != nil {
			return errorsmod.Wrap(err, "failed to distribute referral rewards")
		}

		results = append(results, types.TxResult{
			TxHash:     txhash,
			Address:    address,
			ReferralId: referralRewards.ReferralId,
		})

		remaining.Sub(remaining, referralsRewardAmount) // total amount - referral rewards
	}
	if remaining.Sign() < 0 {
		return errorsmod.Wrapf(types.ErrInvalidAmount, "referral rewards are negative: %s", remaining.String())
	}

	if remaining.Sign() == 0 {
		return nil
	}

	share := new(big.Int).Div(remaining, big.NewInt(int64(len(epoch.Parties))))
	if share.Sign() == 0 {
		return nil
	}

	for _, party := range epoch.Parties {
		partyAddress, err := sdk.AccAddressFromBech32(party.Address)
		if err != nil {
			return errorsmod.Wrapf(err, "invalid party address %s", party.Address)
		}

		evmAddress, err := utils.CosmosToEVM(partyAddress)
		if err != nil {
			return errorsmod.Wrap(err, "failed to convert address")
		}

		txhash, err := k.sendTokens(ctx, share, tokenAddress, evmAddress)
		if err != nil {
			return errorsmod.Wrap(err, "failed to send tokens")
		}

		results = append(results, types.TxResult{
			TxHash:     txhash,
			Address:    evmAddress.String(),
			ReferralId: 0,
		})
	}
	withdrawal.Result = results

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

	toClaim, ok := new(big.Int).SetString(rewards.ToClaim, 10)
	if !ok {
		return nil, "", "", errors.New("invalid referral reward claim")
	}

	if toClaimFromStore.Cmp(toClaim) == -1 {
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

	change := toClaimFromStore.Sub(toClaimFromStore, toClaim)
	rewards.ToClaim = change.String()
	rewards.TotalClaimedAmount = toClaim.String()

	k.InsertReferralRewards(ctx, referral.Id, rewards.TokenId, rewards)

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
