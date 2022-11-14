// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// NFTNFTParams is an auto generated low-level Go binding around an user-defined struct.
type NFTNFTParams struct {
	TokenId  *big.Int
	Amount   *big.Int
	TokenURI string
}

// ERC1155MetaData contains all meta data concerning the ERC1155 contract.
var ERC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"config_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BatchTransferNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BurnNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"MintNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransferNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"batchTransferNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structNFT.NFTParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burnNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"typeID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytesToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"contractIGlobalConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAdminRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"internalType\":\"structNFT.NFTParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintNFT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155MetaData.ABI instead.
var ERC1155ABI = ERC1155MetaData.ABI

// ERC1155 is an auto generated Go binding around an Ethereum contract.
type ERC1155 struct {
	ERC1155Caller     // Read-only binding to the contract
	ERC1155Transactor // Write-only binding to the contract
	ERC1155Filterer   // Log filterer for contract events
}

// ERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155Session struct {
	Contract     *ERC1155          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155CallerSession struct {
	Contract *ERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155TransactorSession struct {
	Contract     *ERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155Raw struct {
	Contract *ERC1155 // Generic contract binding to access the raw methods on
}

// ERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155CallerRaw struct {
	Contract *ERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// ERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155TransactorRaw struct {
	Contract *ERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155 creates a new instance of ERC1155, bound to a specific deployed contract.
func NewERC1155(address common.Address, backend bind.ContractBackend) (*ERC1155, error) {
	contract, err := bindERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155{ERC1155Caller: ERC1155Caller{contract: contract}, ERC1155Transactor: ERC1155Transactor{contract: contract}, ERC1155Filterer: ERC1155Filterer{contract: contract}}, nil
}

// NewERC1155Caller creates a new read-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Caller(address common.Address, caller bind.ContractCaller) (*ERC1155Caller, error) {
	contract, err := bindERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Caller{contract: contract}, nil
}

// NewERC1155Transactor creates a new write-only instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155Transactor, error) {
	contract, err := bindERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155Transactor{contract: contract}, nil
}

// NewERC1155Filterer creates a new log filterer instance of ERC1155, bound to a specific deployed contract.
func NewERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155Filterer, error) {
	contract, err := bindERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155Filterer{contract: contract}, nil
}

// bindERC1155 binds a generic wrapper to an already deployed contract.
func bindERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.ERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.ERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155 *ERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155 *ERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155 *ERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSHASH is a free data retrieval call binding the contract method 0xc87d2cd7.
//
// Solidity: function ADDRESS_HASH() view returns(bytes32)
func (_ERC1155 *ERC1155Caller) ADDRESSHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "ADDRESS_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADDRESSHASH is a free data retrieval call binding the contract method 0xc87d2cd7.
//
// Solidity: function ADDRESS_HASH() view returns(bytes32)
func (_ERC1155 *ERC1155Session) ADDRESSHASH() ([32]byte, error) {
	return _ERC1155.Contract.ADDRESSHASH(&_ERC1155.CallOpts)
}

