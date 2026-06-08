// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationToken\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"BridgedNativeAndSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"DepositedERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"DepositedERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destinationToken\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDestinationAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"DepositedERC20AndSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"DepositedERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralId\",\"type\":\"uint16\"}],\"name\":\"DepositedNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC1155\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"name\":\"WithdrawnERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"WithdrawnNative\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"signaturesThreshold_\",\"type\":\"uint256\"}],\"name\":\"__Bridge_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"signaturesThreshold_\",\"type\":\"uint256\"}],\"name\":\"__Signers_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"}],\"name\":\"addHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"}],\"name\":\"addSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"}],\"name\":\"containsHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"receiver_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId_\",\"type\":\"uint16\"}],\"name\":\"depositERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"receiver_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId_\",\"type\":\"uint16\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"destinationToken_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"receiver_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId_\",\"type\":\"uint16\"}],\"name\":\"depositERC20AndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"receiver_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"referralId_\",\"type\":\"uint16\"}],\"name\":\"depositERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"receiver_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"referralId_\",\"type\":\"uint16\"}],\"name\":\"depositNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"destinationToken_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minDestinationAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapDeadline_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"receiver_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network_\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"referralId_\",\"type\":\"uint16\"}],\"name\":\"depositNativeAndSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"}],\"name\":\"getERC1155SignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"}],\"name\":\"getERC20SignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"}],\"name\":\"getERC721SignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId_\",\"type\":\"uint256\"}],\"name\":\"getNativeSignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signerToUpdate_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAdding_\",\"type\":\"bool\"}],\"name\":\"getUpdateSignersSignHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers_\",\"type\":\"address[]\"}],\"name\":\"removeSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"signaturesThreshold_\",\"type\":\"uint256\"}],\"name\":\"setSignaturesThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signaturesThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signerToUpdate_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAdding_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof_\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC20Merkelized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped_\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"txNonce_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof_\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures_\",\"type\":\"bytes[]\"}],\"name\":\"withdrawNativeMerkelized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ContainsHash is a free data retrieval call binding the contract method 0x0430285a.
//
// Solidity: function containsHash(bytes32 txHash_, uint256 txNonce_) view returns(bool)
func (_Bridge *BridgeCaller) ContainsHash(opts *bind.CallOpts, txHash_ [32]byte, txNonce_ *big.Int) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "containsHash", txHash_, txNonce_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContainsHash is a free data retrieval call binding the contract method 0x0430285a.
//
// Solidity: function containsHash(bytes32 txHash_, uint256 txNonce_) view returns(bool)
func (_Bridge *BridgeSession) ContainsHash(txHash_ [32]byte, txNonce_ *big.Int) (bool, error) {
	return _Bridge.Contract.ContainsHash(&_Bridge.CallOpts, txHash_, txNonce_)
}

// ContainsHash is a free data retrieval call binding the contract method 0x0430285a.
//
// Solidity: function containsHash(bytes32 txHash_, uint256 txNonce_) view returns(bool)
func (_Bridge *BridgeCallerSession) ContainsHash(txHash_ [32]byte, txNonce_ *big.Int) (bool, error) {
	return _Bridge.Contract.ContainsHash(&_Bridge.CallOpts, txHash_, txNonce_)
}

// GetERC1155SignHash is a free data retrieval call binding the contract method 0xb427d67c.
//
// Solidity: function getERC1155SignHash(address token_, uint256 tokenId_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetERC1155SignHash(opts *bind.CallOpts, token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isWrapped_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getERC1155SignHash", token_, tokenId_, amount_, receiver_, txHash_, txNonce_, chainId_, tokenURI_, isWrapped_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC1155SignHash is a free data retrieval call binding the contract method 0xb427d67c.
//
// Solidity: function getERC1155SignHash(address token_, uint256 tokenId_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetERC1155SignHash(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isWrapped_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC1155SignHash(&_Bridge.CallOpts, token_, tokenId_, amount_, receiver_, txHash_, txNonce_, chainId_, tokenURI_, isWrapped_)
}

// GetERC1155SignHash is a free data retrieval call binding the contract method 0xb427d67c.
//
// Solidity: function getERC1155SignHash(address token_, uint256 tokenId_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetERC1155SignHash(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isWrapped_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC1155SignHash(&_Bridge.CallOpts, token_, tokenId_, amount_, receiver_, txHash_, txNonce_, chainId_, tokenURI_, isWrapped_)
}

// GetERC20SignHash is a free data retrieval call binding the contract method 0xaaba091e.
//
// Solidity: function getERC20SignHash(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetERC20SignHash(opts *bind.CallOpts, token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, isWrapped_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getERC20SignHash", token_, amount_, receiver_, txHash_, txNonce_, chainId_, isWrapped_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC20SignHash is a free data retrieval call binding the contract method 0xaaba091e.
//
// Solidity: function getERC20SignHash(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetERC20SignHash(token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, isWrapped_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC20SignHash(&_Bridge.CallOpts, token_, amount_, receiver_, txHash_, txNonce_, chainId_, isWrapped_)
}

// GetERC20SignHash is a free data retrieval call binding the contract method 0xaaba091e.
//
// Solidity: function getERC20SignHash(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetERC20SignHash(token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, isWrapped_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC20SignHash(&_Bridge.CallOpts, token_, amount_, receiver_, txHash_, txNonce_, chainId_, isWrapped_)
}

// GetERC721SignHash is a free data retrieval call binding the contract method 0xaf94570d.
//
// Solidity: function getERC721SignHash(address token_, uint256 tokenId_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetERC721SignHash(opts *bind.CallOpts, token_ common.Address, tokenId_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isWrapped_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getERC721SignHash", token_, tokenId_, receiver_, txHash_, txNonce_, chainId_, tokenURI_, isWrapped_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetERC721SignHash is a free data retrieval call binding the contract method 0xaf94570d.
//
// Solidity: function getERC721SignHash(address token_, uint256 tokenId_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetERC721SignHash(token_ common.Address, tokenId_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isWrapped_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC721SignHash(&_Bridge.CallOpts, token_, tokenId_, receiver_, txHash_, txNonce_, chainId_, tokenURI_, isWrapped_)
}

// GetERC721SignHash is a free data retrieval call binding the contract method 0xaf94570d.
//
// Solidity: function getERC721SignHash(address token_, uint256 tokenId_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_, string tokenURI_, bool isWrapped_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetERC721SignHash(token_ common.Address, tokenId_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int, tokenURI_ string, isWrapped_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetERC721SignHash(&_Bridge.CallOpts, token_, tokenId_, receiver_, txHash_, txNonce_, chainId_, tokenURI_, isWrapped_)
}

