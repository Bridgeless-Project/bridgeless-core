package contracts

import (
	errorsmod "cosmossdk.io/errors"
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
)

var (
	//go:embed compiled_contracts/ERC20Burnable.json
	erc20BurnableJSON []byte

	// ERC20BurnableContract is the compiled ERC20Burnable contract
	ERC20BurnableContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(erc20BurnableJSON, &ERC20BurnableContract)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to init erc20 burnable"))
	}
}
