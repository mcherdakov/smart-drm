// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartdrm

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

// SmartdrmMetaData contains all meta data concerning the Smartdrm contract.
var SmartdrmMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"SmartDRMOutOfBounds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SmartdrmABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartdrmMetaData.ABI instead.
var SmartdrmABI = SmartdrmMetaData.ABI

// Smartdrm is an auto generated Go binding around an Ethereum contract.
type Smartdrm struct {
	SmartdrmCaller     // Read-only binding to the contract
	SmartdrmTransactor // Write-only binding to the contract
	SmartdrmFilterer   // Log filterer for contract events
}

// SmartdrmCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartdrmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartdrmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartdrmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartdrmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartdrmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartdrmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartdrmSession struct {
	Contract     *Smartdrm         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartdrmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartdrmCallerSession struct {
	Contract *SmartdrmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SmartdrmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartdrmTransactorSession struct {
	Contract     *SmartdrmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SmartdrmRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartdrmRaw struct {
	Contract *Smartdrm // Generic contract binding to access the raw methods on
}

// SmartdrmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartdrmCallerRaw struct {
	Contract *SmartdrmCaller // Generic read-only contract binding to access the raw methods on
}

// SmartdrmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartdrmTransactorRaw struct {
	Contract *SmartdrmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartdrm creates a new instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrm(address common.Address, backend bind.ContractBackend) (*Smartdrm, error) {
	contract, err := bindSmartdrm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartdrm{SmartdrmCaller: SmartdrmCaller{contract: contract}, SmartdrmTransactor: SmartdrmTransactor{contract: contract}, SmartdrmFilterer: SmartdrmFilterer{contract: contract}}, nil
}

// NewSmartdrmCaller creates a new read-only instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrmCaller(address common.Address, caller bind.ContractCaller) (*SmartdrmCaller, error) {
	contract, err := bindSmartdrm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartdrmCaller{contract: contract}, nil
}

// NewSmartdrmTransactor creates a new write-only instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrmTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartdrmTransactor, error) {
	contract, err := bindSmartdrm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartdrmTransactor{contract: contract}, nil
}

// NewSmartdrmFilterer creates a new log filterer instance of Smartdrm, bound to a specific deployed contract.
func NewSmartdrmFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartdrmFilterer, error) {
	contract, err := bindSmartdrm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartdrmFilterer{contract: contract}, nil
}

// bindSmartdrm binds a generic wrapper to an already deployed contract.
func bindSmartdrm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmartdrmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartdrm *SmartdrmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartdrm.Contract.SmartdrmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartdrm *SmartdrmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartdrm.Contract.SmartdrmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartdrm *SmartdrmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartdrm.Contract.SmartdrmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartdrm *SmartdrmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartdrm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartdrm *SmartdrmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartdrm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartdrm *SmartdrmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartdrm.Contract.contract.Transact(opts, method, params...)
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 index) view returns(address)
func (_Smartdrm *SmartdrmCaller) GetCreator(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Smartdrm.contract.Call(opts, &out, "getCreator", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 index) view returns(address)
func (_Smartdrm *SmartdrmSession) GetCreator(index *big.Int) (common.Address, error) {
	return _Smartdrm.Contract.GetCreator(&_Smartdrm.CallOpts, index)
}

// GetCreator is a free data retrieval call binding the contract method 0xd48e638a.
//
// Solidity: function getCreator(uint256 index) view returns(address)
func (_Smartdrm *SmartdrmCallerSession) GetCreator(index *big.Int) (common.Address, error) {
	return _Smartdrm.Contract.GetCreator(&_Smartdrm.CallOpts, index)
}

// RegisterCreator is a paid mutator transaction binding the contract method 0x0c6c8fe8.
//
// Solidity: function registerCreator() returns()
func (_Smartdrm *SmartdrmTransactor) RegisterCreator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartdrm.contract.Transact(opts, "registerCreator")
}

// RegisterCreator is a paid mutator transaction binding the contract method 0x0c6c8fe8.
//
// Solidity: function registerCreator() returns()
func (_Smartdrm *SmartdrmSession) RegisterCreator() (*types.Transaction, error) {
	return _Smartdrm.Contract.RegisterCreator(&_Smartdrm.TransactOpts)
}

// RegisterCreator is a paid mutator transaction binding the contract method 0x0c6c8fe8.
//
// Solidity: function registerCreator() returns()
func (_Smartdrm *SmartdrmTransactorSession) RegisterCreator() (*types.Transaction, error) {
	return _Smartdrm.Contract.RegisterCreator(&_Smartdrm.TransactOpts)
}