// GetNativeSignHash is a free data retrieval call binding the contract method 0x337e03a9.
//
// Solidity: function getNativeSignHash(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetNativeSignHash(opts *bind.CallOpts, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getNativeSignHash", amount_, receiver_, txHash_, txNonce_, chainId_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNativeSignHash is a free data retrieval call binding the contract method 0x337e03a9.
//
// Solidity: function getNativeSignHash(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetNativeSignHash(amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int) ([32]byte, error) {
	return _Bridge.Contract.GetNativeSignHash(&_Bridge.CallOpts, amount_, receiver_, txHash_, txNonce_, chainId_)
}

// GetNativeSignHash is a free data retrieval call binding the contract method 0x337e03a9.
//
// Solidity: function getNativeSignHash(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, uint256 chainId_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetNativeSignHash(amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, chainId_ *big.Int) ([32]byte, error) {
	return _Bridge.Contract.GetNativeSignHash(&_Bridge.CallOpts, amount_, receiver_, txHash_, txNonce_, chainId_)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Bridge *BridgeCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Bridge *BridgeSession) GetSigners() ([]common.Address, error) {
	return _Bridge.Contract.GetSigners(&_Bridge.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_Bridge *BridgeCallerSession) GetSigners() ([]common.Address, error) {
	return _Bridge.Contract.GetSigners(&_Bridge.CallOpts)
}

// GetUpdateSignersSignHash is a free data retrieval call binding the contract method 0x68d44771.
//
// Solidity: function getUpdateSignersSignHash(address signerToUpdate_, uint256 startTime_, uint256 deadline_, uint256 txNonce_, bool isAdding_) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetUpdateSignersSignHash(opts *bind.CallOpts, signerToUpdate_ common.Address, startTime_ *big.Int, deadline_ *big.Int, txNonce_ *big.Int, isAdding_ bool) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getUpdateSignersSignHash", signerToUpdate_, startTime_, deadline_, txNonce_, isAdding_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUpdateSignersSignHash is a free data retrieval call binding the contract method 0x68d44771.
//
// Solidity: function getUpdateSignersSignHash(address signerToUpdate_, uint256 startTime_, uint256 deadline_, uint256 txNonce_, bool isAdding_) pure returns(bytes32)
func (_Bridge *BridgeSession) GetUpdateSignersSignHash(signerToUpdate_ common.Address, startTime_ *big.Int, deadline_ *big.Int, txNonce_ *big.Int, isAdding_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetUpdateSignersSignHash(&_Bridge.CallOpts, signerToUpdate_, startTime_, deadline_, txNonce_, isAdding_)
}

// GetUpdateSignersSignHash is a free data retrieval call binding the contract method 0x68d44771.
//
// Solidity: function getUpdateSignersSignHash(address signerToUpdate_, uint256 startTime_, uint256 deadline_, uint256 txNonce_, bool isAdding_) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetUpdateSignersSignHash(signerToUpdate_ common.Address, startTime_ *big.Int, deadline_ *big.Int, txNonce_ *big.Int, isAdding_ bool) ([32]byte, error) {
	return _Bridge.Contract.GetUpdateSignersSignHash(&_Bridge.CallOpts, signerToUpdate_, startTime_, deadline_, txNonce_, isAdding_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeSession) ProxiableUUID() ([32]byte, error) {
	return _Bridge.Contract.ProxiableUUID(&_Bridge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bridge *BridgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bridge.Contract.ProxiableUUID(&_Bridge.CallOpts)
}

// SignaturesThreshold is a free data retrieval call binding the contract method 0x39ce73c7.
//
// Solidity: function signaturesThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) SignaturesThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "signaturesThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignaturesThreshold is a free data retrieval call binding the contract method 0x39ce73c7.
//
// Solidity: function signaturesThreshold() view returns(uint256)
func (_Bridge *BridgeSession) SignaturesThreshold() (*big.Int, error) {
	return _Bridge.Contract.SignaturesThreshold(&_Bridge.CallOpts)
}

// SignaturesThreshold is a free data retrieval call binding the contract method 0x39ce73c7.
//
// Solidity: function signaturesThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) SignaturesThreshold() (*big.Int, error) {
	return _Bridge.Contract.SignaturesThreshold(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) UsedHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "usedHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.UsedHashes(&_Bridge.CallOpts, arg0)
}

// UsedHashes is a free data retrieval call binding the contract method 0xaef18bf7.
//
// Solidity: function usedHashes(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) UsedHashes(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.UsedHashes(&_Bridge.CallOpts, arg0)
}

// BridgeInit is a paid mutator transaction binding the contract method 0x8338fcd8.
//
// Solidity: function __Bridge_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeTransactor) BridgeInit(opts *bind.TransactOpts, signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "__Bridge_init", signers_, signaturesThreshold_)
}

// BridgeInit is a paid mutator transaction binding the contract method 0x8338fcd8.
//
// Solidity: function __Bridge_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeSession) BridgeInit(signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeInit(&_Bridge.TransactOpts, signers_, signaturesThreshold_)
}

// BridgeInit is a paid mutator transaction binding the contract method 0x8338fcd8.
//
// Solidity: function __Bridge_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeTransactorSession) BridgeInit(signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeInit(&_Bridge.TransactOpts, signers_, signaturesThreshold_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeTransactor) SignersInit(opts *bind.TransactOpts, signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "__Signers_init", signers_, signaturesThreshold_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeSession) SignersInit(signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SignersInit(&_Bridge.TransactOpts, signers_, signaturesThreshold_)
}

// SignersInit is a paid mutator transaction binding the contract method 0x09a55841.
//
// Solidity: function __Signers_init(address[] signers_, uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeTransactorSession) SignersInit(signers_ []common.Address, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SignersInit(&_Bridge.TransactOpts, signers_, signaturesThreshold_)
}

// AddHash is a paid mutator transaction binding the contract method 0x5bd5429d.
//
// Solidity: function addHash(bytes32 txHash_, uint256 txNonce_) returns()
func (_Bridge *BridgeTransactor) AddHash(opts *bind.TransactOpts, txHash_ [32]byte, txNonce_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addHash", txHash_, txNonce_)
}

