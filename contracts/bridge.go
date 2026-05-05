package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	errorsmod "cosmossdk.io/errors"

	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
)

var (
	//go:embed compiled_contracts/Bridge.json
	bridgeJSON []byte

	// BridgeContract is the compiled Bridge contract
	BridgeContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(bridgeJSON, &BridgeContract)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to init bridge"))
	}
}
