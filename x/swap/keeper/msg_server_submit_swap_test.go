package keeper_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/Bridgeless-Project/bridgeless-core/v12/contracts"
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

func (m mockBridgeKeeper) GetChain(ctx sdk.Context, id string) (bridgetypes.Chain, bool) {
	chain, found := m.chains[id]
	return chain, found
}

func (m mockBridgeKeeper) GetTokenInfo(ctx sdk.Context, chain, address string) (bridgetypes.TokenInfo, bool) {
	info, found := m.tokenInfo[chain+"|"+address]
	return info, found
}

func (m mockBridgeKeeper) GetDstToken(ctx sdk.Context, srcAddr, srcChain, dscChain string) (bridgetypes.TokenInfo, bool) {
	info, found := m.dstTokens[srcAddr+"|"+srcChain+"|"+dscChain]
	return info, found
}

type mockERC20Keeper struct {
	calls       []evmCall
	routerHash  string
	routerRet   []byte
	defaultHash string
}

func (m *mockERC20Keeper) CallEVM(
	ctx sdk.Context,
	contractABI abi.ABI,
	from, contract common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {
	m.calls = append(m.calls, evmCall{
		method:   method,
		from:     from,
		contract: contract,
		commit:   commit,
		args:     args,
	})

	resp := &evmtypes.MsgEthereumTxResponse{Hash: m.defaultHash}
	if method == "swapExactTokensForTokens" {
		resp.Hash = m.routerHash
		resp.Ret = m.routerRet
	}

	return resp, nil
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
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}

func newRouterReturn(t *testing.T, amounts ...int64) []byte {
	t.Helper()

	values := make([]*big.Int, 0, len(amounts))
	for _, amount := range amounts {
		values = append(values, big.NewInt(amount))
	}

	ret, err := contracts.UniswapV2RouterV2Contract.ABI.Methods["swapExactTokensForTokens"].Outputs.Pack(values)
	require.NoError(t, err)
	return ret
}

func sampleSubmitSwapMsg(creator string, isBridge bool) *swaptypes.MsgSubmitSwapTx {
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
			},
			FinalReceiver: "dest-recipient",
			AmountOutMin:  "90",
		},
		IsBridgeTx: isBridge,
	}
}

func TestSubmitSwapTxRequiresParty(t *testing.T) {
	bridge := mockBridgeKeeper{
		params: bridgetypes.Params{},
	}
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(sdk.AccAddress([]byte("not-party-addr-0001")).String(), false))
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
	k.SetParams(ctx, swaptypes.NewParams(party, "0x9999999999999999999999999999999999999999", "0x2000000000000000000000000000000000000002", swaptypes.DefaultSwapDeadlineSeconds))
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party, false))
	require.NoError(t, err)

	_, err = ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party, false))
	require.Error(t, err)
}

func TestSubmitSwapTxThresholdAndExecutionWithoutBridgeDeposit(t *testing.T) {
	partyA := sdk.AccAddress(make([]byte, 20)).String()
	partyB := sdk.AccAddress(bytesOf(20, 2)).String()
	bridge := mockBridgeKeeper{
		params: bridgetypes.Params{
			TssThreshold: 1,
			Parties: []*bridgetypes.Party{
				{Address: partyA},
				{Address: partyB},
			},
		},
		chains: map[string]bridgetypes.Chain{
			"2": {
				Id:            "2",
				BridgeAddress: "0x2000000000000000000000000000000000000002",
				Operator:      "0x2222222222222222222222222222222222222222",
				Confirmations: 1,
				Name:          "core-evm",
			},
		},
		dstTokens: map[string]bridgetypes.TokenInfo{
			"0x3000000000000000000000000000000000000003|2|": {
				Address: "0x3000000000000000000000000000000000000003",
				ChainId: "",
				TokenId: 3,
			},
		},
	}
	erc20 := &mockERC20Keeper{
		defaultHash: "0x01",
		routerHash:  "0x02",
		routerRet:   newRouterReturn(t, 100, 95),
	}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	customDeadline := uint64(42)
	k.SetParams(ctx, swaptypes.NewParams(
		partyA,
		"0x9999999999999999999999999999999999999999",
		"0x2000000000000000000000000000000000000002",
		customDeadline,
	))
	ms := swapkeeper.NewMsgServerImpl(*k)

	msgA := sampleSubmitSwapMsg(partyA, false)
	msgB := sampleSubmitSwapMsg(partyB, false)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), msgA)
	require.NoError(t, err)

	_, found := k.GetSwap(ctx, msgA.Tx.Tx.DepositTxHash, msgA.Tx.Tx.DepositTxIndex, msgA.Tx.Tx.DepositChainId)
	require.False(t, found)
	require.Len(t, erc20.calls, 0)

	_, err = ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), msgB)
	require.NoError(t, err)

	stored, found := k.GetSwap(ctx, msgA.Tx.Tx.DepositTxHash, msgA.Tx.Tx.DepositTxIndex, msgA.Tx.Tx.DepositChainId)
	require.True(t, found)
	require.Equal(t, "95", stored.FinalAmount)
	require.Empty(t, stored.FinalDepositTxHash)

	require.Len(t, erc20.calls, 2)
	require.Equal(t, "withdrawERC20", erc20.calls[0].method)
	require.Equal(t, "swapExactTokensForTokens", erc20.calls[1].method)
	require.Equal(t, [][]byte{{0x12, 0x34}}, erc20.calls[0].args[6].([][]byte))

	path, ok := erc20.calls[1].args[2].([]common.Address)
	require.True(t, ok)
	require.Equal(t, []common.Address{
		common.HexToAddress(msgA.Tx.Tx.DepositToken),
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
		common.HexToAddress(msgA.Tx.Tx.WithdrawalToken),
	}, path)

	require.Equal(t, swaptypes.ModuleAddress, erc20.calls[1].args[3].(common.Address))
	require.Equal(t, "100", erc20.calls[1].args[0].(*big.Int).String())
	require.Equal(t, "90", erc20.calls[1].args[1].(*big.Int).String())
	require.Equal(t, ctx.BlockTime().Add(time.Duration(customDeadline)*time.Second).Unix(), erc20.calls[1].args[4].(*big.Int).Int64())
}