// AddHash is a paid mutator transaction binding the contract method 0x5bd5429d.
//
// Solidity: function addHash(bytes32 txHash_, uint256 txNonce_) returns()
func (_Bridge *BridgeSession) AddHash(txHash_ [32]byte, txNonce_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddHash(&_Bridge.TransactOpts, txHash_, txNonce_)
}

// AddHash is a paid mutator transaction binding the contract method 0x5bd5429d.
//
// Solidity: function addHash(bytes32 txHash_, uint256 txNonce_) returns()
func (_Bridge *BridgeTransactorSession) AddHash(txHash_ [32]byte, txNonce_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddHash(&_Bridge.TransactOpts, txHash_, txNonce_)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers_) returns()
func (_Bridge *BridgeTransactor) AddSigners(opts *bind.TransactOpts, signers_ []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addSigners", signers_)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers_) returns()
func (_Bridge *BridgeSession) AddSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddSigners(&_Bridge.TransactOpts, signers_)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signers_) returns()
func (_Bridge *BridgeTransactorSession) AddSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddSigners(&_Bridge.TransactOpts, signers_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x5fe27756.
//
// Solidity: function depositERC1155(address token_, uint256 tokenId_, uint256 amount_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactor) DepositERC1155(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC1155", token_, tokenId_, amount_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x5fe27756.
//
// Solidity: function depositERC1155(address token_, uint256 tokenId_, uint256 amount_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeSession) DepositERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC1155 is a paid mutator transaction binding the contract method 0x5fe27756.
//
// Solidity: function depositERC1155(address token_, uint256 tokenId_, uint256 amount_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbe6f93d4.
//
// Solidity: function depositERC20(address token_, uint256 amount_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactor) DepositERC20(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20", token_, amount_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbe6f93d4.
//
// Solidity: function depositERC20(address token_, uint256 amount_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeSession) DepositERC20(token_ common.Address, amount_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, token_, amount_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xbe6f93d4.
//
// Solidity: function depositERC20(address token_, uint256 amount_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC20(token_ common.Address, amount_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, token_, amount_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC20AndSwap is a paid mutator transaction binding the contract method 0x4310578b.
//
// Solidity: function depositERC20AndSwap(address token_, uint256 amount_, string destinationToken_, uint256 minDestinationAmount_, uint256 swapDeadline_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactor) DepositERC20AndSwap(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, destinationToken_ string, minDestinationAmount_ *big.Int, swapDeadline_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20AndSwap", token_, amount_, destinationToken_, minDestinationAmount_, swapDeadline_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC20AndSwap is a paid mutator transaction binding the contract method 0x4310578b.
//
// Solidity: function depositERC20AndSwap(address token_, uint256 amount_, string destinationToken_, uint256 minDestinationAmount_, uint256 swapDeadline_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeSession) DepositERC20AndSwap(token_ common.Address, amount_ *big.Int, destinationToken_ string, minDestinationAmount_ *big.Int, swapDeadline_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20AndSwap(&_Bridge.TransactOpts, token_, amount_, destinationToken_, minDestinationAmount_, swapDeadline_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC20AndSwap is a paid mutator transaction binding the contract method 0x4310578b.
//
// Solidity: function depositERC20AndSwap(address token_, uint256 amount_, string destinationToken_, uint256 minDestinationAmount_, uint256 swapDeadline_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC20AndSwap(token_ common.Address, amount_ *big.Int, destinationToken_ string, minDestinationAmount_ *big.Int, swapDeadline_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20AndSwap(&_Bridge.TransactOpts, token_, amount_, destinationToken_, minDestinationAmount_, swapDeadline_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x90946c6e.
//
// Solidity: function depositERC721(address token_, uint256 tokenId_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactor) DepositERC721(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC721", token_, tokenId_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x90946c6e.
//
// Solidity: function depositERC721(address token_, uint256 tokenId_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeSession) DepositERC721(token_ common.Address, tokenId_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, token_, tokenId_, receiver_, network_, isWrapped_, referralId_)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0x90946c6e.
//
// Solidity: function depositERC721(address token_, uint256 tokenId_, string receiver_, string network_, bool isWrapped_, uint16 referralId_) returns()
func (_Bridge *BridgeTransactorSession) DepositERC721(token_ common.Address, tokenId_ *big.Int, receiver_ string, network_ string, isWrapped_ bool, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC721(&_Bridge.TransactOpts, token_, tokenId_, receiver_, network_, isWrapped_, referralId_)
}

// DepositNative is a paid mutator transaction binding the contract method 0x90e208ee.
//
// Solidity: function depositNative(string receiver_, string network_, uint16 referralId_) payable returns()
func (_Bridge *BridgeTransactor) DepositNative(opts *bind.TransactOpts, receiver_ string, network_ string, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositNative", receiver_, network_, referralId_)
}

// DepositNative is a paid mutator transaction binding the contract method 0x90e208ee.
//
// Solidity: function depositNative(string receiver_, string network_, uint16 referralId_) payable returns()
func (_Bridge *BridgeSession) DepositNative(receiver_ string, network_ string, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, receiver_, network_, referralId_)
}

// DepositNative is a paid mutator transaction binding the contract method 0x90e208ee.
//
// Solidity: function depositNative(string receiver_, string network_, uint16 referralId_) payable returns()
func (_Bridge *BridgeTransactorSession) DepositNative(receiver_ string, network_ string, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNative(&_Bridge.TransactOpts, receiver_, network_, referralId_)
}

// DepositNativeAndSwap is a paid mutator transaction binding the contract method 0xe7bdb15c.
//
// Solidity: function depositNativeAndSwap(string destinationToken_, uint256 minDestinationAmount_, uint256 swapDeadline_, string receiver_, string network_, uint16 referralId_) payable returns()
func (_Bridge *BridgeTransactor) DepositNativeAndSwap(opts *bind.TransactOpts, destinationToken_ string, minDestinationAmount_ *big.Int, swapDeadline_ *big.Int, receiver_ string, network_ string, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositNativeAndSwap", destinationToken_, minDestinationAmount_, swapDeadline_, receiver_, network_, referralId_)
}

// DepositNativeAndSwap is a paid mutator transaction binding the contract method 0xe7bdb15c.
//
// Solidity: function depositNativeAndSwap(string destinationToken_, uint256 minDestinationAmount_, uint256 swapDeadline_, string receiver_, string network_, uint16 referralId_) payable returns()
func (_Bridge *BridgeSession) DepositNativeAndSwap(destinationToken_ string, minDestinationAmount_ *big.Int, swapDeadline_ *big.Int, receiver_ string, network_ string, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNativeAndSwap(&_Bridge.TransactOpts, destinationToken_, minDestinationAmount_, swapDeadline_, receiver_, network_, referralId_)
}

