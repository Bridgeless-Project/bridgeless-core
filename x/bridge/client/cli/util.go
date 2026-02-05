package cli

import (
	"encoding/json"
	"os"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	"github.com/pkg/errors"
)

func parseTxs(path string) ([]types.Transaction, error) {
	var txs []types.Transaction
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &txs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal transaction")
	}
	return txs, nil
}

func parseTx(path string) (*types.Transaction, error) {
	var tx *types.Transaction
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &tx); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal transaction")
	}
	return tx, nil
}

func parseInsertToken(path string) (*types.Token, error) {
	return readFromJSON[types.Token](path)
}

func parseInsertChain(path string) (*types.Chain, error) {
	return readFromJSON[types.Chain](path)
}

func readFromJSON[T any](path string) (*T, error) {
	var result T
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal JSON")
	}

	return &result, nil
}

// StartEpochInput represents the JSON input structure for starting an epoch
type StartEpochInput struct {
	EpochID      uint32          `json:"epoch_id"`
	TSSInfo      []types.TSSInfo `json:"tss_info"`
	TSSThreshold uint32          `json:"tss_threshold"`
}

// parseStartEpochInput parses the start epoch input from a JSON file
func parseStartEpochInput(path string) (*StartEpochInput, error) {
	return readFromJSON[StartEpochInput](path)
}

// parseEpochChainSignatures parses the epoch chain signatures from a JSON file
func parseEpochChainSignatures(path string) ([]types.EpochChainSignatures, error) {
	var signatures []types.EpochChainSignatures
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &signatures); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal epoch chain signatures")
	}
	return signatures, nil
}
