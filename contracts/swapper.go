package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	errorsmod "cosmossdk.io/errors"

	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
)

var (
	//go:embed compiled_contracts/Swapper.json
	swapperJSON []byte

	// SwapperContract is the compiled swapper contract.
	SwapperContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(swapperJSON, &SwapperContract)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to init swapper"))
	}
}
