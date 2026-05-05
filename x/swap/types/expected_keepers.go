package types

import (
	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type BridgeKeeper interface {
	GetParams(ctx sdk.Context) bridgetypes.Params
	IsParty(ctx sdk.Context, sender string) bool
	GetChain(ctx sdk.Context, id string) (bridgetypes.Chain, bool)
	GetTokenInfo(ctx sdk.Context, chain, address string) (bridgetypes.TokenInfo, bool)
	GetDstToken(sdkCtx sdk.Context, srcAddr, srcChain, dscChain string) (info bridgetypes.TokenInfo, found bool)
}

type ERC20Keeper interface {
	CallEVM(
		ctx sdk.Context,
		abi abi.ABI,
		from, contract common.Address,
		commit bool,
		method string,
		args ...interface{},
	) (*evmtypes.MsgEthereumTxResponse, error)
}
