package backend

import (
	"bufio"
	errorsmod "cosmossdk.io/errors"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/Bridgeless-Project/bridgeless-core/v12/app"
	"github.com/Bridgeless-Project/bridgeless-core/v12/crypto/hd"
	"github.com/Bridgeless-Project/bridgeless-core/v12/encoding"
	"github.com/Bridgeless-Project/bridgeless-core/v12/indexer"
	"github.com/Bridgeless-Project/bridgeless-core/v12/rpc/backend/mocks"
	rpctypes "github.com/Bridgeless-Project/bridgeless-core/v12/rpc/types"
	utiltx "github.com/Bridgeless-Project/bridgeless-core/v12/testutil/tx"
	"github.com/Bridgeless-Project/bridgeless-core/v12/utils"
	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
)

type BackendTestSuite struct {
	suite.Suite

	backend *Backend
	from    common.Address
	acc     sdk.AccAddress
	signer  keyring.Signer
}

func TestBackendTestSuite(t *testing.T) {
	suite.Run(t, new(BackendTestSuite))
}

const ChainID = utils.TestnetChainID + "-1"

// SetupTest is executed before every BackendTestSuite test
func (suite *BackendTestSuite) SetupTest() {
	ctx := server.NewDefaultContext()
	ctx.Viper.Set("telemetry.global-labels", []interface{}{})

	baseDir := suite.T().TempDir()
	nodeDirName := "node"
	clientDir := filepath.Join(baseDir, nodeDirName, "evmoscli")
	keyRing, err := suite.generateTestKeyring(clientDir)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to generate test keyring"))
	}

	// Create Account with set sequence
	suite.acc = sdk.AccAddress(utiltx.GenerateAddress().Bytes())
	accounts := map[string]client.TestAccount{}
	accounts[suite.acc.String()] = client.TestAccount{
		Address: suite.acc,
		Num:     uint64(1),
		Seq:     uint64(1),
	}

	from, priv := utiltx.NewAddrKey()
	suite.from = from
	suite.signer = utiltx.NewSigner(priv)
	suite.Require().NoError(err)

	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	clientCtx := client.Context{}.WithChainID(ChainID).
		WithHeight(1).
		WithTxConfig(encodingConfig.TxConfig).
		WithKeyringDir(clientDir).
		WithKeyring(keyRing).
		WithAccountRetriever(client.TestAccountRetriever{Accounts: accounts})

	allowUnprotectedTxs := false
	idxer := indexer.NewKVIndexer(dbm.NewMemDB(), ctx.Logger, clientCtx)

	suite.backend = NewBackend(ctx, ctx.Logger, clientCtx, allowUnprotectedTxs, idxer)
	suite.backend.queryClient.QueryClient = mocks.NewEVMQueryClient(suite.T())
	suite.backend.clientCtx.Client = mocks.NewClient(suite.T())
	suite.backend.queryClient.FeeMarket = mocks.NewFeeMarketQueryClient(suite.T())
	suite.backend.ctx = rpctypes.ContextWithHeight(1)

	// Add codec
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	suite.backend.clientCtx.Codec = encCfg.Codec
}

// buildEthereumTx returns an example legacy Ethereum transaction
func (suite *BackendTestSuite) buildEthereumTx() (*evmtypes.MsgEthereumTx, []byte) {
	ethTxParams := evmtypes.EvmTxArgs{
		ChainID:  suite.backend.chainID,
		Nonce:    uint64(0),
		To:       &common.Address{},
		Amount:   big.NewInt(0),
		GasLimit: 100000,
		GasPrice: big.NewInt(1),
	}
	msgEthereumTx := evmtypes.NewTx(&ethTxParams)

	// A valid msg should have empty `From`
	msgEthereumTx.From = suite.from.Hex()

	txBuilder := suite.backend.clientCtx.TxConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msgEthereumTx)
	suite.Require().NoError(err)

	bz, err := suite.backend.clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	suite.Require().NoError(err)
	return msgEthereumTx, bz
}

// buildFormattedBlock returns a formatted block for testing
func (suite *BackendTestSuite) buildFormattedBlock(
	blockRes *tmrpctypes.ResultBlockResults,
	resBlock *tmrpctypes.ResultBlock,
	fullTx bool,
	tx *evmtypes.MsgEthereumTx,
	validator sdk.AccAddress,
	baseFee *big.Int,
) map[string]interface{} {
	header := resBlock.Block.Header
	gasLimit := int64(^uint32(0)) // for `MaxGas = -1` (DefaultConsensusParams)
	gasUsed := new(big.Int).SetUint64(uint64(blockRes.TxsResults[0].GasUsed))

	root := common.Hash{}.Bytes()
	receipt := ethtypes.NewReceipt(root, false, gasUsed.Uint64())
	bloom := ethtypes.CreateBloom(ethtypes.Receipts{receipt})

	ethRPCTxs := []interface{}{}
	if tx != nil {
		if fullTx {
			rpcTx, err := rpctypes.NewRPCTransaction(
				tx.AsTransaction(),
				common.BytesToHash(header.Hash()),
				uint64(header.Height),
				uint64(0),
				baseFee,
				suite.backend.chainID,
			)
			suite.Require().NoError(err)
			ethRPCTxs = []interface{}{rpcTx}
		} else {
			ethRPCTxs = []interface{}{common.HexToHash(tx.Hash)}
		}
	}

	return rpctypes.FormatBlock(
		header,
		resBlock.Block.Size(),
		gasLimit,
		gasUsed,
		ethRPCTxs,
		bloom,
		common.BytesToAddress(validator.Bytes()),
		baseFee,
	)
}

func (suite *BackendTestSuite) generateTestKeyring(clientDir string) (keyring.Keyring, error) {
	buf := bufio.NewReader(os.Stdin)
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	return keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, clientDir, buf, encCfg.Codec, []keyring.Option{hd.EthSecp256k1Option()}...)
}

func (suite *BackendTestSuite) signAndEncodeEthTx(msgEthereumTx *evmtypes.MsgEthereumTx) []byte {
	from, priv := utiltx.NewAddrKey()
	signer := utiltx.NewSigner(priv)

	queryClient := suite.backend.queryClient.QueryClient.(*mocks.EVMQueryClient)
	RegisterParamsWithoutHeader(queryClient, 1)

	ethSigner := ethtypes.LatestSigner(suite.backend.ChainConfig())
	msgEthereumTx.From = from.String()
	err := msgEthereumTx.Sign(ethSigner, signer)
	suite.Require().NoError(err)

	tx, err := msgEthereumTx.BuildTx(suite.backend.clientCtx.TxConfig.NewTxBuilder(), utils.BaseDenom)
	suite.Require().NoError(err)

	txEncoder := suite.backend.clientCtx.TxConfig.TxEncoder()
	txBz, err := txEncoder(tx)
	suite.Require().NoError(err)

	return txBz
}
