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

package cli

import (
	errorsmod "cosmossdk.io/errors"
	"os"
	"path/filepath"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/erc20/types"
	"github.com/cosmos/cosmos-sdk/codec"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// ParseRegisterCoinProposal reads and parses a ParseRegisterCoinProposal from a file.
func ParseMetadata(cdc codec.JSONCodec, metadataFile string) ([]banktypes.Metadata, error) {
	proposalMetadata := types.ProposalMetadata{}

	contents, err := os.ReadFile(filepath.Clean(metadataFile))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to read metadata file")
	}

	if err = cdc.UnmarshalJSON(contents, &proposalMetadata); err != nil {
		return nil, errorsmod.Wrap(err, "failed to unmarshal proposal metadata")
	}

	return proposalMetadata.Metadata, nil
}