// ADDRESSHASH is a free data retrieval call binding the contract method 0xc87d2cd7.
//
// Solidity: function ADDRESS_HASH() view returns(bytes32)
func (_ERC1155 *ERC1155CallerSession) ADDRESSHASH() ([32]byte, error) {
	return _ERC1155.Contract.ADDRESSHASH(&_ERC1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.BalanceOf(&_ERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ERC1155 *ERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ERC1155.Contract.BalanceOfBatch(&_ERC1155.CallOpts, accounts, ids)
}

// BytesToAddress is a free data retrieval call binding the contract method 0xf38c363f.
//
// Solidity: function bytesToAddress(bytes32 typeID, bytes data) pure returns(address addr)
func (_ERC1155 *ERC1155Caller) BytesToAddress(opts *bind.CallOpts, typeID [32]byte, data []byte) (common.Address, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "bytesToAddress", typeID, data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BytesToAddress is a free data retrieval call binding the contract method 0xf38c363f.
//
// Solidity: function bytesToAddress(bytes32 typeID, bytes data) pure returns(address addr)
func (_ERC1155 *ERC1155Session) BytesToAddress(typeID [32]byte, data []byte) (common.Address, error) {
	return _ERC1155.Contract.BytesToAddress(&_ERC1155.CallOpts, typeID, data)
}

// BytesToAddress is a free data retrieval call binding the contract method 0xf38c363f.
//
// Solidity: function bytesToAddress(bytes32 typeID, bytes data) pure returns(address addr)
func (_ERC1155 *ERC1155CallerSession) BytesToAddress(typeID [32]byte, data []byte) (common.Address, error) {
	return _ERC1155.Contract.BytesToAddress(&_ERC1155.CallOpts, typeID, data)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_ERC1155 *ERC1155Caller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_ERC1155 *ERC1155Session) Config() (common.Address, error) {
	return _ERC1155.Contract.Config(&_ERC1155.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_ERC1155 *ERC1155CallerSession) Config() (common.Address, error) {
	return _ERC1155.Contract.Config(&_ERC1155.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_ERC1155 *ERC1155Caller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_ERC1155 *ERC1155Session) Exists(id *big.Int) (bool, error) {
	return _ERC1155.Contract.Exists(&_ERC1155.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) Exists(id *big.Int) (bool, error) {
	return _ERC1155.Contract.Exists(&_ERC1155.CallOpts, id)
}

// IsAdminRole is a free data retrieval call binding the contract method 0x6cd1093c.
//
// Solidity: function isAdminRole() view returns(bool)
func (_ERC1155 *ERC1155Caller) IsAdminRole(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isAdminRole")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdminRole is a free data retrieval call binding the contract method 0x6cd1093c.
//
// Solidity: function isAdminRole() view returns(bool)
func (_ERC1155 *ERC1155Session) IsAdminRole() (bool, error) {
	return _ERC1155.Contract.IsAdminRole(&_ERC1155.CallOpts)
}

// IsAdminRole is a free data retrieval call binding the contract method 0x6cd1093c.
//
// Solidity: function isAdminRole() view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsAdminRole() (bool, error) {
	return _ERC1155.Contract.IsAdminRole(&_ERC1155.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ERC1155.Contract.IsApprovedForAll(&_ERC1155.CallOpts, account, operator)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC1155 *ERC1155Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC1155 *ERC1155Session) Paused() (bool, error) {
	return _ERC1155.Contract.Paused(&_ERC1155.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC1155 *ERC1155CallerSession) Paused() (bool, error) {
	return _ERC1155.Contract.Paused(&_ERC1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155 *ERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155.Contract.SupportsInterface(&_ERC1155.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Caller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155Session) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TotalSupply(&_ERC1155.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ERC1155 *ERC1155CallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ERC1155.Contract.TotalSupply(&_ERC1155.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155 *ERC1155Caller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC1155.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155 *ERC1155Session) Uri(tokenId *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_ERC1155 *ERC1155CallerSession) Uri(tokenId *big.Int) (string, error) {
	return _ERC1155.Contract.Uri(&_ERC1155.CallOpts, tokenId)
}

// BatchTransferNFT is a paid mutator transaction binding the contract method 0x777b71b3.
//
// Solidity: function batchTransferNFT(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) BatchTransferNFT(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "batchTransferNFT", to, ids, amounts, data)
}

// BatchTransferNFT is a paid mutator transaction binding the contract method 0x777b71b3.
//
// Solidity: function batchTransferNFT(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Session) BatchTransferNFT(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchTransferNFT(&_ERC1155.TransactOpts, to, ids, amounts, data)
}

// BatchTransferNFT is a paid mutator transaction binding the contract method 0x777b71b3.
//
// Solidity: function batchTransferNFT(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) BatchTransferNFT(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.BatchTransferNFT(&_ERC1155.TransactOpts, to, ids, amounts, data)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenId, uint256 amount) returns()
func (_ERC1155 *ERC1155Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "burn", tokenId, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenId, uint256 amount) returns()
func (_ERC1155 *ERC1155Session) Burn(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Burn(&_ERC1155.TransactOpts, tokenId, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xb390c0ab.
//
// Solidity: function burn(uint256 tokenId, uint256 amount) returns()
func (_ERC1155 *ERC1155TransactorSession) Burn(tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ERC1155.Contract.Burn(&_ERC1155.TransactOpts, tokenId, amount)
}

// BurnNFT is a paid mutator transaction binding the contract method 0xf7926866.
//
// Solidity: function burnNFT((uint256,uint256,string)[] params, bytes data) returns(bool)
func (_ERC1155 *ERC1155Transactor) BurnNFT(opts *bind.TransactOpts, params []NFTNFTParams, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "burnNFT", params, data)
}

// BurnNFT is a paid mutator transaction binding the contract method 0xf7926866.
//
// Solidity: function burnNFT((uint256,uint256,string)[] params, bytes data) returns(bool)
func (_ERC1155 *ERC1155Session) BurnNFT(params []NFTNFTParams, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.BurnNFT(&_ERC1155.TransactOpts, params, data)
}

// BurnNFT is a paid mutator transaction binding the contract method 0xf7926866.
//
// Solidity: function burnNFT((uint256,uint256,string)[] params, bytes data) returns(bool)
func (_ERC1155 *ERC1155TransactorSession) BurnNFT(params []NFTNFTParams, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.BurnNFT(&_ERC1155.TransactOpts, params, data)
}

// MintNFT is a paid mutator transaction binding the contract method 0xb5f14930.
//
// Solidity: function mintNFT((uint256,uint256,string)[] params, bytes data) returns(bool)
func (_ERC1155 *ERC1155Transactor) MintNFT(opts *bind.TransactOpts, params []NFTNFTParams, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "mintNFT", params, data)
}

// MintNFT is a paid mutator transaction binding the contract method 0xb5f14930.
//
// Solidity: function mintNFT((uint256,uint256,string)[] params, bytes data) returns(bool)
func (_ERC1155 *ERC1155Session) MintNFT(params []NFTNFTParams, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.MintNFT(&_ERC1155.TransactOpts, params, data)
}

// MintNFT is a paid mutator transaction binding the contract method 0xb5f14930.
//
// Solidity: function mintNFT((uint256,uint256,string)[] params, bytes data) returns(bool)
func (_ERC1155 *ERC1155TransactorSession) MintNFT(params []NFTNFTParams, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.MintNFT(&_ERC1155.TransactOpts, params, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeBatchTransferFrom(&_ERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.SafeTransferFrom(&_ERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC1155 *ERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetApprovalForAll(&_ERC1155.TransactOpts, operator, approved)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool isPaused) returns()
func (_ERC1155 *ERC1155Transactor) SetPaused(opts *bind.TransactOpts, isPaused bool) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setPaused", isPaused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool isPaused) returns()
func (_ERC1155 *ERC1155Session) SetPaused(isPaused bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetPaused(&_ERC1155.TransactOpts, isPaused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool isPaused) returns()
func (_ERC1155 *ERC1155TransactorSession) SetPaused(isPaused bool) (*types.Transaction, error) {
	return _ERC1155.Contract.SetPaused(&_ERC1155.TransactOpts, isPaused)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string tokenUri) returns()
func (_ERC1155 *ERC1155Transactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, tokenUri string) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "setTokenURI", tokenId, tokenUri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string tokenUri) returns()
func (_ERC1155 *ERC1155Session) SetTokenURI(tokenId *big.Int, tokenUri string) (*types.Transaction, error) {
	return _ERC1155.Contract.SetTokenURI(&_ERC1155.TransactOpts, tokenId, tokenUri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string tokenUri) returns()
func (_ERC1155 *ERC1155TransactorSession) SetTokenURI(tokenId *big.Int, tokenUri string) (*types.Transaction, error) {
	return _ERC1155.Contract.SetTokenURI(&_ERC1155.TransactOpts, tokenId, tokenUri)
}

// TransferNFT is a paid mutator transaction binding the contract method 0x96f6b2fa.
//
// Solidity: function transferNFT(address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Transactor) TransferNFT(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.contract.Transact(opts, "transferNFT", to, id, amount, data)
}

// TransferNFT is a paid mutator transaction binding the contract method 0x96f6b2fa.
//
// Solidity: function transferNFT(address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155Session) TransferNFT(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferNFT(&_ERC1155.TransactOpts, to, id, amount, data)
}

// TransferNFT is a paid mutator transaction binding the contract method 0x96f6b2fa.
//
// Solidity: function transferNFT(address to, uint256 id, uint256 amount, bytes data) returns()
func (_ERC1155 *ERC1155TransactorSession) TransferNFT(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC1155.Contract.TransferNFT(&_ERC1155.TransactOpts, to, id, amount, data)
}

// ERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC1155 contract.
type ERC1155ApprovalForAllIterator struct {
	Event *ERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155ApprovalForAll)
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
		it.Event = new(ERC1155ApprovalForAll)
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
func (it *ERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155ApprovalForAll represents a ApprovalForAll event raised by the ERC1155 contract.
type ERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155ApprovalForAllIterator{contract: _ERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155ApprovalForAll)
				if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ERC1155 *ERC1155Filterer) ParseApprovalForAll(log types.Log) (*ERC1155ApprovalForAll, error) {
	event := new(ERC1155ApprovalForAll)
	if err := _ERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155BatchTransferNFTIterator is returned from FilterBatchTransferNFT and is used to iterate over the raw logs and unpacked data for BatchTransferNFT events raised by the ERC1155 contract.
type ERC1155BatchTransferNFTIterator struct {
	Event *ERC1155BatchTransferNFT // Event containing the contract specifics and raw log

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
func (it *ERC1155BatchTransferNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155BatchTransferNFT)
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
		it.Event = new(ERC1155BatchTransferNFT)
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
func (it *ERC1155BatchTransferNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155BatchTransferNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155BatchTransferNFT represents a BatchTransferNFT event raised by the ERC1155 contract.
type ERC1155BatchTransferNFT struct {
	From     common.Address
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Data     common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferNFT is a free log retrieval operation binding the contract event 0x2a1c80b42dd482d3c706809dedfdb84e8ac662fbc11d1652cb73405573207d73.
//
// Solidity: event BatchTransferNFT(address indexed from, address indexed to, uint256[] tokenIds, uint256[] amounts, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) FilterBatchTransferNFT(opts *bind.FilterOpts, from []common.Address, to []common.Address, data [][]byte) (*ERC1155BatchTransferNFTIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "BatchTransferNFT", fromRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155BatchTransferNFTIterator{contract: _ERC1155.contract, event: "BatchTransferNFT", logs: logs, sub: sub}, nil
}

// WatchBatchTransferNFT is a free log subscription operation binding the contract event 0x2a1c80b42dd482d3c706809dedfdb84e8ac662fbc11d1652cb73405573207d73.
//
// Solidity: event BatchTransferNFT(address indexed from, address indexed to, uint256[] tokenIds, uint256[] amounts, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) WatchBatchTransferNFT(opts *bind.WatchOpts, sink chan<- *ERC1155BatchTransferNFT, from []common.Address, to []common.Address, data [][]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "BatchTransferNFT", fromRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155BatchTransferNFT)
				if err := _ERC1155.contract.UnpackLog(event, "BatchTransferNFT", log); err != nil {
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

// ParseBatchTransferNFT is a log parse operation binding the contract event 0x2a1c80b42dd482d3c706809dedfdb84e8ac662fbc11d1652cb73405573207d73.
//
// Solidity: event BatchTransferNFT(address indexed from, address indexed to, uint256[] tokenIds, uint256[] amounts, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) ParseBatchTransferNFT(log types.Log) (*ERC1155BatchTransferNFT, error) {
	event := new(ERC1155BatchTransferNFT)
	if err := _ERC1155.contract.UnpackLog(event, "BatchTransferNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155BurnNFTIterator is returned from FilterBurnNFT and is used to iterate over the raw logs and unpacked data for BurnNFT events raised by the ERC1155 contract.
type ERC1155BurnNFTIterator struct {
	Event *ERC1155BurnNFT // Event containing the contract specifics and raw log

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
func (it *ERC1155BurnNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155BurnNFT)
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
		it.Event = new(ERC1155BurnNFT)
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
func (it *ERC1155BurnNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155BurnNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155BurnNFT represents a BurnNFT event raised by the ERC1155 contract.
type ERC1155BurnNFT struct {
	Sender  common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Data    common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurnNFT is a free log retrieval operation binding the contract event 0x29bf646a42408763909cd50a4b045ec8fa106fe8e83f9e6e50136ac17838bb44.
//
// Solidity: event BurnNFT(address indexed sender, address indexed to, uint256 tokenId, uint256 amount, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) FilterBurnNFT(opts *bind.FilterOpts, sender []common.Address, to []common.Address, data [][]byte) (*ERC1155BurnNFTIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "BurnNFT", senderRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155BurnNFTIterator{contract: _ERC1155.contract, event: "BurnNFT", logs: logs, sub: sub}, nil
}

// WatchBurnNFT is a free log subscription operation binding the contract event 0x29bf646a42408763909cd50a4b045ec8fa106fe8e83f9e6e50136ac17838bb44.
//
// Solidity: event BurnNFT(address indexed sender, address indexed to, uint256 tokenId, uint256 amount, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) WatchBurnNFT(opts *bind.WatchOpts, sink chan<- *ERC1155BurnNFT, sender []common.Address, to []common.Address, data [][]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "BurnNFT", senderRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155BurnNFT)
				if err := _ERC1155.contract.UnpackLog(event, "BurnNFT", log); err != nil {
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

// ParseBurnNFT is a log parse operation binding the contract event 0x29bf646a42408763909cd50a4b045ec8fa106fe8e83f9e6e50136ac17838bb44.
//
// Solidity: event BurnNFT(address indexed sender, address indexed to, uint256 tokenId, uint256 amount, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) ParseBurnNFT(log types.Log) (*ERC1155BurnNFT, error) {
	event := new(ERC1155BurnNFT)
	if err := _ERC1155.contract.UnpackLog(event, "BurnNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1155 contract.
type ERC1155InitializedIterator struct {
	Event *ERC1155Initialized // Event containing the contract specifics and raw log

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
func (it *ERC1155InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Initialized)
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
		it.Event = new(ERC1155Initialized)
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
func (it *ERC1155InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Initialized represents a Initialized event raised by the ERC1155 contract.
type ERC1155Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155 *ERC1155Filterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1155InitializedIterator, error) {

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1155InitializedIterator{contract: _ERC1155.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1155 *ERC1155Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1155Initialized) (event.Subscription, error) {

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Initialized)
				if err := _ERC1155.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC1155 *ERC1155Filterer) ParseInitialized(log types.Log) (*ERC1155Initialized, error) {
	event := new(ERC1155Initialized)
	if err := _ERC1155.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MintNFTIterator is returned from FilterMintNFT and is used to iterate over the raw logs and unpacked data for MintNFT events raised by the ERC1155 contract.
type ERC1155MintNFTIterator struct {
	Event *ERC1155MintNFT // Event containing the contract specifics and raw log

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
func (it *ERC1155MintNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MintNFT)
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
		it.Event = new(ERC1155MintNFT)
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
func (it *ERC1155MintNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MintNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MintNFT represents a MintNFT event raised by the ERC1155 contract.
type ERC1155MintNFT struct {
	Sender   common.Address
	To       common.Address
	TokenId  *big.Int
	Amount   *big.Int
	TokenURI string
	Data     common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintNFT is a free log retrieval operation binding the contract event 0xf8c14845547984d368bfc02f2341a32c0c92535c2b6557fc63871d56e4899817.
//
// Solidity: event MintNFT(address indexed sender, address indexed to, uint256 tokenId, uint256 amount, string tokenURI, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) FilterMintNFT(opts *bind.FilterOpts, sender []common.Address, to []common.Address, data [][]byte) (*ERC1155MintNFTIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "MintNFT", senderRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MintNFTIterator{contract: _ERC1155.contract, event: "MintNFT", logs: logs, sub: sub}, nil
}

// WatchMintNFT is a free log subscription operation binding the contract event 0xf8c14845547984d368bfc02f2341a32c0c92535c2b6557fc63871d56e4899817.
//
// Solidity: event MintNFT(address indexed sender, address indexed to, uint256 tokenId, uint256 amount, string tokenURI, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) WatchMintNFT(opts *bind.WatchOpts, sink chan<- *ERC1155MintNFT, sender []common.Address, to []common.Address, data [][]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "MintNFT", senderRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MintNFT)
				if err := _ERC1155.contract.UnpackLog(event, "MintNFT", log); err != nil {
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

// ParseMintNFT is a log parse operation binding the contract event 0xf8c14845547984d368bfc02f2341a32c0c92535c2b6557fc63871d56e4899817.
//
// Solidity: event MintNFT(address indexed sender, address indexed to, uint256 tokenId, uint256 amount, string tokenURI, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) ParseMintNFT(log types.Log) (*ERC1155MintNFT, error) {
	event := new(ERC1155MintNFT)
	if err := _ERC1155.contract.UnpackLog(event, "MintNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC1155 contract.
type ERC1155PausedIterator struct {
	Event *ERC1155Paused // Event containing the contract specifics and raw log

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
func (it *ERC1155PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Paused)
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
		it.Event = new(ERC1155Paused)
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
func (it *ERC1155PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Paused represents a Paused event raised by the ERC1155 contract.
type ERC1155Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC1155 *ERC1155Filterer) FilterPaused(opts *bind.FilterOpts) (*ERC1155PausedIterator, error) {

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC1155PausedIterator{contract: _ERC1155.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC1155 *ERC1155Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC1155Paused) (event.Subscription, error) {

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Paused)
				if err := _ERC1155.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC1155 *ERC1155Filterer) ParsePaused(log types.Log) (*ERC1155Paused, error) {
	event := new(ERC1155Paused)
	if err := _ERC1155.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ERC1155 contract.
type ERC1155TransferBatchIterator struct {
	Event *ERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferBatch)
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
		it.Event = new(ERC1155TransferBatch)
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
func (it *ERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferBatch represents a TransferBatch event raised by the ERC1155 contract.
type ERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferBatchIterator{contract: _ERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferBatch)
				if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ERC1155 *ERC1155Filterer) ParseTransferBatch(log types.Log) (*ERC1155TransferBatch, error) {
	event := new(ERC1155TransferBatch)
	if err := _ERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferNFTIterator is returned from FilterTransferNFT and is used to iterate over the raw logs and unpacked data for TransferNFT events raised by the ERC1155 contract.
type ERC1155TransferNFTIterator struct {
	Event *ERC1155TransferNFT // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferNFT)
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
		it.Event = new(ERC1155TransferNFT)
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
func (it *ERC1155TransferNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferNFT represents a TransferNFT event raised by the ERC1155 contract.
type ERC1155TransferNFT struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Amount  *big.Int
	Data    common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferNFT is a free log retrieval operation binding the contract event 0x602e5c8aa5c8dff82bb239f42a524de238e18b9e11bcceecfef63f54afa4555b.
//
// Solidity: event TransferNFT(address indexed from, address indexed to, uint256 tokenId, uint256 amount, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) FilterTransferNFT(opts *bind.FilterOpts, from []common.Address, to []common.Address, data [][]byte) (*ERC1155TransferNFTIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferNFT", fromRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferNFTIterator{contract: _ERC1155.contract, event: "TransferNFT", logs: logs, sub: sub}, nil
}

// WatchTransferNFT is a free log subscription operation binding the contract event 0x602e5c8aa5c8dff82bb239f42a524de238e18b9e11bcceecfef63f54afa4555b.
//
// Solidity: event TransferNFT(address indexed from, address indexed to, uint256 tokenId, uint256 amount, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) WatchTransferNFT(opts *bind.WatchOpts, sink chan<- *ERC1155TransferNFT, from []common.Address, to []common.Address, data [][]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferNFT", fromRule, toRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferNFT)
				if err := _ERC1155.contract.UnpackLog(event, "TransferNFT", log); err != nil {
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

// ParseTransferNFT is a log parse operation binding the contract event 0x602e5c8aa5c8dff82bb239f42a524de238e18b9e11bcceecfef63f54afa4555b.
//
// Solidity: event TransferNFT(address indexed from, address indexed to, uint256 tokenId, uint256 amount, bytes indexed data)
func (_ERC1155 *ERC1155Filterer) ParseTransferNFT(log types.Log) (*ERC1155TransferNFT, error) {
	event := new(ERC1155TransferNFT)
	if err := _ERC1155.contract.UnpackLog(event, "TransferNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ERC1155 contract.
type ERC1155TransferSingleIterator struct {
	Event *ERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *ERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155TransferSingle)
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
		it.Event = new(ERC1155TransferSingle)
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
func (it *ERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155TransferSingle represents a TransferSingle event raised by the ERC1155 contract.
type ERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155TransferSingleIterator{contract: _ERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155TransferSingle)
				if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ERC1155 *ERC1155Filterer) ParseTransferSingle(log types.Log) (*ERC1155TransferSingle, error) {
	event := new(ERC1155TransferSingle)
	if err := _ERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ERC1155 contract.
type ERC1155URIIterator struct {
	Event *ERC1155URI // Event containing the contract specifics and raw log

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
func (it *ERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155URI)
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
		it.Event = new(ERC1155URI)
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
func (it *ERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155URI represents a URI event raised by the ERC1155 contract.
type ERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155URIIterator{contract: _ERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155URI)
				if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ERC1155 *ERC1155Filterer) ParseURI(log types.Log) (*ERC1155URI, error) {
	event := new(ERC1155URI)
	if err := _ERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC1155 contract.
type ERC1155UnpausedIterator struct {
	Event *ERC1155Unpaused // Event containing the contract specifics and raw log

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
func (it *ERC1155UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155Unpaused)
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
		it.Event = new(ERC1155Unpaused)
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
func (it *ERC1155UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155Unpaused represents a Unpaused event raised by the ERC1155 contract.
type ERC1155Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC1155 *ERC1155Filterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC1155UnpausedIterator, error) {

	logs, sub, err := _ERC1155.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC1155UnpausedIterator{contract: _ERC1155.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC1155 *ERC1155Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC1155Unpaused) (event.Subscription, error) {

	logs, sub, err := _ERC1155.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155Unpaused)
				if err := _ERC1155.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC1155 *ERC1155Filterer) ParseUnpaused(log types.Log) (*ERC1155Unpaused, error) {
	event := new(ERC1155Unpaused)
	if err := _ERC1155.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
