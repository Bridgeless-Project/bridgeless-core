package types

import (
	"math/big"

	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Zano-Execution-Layer/go-ethereum/accounts/abi"
	"github.com/Zano-Execution-Layer/go-ethereum/common"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins // Methods imported from bank should be defined here
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type BridgeHook interface {
	GetTokenPrice(ctx sdk.Context, tokenAddress string, amountIn *big.Int) (*big.Int, []common.Address, error)
}

type ERC20Keeper interface {
	CallEVMAsTx(
		ctx sdk.Context,
		abi abi.ABI,
		from, contract common.Address,
		commit bool,
		method string,
		args ...interface{},
	) (*evmtypes.MsgEthereumTxResponse, error)
}
