// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CDABI is the input ABI used to generate the binding from.
const CDABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"setChainID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CDFuncSigs maps the 4-byte function signature to its string representation.
var CDFuncSigs = map[string]string{
	"bff0af6a": "getchainID()",
	"ed8d47e6": "setChainID(uint256)",
}

// CDBin is the compiled bytecode used for deploying new contracts.
var CDBin = "0x608060405234801561001057600080fd5b5060ab8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c8063bff0af6a146037578063ed8d47e614604f575b600080fd5b603d606b565b60408051918252519081900360200190f35b606960048036036020811015606357600080fd5b50356071565b005b60005490565b60005556fea265627a7a72315820d4feb22fc97c9aafa15c3cf1bdff93fe63918bab31d8e54ce5fbf203d9cae48464736f6c63430005100032"

// DeployCD deploys a new Ethereum contract, binding an instance of CD to it.
func DeployCD(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CD, error) {
	parsed, err := abi.JSON(strings.NewReader(CDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CD{CDCaller: CDCaller{contract: contract}, CDTransactor: CDTransactor{contract: contract}, CDFilterer: CDFilterer{contract: contract}}, nil
}

// CD is an auto generated Go binding around an Ethereum contract.
type CD struct {
	CDCaller     // Read-only binding to the contract
	CDTransactor // Write-only binding to the contract
	CDFilterer   // Log filterer for contract events
}

// CDCaller is an auto generated read-only Go binding around an Ethereum contract.
type CDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CDSession struct {
	Contract     *CD               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CDCallerSession struct {
	Contract *CDCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CDTransactorSession struct {
	Contract     *CDTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CDRaw is an auto generated low-level Go binding around an Ethereum contract.
type CDRaw struct {
	Contract *CD // Generic contract binding to access the raw methods on
}

// CDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CDCallerRaw struct {
	Contract *CDCaller // Generic read-only contract binding to access the raw methods on
}

// CDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CDTransactorRaw struct {
	Contract *CDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCD creates a new instance of CD, bound to a specific deployed contract.
func NewCD(address common.Address, backend bind.ContractBackend) (*CD, error) {
	contract, err := bindCD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CD{CDCaller: CDCaller{contract: contract}, CDTransactor: CDTransactor{contract: contract}, CDFilterer: CDFilterer{contract: contract}}, nil
}

// NewCDCaller creates a new read-only instance of CD, bound to a specific deployed contract.
func NewCDCaller(address common.Address, caller bind.ContractCaller) (*CDCaller, error) {
	contract, err := bindCD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CDCaller{contract: contract}, nil
}

// NewCDTransactor creates a new write-only instance of CD, bound to a specific deployed contract.
func NewCDTransactor(address common.Address, transactor bind.ContractTransactor) (*CDTransactor, error) {
	contract, err := bindCD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CDTransactor{contract: contract}, nil
}

// NewCDFilterer creates a new log filterer instance of CD, bound to a specific deployed contract.
func NewCDFilterer(address common.Address, filterer bind.ContractFilterer) (*CDFilterer, error) {
	contract, err := bindCD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CDFilterer{contract: contract}, nil
}

// bindCD binds a generic wrapper to an already deployed contract.
func bindCD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CD *CDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CD.Contract.CDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CD *CDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CD.Contract.CDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CD *CDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CD.Contract.CDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CD *CDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CD *CDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CD *CDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CD.Contract.contract.Transact(opts, method, params...)
}

// GetchainID is a free data retrieval call binding the contract method 0xbff0af6a.
//
// Solidity: function getchainID() view returns(uint256)
func (_CD *CDCaller) GetchainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CD.contract.Call(opts, &out, "getchainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetchainID is a free data retrieval call binding the contract method 0xbff0af6a.
//
// Solidity: function getchainID() view returns(uint256)
func (_CD *CDSession) GetchainID() (*big.Int, error) {
	return _CD.Contract.GetchainID(&_CD.CallOpts)
}

// GetchainID is a free data retrieval call binding the contract method 0xbff0af6a.
//
// Solidity: function getchainID() view returns(uint256)
func (_CD *CDCallerSession) GetchainID() (*big.Int, error) {
	return _CD.Contract.GetchainID(&_CD.CallOpts)
}

// SetChainID is a paid mutator transaction binding the contract method 0xed8d47e6.
//
// Solidity: function setChainID(uint256 x) returns()
func (_CD *CDTransactor) SetChainID(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _CD.contract.Transact(opts, "setChainID", x)
}

// SetChainID is a paid mutator transaction binding the contract method 0xed8d47e6.
//
// Solidity: function setChainID(uint256 x) returns()
func (_CD *CDSession) SetChainID(x *big.Int) (*types.Transaction, error) {
	return _CD.Contract.SetChainID(&_CD.TransactOpts, x)
}

// SetChainID is a paid mutator transaction binding the contract method 0xed8d47e6.
//
// Solidity: function setChainID(uint256 x) returns()
func (_CD *CDTransactorSession) SetChainID(x *big.Int) (*types.Transaction, error) {
	return _CD.Contract.SetChainID(&_CD.TransactOpts, x)
}
