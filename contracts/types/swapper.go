// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/Zano-Execution-Layer/go-ethereum"
	"github.com/Zano-Execution-Layer/go-ethereum/accounts/abi"
	"github.com/Zano-Execution-Layer/go-ethereum/accounts/abi/bind"
	"github.com/Zano-Execution-Layer/go-ethereum/common"
	"github.com/Zano-Execution-Layer/go-ethereum/core/types"
	"github.com/Zano-Execution-Layer/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISwapperDepositParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapperDepositParams struct {
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
}

// ISwapperSwapParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapperSwapParams struct {
	AmountIn                 *big.Int
	MinDestinationAmount     *big.Int
	SwapDeadline             *big.Int
	Path                     []common.Address
	IsDestinationTokenNative bool
}

// ISwapperWithdrawParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapperWithdrawParams struct {
	Token      common.Address
	Amount     *big.Int
	TxHash     [32]byte
	TxNonce    *big.Int
	IsWrapped  bool
	Signatures [][]byte
}

// SwapperMetaData contains all meta data concerning the Swapper contract.
var SwapperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"CrossChainERC20Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"CrossChainERC20FallbackDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"CrossChainNativeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LocalERC20Transferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"LocalNativeTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isDestinationTokenNative\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISwapper.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structISwapper.DepositParams\",\"name\":\"destinationDepositParams\",\"type\":\"tuple\"}],\"name\":\"SwappedAndRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isDestinationTokenNative\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structISwapper.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structISwapper.DepositParams\",\"name\":\"destinationDepositParams\",\"type\":\"tuple\"}],\"name\":\"SwappedETHAndRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2RouterAddress_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"operators_\",\"type\":\"address[]\"}],\"name\":\"__Swapper_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"}],\"name\":\"isCurrentNetwork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isDestinationTokenNative\",\"type\":\"bool\"}],\"internalType\":\"structISwapper.SwapParams\",\"name\":\"swapParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"internalType\":\"structISwapper.DepositParams\",\"name\":\"destinationDepositParams_\",\"type\":\"tuple\"}],\"name\":\"swapAndRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isDestinationTokenNative\",\"type\":\"bool\"}],\"internalType\":\"structISwapper.SwapParams\",\"name\":\"swapParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"internalType\":\"structISwapper.DepositParams\",\"name\":\"destinationDepositParams_\",\"type\":\"tuple\"}],\"name\":\"swapETHAndRoute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structISwapper.WithdrawParams\",\"name\":\"withdrawParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"isDestinationTokenNative\",\"type\":\"bool\"}],\"internalType\":\"structISwapper.SwapParams\",\"name\":\"swapParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"internalType\":\"structISwapper.DepositParams\",\"name\":\"destinationDepositParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"internalType\":\"structISwapper.DepositParams\",\"name\":\"fallbackDepositParams_\",\"type\":\"tuple\"}],\"name\":\"withdrawSwapAndRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SwapperABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapperMetaData.ABI instead.
var SwapperABI = SwapperMetaData.ABI

// Swapper is an auto generated Go binding around an Ethereum contract.
type Swapper struct {
	SwapperCaller     // Read-only binding to the contract
	SwapperTransactor // Write-only binding to the contract
	SwapperFilterer   // Log filterer for contract events
}

// SwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapperSession struct {
	Contract     *Swapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapperCallerSession struct {
	Contract *SwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapperTransactorSession struct {
	Contract     *SwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapperRaw struct {
	Contract *Swapper // Generic contract binding to access the raw methods on
}

// SwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapperCallerRaw struct {
	Contract *SwapperCaller // Generic read-only contract binding to access the raw methods on
}

// SwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapperTransactorRaw struct {
	Contract *SwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapper creates a new instance of Swapper, bound to a specific deployed contract.
func NewSwapper(address common.Address, backend bind.ContractBackend) (*Swapper, error) {
	contract, err := bindSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// NewSwapperCaller creates a new read-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperCaller(address common.Address, caller bind.ContractCaller) (*SwapperCaller, error) {
	contract, err := bindSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperCaller{contract: contract}, nil
}

// NewSwapperTransactor creates a new write-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapperTransactor, error) {
	contract, err := bindSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperTransactor{contract: contract}, nil
}

// NewSwapperFilterer creates a new log filterer instance of Swapper, bound to a specific deployed contract.
func NewSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapperFilterer, error) {
	contract, err := bindSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapperFilterer{contract: contract}, nil
}

