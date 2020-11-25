// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balance

import (
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BalanceCheckerABI is the input ABI used to generate the binding from.
const BalanceCheckerABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BalanceChecker is an auto generated Go binding around an Ethereum contract.
type BalanceChecker struct {
	BalanceCheckerCaller     // Read-only binding to the contract
	BalanceCheckerTransactor // Write-only binding to the contract
	BalanceCheckerFilterer   // Log filterer for contract events
}

// BalanceCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceCheckerSession struct {
	Contract     *BalanceChecker   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceCheckerCallerSession struct {
	Contract *BalanceCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalanceCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceCheckerTransactorSession struct {
	Contract     *BalanceCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalanceCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceCheckerRaw struct {
	Contract *BalanceChecker // Generic contract binding to access the raw methods on
}

// BalanceCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceCheckerCallerRaw struct {
	Contract *BalanceCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceCheckerTransactorRaw struct {
	Contract *BalanceCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceChecker creates a new instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceChecker(address common.Address, backend bind.ContractBackend) (*BalanceChecker, error) {
	contract, err := bindBalanceChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceChecker{BalanceCheckerCaller: BalanceCheckerCaller{contract: contract}, BalanceCheckerTransactor: BalanceCheckerTransactor{contract: contract}, BalanceCheckerFilterer: BalanceCheckerFilterer{contract: contract}}, nil
}

// NewBalanceCheckerCaller creates a new read-only instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceCheckerCaller(address common.Address, caller bind.ContractCaller) (*BalanceCheckerCaller, error) {
	contract, err := bindBalanceChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCheckerCaller{contract: contract}, nil
}

// NewBalanceCheckerTransactor creates a new write-only instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceCheckerTransactor, error) {
	contract, err := bindBalanceChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCheckerTransactor{contract: contract}, nil
}

// NewBalanceCheckerFilterer creates a new log filterer instance of BalanceChecker, bound to a specific deployed contract.
func NewBalanceCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceCheckerFilterer, error) {
	contract, err := bindBalanceChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceCheckerFilterer{contract: contract}, nil
}

// bindBalanceChecker binds a generic wrapper to an already deployed contract.
func bindBalanceChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceCheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceChecker *BalanceCheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceChecker.Contract.BalanceCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceChecker *BalanceCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceChecker.Contract.BalanceCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceChecker *BalanceCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceChecker.Contract.BalanceCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceChecker *BalanceCheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceChecker *BalanceCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceChecker *BalanceCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceChecker.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0xf0002ea9.
//
// Solidity: function balances(address[] users, address[] tokens) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerCaller) Balances(opts *bind.CallOpts, users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BalanceChecker.contract.Call(opts, &out, "balances", users, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0xf0002ea9.
//
// Solidity: function balances(address[] users, address[] tokens) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerSession) Balances(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _BalanceChecker.Contract.Balances(&_BalanceChecker.CallOpts, users, tokens)
}

// Balances is a free data retrieval call binding the contract method 0xf0002ea9.
//
// Solidity: function balances(address[] users, address[] tokens) view returns(uint256[])
func (_BalanceChecker *BalanceCheckerCallerSession) Balances(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _BalanceChecker.Contract.Balances(&_BalanceChecker.CallOpts, users, tokens)
}
