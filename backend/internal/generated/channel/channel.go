// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package channel

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

// ChannelMetaData contains all meta data concerning the Channel contract.
var ChannelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"channelTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"closeChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChannelRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChannelSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChannelTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStartDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use ChannelMetaData.ABI instead.
var ChannelABI = ChannelMetaData.ABI

// Channel is an auto generated Go binding around an Ethereum contract.
type Channel struct {
	ChannelCaller     // Read-only binding to the contract
	ChannelTransactor // Write-only binding to the contract
	ChannelFilterer   // Log filterer for contract events
}

// ChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelSession struct {
	Contract     *Channel          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelCallerSession struct {
	Contract *ChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelTransactorSession struct {
	Contract     *ChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelRaw struct {
	Contract *Channel // Generic contract binding to access the raw methods on
}

// ChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelCallerRaw struct {
	Contract *ChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelTransactorRaw struct {
	Contract *ChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannel creates a new instance of Channel, bound to a specific deployed contract.
func NewChannel(address common.Address, backend bind.ContractBackend) (*Channel, error) {
	contract, err := bindChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// NewChannelCaller creates a new read-only instance of Channel, bound to a specific deployed contract.
func NewChannelCaller(address common.Address, caller bind.ContractCaller) (*ChannelCaller, error) {
	contract, err := bindChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelCaller{contract: contract}, nil
}

// NewChannelTransactor creates a new write-only instance of Channel, bound to a specific deployed contract.
func NewChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelTransactor, error) {
	contract, err := bindChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelTransactor{contract: contract}, nil
}

// NewChannelFilterer creates a new log filterer instance of Channel, bound to a specific deployed contract.
func NewChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelFilterer, error) {
	contract, err := bindChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelFilterer{contract: contract}, nil
}

// bindChannel binds a generic wrapper to an already deployed contract.
func bindChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChannelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.ChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transact(opts, method, params...)
}

// GetChannelRecipient is a free data retrieval call binding the contract method 0x7b909318.
//
// Solidity: function getChannelRecipient() view returns(address)
func (_Channel *ChannelCaller) GetChannelRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "getChannelRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChannelRecipient is a free data retrieval call binding the contract method 0x7b909318.
//
// Solidity: function getChannelRecipient() view returns(address)
func (_Channel *ChannelSession) GetChannelRecipient() (common.Address, error) {
	return _Channel.Contract.GetChannelRecipient(&_Channel.CallOpts)
}

// GetChannelRecipient is a free data retrieval call binding the contract method 0x7b909318.
//
// Solidity: function getChannelRecipient() view returns(address)
func (_Channel *ChannelCallerSession) GetChannelRecipient() (common.Address, error) {
	return _Channel.Contract.GetChannelRecipient(&_Channel.CallOpts)
}

// GetChannelSender is a free data retrieval call binding the contract method 0x08ca3871.
//
// Solidity: function getChannelSender() view returns(address)
func (_Channel *ChannelCaller) GetChannelSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "getChannelSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChannelSender is a free data retrieval call binding the contract method 0x08ca3871.
//
// Solidity: function getChannelSender() view returns(address)
func (_Channel *ChannelSession) GetChannelSender() (common.Address, error) {
	return _Channel.Contract.GetChannelSender(&_Channel.CallOpts)
}

// GetChannelSender is a free data retrieval call binding the contract method 0x08ca3871.
//
// Solidity: function getChannelSender() view returns(address)
func (_Channel *ChannelCallerSession) GetChannelSender() (common.Address, error) {
	return _Channel.Contract.GetChannelSender(&_Channel.CallOpts)
}

// GetChannelTimeout is a free data retrieval call binding the contract method 0x9d9ecd78.
//
// Solidity: function getChannelTimeout() view returns(uint256)
func (_Channel *ChannelCaller) GetChannelTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "getChannelTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChannelTimeout is a free data retrieval call binding the contract method 0x9d9ecd78.
//
// Solidity: function getChannelTimeout() view returns(uint256)
func (_Channel *ChannelSession) GetChannelTimeout() (*big.Int, error) {
	return _Channel.Contract.GetChannelTimeout(&_Channel.CallOpts)
}

// GetChannelTimeout is a free data retrieval call binding the contract method 0x9d9ecd78.
//
// Solidity: function getChannelTimeout() view returns(uint256)
func (_Channel *ChannelCallerSession) GetChannelTimeout() (*big.Int, error) {
	return _Channel.Contract.GetChannelTimeout(&_Channel.CallOpts)
}

// GetStartDate is a free data retrieval call binding the contract method 0x78f305c6.
//
// Solidity: function getStartDate() view returns(uint256)
func (_Channel *ChannelCaller) GetStartDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "getStartDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartDate is a free data retrieval call binding the contract method 0x78f305c6.
//
// Solidity: function getStartDate() view returns(uint256)
func (_Channel *ChannelSession) GetStartDate() (*big.Int, error) {
	return _Channel.Contract.GetStartDate(&_Channel.CallOpts)
}

// GetStartDate is a free data retrieval call binding the contract method 0x78f305c6.
//
// Solidity: function getStartDate() view returns(uint256)
func (_Channel *ChannelCallerSession) GetStartDate() (*big.Int, error) {
	return _Channel.Contract.GetStartDate(&_Channel.CallOpts)
}

// ChannelTimeout is a paid mutator transaction binding the contract method 0x2ef2d55e.
//
// Solidity: function channelTimeout() returns()
func (_Channel *ChannelTransactor) ChannelTimeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "channelTimeout")
}

// ChannelTimeout is a paid mutator transaction binding the contract method 0x2ef2d55e.
//
// Solidity: function channelTimeout() returns()
func (_Channel *ChannelSession) ChannelTimeout() (*types.Transaction, error) {
	return _Channel.Contract.ChannelTimeout(&_Channel.TransactOpts)
}

// ChannelTimeout is a paid mutator transaction binding the contract method 0x2ef2d55e.
//
// Solidity: function channelTimeout() returns()
func (_Channel *ChannelTransactorSession) ChannelTimeout() (*types.Transaction, error) {
	return _Channel.Contract.ChannelTimeout(&_Channel.TransactOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xa5e84c26.
//
// Solidity: function closeChannel(bytes32 h, uint8 v, bytes32 r, bytes32 s, uint256 value) returns()
func (_Channel *ChannelTransactor) CloseChannel(opts *bind.TransactOpts, h [32]byte, v uint8, r [32]byte, s [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "closeChannel", h, v, r, s, value)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xa5e84c26.
//
// Solidity: function closeChannel(bytes32 h, uint8 v, bytes32 r, bytes32 s, uint256 value) returns()
func (_Channel *ChannelSession) CloseChannel(h [32]byte, v uint8, r [32]byte, s [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Channel.Contract.CloseChannel(&_Channel.TransactOpts, h, v, r, s, value)
}

// CloseChannel is a paid mutator transaction binding the contract method 0xa5e84c26.
//
// Solidity: function closeChannel(bytes32 h, uint8 v, bytes32 r, bytes32 s, uint256 value) returns()
func (_Channel *ChannelTransactorSession) CloseChannel(h [32]byte, v uint8, r [32]byte, s [32]byte, value *big.Int) (*types.Transaction, error) {
	return _Channel.Contract.CloseChannel(&_Channel.TransactOpts, h, v, r, s, value)
}