// DepositNativeAndSwap is a paid mutator transaction binding the contract method 0xe7bdb15c.
//
// Solidity: function depositNativeAndSwap(string destinationToken_, uint256 minDestinationAmount_, uint256 swapDeadline_, string receiver_, string network_, uint16 referralId_) payable returns()
func (_Bridge *BridgeTransactorSession) DepositNativeAndSwap(destinationToken_ string, minDestinationAmount_ *big.Int, swapDeadline_ *big.Int, receiver_ string, network_ string, referralId_ uint16) (*types.Transaction, error) {
	return _Bridge.Contract.DepositNativeAndSwap(&_Bridge.TransactOpts, destinationToken_, minDestinationAmount_, swapDeadline_, receiver_, network_, referralId_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155BatchReceived(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC1155Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Bridge *BridgeTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Bridge.Contract.OnERC721Received(&_Bridge.TransactOpts, arg0, arg1, arg2, arg3)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers_) returns()
func (_Bridge *BridgeTransactor) RemoveSigners(opts *bind.TransactOpts, signers_ []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeSigners", signers_)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers_) returns()
func (_Bridge *BridgeSession) RemoveSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveSigners(&_Bridge.TransactOpts, signers_)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signers_) returns()
func (_Bridge *BridgeTransactorSession) RemoveSigners(signers_ []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveSigners(&_Bridge.TransactOpts, signers_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeTransactor) SetSignaturesThreshold(opts *bind.TransactOpts, signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setSignaturesThreshold", signaturesThreshold_)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeSession) SetSignaturesThreshold(signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetSignaturesThreshold(&_Bridge.TransactOpts, signaturesThreshold_)
}

// SetSignaturesThreshold is a paid mutator transaction binding the contract method 0xbf1fe08f.
//
// Solidity: function setSignaturesThreshold(uint256 signaturesThreshold_) returns()
func (_Bridge *BridgeTransactorSession) SetSignaturesThreshold(signaturesThreshold_ *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetSignaturesThreshold(&_Bridge.TransactOpts, signaturesThreshold_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8f826791.
//
// Solidity: function updateSigner(address signerToUpdate_, uint256 startTime_, uint256 deadline_, uint256 txNonce_, bool isAdding_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) UpdateSigner(opts *bind.TransactOpts, signerToUpdate_ common.Address, startTime_ *big.Int, deadline_ *big.Int, txNonce_ *big.Int, isAdding_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateSigner", signerToUpdate_, startTime_, deadline_, txNonce_, isAdding_, signatures_)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8f826791.
//
// Solidity: function updateSigner(address signerToUpdate_, uint256 startTime_, uint256 deadline_, uint256 txNonce_, bool isAdding_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) UpdateSigner(signerToUpdate_ common.Address, startTime_ *big.Int, deadline_ *big.Int, txNonce_ *big.Int, isAdding_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateSigner(&_Bridge.TransactOpts, signerToUpdate_, startTime_, deadline_, txNonce_, isAdding_, signatures_)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0x8f826791.
//
// Solidity: function updateSigner(address signerToUpdate_, uint256 startTime_, uint256 deadline_, uint256 txNonce_, bool isAdding_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) UpdateSigner(signerToUpdate_ common.Address, startTime_ *big.Int, deadline_ *big.Int, txNonce_ *big.Int, isAdding_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateSigner(&_Bridge.TransactOpts, signerToUpdate_, startTime_, deadline_, txNonce_, isAdding_, signatures_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Bridge *BridgeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Bridge *BridgeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeTo(&_Bridge.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Bridge *BridgeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeTo(&_Bridge.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToAndCall(&_Bridge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.UpgradeToAndCall(&_Bridge.TransactOpts, newImplementation, data)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xb3953d44.
//
// Solidity: function withdrawERC1155(address token_, uint256 tokenId_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC1155(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC1155", token_, tokenId_, amount_, receiver_, txHash_, txNonce_, tokenURI_, isWrapped_, signatures_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xb3953d44.
//
// Solidity: function withdrawERC1155(address token_, uint256 tokenId_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, receiver_, txHash_, txNonce_, tokenURI_, isWrapped_, signatures_)
}

// WithdrawERC1155 is a paid mutator transaction binding the contract method 0xb3953d44.
//
// Solidity: function withdrawERC1155(address token_, uint256 tokenId_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC1155(token_ common.Address, tokenId_ *big.Int, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC1155(&_Bridge.TransactOpts, token_, tokenId_, amount_, receiver_, txHash_, txNonce_, tokenURI_, isWrapped_, signatures_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x94995fc4.
//
// Solidity: function withdrawERC20(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC20(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20", token_, amount_, receiver_, txHash_, txNonce_, isWrapped_, signatures_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x94995fc4.
//
// Solidity: function withdrawERC20(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC20(token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, token_, amount_, receiver_, txHash_, txNonce_, isWrapped_, signatures_)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x94995fc4.
//
// Solidity: function withdrawERC20(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC20(token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20(&_Bridge.TransactOpts, token_, amount_, receiver_, txHash_, txNonce_, isWrapped_, signatures_)
}

// WithdrawERC20Merkelized is a paid mutator transaction binding the contract method 0xfa242299.
//
// Solidity: function withdrawERC20Merkelized(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bool isWrapped_, bytes32[] merkleProof_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC20Merkelized(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isWrapped_ bool, merkleProof_ [][32]byte, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC20Merkelized", token_, amount_, receiver_, txHash_, txNonce_, isWrapped_, merkleProof_, signatures_)
}

// WithdrawERC20Merkelized is a paid mutator transaction binding the contract method 0xfa242299.
//
// Solidity: function withdrawERC20Merkelized(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bool isWrapped_, bytes32[] merkleProof_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC20Merkelized(token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isWrapped_ bool, merkleProof_ [][32]byte, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20Merkelized(&_Bridge.TransactOpts, token_, amount_, receiver_, txHash_, txNonce_, isWrapped_, merkleProof_, signatures_)
}

// WithdrawERC20Merkelized is a paid mutator transaction binding the contract method 0xfa242299.
//
// Solidity: function withdrawERC20Merkelized(address token_, uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bool isWrapped_, bytes32[] merkleProof_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC20Merkelized(token_ common.Address, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, isWrapped_ bool, merkleProof_ [][32]byte, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC20Merkelized(&_Bridge.TransactOpts, token_, amount_, receiver_, txHash_, txNonce_, isWrapped_, merkleProof_, signatures_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x7eb9d447.
//
// Solidity: function withdrawERC721(address token_, uint256 tokenId_, address receiver_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawERC721(opts *bind.TransactOpts, token_ common.Address, tokenId_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawERC721", token_, tokenId_, receiver_, txHash_, txNonce_, tokenURI_, isWrapped_, signatures_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x7eb9d447.
//
// Solidity: function withdrawERC721(address token_, uint256 tokenId_, address receiver_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawERC721(token_ common.Address, tokenId_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, token_, tokenId_, receiver_, txHash_, txNonce_, tokenURI_, isWrapped_, signatures_)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x7eb9d447.
//
// Solidity: function withdrawERC721(address token_, uint256 tokenId_, address receiver_, bytes32 txHash_, uint256 txNonce_, string tokenURI_, bool isWrapped_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawERC721(token_ common.Address, tokenId_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, tokenURI_ string, isWrapped_ bool, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawERC721(&_Bridge.TransactOpts, token_, tokenId_, receiver_, txHash_, txNonce_, tokenURI_, isWrapped_, signatures_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x1c3d9c87.
//
// Solidity: function withdrawNative(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawNative(opts *bind.TransactOpts, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNative", amount_, receiver_, txHash_, txNonce_, signatures_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x1c3d9c87.
//
// Solidity: function withdrawNative(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawNative(amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, amount_, receiver_, txHash_, txNonce_, signatures_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x1c3d9c87.
//
// Solidity: function withdrawNative(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNative(amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNative(&_Bridge.TransactOpts, amount_, receiver_, txHash_, txNonce_, signatures_)
}

