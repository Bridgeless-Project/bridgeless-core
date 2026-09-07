package keeper_test

import (
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/Zano-Execution-Layer/go-ethereum/common"
	"github.com/Zano-Execution-Layer/go-ethereum/core"
	ethtypes "github.com/Zano-Execution-Layer/go-ethereum/core/types"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/keeper"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/statedb"
	"github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
)

// LogRecordHook records all the logs
type LogRecordHook struct {
	Logs []*ethtypes.Log
}

func (dh *LogRecordHook) PostTxProcessing(_ sdk.Context, _ core.Message, receipt *ethtypes.Receipt) error {
	dh.Logs = receipt.Logs
	return nil
}

// FailureHook always fail
type FailureHook struct{}

func (dh FailureHook) PostTxProcessing(_ sdk.Context, _ core.Message, _ *ethtypes.Receipt) error {
	return errors.New("post tx processing failed")
}

type TransientRecordHook struct {
	keeper *keeper.Keeper

	TxIndex uint64
	LogSize uint64
}

func (h *TransientRecordHook) PostTxProcessing(ctx sdk.Context, _ core.Message, _ *ethtypes.Receipt) error {
	h.TxIndex = h.keeper.GetTxIndexTransient(ctx)
	h.LogSize = h.keeper.GetLogSizeTransient(ctx)
	return nil
}

func (suite *KeeperTestSuite) TestEvmHooks() {
	testCases := []struct {
		msg       string
		setupHook func() types.EvmHooks
		expFunc   func(hook types.EvmHooks, result error)
	}{
		{
			"log collect hook",
			func() types.EvmHooks {
				return &LogRecordHook{}
			},
			func(hook types.EvmHooks, result error) {
				suite.Require().NoError(result)
				suite.Require().Equal(1, len(hook.(*LogRecordHook).Logs))
			},
		},
		{
			"always fail hook",
			func() types.EvmHooks {
				return &FailureHook{}
			},
			func(hook types.EvmHooks, result error) {
				suite.Require().Error(result)
			},
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.app.EvmKeeper = suite.app.EvmKeeper.CleanHooks()
		hook := tc.setupHook()
		suite.app.EvmKeeper.SetHooks(keeper.NewMultiEvmHooks(hook))

		k := suite.app.EvmKeeper
		ctx := suite.ctx
		txHash := common.BigToHash(big.NewInt(1))
		vmdb := statedb.New(ctx, k, statedb.NewTxConfig(
			common.BytesToHash(ctx.HeaderHash().Bytes()),
			txHash,
			0,
			0,
		))

		vmdb.AddLog(&ethtypes.Log{
			Topics:  []common.Hash{},
			Address: suite.address,
		})
		logs := vmdb.Logs()
		receipt := &ethtypes.Receipt{
			TxHash: txHash,
			Logs:   logs,
		}
		result := k.PostTxProcessing(ctx, ethtypes.Message{}, receipt)

		tc.expFunc(hook, result)
	}
}

func (suite *KeeperTestSuite) TestApplyTransactionReservesTransientMetadataBeforeHooks() {
	suite.SetupTest()
	suite.app.EvmKeeper = suite.app.EvmKeeper.CleanHooks()

	hook := &TransientRecordHook{keeper: suite.app.EvmKeeper}
	suite.app.EvmKeeper.SetHooks(keeper.NewMultiEvmHooks(hook))

	to := common.Address{}
	nonce := suite.app.EvmKeeper.GetNonce(suite.ctx, suite.address)
	tx, err := newSignedEthTx(&ethtypes.AccessListTx{
		GasPrice: big.NewInt(1),
		Gas:      21_000,
		To:       &to,
		Value:    big.NewInt(0),
		Data:     []byte{},
	}, nonce, sdk.AccAddress(suite.address.Bytes()), suite.signer, suite.ethSigner)
	suite.Require().NoError(err)

	res, err := suite.app.EvmKeeper.ApplyTransaction(suite.ctx, tx)
	suite.Require().NoError(err)
	suite.Require().False(res.Failed())
	suite.Require().Equal(uint64(1), hook.TxIndex)
	suite.Require().Equal(uint64(0), hook.LogSize)
	suite.Require().Equal(uint64(1), suite.app.EvmKeeper.GetTxIndexTransient(suite.ctx))
}
