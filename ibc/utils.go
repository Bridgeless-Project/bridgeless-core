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

package ibc

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"

	"github.com/Bridgeless-Project/bridgeless-core/v12/utils"
)

// GetTransferSenderRecipient returns the sender and recipient sdk.AccAddresses
// from an ICS20 FungibleTokenPacketData as well as the original sender bech32
// address from the packet data. This function fails if:
//   - the packet data is not FungibleTokenPacketData
//   - sender address is invalid
//   - recipient address is invalid
func GetTransferSenderRecipient(packet channeltypes.Packet) (
	sender, recipient sdk.AccAddress,
	senderBech32, recipientBech32 string,
	err error,
) {
	// unmarshal packet data to obtain the sender and recipient
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return nil, nil, "", "", errorsmod.Wrapf(errortypes.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data")
	}

	// validate the sender bech32 address from the counterparty chain
	// and change the bech32 human readable prefix (HRP) of the sender to `evmos`
	sender, err = utils.GetEvmosAddressFromBech32(data.Sender)
	if err != nil {
		return nil, nil, "", "", errorsmod.Wrap(err, "invalid sender")
	}

	// validate the recipient bech32 address from the counterparty chain
	// and change the bech32 human readable prefix (HRP) of the recipient to `evmos`
	recipient, err = utils.GetEvmosAddressFromBech32(data.Receiver)
	if err != nil {
		return nil, nil, "", "", errorsmod.Wrap(err, "invalid recipient")
	}

	return sender, recipient, data.Sender, data.Receiver, nil
}

// GetTransferAmount returns the amount from an ICS20 FungibleTokenPacketData as a string.
func GetTransferAmount(packet channeltypes.Packet) (string, error) {
	// unmarshal packet data to obtain the sender and recipient
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return "", errorsmod.Wrapf(errortypes.ErrUnknownRequest, "cannot unmarshal ICS-20 transfer packet data")
	}

	if data.Amount == "" {
		return "", errorsmod.Wrapf(errortypes.ErrInvalidCoins, "empty amount")
	}

	if _, ok := sdk.NewIntFromString(data.Amount); !ok {
		return "", errorsmod.Wrapf(errortypes.ErrInvalidCoins, "invalid amount")
	}

	return data.Amount, nil
}

// GetReceivedCoin returns the transferred coin from an ICS20 FungibleTokenPacketData
// as seen from the destination chain.
// If the receiving chain is the source chain of the tokens, it removes the prefix
// path added by source (i.e sender) chain to the denom. Otherwise, it adds the
// prefix path from the destination chain to the denom.
func GetReceivedCoin(srcPort, srcChannel, dstPort, dstChannel, rawDenom, rawAmt string) sdk.Coin {
	// NOTE: Denom and amount are already validated
	amount, _ := sdk.NewIntFromString(rawAmt)

	if transfertypes.ReceiverChainIsSource(srcPort, srcChannel, rawDenom) {
		// remove prefix added by sender chain
		voucherPrefix := transfertypes.GetDenomPrefix(srcPort, srcChannel)
		unprefixedDenom := rawDenom[len(voucherPrefix):]

		// coin denomination used in sending from the escrow address
		denom := unprefixedDenom

		// The denomination used to send the coins is either the native denom or the hash of the path
		// if the denomination is not native.
		denomTrace := transfertypes.ParseDenomTrace(unprefixedDenom)
		if denomTrace.Path != "" {
			denom = denomTrace.IBCDenom()
		}

		return sdk.Coin{
			Denom:  denom,
			Amount: amount,
		}
	}

	// since SendPacket did not prefix the denomination, we must prefix denomination here
	sourcePrefix := transfertypes.GetDenomPrefix(dstPort, dstChannel)
	// NOTE: sourcePrefix contains the trailing "/"
	prefixedDenom := sourcePrefix + rawDenom

	// construct the denomination trace from the full raw denomination
	denomTrace := transfertypes.ParseDenomTrace(prefixedDenom)
	voucherDenom := denomTrace.IBCDenom()

	return sdk.Coin{
		Denom:  voucherDenom,
		Amount: amount,
	}
}

// GetSentCoin returns the sent coin from an ICS20 FungibleTokenPacketData.
func GetSentCoin(rawDenom, rawAmt string) sdk.Coin {
	// NOTE: Denom and amount are already validated
	amount, _ := sdk.NewIntFromString(rawAmt)
	trace := transfertypes.ParseDenomTrace(rawDenom)

	return sdk.Coin{
		Denom:  trace.IBCDenom(),
		Amount: amount,
	}
}