// bindSwapper binds a generic wrapper to an already deployed contract.
func bindSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.SwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Swapper *SwapperCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Swapper *SwapperSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Swapper.Contract.DEFAULTADMINROLE(&_Swapper.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Swapper *SwapperCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Swapper.Contract.DEFAULTADMINROLE(&_Swapper.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Swapper *SwapperCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Swapper *SwapperSession) OPERATORROLE() ([32]byte, error) {
	return _Swapper.Contract.OPERATORROLE(&_Swapper.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Swapper *SwapperCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Swapper.Contract.OPERATORROLE(&_Swapper.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Swapper *SwapperCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Swapper *SwapperSession) Bridge() (common.Address, error) {
	return _Swapper.Contract.Bridge(&_Swapper.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Swapper *SwapperCallerSession) Bridge() (common.Address, error) {
	return _Swapper.Contract.Bridge(&_Swapper.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Swapper *SwapperCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Swapper *SwapperSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Swapper.Contract.GetRoleAdmin(&_Swapper.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Swapper *SwapperCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Swapper.Contract.GetRoleAdmin(&_Swapper.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Swapper *SwapperCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Swapper *SwapperSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Swapper.Contract.GetRoleMember(&_Swapper.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Swapper *SwapperCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Swapper.Contract.GetRoleMember(&_Swapper.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Swapper *SwapperCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Swapper *SwapperSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Swapper.Contract.GetRoleMemberCount(&_Swapper.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Swapper *SwapperCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Swapper.Contract.GetRoleMemberCount(&_Swapper.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Swapper *SwapperCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Swapper *SwapperSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Swapper.Contract.HasRole(&_Swapper.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Swapper *SwapperCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Swapper.Contract.HasRole(&_Swapper.CallOpts, role, account)
}

// IsCurrentNetwork is a free data retrieval call binding the contract method 0xa771fc5e.
//
// Solidity: function isCurrentNetwork(string network_) view returns(bool)
func (_Swapper *SwapperCaller) IsCurrentNetwork(opts *bind.CallOpts, network_ string) (bool, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "isCurrentNetwork", network_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentNetwork is a free data retrieval call binding the contract method 0xa771fc5e.
//
// Solidity: function isCurrentNetwork(string network_) view returns(bool)
func (_Swapper *SwapperSession) IsCurrentNetwork(network_ string) (bool, error) {
	return _Swapper.Contract.IsCurrentNetwork(&_Swapper.CallOpts, network_)
}

// IsCurrentNetwork is a free data retrieval call binding the contract method 0xa771fc5e.
//
// Solidity: function isCurrentNetwork(string network_) view returns(bool)
func (_Swapper *SwapperCallerSession) IsCurrentNetwork(network_ string) (bool, error) {
	return _Swapper.Contract.IsCurrentNetwork(&_Swapper.CallOpts, network_)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(string)
func (_Swapper *SwapperCaller) Network(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(string)
func (_Swapper *SwapperSession) Network() (string, error) {
	return _Swapper.Contract.Network(&_Swapper.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(string)
func (_Swapper *SwapperCallerSession) Network() (string, error) {
	return _Swapper.Contract.Network(&_Swapper.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Swapper *SwapperCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Swapper *SwapperSession) ProxiableUUID() ([32]byte, error) {
	return _Swapper.Contract.ProxiableUUID(&_Swapper.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Swapper *SwapperCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Swapper.Contract.ProxiableUUID(&_Swapper.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Swapper *SwapperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Swapper *SwapperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Swapper.Contract.SupportsInterface(&_Swapper.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Swapper *SwapperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Swapper.Contract.SupportsInterface(&_Swapper.CallOpts, interfaceId)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Swapper *SwapperCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Swapper *SwapperSession) UniswapV2Router() (common.Address, error) {
	return _Swapper.Contract.UniswapV2Router(&_Swapper.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Swapper *SwapperCallerSession) UniswapV2Router() (common.Address, error) {
	return _Swapper.Contract.UniswapV2Router(&_Swapper.CallOpts)
}

// SwapperInit is a paid mutator transaction binding the contract method 0x6c991123.
//
// Solidity: function __Swapper_init(string network_, address bridgeAddress_, address uniswapV2RouterAddress_, address[] operators_) returns()
func (_Swapper *SwapperTransactor) SwapperInit(opts *bind.TransactOpts, network_ string, bridgeAddress_ common.Address, uniswapV2RouterAddress_ common.Address, operators_ []common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "__Swapper_init", network_, bridgeAddress_, uniswapV2RouterAddress_, operators_)
}

// SwapperInit is a paid mutator transaction binding the contract method 0x6c991123.
//
// Solidity: function __Swapper_init(string network_, address bridgeAddress_, address uniswapV2RouterAddress_, address[] operators_) returns()
func (_Swapper *SwapperSession) SwapperInit(network_ string, bridgeAddress_ common.Address, uniswapV2RouterAddress_ common.Address, operators_ []common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperInit(&_Swapper.TransactOpts, network_, bridgeAddress_, uniswapV2RouterAddress_, operators_)
}

// SwapperInit is a paid mutator transaction binding the contract method 0x6c991123.
//
// Solidity: function __Swapper_init(string network_, address bridgeAddress_, address uniswapV2RouterAddress_, address[] operators_) returns()
func (_Swapper *SwapperTransactorSession) SwapperInit(network_ string, bridgeAddress_ common.Address, uniswapV2RouterAddress_ common.Address, operators_ []common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperInit(&_Swapper.TransactOpts, network_, bridgeAddress_, uniswapV2RouterAddress_, operators_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Swapper *SwapperTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Swapper *SwapperSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.GrantRole(&_Swapper.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Swapper *SwapperTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.GrantRole(&_Swapper.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Swapper *SwapperTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Swapper *SwapperSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.RenounceRole(&_Swapper.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Swapper *SwapperTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.RenounceRole(&_Swapper.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Swapper *SwapperTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Swapper *SwapperSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.RevokeRole(&_Swapper.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Swapper *SwapperTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.RevokeRole(&_Swapper.TransactOpts, role, account)
}

// SwapAndRoute is a paid mutator transaction binding the contract method 0x8b2cec21.
//
// Solidity: function swapAndRoute((uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_) returns()
func (_Swapper *SwapperTransactor) SwapAndRoute(opts *bind.TransactOpts, swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swapAndRoute", swapParams_, destinationDepositParams_)
}

// SwapAndRoute is a paid mutator transaction binding the contract method 0x8b2cec21.
//
// Solidity: function swapAndRoute((uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_) returns()
func (_Swapper *SwapperSession) SwapAndRoute(swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.Contract.SwapAndRoute(&_Swapper.TransactOpts, swapParams_, destinationDepositParams_)
}

// SwapAndRoute is a paid mutator transaction binding the contract method 0x8b2cec21.
//
// Solidity: function swapAndRoute((uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_) returns()
func (_Swapper *SwapperTransactorSession) SwapAndRoute(swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.Contract.SwapAndRoute(&_Swapper.TransactOpts, swapParams_, destinationDepositParams_)
}

// SwapETHAndRoute is a paid mutator transaction binding the contract method 0x5e35acc2.
//
// Solidity: function swapETHAndRoute((uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_) payable returns()
func (_Swapper *SwapperTransactor) SwapETHAndRoute(opts *bind.TransactOpts, swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swapETHAndRoute", swapParams_, destinationDepositParams_)
}

// SwapETHAndRoute is a paid mutator transaction binding the contract method 0x5e35acc2.
//
// Solidity: function swapETHAndRoute((uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_) payable returns()
func (_Swapper *SwapperSession) SwapETHAndRoute(swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.Contract.SwapETHAndRoute(&_Swapper.TransactOpts, swapParams_, destinationDepositParams_)
}

// SwapETHAndRoute is a paid mutator transaction binding the contract method 0x5e35acc2.
//
// Solidity: function swapETHAndRoute((uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_) payable returns()
func (_Swapper *SwapperTransactorSession) SwapETHAndRoute(swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.Contract.SwapETHAndRoute(&_Swapper.TransactOpts, swapParams_, destinationDepositParams_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Swapper *SwapperTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Swapper *SwapperSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.UpgradeTo(&_Swapper.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Swapper *SwapperTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Swapper.Contract.UpgradeTo(&_Swapper.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Swapper *SwapperTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Swapper *SwapperSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Swapper.Contract.UpgradeToAndCall(&_Swapper.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Swapper *SwapperTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Swapper.Contract.UpgradeToAndCall(&_Swapper.TransactOpts, newImplementation, data)
}

// WithdrawSwapAndRoute is a paid mutator transaction binding the contract method 0x08739b6e.
//
// Solidity: function withdrawSwapAndRoute((address,uint256,bytes32,uint256,bool,bytes[]) withdrawParams_, (uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_, (string,string,bool,uint16) fallbackDepositParams_) returns()
func (_Swapper *SwapperTransactor) WithdrawSwapAndRoute(opts *bind.TransactOpts, withdrawParams_ ISwapperWithdrawParams, swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams, fallbackDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "withdrawSwapAndRoute", withdrawParams_, swapParams_, destinationDepositParams_, fallbackDepositParams_)
}

// WithdrawSwapAndRoute is a paid mutator transaction binding the contract method 0x08739b6e.
//
// Solidity: function withdrawSwapAndRoute((address,uint256,bytes32,uint256,bool,bytes[]) withdrawParams_, (uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_, (string,string,bool,uint16) fallbackDepositParams_) returns()
func (_Swapper *SwapperSession) WithdrawSwapAndRoute(withdrawParams_ ISwapperWithdrawParams, swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams, fallbackDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.Contract.WithdrawSwapAndRoute(&_Swapper.TransactOpts, withdrawParams_, swapParams_, destinationDepositParams_, fallbackDepositParams_)
}

// WithdrawSwapAndRoute is a paid mutator transaction binding the contract method 0x08739b6e.
//
// Solidity: function withdrawSwapAndRoute((address,uint256,bytes32,uint256,bool,bytes[]) withdrawParams_, (uint256,uint256,uint256,address[],bool) swapParams_, (string,string,bool,uint16) destinationDepositParams_, (string,string,bool,uint16) fallbackDepositParams_) returns()
func (_Swapper *SwapperTransactorSession) WithdrawSwapAndRoute(withdrawParams_ ISwapperWithdrawParams, swapParams_ ISwapperSwapParams, destinationDepositParams_ ISwapperDepositParams, fallbackDepositParams_ ISwapperDepositParams) (*types.Transaction, error) {
	return _Swapper.Contract.WithdrawSwapAndRoute(&_Swapper.TransactOpts, withdrawParams_, swapParams_, destinationDepositParams_, fallbackDepositParams_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swapper *SwapperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swapper *SwapperSession) Receive() (*types.Transaction, error) {
	return _Swapper.Contract.Receive(&_Swapper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Swapper *SwapperTransactorSession) Receive() (*types.Transaction, error) {
	return _Swapper.Contract.Receive(&_Swapper.TransactOpts)
}

// SwapperAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Swapper contract.
type SwapperAdminChangedIterator struct {
	Event *SwapperAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperAdminChanged represents a AdminChanged event raised by the Swapper contract.
type SwapperAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Swapper *SwapperFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SwapperAdminChangedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SwapperAdminChangedIterator{contract: _Swapper.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Swapper *SwapperFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SwapperAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperAdminChanged)
				if err := _Swapper.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Swapper *SwapperFilterer) ParseAdminChanged(log types.Log) (*SwapperAdminChanged, error) {
	event := new(SwapperAdminChanged)
	if err := _Swapper.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Swapper contract.
type SwapperBeaconUpgradedIterator struct {
	Event *SwapperBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperBeaconUpgraded represents a BeaconUpgraded event raised by the Swapper contract.
type SwapperBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Swapper *SwapperFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SwapperBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SwapperBeaconUpgradedIterator{contract: _Swapper.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Swapper *SwapperFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SwapperBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperBeaconUpgraded)
				if err := _Swapper.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Swapper *SwapperFilterer) ParseBeaconUpgraded(log types.Log) (*SwapperBeaconUpgraded, error) {
	event := new(SwapperBeaconUpgraded)
	if err := _Swapper.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperCrossChainERC20DepositedIterator is returned from FilterCrossChainERC20Deposited and is used to iterate over the raw logs and unpacked data for CrossChainERC20Deposited events raised by the Swapper contract.
type SwapperCrossChainERC20DepositedIterator struct {
	Event *SwapperCrossChainERC20Deposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperCrossChainERC20DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperCrossChainERC20Deposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperCrossChainERC20Deposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperCrossChainERC20DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperCrossChainERC20DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperCrossChainERC20Deposited represents a CrossChainERC20Deposited event raised by the Swapper contract.
type SwapperCrossChainERC20Deposited struct {
	Token      common.Address
	Amount     *big.Int
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainERC20Deposited is a free log retrieval operation binding the contract event 0x0ae7cd8ebafa6db31828b578024ffa96e53e2135460291ab6609e5d605ad7112.
//
// Solidity: event CrossChainERC20Deposited(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Swapper *SwapperFilterer) FilterCrossChainERC20Deposited(opts *bind.FilterOpts) (*SwapperCrossChainERC20DepositedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "CrossChainERC20Deposited")
	if err != nil {
		return nil, err
	}
	return &SwapperCrossChainERC20DepositedIterator{contract: _Swapper.contract, event: "CrossChainERC20Deposited", logs: logs, sub: sub}, nil
}

// WatchCrossChainERC20Deposited is a free log subscription operation binding the contract event 0x0ae7cd8ebafa6db31828b578024ffa96e53e2135460291ab6609e5d605ad7112.
//
// Solidity: event CrossChainERC20Deposited(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Swapper *SwapperFilterer) WatchCrossChainERC20Deposited(opts *bind.WatchOpts, sink chan<- *SwapperCrossChainERC20Deposited) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "CrossChainERC20Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperCrossChainERC20Deposited)
				if err := _Swapper.contract.UnpackLog(event, "CrossChainERC20Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainERC20Deposited is a log parse operation binding the contract event 0x0ae7cd8ebafa6db31828b578024ffa96e53e2135460291ab6609e5d605ad7112.
//
// Solidity: event CrossChainERC20Deposited(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Swapper *SwapperFilterer) ParseCrossChainERC20Deposited(log types.Log) (*SwapperCrossChainERC20Deposited, error) {
	event := new(SwapperCrossChainERC20Deposited)
	if err := _Swapper.contract.UnpackLog(event, "CrossChainERC20Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperCrossChainERC20FallbackDepositedIterator is returned from FilterCrossChainERC20FallbackDeposited and is used to iterate over the raw logs and unpacked data for CrossChainERC20FallbackDeposited events raised by the Swapper contract.
type SwapperCrossChainERC20FallbackDepositedIterator struct {
	Event *SwapperCrossChainERC20FallbackDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperCrossChainERC20FallbackDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperCrossChainERC20FallbackDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperCrossChainERC20FallbackDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperCrossChainERC20FallbackDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperCrossChainERC20FallbackDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperCrossChainERC20FallbackDeposited represents a CrossChainERC20FallbackDeposited event raised by the Swapper contract.
type SwapperCrossChainERC20FallbackDeposited struct {
	Token      common.Address
	Amount     *big.Int
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainERC20FallbackDeposited is a free log retrieval operation binding the contract event 0xb58dab02f158fba6699f59c338ef4c35a4ce4671dac982f5d455353a74a1100a.
//
// Solidity: event CrossChainERC20FallbackDeposited(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Swapper *SwapperFilterer) FilterCrossChainERC20FallbackDeposited(opts *bind.FilterOpts) (*SwapperCrossChainERC20FallbackDepositedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "CrossChainERC20FallbackDeposited")
	if err != nil {
		return nil, err
	}
	return &SwapperCrossChainERC20FallbackDepositedIterator{contract: _Swapper.contract, event: "CrossChainERC20FallbackDeposited", logs: logs, sub: sub}, nil
}

// WatchCrossChainERC20FallbackDeposited is a free log subscription operation binding the contract event 0xb58dab02f158fba6699f59c338ef4c35a4ce4671dac982f5d455353a74a1100a.
//
// Solidity: event CrossChainERC20FallbackDeposited(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Swapper *SwapperFilterer) WatchCrossChainERC20FallbackDeposited(opts *bind.WatchOpts, sink chan<- *SwapperCrossChainERC20FallbackDeposited) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "CrossChainERC20FallbackDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperCrossChainERC20FallbackDeposited)
				if err := _Swapper.contract.UnpackLog(event, "CrossChainERC20FallbackDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainERC20FallbackDeposited is a log parse operation binding the contract event 0xb58dab02f158fba6699f59c338ef4c35a4ce4671dac982f5d455353a74a1100a.
//
// Solidity: event CrossChainERC20FallbackDeposited(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Swapper *SwapperFilterer) ParseCrossChainERC20FallbackDeposited(log types.Log) (*SwapperCrossChainERC20FallbackDeposited, error) {
	event := new(SwapperCrossChainERC20FallbackDeposited)
	if err := _Swapper.contract.UnpackLog(event, "CrossChainERC20FallbackDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperCrossChainNativeDepositedIterator is returned from FilterCrossChainNativeDeposited and is used to iterate over the raw logs and unpacked data for CrossChainNativeDeposited events raised by the Swapper contract.
type SwapperCrossChainNativeDepositedIterator struct {
	Event *SwapperCrossChainNativeDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperCrossChainNativeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperCrossChainNativeDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperCrossChainNativeDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperCrossChainNativeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperCrossChainNativeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperCrossChainNativeDeposited represents a CrossChainNativeDeposited event raised by the Swapper contract.
type SwapperCrossChainNativeDeposited struct {
	Amount     *big.Int
	Receiver   string
	Network    string
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCrossChainNativeDeposited is a free log retrieval operation binding the contract event 0xa9ae83ac6f5adde54b73eac4522a9c8830486d6ee498a8d8158b07428c9cfbfb.
//
// Solidity: event CrossChainNativeDeposited(uint256 amount, string receiver, string network, uint16 referralId)
func (_Swapper *SwapperFilterer) FilterCrossChainNativeDeposited(opts *bind.FilterOpts) (*SwapperCrossChainNativeDepositedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "CrossChainNativeDeposited")
	if err != nil {
		return nil, err
	}
	return &SwapperCrossChainNativeDepositedIterator{contract: _Swapper.contract, event: "CrossChainNativeDeposited", logs: logs, sub: sub}, nil
}

// WatchCrossChainNativeDeposited is a free log subscription operation binding the contract event 0xa9ae83ac6f5adde54b73eac4522a9c8830486d6ee498a8d8158b07428c9cfbfb.
//
// Solidity: event CrossChainNativeDeposited(uint256 amount, string receiver, string network, uint16 referralId)
func (_Swapper *SwapperFilterer) WatchCrossChainNativeDeposited(opts *bind.WatchOpts, sink chan<- *SwapperCrossChainNativeDeposited) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "CrossChainNativeDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperCrossChainNativeDeposited)
				if err := _Swapper.contract.UnpackLog(event, "CrossChainNativeDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainNativeDeposited is a log parse operation binding the contract event 0xa9ae83ac6f5adde54b73eac4522a9c8830486d6ee498a8d8158b07428c9cfbfb.
//
// Solidity: event CrossChainNativeDeposited(uint256 amount, string receiver, string network, uint16 referralId)
func (_Swapper *SwapperFilterer) ParseCrossChainNativeDeposited(log types.Log) (*SwapperCrossChainNativeDeposited, error) {
	event := new(SwapperCrossChainNativeDeposited)
	if err := _Swapper.contract.UnpackLog(event, "CrossChainNativeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Swapper contract.
type SwapperInitializedIterator struct {
	Event *SwapperInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperInitialized represents a Initialized event raised by the Swapper contract.
type SwapperInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swapper *SwapperFilterer) FilterInitialized(opts *bind.FilterOpts) (*SwapperInitializedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SwapperInitializedIterator{contract: _Swapper.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swapper *SwapperFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SwapperInitialized) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperInitialized)
				if err := _Swapper.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Swapper *SwapperFilterer) ParseInitialized(log types.Log) (*SwapperInitialized, error) {
	event := new(SwapperInitialized)
	if err := _Swapper.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperLocalERC20TransferredIterator is returned from FilterLocalERC20Transferred and is used to iterate over the raw logs and unpacked data for LocalERC20Transferred events raised by the Swapper contract.
type SwapperLocalERC20TransferredIterator struct {
	Event *SwapperLocalERC20Transferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperLocalERC20TransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperLocalERC20Transferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperLocalERC20Transferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperLocalERC20TransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperLocalERC20TransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperLocalERC20Transferred represents a LocalERC20Transferred event raised by the Swapper contract.
type SwapperLocalERC20Transferred struct {
	Amount   *big.Int
	Receiver common.Address
	Token    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLocalERC20Transferred is a free log retrieval operation binding the contract event 0x24c6bdf20618505e3f9ae9ea9842038c8746005a739116bde40b8524ab05ce74.
//
// Solidity: event LocalERC20Transferred(uint256 amount, address receiver, address token)
func (_Swapper *SwapperFilterer) FilterLocalERC20Transferred(opts *bind.FilterOpts) (*SwapperLocalERC20TransferredIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "LocalERC20Transferred")
	if err != nil {
		return nil, err
	}
	return &SwapperLocalERC20TransferredIterator{contract: _Swapper.contract, event: "LocalERC20Transferred", logs: logs, sub: sub}, nil
}

// WatchLocalERC20Transferred is a free log subscription operation binding the contract event 0x24c6bdf20618505e3f9ae9ea9842038c8746005a739116bde40b8524ab05ce74.
//
// Solidity: event LocalERC20Transferred(uint256 amount, address receiver, address token)
func (_Swapper *SwapperFilterer) WatchLocalERC20Transferred(opts *bind.WatchOpts, sink chan<- *SwapperLocalERC20Transferred) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "LocalERC20Transferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperLocalERC20Transferred)
				if err := _Swapper.contract.UnpackLog(event, "LocalERC20Transferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLocalERC20Transferred is a log parse operation binding the contract event 0x24c6bdf20618505e3f9ae9ea9842038c8746005a739116bde40b8524ab05ce74.
//
// Solidity: event LocalERC20Transferred(uint256 amount, address receiver, address token)
func (_Swapper *SwapperFilterer) ParseLocalERC20Transferred(log types.Log) (*SwapperLocalERC20Transferred, error) {
	event := new(SwapperLocalERC20Transferred)
	if err := _Swapper.contract.UnpackLog(event, "LocalERC20Transferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperLocalNativeTransferredIterator is returned from FilterLocalNativeTransferred and is used to iterate over the raw logs and unpacked data for LocalNativeTransferred events raised by the Swapper contract.
type SwapperLocalNativeTransferredIterator struct {
	Event *SwapperLocalNativeTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperLocalNativeTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperLocalNativeTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperLocalNativeTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperLocalNativeTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperLocalNativeTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperLocalNativeTransferred represents a LocalNativeTransferred event raised by the Swapper contract.
type SwapperLocalNativeTransferred struct {
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLocalNativeTransferred is a free log retrieval operation binding the contract event 0x75f0c499156ab8a52bfd1492602165a82839b868a1da601439b7fad5a7c73caf.
//
// Solidity: event LocalNativeTransferred(uint256 amount, address receiver)
func (_Swapper *SwapperFilterer) FilterLocalNativeTransferred(opts *bind.FilterOpts) (*SwapperLocalNativeTransferredIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "LocalNativeTransferred")
	if err != nil {
		return nil, err
	}
	return &SwapperLocalNativeTransferredIterator{contract: _Swapper.contract, event: "LocalNativeTransferred", logs: logs, sub: sub}, nil
}

// WatchLocalNativeTransferred is a free log subscription operation binding the contract event 0x75f0c499156ab8a52bfd1492602165a82839b868a1da601439b7fad5a7c73caf.
//
// Solidity: event LocalNativeTransferred(uint256 amount, address receiver)
func (_Swapper *SwapperFilterer) WatchLocalNativeTransferred(opts *bind.WatchOpts, sink chan<- *SwapperLocalNativeTransferred) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "LocalNativeTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperLocalNativeTransferred)
				if err := _Swapper.contract.UnpackLog(event, "LocalNativeTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLocalNativeTransferred is a log parse operation binding the contract event 0x75f0c499156ab8a52bfd1492602165a82839b868a1da601439b7fad5a7c73caf.
//
// Solidity: event LocalNativeTransferred(uint256 amount, address receiver)
func (_Swapper *SwapperFilterer) ParseLocalNativeTransferred(log types.Log) (*SwapperLocalNativeTransferred, error) {
	event := new(SwapperLocalNativeTransferred)
	if err := _Swapper.contract.UnpackLog(event, "LocalNativeTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Swapper contract.
type SwapperRoleAdminChangedIterator struct {
	Event *SwapperRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperRoleAdminChanged represents a RoleAdminChanged event raised by the Swapper contract.
type SwapperRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Swapper *SwapperFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SwapperRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SwapperRoleAdminChangedIterator{contract: _Swapper.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Swapper *SwapperFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SwapperRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperRoleAdminChanged)
				if err := _Swapper.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Swapper *SwapperFilterer) ParseRoleAdminChanged(log types.Log) (*SwapperRoleAdminChanged, error) {
	event := new(SwapperRoleAdminChanged)
	if err := _Swapper.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Swapper contract.
type SwapperRoleGrantedIterator struct {
	Event *SwapperRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperRoleGranted represents a RoleGranted event raised by the Swapper contract.
type SwapperRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Swapper *SwapperFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SwapperRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SwapperRoleGrantedIterator{contract: _Swapper.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Swapper *SwapperFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SwapperRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperRoleGranted)
				if err := _Swapper.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Swapper *SwapperFilterer) ParseRoleGranted(log types.Log) (*SwapperRoleGranted, error) {
	event := new(SwapperRoleGranted)
	if err := _Swapper.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Swapper contract.
type SwapperRoleRevokedIterator struct {
	Event *SwapperRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperRoleRevoked represents a RoleRevoked event raised by the Swapper contract.
type SwapperRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Swapper *SwapperFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SwapperRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SwapperRoleRevokedIterator{contract: _Swapper.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Swapper *SwapperFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SwapperRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperRoleRevoked)
				if err := _Swapper.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Swapper *SwapperFilterer) ParseRoleRevoked(log types.Log) (*SwapperRoleRevoked, error) {
	event := new(SwapperRoleRevoked)
	if err := _Swapper.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperSwappedAndRoutedIterator is returned from FilterSwappedAndRouted and is used to iterate over the raw logs and unpacked data for SwappedAndRouted events raised by the Swapper contract.
type SwapperSwappedAndRoutedIterator struct {
	Event *SwapperSwappedAndRouted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperSwappedAndRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperSwappedAndRouted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperSwappedAndRouted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperSwappedAndRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperSwappedAndRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperSwappedAndRouted represents a SwappedAndRouted event raised by the Swapper contract.
type SwapperSwappedAndRouted struct {
	Sender                   common.Address
	SwapParams               ISwapperSwapParams
	DestinationDepositParams ISwapperDepositParams
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSwappedAndRouted is a free log retrieval operation binding the contract event 0x5a34d0edd3a46825bf9d5578bb926ae347d6bb174d894b60f05122f877e7d70d.
//
// Solidity: event SwappedAndRouted(address sender, (uint256,uint256,uint256,address[],bool) swapParams, (string,string,bool,uint16) destinationDepositParams)
func (_Swapper *SwapperFilterer) FilterSwappedAndRouted(opts *bind.FilterOpts) (*SwapperSwappedAndRoutedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "SwappedAndRouted")
	if err != nil {
		return nil, err
	}
	return &SwapperSwappedAndRoutedIterator{contract: _Swapper.contract, event: "SwappedAndRouted", logs: logs, sub: sub}, nil
}

// WatchSwappedAndRouted is a free log subscription operation binding the contract event 0x5a34d0edd3a46825bf9d5578bb926ae347d6bb174d894b60f05122f877e7d70d.
//
// Solidity: event SwappedAndRouted(address sender, (uint256,uint256,uint256,address[],bool) swapParams, (string,string,bool,uint16) destinationDepositParams)
func (_Swapper *SwapperFilterer) WatchSwappedAndRouted(opts *bind.WatchOpts, sink chan<- *SwapperSwappedAndRouted) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "SwappedAndRouted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperSwappedAndRouted)
				if err := _Swapper.contract.UnpackLog(event, "SwappedAndRouted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwappedAndRouted is a log parse operation binding the contract event 0x5a34d0edd3a46825bf9d5578bb926ae347d6bb174d894b60f05122f877e7d70d.
//
// Solidity: event SwappedAndRouted(address sender, (uint256,uint256,uint256,address[],bool) swapParams, (string,string,bool,uint16) destinationDepositParams)
func (_Swapper *SwapperFilterer) ParseSwappedAndRouted(log types.Log) (*SwapperSwappedAndRouted, error) {
	event := new(SwapperSwappedAndRouted)
	if err := _Swapper.contract.UnpackLog(event, "SwappedAndRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperSwappedETHAndRoutedIterator is returned from FilterSwappedETHAndRouted and is used to iterate over the raw logs and unpacked data for SwappedETHAndRouted events raised by the Swapper contract.
type SwapperSwappedETHAndRoutedIterator struct {
	Event *SwapperSwappedETHAndRouted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperSwappedETHAndRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperSwappedETHAndRouted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperSwappedETHAndRouted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperSwappedETHAndRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperSwappedETHAndRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperSwappedETHAndRouted represents a SwappedETHAndRouted event raised by the Swapper contract.
type SwapperSwappedETHAndRouted struct {
	Sender                   common.Address
	SwapParams               ISwapperSwapParams
	DestinationDepositParams ISwapperDepositParams
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSwappedETHAndRouted is a free log retrieval operation binding the contract event 0x4cdb39bf70dc36e2ae287ec175b1c3759b318725edb4080130bae004a6dfaed1.
//
// Solidity: event SwappedETHAndRouted(address sender, (uint256,uint256,uint256,address[],bool) swapParams, (string,string,bool,uint16) destinationDepositParams)
func (_Swapper *SwapperFilterer) FilterSwappedETHAndRouted(opts *bind.FilterOpts) (*SwapperSwappedETHAndRoutedIterator, error) {

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "SwappedETHAndRouted")
	if err != nil {
		return nil, err
	}
	return &SwapperSwappedETHAndRoutedIterator{contract: _Swapper.contract, event: "SwappedETHAndRouted", logs: logs, sub: sub}, nil
}

// WatchSwappedETHAndRouted is a free log subscription operation binding the contract event 0x4cdb39bf70dc36e2ae287ec175b1c3759b318725edb4080130bae004a6dfaed1.
//
// Solidity: event SwappedETHAndRouted(address sender, (uint256,uint256,uint256,address[],bool) swapParams, (string,string,bool,uint16) destinationDepositParams)
func (_Swapper *SwapperFilterer) WatchSwappedETHAndRouted(opts *bind.WatchOpts, sink chan<- *SwapperSwappedETHAndRouted) (event.Subscription, error) {

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "SwappedETHAndRouted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperSwappedETHAndRouted)
				if err := _Swapper.contract.UnpackLog(event, "SwappedETHAndRouted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwappedETHAndRouted is a log parse operation binding the contract event 0x4cdb39bf70dc36e2ae287ec175b1c3759b318725edb4080130bae004a6dfaed1.
//
// Solidity: event SwappedETHAndRouted(address sender, (uint256,uint256,uint256,address[],bool) swapParams, (string,string,bool,uint16) destinationDepositParams)
func (_Swapper *SwapperFilterer) ParseSwappedETHAndRouted(log types.Log) (*SwapperSwappedETHAndRouted, error) {
	event := new(SwapperSwappedETHAndRouted)
	if err := _Swapper.contract.UnpackLog(event, "SwappedETHAndRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapperUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Swapper contract.
type SwapperUpgradedIterator struct {
	Event *SwapperUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapperUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapperUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapperUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapperUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapperUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapperUpgraded represents a Upgraded event raised by the Swapper contract.
type SwapperUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Swapper *SwapperFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SwapperUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Swapper.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SwapperUpgradedIterator{contract: _Swapper.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Swapper *SwapperFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SwapperUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Swapper.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapperUpgraded)
				if err := _Swapper.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Swapper *SwapperFilterer) ParseUpgraded(log types.Log) (*SwapperUpgraded, error) {
	event := new(SwapperUpgraded)
	if err := _Swapper.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