func TestSubmitSwapTxBridgeDepositStoresMockHash(t *testing.T) {
	partyA := sdk.AccAddress(make([]byte, 20)).String()
	partyB := sdk.AccAddress(bytesOf(20, 3)).String()
	msg := sampleSubmitSwapMsg(partyA, true)
	bridge := mockBridgeKeeper{
		params: bridgetypes.Params{
			TssThreshold: 1,
			Parties: []*bridgetypes.Party{
				{Address: partyA},
				{Address: partyB},
			},
		},
		chains: map[string]bridgetypes.Chain{
			"2": {
				Id:            "2",
				BridgeAddress: "0x2000000000000000000000000000000000000002",
				Operator:      "0x2222222222222222222222222222222222222222",
				Confirmations: 1,
				Name:          "core-evm",
			},
		},
		tokenInfo: map[string]bridgetypes.TokenInfo{
			msg.Tx.Tx.WithdrawalChainId + "|" + msg.Tx.Tx.WithdrawalToken: {
				Address:   msg.Tx.Tx.WithdrawalToken,
				ChainId:   msg.Tx.Tx.WithdrawalChainId,
				TokenId:   3,
				IsWrapped: true,
			},
		},
		dstTokens: map[string]bridgetypes.TokenInfo{
			msg.Tx.Tx.WithdrawalToken + "|" + msg.Tx.Tx.WithdrawalChainId + "|": {
				Address: msg.Tx.Tx.WithdrawalToken,
				ChainId: "",
				TokenId: 3,
			},
		},
	}
	erc20 := &mockERC20Keeper{
		defaultHash: "0x01",
		routerHash:  "0x02",
		routerRet:   newRouterReturn(t, 100, 99),
	}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	k.SetParams(ctx, swaptypes.NewParams(
		partyA,
		"0x9999999999999999999999999999999999999999",
		"0x2000000000000000000000000000000000000002",
		swaptypes.DefaultSwapDeadlineSeconds,
	))
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)

	_, err = ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), &swaptypes.MsgSubmitSwapTx{
		Creator:    partyB,
		Tx:         msg.Tx,
		IsBridgeTx: true,
	})
	require.NoError(t, err)

	stored, found := k.GetSwap(ctx, msg.Tx.Tx.DepositTxHash, msg.Tx.Tx.DepositTxIndex, msg.Tx.Tx.DepositChainId)
	require.True(t, found)
	require.Equal(t, "99", stored.FinalAmount)
	require.NotEmpty(t, stored.FinalDepositTxHash)
	require.Len(t, erc20.calls, 3)
	require.Equal(t, "depositERC20", erc20.calls[2].method)
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
	k.SetParams(ctx, swaptypes.NewParams(
		party,
		"0x9999999999999999999999999999999999999999",
		"0x2000000000000000000000000000000000000002",
		swaptypes.DefaultSwapDeadlineSeconds,
	))
	existing := sampleSubmitSwapMsg(party, false).Tx
	existing.FinalAmount = "42"
	k.SetSwap(ctx, *existing)
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party, false))
	require.Error(t, err)
}

func TestSubmitSwapTxRejectsMissingRouterConfig(t *testing.T) {
	party := sdk.AccAddress(make([]byte, 20)).String()
	bridge := mockBridgeKeeper{
		params: bridgetypes.Params{
			Parties: []*bridgetypes.Party{{Address: party}},
		},
		chains: map[string]bridgetypes.Chain{
			"2": {
				Id:            "2",
				BridgeAddress: "0x2000000000000000000000000000000000000002",
				Confirmations: 1,
				Name:          "core-evm",
			},
		},
	}
	erc20 := &mockERC20Keeper{}
	k, ctx := newSubmitSwapKeeper(t, bridge, erc20)
	k.SetParams(ctx, swaptypes.DefaultParams())
	ms := swapkeeper.NewMsgServerImpl(*k)

	_, err := ms.SubmitSwapTx(sdk.WrapSDKContext(ctx), sampleSubmitSwapMsg(party, false))
	require.Error(t, err)
}
