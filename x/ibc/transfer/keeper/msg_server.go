// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package keeper

import (
	"context"
	"strings"

	"github.com/armon/go-metrics"
	"github.com/ethereum/go-ethereum/common"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	erc20types "github.com/Bridgeless-Project/bridgeless-core/v12/x/erc20/types"
	"github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
)

var _ types.MsgServer = Keeper{}

// Transfer defines a gRPC msg server method for MsgTransfer.
// This implementation overrides the default ICS20 transfer's by converting
// the ERC20 tokens to their Cosmos representation if the token pair has been
// registered through governance.
// If user doesn't have enough balance of coin, it will attempt to convert
// ERC20 tokens to the coin denomination, and continue with a regular transfer.
func (k Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// use a zero gas config to avoid extra costs for the relayers
	kvGasCfg := ctx.KVGasConfig()
	transientKVGasCfg := ctx.TransientKVGasConfig()

	// use a zero gas config to avoid extra costs for the relayers
	ctx = ctx.
		WithKVGasConfig(storetypes.GasConfig{}).
		WithTransientKVGasConfig(storetypes.GasConfig{})

	defer func() {
		// return the KV gas config to initial values
		ctx = ctx.
			WithKVGasConfig(kvGasCfg).
			WithTransientKVGasConfig(transientKVGasCfg)
	}()

	// use native denom or contract address
	denom := strings.TrimPrefix(msg.Token.Denom, erc20types.ModuleName+"/")

	pairID := k.erc20Keeper.GetTokenPairID(ctx, denom)
	if len(pairID) == 0 {
		// no-op: token is not registered so we can proceed with regular transfer
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	pair, _ := k.erc20Keeper.GetTokenPair(ctx, pairID)
	if !pair.Enabled {
		// no-op: pair is not enabled so we can proceed with regular transfer
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	senderAcc := k.accountKeeper.GetAccount(ctx, sender)

	if erc20types.IsModuleAccount(senderAcc) {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	if !k.erc20Keeper.IsERC20Enabled(ctx) {
		// no-op: continue with regular transfer
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	// update the msg denom to the token pair denom
	msg.Token.Denom = pair.Denom

	// if the user has enough balance of the Cosmos representation, then we don't need to Convert
	balance := k.bankKeeper.GetBalance(ctx, sender, pair.Denom)
	if balance.Amount.GTE(msg.Token.Amount) {

		defer func() {
			telemetry.IncrCounterWithLabels(
				[]string{"erc20", "ibc", "transfer", "total"},
				1,
				[]metrics.Label{
					telemetry.NewLabel("denom", pair.Denom),
				},
			)
		}()

		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	// only convert the remaining difference
	difference := msg.Token.Amount.Sub(balance.Amount)

	msgConvertERC20 := erc20types.NewMsgConvertERC20(
		difference,
		sender,
		pair.GetERC20Contract(),
		common.BytesToAddress(sender.Bytes()),
	)

	// Use MsgConvertERC20 to convert the ERC20 to a Cosmos IBC Coin
	if _, err := k.erc20Keeper.ConvertERC20(sdk.WrapSDKContext(ctx), msgConvertERC20); err != nil {
		return nil, err
	}

	defer func() {
		telemetry.IncrCounterWithLabels(
			[]string{"erc20", "ibc", "transfer", "total"},
			1,
			[]metrics.Label{
				telemetry.NewLabel("denom", pair.Denom),
			},
		)
	}()

	return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
}
