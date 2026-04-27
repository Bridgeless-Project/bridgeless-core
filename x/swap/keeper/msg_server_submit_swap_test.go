package keeper_test

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	bridgetypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/bridge/types"
	evmtypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/evm/types"
	swapkeeper "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/keeper"
	swaptypes "github.com/Bridgeless-Project/bridgeless-core/v12/x/swap/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

const (
	wrappedBridgeAddress = "0x2000000000000000000000000000000000000002"
	swapperAddress       = "0x9999999999999999999999999999999999999999"
	localDstTokenAddress = "0x4000000000000000000000000000000000000004"
)

type evmCall struct {
	method   string
	from     common.Address
	contract common.Address
	commit   bool
	args     []interface{}
}

type mockBridgeKeeper struct {
	params    bridgetypes.Params
	chains    map[string]bridgetypes.Chain
	tokenInfo map[string]bridgetypes.TokenInfo
	dstTokens map[string]bridgetypes.TokenInfo
}

func (m mockBridgeKeeper) GetParams(_ sdk.Context) bridgetypes.Params {
	return m.params
}

func (m mockBridgeKeeper) IsParty(_ sdk.Context, sender string) bool {
	for _, party := range m.params.Parties {
		if party.Address == sender {
			return true
		}
	}

	return false
}

func (m mockBridgeKeeper) GetChain(_ sdk.Context, id string) (bridgetypes.Chain, bool) {
	chain, found := m.chains[id]
	return chain, found
}

func (m mockBridgeKeeper) GetTokenInfo(_ sdk.Context, chain, address string) (bridgetypes.TokenInfo, bool) {
	info, found := m.tokenInfo[chain+"|"+address]
	return info, found
}

func (m mockBridgeKeeper) GetDstToken(_ sdk.Context, srcAddr, srcChain, dstChain string) (bridgetypes.TokenInfo, bool) {
	info, found := m.dstTokens[srcAddr+"|"+srcChain+"|"+dstChain]
	return info, found
}

type mockERC20Keeper struct {
	calls       []evmCall
	defaultHash string
}

