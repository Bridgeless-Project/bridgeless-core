package keeper_test

import (
	"context"
	"math/big"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/erc20/types"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/statedb"
	evm "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/stretchr/testify/mock"
)

var _ types.EVMKeeper = &MockEVMKeeper{}

type MockEVMKeeper struct {
	mock.Mock
}

func (m *MockEVMKeeper) GetParams(_ sdk.Context) evm.Params {
	args := m.Called(mock.Anything)
	return args.Get(0).(evm.Params)
}

func (m *MockEVMKeeper) ChainID() *big.Int {
	args := m.Called()
	return args.Get(0).(*big.Int)
}

func (m *MockEVMKeeper) GetAccountWithoutBalance(_ sdk.Context, _ common.Address) *statedb.Account {
	args := m.Called(mock.Anything, mock.Anything)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*statedb.Account)
}

func (m *MockEVMKeeper) EVMConfig(_ sdk.Context, _ sdk.ConsAddress, _ *big.Int) (*statedb.EVMConfig, error) {
	args := m.Called(mock.Anything, mock.Anything, mock.Anything)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*statedb.EVMConfig), args.Error(1)
}

func (m *MockEVMKeeper) TxConfig(_ sdk.Context, _ common.Hash) statedb.TxConfig {
	args := m.Called(mock.Anything, mock.Anything)
	return args.Get(0).(statedb.TxConfig)
}

func (m *MockEVMKeeper) EstimateGas(_ context.Context, _ *evm.EthCallRequest) (*evm.EstimateGasResponse, error) {
	args := m.Called(mock.Anything, mock.Anything)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*evm.EstimateGasResponse), args.Error(1)
}

func (m *MockEVMKeeper) ApplyMessage(_ sdk.Context, _ core.Message, _ vm.EVMLogger, _ bool) (*evm.MsgEthereumTxResponse, error) {
	args := m.Called(mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*evm.MsgEthereumTxResponse), args.Error(1)
}

func (m *MockEVMKeeper) ApplyMessageWithConfig(
	_ sdk.Context,
	_ core.Message,
	_ vm.EVMLogger,
	_ bool,
	_ *statedb.EVMConfig,
	_ statedb.TxConfig,
) (*evm.MsgEthereumTxResponse, error) {
	args := m.Called(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*evm.MsgEthereumTxResponse), args.Error(1)
}

func (m *MockEVMKeeper) BroadcastTxResponse(
	_ sdk.Context,
	_ string,
	_ string,
	_ string,
	_ uint8,
	_ uint64,
	_ *evm.MsgEthereumTxResponse,
) error {
	args := m.Called(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	return args.Error(0)
}

var _ types.BankKeeper = &MockBankKeeper{}

type MockBankKeeper struct {
	mock.Mock
}

func (b *MockBankKeeper) SendCoinsFromModuleToAccount(_ sdk.Context, _ string, _ sdk.AccAddress, _ sdk.Coins) error {
	args := b.Called(mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	return args.Error(0)
}

func (b *MockBankKeeper) SendCoinsFromAccountToModule(_ sdk.Context, _ sdk.AccAddress, _ string, _ sdk.Coins) error {
	args := b.Called(mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	return args.Error(0)
}

func (b *MockBankKeeper) MintCoins(_ sdk.Context, _ string, _ sdk.Coins) error {
	args := b.Called(mock.Anything, mock.Anything, mock.Anything)
	return args.Error(0)
}

func (b *MockBankKeeper) BurnCoins(_ sdk.Context, _ string, _ sdk.Coins) error {
	args := b.Called(mock.Anything, mock.Anything, mock.Anything)
	return args.Error(0)
}

func (b *MockBankKeeper) IsSendEnabledCoin(_ sdk.Context, _ sdk.Coin) bool {
	args := b.Called(mock.Anything, mock.Anything)
	return args.Bool(0)
}

func (b *MockBankKeeper) BlockedAddr(_ sdk.AccAddress) bool {
	args := b.Called(mock.Anything)
	return args.Bool(0)
}

//nolint:all
func (b *MockBankKeeper) GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool) {
	args := b.Called(mock.Anything, mock.Anything)
	return args.Get(0).(banktypes.Metadata), args.Bool(1)
}

func (b *MockBankKeeper) SetDenomMetaData(_ sdk.Context, _ banktypes.Metadata) {
}

func (b *MockBankKeeper) HasSupply(_ sdk.Context, _ string) bool {
	args := b.Called(mock.Anything, mock.Anything)
	return args.Bool(0)
}

func (b *MockBankKeeper) GetBalance(_ sdk.Context, _ sdk.AccAddress, _ string) sdk.Coin {
	args := b.Called(mock.Anything, mock.Anything)
	return args.Get(0).(sdk.Coin)
}