// WithdrawNativeMerkelized is a paid mutator transaction binding the contract method 0xfb53a90e.
//
// Solidity: function withdrawNativeMerkelized(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bytes32[] merkleProof_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactor) WithdrawNativeMerkelized(opts *bind.TransactOpts, amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, merkleProof_ [][32]byte, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdrawNativeMerkelized", amount_, receiver_, txHash_, txNonce_, merkleProof_, signatures_)
}

// WithdrawNativeMerkelized is a paid mutator transaction binding the contract method 0xfb53a90e.
//
// Solidity: function withdrawNativeMerkelized(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bytes32[] merkleProof_, bytes[] signatures_) returns()
func (_Bridge *BridgeSession) WithdrawNativeMerkelized(amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, merkleProof_ [][32]byte, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNativeMerkelized(&_Bridge.TransactOpts, amount_, receiver_, txHash_, txNonce_, merkleProof_, signatures_)
}

// WithdrawNativeMerkelized is a paid mutator transaction binding the contract method 0xfb53a90e.
//
// Solidity: function withdrawNativeMerkelized(uint256 amount_, address receiver_, bytes32 txHash_, uint256 txNonce_, bytes32[] merkleProof_, bytes[] signatures_) returns()
func (_Bridge *BridgeTransactorSession) WithdrawNativeMerkelized(amount_ *big.Int, receiver_ common.Address, txHash_ [32]byte, txNonce_ *big.Int, merkleProof_ [][32]byte, signatures_ [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.WithdrawNativeMerkelized(&_Bridge.TransactOpts, amount_, receiver_, txHash_, txNonce_, merkleProof_, signatures_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Bridge contract.
type BridgeAdminChangedIterator struct {
	Event *BridgeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminChanged)
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
		it.Event = new(BridgeAdminChanged)
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
func (it *BridgeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminChanged represents a AdminChanged event raised by the Bridge contract.
type BridgeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Bridge *BridgeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BridgeAdminChangedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeAdminChangedIterator{contract: _Bridge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Bridge *BridgeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseAdminChanged(log types.Log) (*BridgeAdminChanged, error) {
	event := new(BridgeAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Bridge contract.
type BridgeBeaconUpgradedIterator struct {
	Event *BridgeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBeaconUpgraded)
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
		it.Event = new(BridgeBeaconUpgraded)
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
func (it *BridgeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBeaconUpgraded represents a BeaconUpgraded event raised by the Bridge contract.
type BridgeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Bridge *BridgeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BridgeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBeaconUpgradedIterator{contract: _Bridge.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Bridge *BridgeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBeaconUpgraded)
				if err := _Bridge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseBeaconUpgraded(log types.Log) (*BridgeBeaconUpgraded, error) {
	event := new(BridgeBeaconUpgraded)
	if err := _Bridge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBridgedNativeAndSwappedIterator is returned from FilterBridgedNativeAndSwapped and is used to iterate over the raw logs and unpacked data for BridgedNativeAndSwapped events raised by the Bridge contract.
type BridgeBridgedNativeAndSwappedIterator struct {
	Event *BridgeBridgedNativeAndSwapped // Event containing the contract specifics and raw log

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
func (it *BridgeBridgedNativeAndSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgedNativeAndSwapped)
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
		it.Event = new(BridgeBridgedNativeAndSwapped)
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
func (it *BridgeBridgedNativeAndSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgedNativeAndSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgedNativeAndSwapped represents a BridgedNativeAndSwapped event raised by the Bridge contract.
type BridgeBridgedNativeAndSwapped struct {
	Amount               *big.Int
	DestinationToken     string
	MinDestinationAmount *big.Int
	SwapDeadline         *big.Int
	Receiver             string
	Network              string
	ReferralId           uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBridgedNativeAndSwapped is a free log retrieval operation binding the contract event 0xbd9b774a1f4510c791f62926503049cdfd6cfd1624c9a083935312f6aeae9dcf.
//
// Solidity: event BridgedNativeAndSwapped(uint256 amount, string destinationToken, uint256 minDestinationAmount, uint256 swapDeadline, string receiver, string network, uint16 referralId)
func (_Bridge *BridgeFilterer) FilterBridgedNativeAndSwapped(opts *bind.FilterOpts) (*BridgeBridgedNativeAndSwappedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgedNativeAndSwapped")
	if err != nil {
		return nil, err
	}
	return &BridgeBridgedNativeAndSwappedIterator{contract: _Bridge.contract, event: "BridgedNativeAndSwapped", logs: logs, sub: sub}, nil
}

// WatchBridgedNativeAndSwapped is a free log subscription operation binding the contract event 0xbd9b774a1f4510c791f62926503049cdfd6cfd1624c9a083935312f6aeae9dcf.
//
// Solidity: event BridgedNativeAndSwapped(uint256 amount, string destinationToken, uint256 minDestinationAmount, uint256 swapDeadline, string receiver, string network, uint16 referralId)
func (_Bridge *BridgeFilterer) WatchBridgedNativeAndSwapped(opts *bind.WatchOpts, sink chan<- *BridgeBridgedNativeAndSwapped) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgedNativeAndSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgedNativeAndSwapped)
				if err := _Bridge.contract.UnpackLog(event, "BridgedNativeAndSwapped", log); err != nil {
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

// ParseBridgedNativeAndSwapped is a log parse operation binding the contract event 0xbd9b774a1f4510c791f62926503049cdfd6cfd1624c9a083935312f6aeae9dcf.
//
// Solidity: event BridgedNativeAndSwapped(uint256 amount, string destinationToken, uint256 minDestinationAmount, uint256 swapDeadline, string receiver, string network, uint16 referralId)
func (_Bridge *BridgeFilterer) ParseBridgedNativeAndSwapped(log types.Log) (*BridgeBridgedNativeAndSwapped, error) {
	event := new(BridgeBridgedNativeAndSwapped)
	if err := _Bridge.contract.UnpackLog(event, "BridgedNativeAndSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC1155Iterator is returned from FilterDepositedERC1155 and is used to iterate over the raw logs and unpacked data for DepositedERC1155 events raised by the Bridge contract.
type BridgeDepositedERC1155Iterator struct {
	Event *BridgeDepositedERC1155 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC1155)
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
		it.Event = new(BridgeDepositedERC1155)
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
func (it *BridgeDepositedERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC1155 represents a DepositedERC1155 event raised by the Bridge contract.
type BridgeDepositedERC1155 struct {
	Token      common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC1155 is a free log retrieval operation binding the contract event 0x6f2f9c82f8808cf5cb4332789c02c0c95fad4f7586899cc1bc426f254debf165.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) FilterDepositedERC1155(opts *bind.FilterOpts) (*BridgeDepositedERC1155Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC1155Iterator{contract: _Bridge.contract, event: "DepositedERC1155", logs: logs, sub: sub}, nil
}

// WatchDepositedERC1155 is a free log subscription operation binding the contract event 0x6f2f9c82f8808cf5cb4332789c02c0c95fad4f7586899cc1bc426f254debf165.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) WatchDepositedERC1155(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC1155) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC1155)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
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

// ParseDepositedERC1155 is a log parse operation binding the contract event 0x6f2f9c82f8808cf5cb4332789c02c0c95fad4f7586899cc1bc426f254debf165.
//
// Solidity: event DepositedERC1155(address token, uint256 tokenId, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) ParseDepositedERC1155(log types.Log) (*BridgeDepositedERC1155, error) {
	event := new(BridgeDepositedERC1155)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC20Iterator is returned from FilterDepositedERC20 and is used to iterate over the raw logs and unpacked data for DepositedERC20 events raised by the Bridge contract.
type BridgeDepositedERC20Iterator struct {
	Event *BridgeDepositedERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC20)
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
		it.Event = new(BridgeDepositedERC20)
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
func (it *BridgeDepositedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC20 represents a DepositedERC20 event raised by the Bridge contract.
type BridgeDepositedERC20 struct {
	Token      common.Address
	Amount     *big.Int
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC20 is a free log retrieval operation binding the contract event 0xfc11ca985085252b65dba84249af6977f9d4cb2acc79ee95fff01985d1b56275.
//
// Solidity: event DepositedERC20(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) FilterDepositedERC20(opts *bind.FilterOpts) (*BridgeDepositedERC20Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC20Iterator{contract: _Bridge.contract, event: "DepositedERC20", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20 is a free log subscription operation binding the contract event 0xfc11ca985085252b65dba84249af6977f9d4cb2acc79ee95fff01985d1b56275.
//
// Solidity: event DepositedERC20(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) WatchDepositedERC20(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC20) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC20)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
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

// ParseDepositedERC20 is a log parse operation binding the contract event 0xfc11ca985085252b65dba84249af6977f9d4cb2acc79ee95fff01985d1b56275.
//
// Solidity: event DepositedERC20(address token, uint256 amount, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) ParseDepositedERC20(log types.Log) (*BridgeDepositedERC20, error) {
	event := new(BridgeDepositedERC20)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC20AndSwappedIterator is returned from FilterDepositedERC20AndSwapped and is used to iterate over the raw logs and unpacked data for DepositedERC20AndSwapped events raised by the Bridge contract.
type BridgeDepositedERC20AndSwappedIterator struct {
	Event *BridgeDepositedERC20AndSwapped // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC20AndSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC20AndSwapped)
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
		it.Event = new(BridgeDepositedERC20AndSwapped)
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
func (it *BridgeDepositedERC20AndSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC20AndSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC20AndSwapped represents a DepositedERC20AndSwapped event raised by the Bridge contract.
type BridgeDepositedERC20AndSwapped struct {
	Token                common.Address
	Amount               *big.Int
	DestinationToken     string
	MinDestinationAmount *big.Int
	SwapDeadline         *big.Int
	Receiver             string
	Network              string
	IsWrapped            bool
	ReferralId           uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC20AndSwapped is a free log retrieval operation binding the contract event 0x11c2030e768e3b5b27824d0d8b672c67f863c9a58ae85cf87d41363bbb125281.
//
// Solidity: event DepositedERC20AndSwapped(address token, uint256 amount, string destinationToken, uint256 minDestinationAmount, uint256 swapDeadline, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) FilterDepositedERC20AndSwapped(opts *bind.FilterOpts) (*BridgeDepositedERC20AndSwappedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC20AndSwapped")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC20AndSwappedIterator{contract: _Bridge.contract, event: "DepositedERC20AndSwapped", logs: logs, sub: sub}, nil
}

// WatchDepositedERC20AndSwapped is a free log subscription operation binding the contract event 0x11c2030e768e3b5b27824d0d8b672c67f863c9a58ae85cf87d41363bbb125281.
//
// Solidity: event DepositedERC20AndSwapped(address token, uint256 amount, string destinationToken, uint256 minDestinationAmount, uint256 swapDeadline, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) WatchDepositedERC20AndSwapped(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC20AndSwapped) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC20AndSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC20AndSwapped)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC20AndSwapped", log); err != nil {
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

// ParseDepositedERC20AndSwapped is a log parse operation binding the contract event 0x11c2030e768e3b5b27824d0d8b672c67f863c9a58ae85cf87d41363bbb125281.
//
// Solidity: event DepositedERC20AndSwapped(address token, uint256 amount, string destinationToken, uint256 minDestinationAmount, uint256 swapDeadline, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) ParseDepositedERC20AndSwapped(log types.Log) (*BridgeDepositedERC20AndSwapped, error) {
	event := new(BridgeDepositedERC20AndSwapped)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC20AndSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedERC721Iterator is returned from FilterDepositedERC721 and is used to iterate over the raw logs and unpacked data for DepositedERC721 events raised by the Bridge contract.
type BridgeDepositedERC721Iterator struct {
	Event *BridgeDepositedERC721 // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedERC721)
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
		it.Event = new(BridgeDepositedERC721)
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
func (it *BridgeDepositedERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedERC721 represents a DepositedERC721 event raised by the Bridge contract.
type BridgeDepositedERC721 struct {
	Token      common.Address
	TokenId    *big.Int
	Receiver   string
	Network    string
	IsWrapped  bool
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedERC721 is a free log retrieval operation binding the contract event 0x5a035a04c3f86dbf0cfe44e37374a1c7d8ad9d3b2542a39acaab6e3cb18e9738.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) FilterDepositedERC721(opts *bind.FilterOpts) (*BridgeDepositedERC721Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedERC721Iterator{contract: _Bridge.contract, event: "DepositedERC721", logs: logs, sub: sub}, nil
}

// WatchDepositedERC721 is a free log subscription operation binding the contract event 0x5a035a04c3f86dbf0cfe44e37374a1c7d8ad9d3b2542a39acaab6e3cb18e9738.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) WatchDepositedERC721(opts *bind.WatchOpts, sink chan<- *BridgeDepositedERC721) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedERC721)
				if err := _Bridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
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

// ParseDepositedERC721 is a log parse operation binding the contract event 0x5a035a04c3f86dbf0cfe44e37374a1c7d8ad9d3b2542a39acaab6e3cb18e9738.
//
// Solidity: event DepositedERC721(address token, uint256 tokenId, string receiver, string network, bool isWrapped, uint16 referralId)
func (_Bridge *BridgeFilterer) ParseDepositedERC721(log types.Log) (*BridgeDepositedERC721, error) {
	event := new(BridgeDepositedERC721)
	if err := _Bridge.contract.UnpackLog(event, "DepositedERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositedNativeIterator is returned from FilterDepositedNative and is used to iterate over the raw logs and unpacked data for DepositedNative events raised by the Bridge contract.
type BridgeDepositedNativeIterator struct {
	Event *BridgeDepositedNative // Event containing the contract specifics and raw log

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
func (it *BridgeDepositedNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositedNative)
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
		it.Event = new(BridgeDepositedNative)
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
func (it *BridgeDepositedNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositedNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositedNative represents a DepositedNative event raised by the Bridge contract.
type BridgeDepositedNative struct {
	Amount     *big.Int
	Receiver   string
	Network    string
	ReferralId uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositedNative is a free log retrieval operation binding the contract event 0x618bd7f3201fe12a7051eb9e8f45ac82a6648b19a07a79c9189ac480546e68c1.
//
// Solidity: event DepositedNative(uint256 amount, string receiver, string network, uint16 referralId)
func (_Bridge *BridgeFilterer) FilterDepositedNative(opts *bind.FilterOpts) (*BridgeDepositedNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositedNativeIterator{contract: _Bridge.contract, event: "DepositedNative", logs: logs, sub: sub}, nil
}

// WatchDepositedNative is a free log subscription operation binding the contract event 0x618bd7f3201fe12a7051eb9e8f45ac82a6648b19a07a79c9189ac480546e68c1.
//
// Solidity: event DepositedNative(uint256 amount, string receiver, string network, uint16 referralId)
func (_Bridge *BridgeFilterer) WatchDepositedNative(opts *bind.WatchOpts, sink chan<- *BridgeDepositedNative) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositedNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositedNative)
				if err := _Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
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

// ParseDepositedNative is a log parse operation binding the contract event 0x618bd7f3201fe12a7051eb9e8f45ac82a6648b19a07a79c9189ac480546e68c1.
//
// Solidity: event DepositedNative(uint256 amount, string receiver, string network, uint16 referralId)
func (_Bridge *BridgeFilterer) ParseDepositedNative(log types.Log) (*BridgeDepositedNative, error) {
	event := new(BridgeDepositedNative)
	if err := _Bridge.contract.UnpackLog(event, "DepositedNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bridge contract.
type BridgeUpgradedIterator struct {
	Event *BridgeUpgraded // Event containing the contract specifics and raw log

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
func (it *BridgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUpgraded)
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
		it.Event = new(BridgeUpgraded)
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
func (it *BridgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUpgraded represents a Upgraded event raised by the Bridge contract.
type BridgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BridgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUpgradedIterator{contract: _Bridge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bridge *BridgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BridgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUpgraded)
				if err := _Bridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUpgraded(log types.Log) (*BridgeUpgraded, error) {
	event := new(BridgeUpgraded)
	if err := _Bridge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnERC1155Iterator is returned from FilterWithdrawnERC1155 and is used to iterate over the raw logs and unpacked data for WithdrawnERC1155 events raised by the Bridge contract.
type BridgeWithdrawnERC1155Iterator struct {
	Event *BridgeWithdrawnERC1155 // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnERC1155Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnERC1155)
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
		it.Event = new(BridgeWithdrawnERC1155)
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
func (it *BridgeWithdrawnERC1155Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnERC1155Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnERC1155 represents a WithdrawnERC1155 event raised by the Bridge contract.
type BridgeWithdrawnERC1155 struct {
	Token     common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Receiver  common.Address
	TokenURI  string
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC1155 is a free log retrieval operation binding the contract event 0x605d13472d10382e722662eac29514069901f8f837b7e41d7beac67da87342f3.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, uint256 amount, address receiver, string tokenURI, bool isWrapped)
func (_Bridge *BridgeFilterer) FilterWithdrawnERC1155(opts *bind.FilterOpts) (*BridgeWithdrawnERC1155Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnERC1155Iterator{contract: _Bridge.contract, event: "WithdrawnERC1155", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC1155 is a free log subscription operation binding the contract event 0x605d13472d10382e722662eac29514069901f8f837b7e41d7beac67da87342f3.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, uint256 amount, address receiver, string tokenURI, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchWithdrawnERC1155(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnERC1155) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnERC1155")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnERC1155)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
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

// ParseWithdrawnERC1155 is a log parse operation binding the contract event 0x605d13472d10382e722662eac29514069901f8f837b7e41d7beac67da87342f3.
//
// Solidity: event WithdrawnERC1155(address token, uint256 tokenId, uint256 amount, address receiver, string tokenURI, bool isWrapped)
func (_Bridge *BridgeFilterer) ParseWithdrawnERC1155(log types.Log) (*BridgeWithdrawnERC1155, error) {
	event := new(BridgeWithdrawnERC1155)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC1155", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnERC20Iterator is returned from FilterWithdrawnERC20 and is used to iterate over the raw logs and unpacked data for WithdrawnERC20 events raised by the Bridge contract.
type BridgeWithdrawnERC20Iterator struct {
	Event *BridgeWithdrawnERC20 // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnERC20)
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
		it.Event = new(BridgeWithdrawnERC20)
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
func (it *BridgeWithdrawnERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnERC20 represents a WithdrawnERC20 event raised by the Bridge contract.
type BridgeWithdrawnERC20 struct {
	Token     common.Address
	Amount    *big.Int
	Receiver  common.Address
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC20 is a free log retrieval operation binding the contract event 0xc872b0b229e2cd5fd7032daefa79bbf979dac748a63ba99d8552c6d0e32299ec.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, address receiver, bool isWrapped)
func (_Bridge *BridgeFilterer) FilterWithdrawnERC20(opts *bind.FilterOpts) (*BridgeWithdrawnERC20Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnERC20Iterator{contract: _Bridge.contract, event: "WithdrawnERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC20 is a free log subscription operation binding the contract event 0xc872b0b229e2cd5fd7032daefa79bbf979dac748a63ba99d8552c6d0e32299ec.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, address receiver, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchWithdrawnERC20(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnERC20) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnERC20)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
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

// ParseWithdrawnERC20 is a log parse operation binding the contract event 0xc872b0b229e2cd5fd7032daefa79bbf979dac748a63ba99d8552c6d0e32299ec.
//
// Solidity: event WithdrawnERC20(address token, uint256 amount, address receiver, bool isWrapped)
func (_Bridge *BridgeFilterer) ParseWithdrawnERC20(log types.Log) (*BridgeWithdrawnERC20, error) {
	event := new(BridgeWithdrawnERC20)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnERC721Iterator is returned from FilterWithdrawnERC721 and is used to iterate over the raw logs and unpacked data for WithdrawnERC721 events raised by the Bridge contract.
type BridgeWithdrawnERC721Iterator struct {
	Event *BridgeWithdrawnERC721 // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnERC721)
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
		it.Event = new(BridgeWithdrawnERC721)
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
func (it *BridgeWithdrawnERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnERC721 represents a WithdrawnERC721 event raised by the Bridge contract.
type BridgeWithdrawnERC721 struct {
	Token     common.Address
	TokenId   *big.Int
	Receiver  common.Address
	TokenURI  string
	IsWrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnERC721 is a free log retrieval operation binding the contract event 0x80e3603e2865de4b313cf490fa04546f36b39ce73a11598521bdd9f75cef11cf.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, address receiver, string tokenURI, bool isWrapped)
func (_Bridge *BridgeFilterer) FilterWithdrawnERC721(opts *bind.FilterOpts) (*BridgeWithdrawnERC721Iterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnERC721Iterator{contract: _Bridge.contract, event: "WithdrawnERC721", logs: logs, sub: sub}, nil
}

// WatchWithdrawnERC721 is a free log subscription operation binding the contract event 0x80e3603e2865de4b313cf490fa04546f36b39ce73a11598521bdd9f75cef11cf.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, address receiver, string tokenURI, bool isWrapped)
func (_Bridge *BridgeFilterer) WatchWithdrawnERC721(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnERC721) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnERC721)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
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

// ParseWithdrawnERC721 is a log parse operation binding the contract event 0x80e3603e2865de4b313cf490fa04546f36b39ce73a11598521bdd9f75cef11cf.
//
// Solidity: event WithdrawnERC721(address token, uint256 tokenId, address receiver, string tokenURI, bool isWrapped)
func (_Bridge *BridgeFilterer) ParseWithdrawnERC721(log types.Log) (*BridgeWithdrawnERC721, error) {
	event := new(BridgeWithdrawnERC721)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawnNativeIterator is returned from FilterWithdrawnNative and is used to iterate over the raw logs and unpacked data for WithdrawnNative events raised by the Bridge contract.
type BridgeWithdrawnNativeIterator struct {
	Event *BridgeWithdrawnNative // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawnNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawnNative)
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
		it.Event = new(BridgeWithdrawnNative)
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
func (it *BridgeWithdrawnNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawnNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawnNative represents a WithdrawnNative event raised by the Bridge contract.
type BridgeWithdrawnNative struct {
	Amount   *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnNative is a free log retrieval operation binding the contract event 0x2aaa8f83c299480568ccc5c6b1cdd781d7ef20d477b5891cf957f7e2ff3ab667.
//
// Solidity: event WithdrawnNative(uint256 amount, address receiver)
func (_Bridge *BridgeFilterer) FilterWithdrawnNative(opts *bind.FilterOpts) (*BridgeWithdrawnNativeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawnNativeIterator{contract: _Bridge.contract, event: "WithdrawnNative", logs: logs, sub: sub}, nil
}

// WatchWithdrawnNative is a free log subscription operation binding the contract event 0x2aaa8f83c299480568ccc5c6b1cdd781d7ef20d477b5891cf957f7e2ff3ab667.
//
// Solidity: event WithdrawnNative(uint256 amount, address receiver)
func (_Bridge *BridgeFilterer) WatchWithdrawnNative(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawnNative) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawnNative")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawnNative)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
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

// ParseWithdrawnNative is a log parse operation binding the contract event 0x2aaa8f83c299480568ccc5c6b1cdd781d7ef20d477b5891cf957f7e2ff3ab667.
//
// Solidity: event WithdrawnNative(uint256 amount, address receiver)
func (_Bridge *BridgeFilterer) ParseWithdrawnNative(log types.Log) (*BridgeWithdrawnNative, error) {
	event := new(BridgeWithdrawnNative)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawnNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