func (m *mockERC20Keeper) CallEVM(
	_ sdk.Context,
	contractABI abi.ABI,
	from, contract common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {
	_, err := contractABI.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	m.calls = append(m.calls, evmCall{
		method:   method,
		from:     from,
		contract: contract,
		commit:   commit,
		args:     args,
	})

	return &evmtypes.MsgEthereumTxResponse{Hash: m.defaultHash}, nil
}

func newSubmitSwapKeeper(t testing.TB, bridge mockBridgeKeeper, erc20 *mockERC20Keeper) (*swapkeeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(swaptypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(swaptypes.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	paramsSubspace := typesparams.NewSubspace(cdc, swaptypes.Amino, storeKey, memStoreKey, "SwapParams")

	k := swapkeeper.NewKeeper(cdc, storeKey, memStoreKey, paramsSubspace, bridge, erc20)
	ctx := sdk.NewContext(stateStore, tmproto.Header{ChainID: "core"}, false, log.NewNopLogger())
	return k, ctx
}

func sampleSubmitSwapMsg(creator string) *swaptypes.MsgSubmitSwapTx {
	return &swaptypes.MsgSubmitSwapTx{
		Creator: creator,
		Tx: &swaptypes.SwapTransaction{
			Tx: bridgetypes.Transaction{
				DepositChainId:    "1",
				DepositTxHash:     "0xabc",
				DepositTxIndex:    7,
				DepositToken:      "0x1000000000000000000000000000000000000001",
				DepositAmount:     "100",
				Depositor:         "0xdepositor",
				Receiver:          "0xreceiver",
				WithdrawalChainId: "2",
				WithdrawalToken:   "0x3000000000000000000000000000000000000003",
				Signature:         "0x1234",
				WithdrawalAmount:  "0",
				CommissionAmount:  "0",
				TxData:            "recovery-recipient",
				ReferralId:        12,
			},
			FinalReceiver: "dest-recipient",
			AmountOutMin:  "90",
		},
	}
}

func submitSwapBridge(parties ...string) mockBridgeKeeper {
	partyRecords := make([]*bridgetypes.Party, 0, len(parties))
	for _, party := range parties {
		partyRecords = append(partyRecords, &bridgetypes.Party{Address: party})
	}

	return mockBridgeKeeper{
		params: bridgetypes.Params{
			TssThreshold: 1,
			Parties:      partyRecords,
		},
		chains: map[string]bridgetypes.Chain{
			"2": {
				Id:            "2",
				BridgeAddress: "0x2222222222222222222222222222222222222222",
				Confirmations: 1,
				Name:          "destination",
			},
		},
		tokenInfo: map[string]bridgetypes.TokenInfo{
			"2|0x3000000000000000000000000000000000000003": {
				Address:   "0x3000000000000000000000000000000000000003",
				ChainId:   "2",
				TokenId:   3,
				IsWrapped: true,
			},
		},
		dstTokens: map[string]bridgetypes.TokenInfo{
			"0x3000000000000000000000000000000000000003|2|core": {
				Address: localDstTokenAddress,
				ChainId: "core",
				TokenId: 3,
			},
		},
	}
}

func setSubmitSwapParams(ctx sdk.Context, k *swapkeeper.Keeper, admin string, deadline uint64) {
	k.SetParams(ctx, swaptypes.NewParams(admin, wrappedBridgeAddress, swapperAddress, deadline))
}

func bytesOf(length int, value byte) []byte {
	bz := make([]byte, length)
	for i := range bz {
		bz[i] = value
	}
	return bz
}

func requireBigIntString(t *testing.T, value interface{}, expected string) {
	t.Helper()

	require.Equal(t, expected, value.(*big.Int).String())
}

func fieldByName(t *testing.T, value interface{}, name string) interface{} {
	t.Helper()

	field := reflect.ValueOf(value).FieldByName(name)
	require.True(t, field.IsValid(), "missing field %s", name)
	return field.Interface()
}

func TestSubmitSwapTxRequiresParty(t *testing.T) {
	bridge := mockBridgeKeeper{params: bridgetypes.Params{}}
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(sdk.AccAddress([]byte("not-party-addr-0001")).String()))
	require.Error(t, err)
}

func TestSubmitSwapTxRejectsNilMessage(t *testing.T) {
	bridge := mockBridgeKeeper{}
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
}

func TestSubmitSwapTxDuplicateSubmitter(t *testing.T) {
	party := sdk.AccAddress(make([]byte, 20)).String()
	bridge := mockBridgeKeeper{
		params: bridgetypes.Params{
			TssThreshold: 1,
			Parties:      []*bridgetypes.Party{{Address: party}},
		},
	}
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	setSubmitSwapParams(ctx, k, party, swaptypes.DefaultSwapDeadlineSeconds)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party))
	require.NoError(t, err)

	_, err = ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party))
	require.Error(t, err)
}

