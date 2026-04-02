package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	errorsmod "cosmossdk.io/errors"

	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
)

var (
	//go:embed compiled_contracts/UniswapV2RouterV2.json
	uniswapV2RouterV2JSON []byte

	// UniswapV2RouterV2Contract is the compiled UniswapV2RouterV2 contract
	UniswapV2RouterV2Contract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(uniswapV2RouterV2JSON, &UniswapV2RouterV2Contract)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to init uniswap v2 router v2"))
	}
}
