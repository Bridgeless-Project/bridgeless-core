package contracts

import (
	errorsmod "cosmossdk.io/errors"
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/erc20/types"
)

// This is an evil token. Whenever an A -> B transfer is called,
// a predefined C is given a massive allowance on B.
var (
	//go:embed compiled_contracts/ERC20MaliciousDelayed.json
	ERC20MaliciousDelayedJSON []byte //nolint: golint

	// ERC20MaliciousDelayedContract is the compiled erc20 contract
	ERC20MaliciousDelayedContract evmtypes.CompiledContract

	// ERC20MaliciousDelayedAddress is the erc20 module address
	ERC20MaliciousDelayedAddress common.Address
)

func init() {
	ERC20MaliciousDelayedAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20MaliciousDelayedJSON, &ERC20MaliciousDelayedContract)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to init erc20 malicious delayed"))
	}

	if len(ERC20MaliciousDelayedContract.Bin) == 0 {
		panic("load contract failed")
	}
}