func TestSubmitSwapTxThresholdAndSwapperExecution(t *testing.T) {
	partyA := sdk.AccAddress(make([]byte, 20)).String()
	partyB := sdk.AccAddress(bytesOf(20, 2)).String()
	bridge := submitSwapBridge(partyA, partyB)
	erc20 := &mockERC20Keeper{defaultHash: "0xswapper"}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	customDeadline := uint64(42)
	setSubmitSwapParams(ctx, k, partyA, customDeadline)
	ms := swapkeeper.NewMsgServerImpl(*k)

	msgA := sampleSubmitSwapMsg(partyA)
	msgB := sampleSubmitSwapMsg(partyB)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), msgA)
	require.NoError(t, err)

	_, found := k.GetSwap(ctx, msgA.Tx.Tx.DepositTxHash, msgA.Tx.Tx.DepositTxIndex, msgA.Tx.Tx.DepositChainId)
	require.False(t, found)
	require.Len(t, erc20.calls, 0)

	_, err = ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), msgB)
	require.NoError(t, err)

	stored, found := k.GetSwap(ctx, msgA.Tx.Tx.DepositTxHash, msgA.Tx.Tx.DepositTxIndex, msgA.Tx.Tx.DepositChainId)
	require.True(t, found)
	require.Empty(t, stored.FinalAmount)
	require.Equal(t, "0xswapper", stored.FinalDepositTxHash)

	require.Len(t, erc20.calls, 1)
	call := erc20.calls[0]
	require.Equal(t, "withdrawSwapAndRoute", call.method)
	require.Equal(t, swaptypes.ModuleAddress, call.from)
	require.Equal(t, common.HexToAddress(swapperAddress), call.contract)
	require.True(t, call.commit)
	require.Len(t, call.args, 4)

	withdrawParams := call.args[0]
	require.Equal(t, common.HexToAddress(msgA.Tx.Tx.DepositToken), fieldByName(t, withdrawParams, "Token"))
	requireBigIntString(t, fieldByName(t, withdrawParams, "Amount"), "100")
	require.Equal(t, common.HexToHash(msgA.Tx.Tx.DepositTxHash), fieldByName(t, withdrawParams, "TxHash"))
	requireBigIntString(t, fieldByName(t, withdrawParams, "TxNonce"), "7")
	require.Equal(t, false, fieldByName(t, withdrawParams, "IsWrapped"))
	require.Equal(t, [][]byte{{0x12, 0x34}}, fieldByName(t, withdrawParams, "Signatures"))

	swapParams := call.args[1]
	requireBigIntString(t, fieldByName(t, swapParams, "AmountIn"), "100")
	requireBigIntString(t, fieldByName(t, swapParams, "MinDestinationAmount"), "90")
	require.Equal(t, ctx.BlockTime().Add(time.Duration(customDeadline)*time.Second).Unix(), fieldByName(t, swapParams, "SwapDeadline").(*big.Int).Int64())
	require.Equal(t, []common.Address{
		common.HexToAddress(msgA.Tx.Tx.DepositToken),
		common.HexToAddress(wrappedBridgeAddress),
		common.HexToAddress(localDstTokenAddress),
	}, fieldByName(t, swapParams, "Path"))
	require.Equal(t, false, fieldByName(t, swapParams, "IsDestinationTokenNative"))

	destinationDepositParams := call.args[2]
	require.Equal(t, "dest-recipient", fieldByName(t, destinationDepositParams, "Receiver"))
	require.Equal(t, "2", fieldByName(t, destinationDepositParams, "Network"))
	require.Equal(t, true, fieldByName(t, destinationDepositParams, "IsWrapped"))
	require.Equal(t, uint16(12), fieldByName(t, destinationDepositParams, "ReferralId"))

	fallbackDepositParams := call.args[3]
	require.Equal(t, "recovery-recipient", fieldByName(t, fallbackDepositParams, "Receiver"))
	require.Equal(t, "1", fieldByName(t, fallbackDepositParams, "Network"))
	require.Equal(t, false, fieldByName(t, fallbackDepositParams, "IsWrapped"))
	require.Equal(t, uint16(12), fieldByName(t, fallbackDepositParams, "ReferralId"))
}

func TestSubmitSwapTxRejectsAlreadyProcessedSwap(t *testing.T) {
	party := sdk.AccAddress(make([]byte, 20)).String()
	bridge := mockBridgeKeeper{
		params: bridgetypes.Params{
			Parties: []*bridgetypes.Party{{Address: party}},
		},
	}
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	setSubmitSwapParams(ctx, k, party, swaptypes.DefaultSwapDeadlineSeconds)
	existing := sampleSubmitSwapMsg(party).Tx
	existing.FinalAmount = "42"
	k.SetSwap(ctx, *existing)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party))
	require.Error(t, err)
}

func TestSubmitSwapTxRejectsMissingSwapperConfig(t *testing.T) {
	party := sdk.AccAddress(make([]byte, 20)).String()
	bridge := submitSwapBridge(party)
	bridge.params.TssThreshold = 0
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	k.SetParams(ctx, swaptypes.DefaultParams())
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party))
	require.Error(t, err)
}

func TestSubmitSwapTxRejectsMissingTokenMapping(t *testing.T) {
	party := sdk.AccAddress(make([]byte, 20)).String()
	bridge := submitSwapBridge(party)
	bridge.params.TssThreshold = 0
	bridge.dstTokens = nil
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	setSubmitSwapParams(ctx, k, party, swaptypes.DefaultSwapDeadlineSeconds)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party))
	require.Error(t, err)
	require.Len(t, erc20.calls, 0)
}
