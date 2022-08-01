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

// DidMetaData contains all meta data concerning the Did contract.
var DidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fieldKey\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fieldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"PIDAttributeChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"status\",\"type\":\"int8\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"createTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"createPid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identify\",\"type\":\"address\"}],\"name\":\"getLatestBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identify\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identify\",\"type\":\"address\"}],\"name\":\"isIdentityExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"fieldKey\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fieldValue\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"setAttribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a95806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715018a611610066578063715018a61461010e5780638da5cb5b1461011657806396d6b63414610131578063e4e9f5c014610144578063f2fde38b1461016557600080fd5b806301dbe85814610098578063170abf9c146100ad57806330ccebb5146100d55780633bcc8a87146100fb575b600080fd5b6100ab6100a63660046107d5565b610178565b005b6100c06100bb3660046107ff565b61024f565b60405190151581526020015b60405180910390f35b6100e86100e33660046107ff565b6102bc565b60405160009190910b81526020016100cc565b6100c06101093660046108cb565b61033e565b6100ab610455565b6033546040516001600160a01b0390911681526020016100cc565b6100c061013f366004610947565b610469565b6101576101523660046107ff565b610635565b6040519081526020016100cc565b6100ab6101733660046107ff565b6106b0565b606654600090815b818110156101da57336001600160a01b0316606682815481106101a5576101a5610996565b6000918252602090912001546001600160a01b0316036101c857600192506101da565b806101d2816109ac565b915050610180565b50816102275760405162461bcd60e51b8152602060048201526017602482015276191bd8dd5b595b9d08191bd95cc81b9bdd08195e1a5cdd604a1b60448201526064015b60405180910390fd5b5050336000908152606560205260409020600101805460ff191660ff92909216919091179055565b6066546000908190815b818110156102b357846001600160a01b03166066828154811061027e5761027e610996565b6000918252602090912001546001600160a01b0316036102a157600192506102b3565b806102ab816109ac565b915050610259565b50909392505050565b60665460009060001990825b818110156102b357846001600160a01b0316606682815481106102ed576102ed610996565b6000918252602090912001546001600160a01b03160361032c576001600160a01b038516600090815260656020526040812060010154900b92506102b3565b80610336816109ac565b9150506102c8565b6066546000908190815b818110156103a257336001600160a01b03166066828154811061036d5761036d610996565b6000918252602090912001546001600160a01b03160361039057600192506103a2565b8061039a816109ac565b915050610348565b50816103ea5760405162461bcd60e51b8152602060048201526017602482015276191bd8dd5b595b9d08191bd95cc81b9bdd08195e1a5cdd604a1b604482015260640161021e565b33600081815260656020526040908190205490517f78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd9161042f918a918a918a90610a20565b60405180910390a250503360009081526065602052604090204390555060019392505050565b61045d610729565b6104676000610783565b565b6066546000908190815b818110156104cd57336001600160a01b03166066828154811061049857610498610996565b6000918252602090912001546001600160a01b0316036104bb57600192506104cd565b806104c5816109ac565b915050610473565b50811561051c5760405162461bcd60e51b815260206004820152601760248201527f646f63756d656e7420616c726561647920657869737473000000000000000000604482015260640161021e565b336001600160a01b03167f78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd60008860008860405161055d9493929190610a20565b60405180910390a2336001600160a01b03167f78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd6001876000886040516105a69493929190610a20565b60405180910390a25050604080518082018252438152600060208083018281523380845260659092529382209251835592516001928301805460ff191660ff90921691909117905560668054808401825591527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b031916909217909155949350505050565b6066546000908190815b818110156102b357846001600160a01b03166066828154811061066457610664610996565b6000918252602090912001546001600160a01b03160361069e576001600160a01b03851660009081526065602052604090205492506102b3565b806106a8816109ac565b91505061063f565b6106b8610729565b6001600160a01b03811661071d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161021e565b61072681610783565b50565b6033546001600160a01b031633146104675760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161021e565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000602082840312156107e757600080fd5b81358060000b81146107f857600080fd5b9392505050565b60006020828403121561081157600080fd5b81356001600160a01b03811681146107f857600080fd5b634e487b7160e01b600052604160045260246000fd5b600082601f83011261084f57600080fd5b813567ffffffffffffffff8082111561086a5761086a610828565b604051601f8301601f19908116603f0116810190828211818310171561089257610892610828565b816040528381528660208588010111156108ab57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156108e057600080fd5b833560ff811681146108f157600080fd5b9250602084013567ffffffffffffffff8082111561090e57600080fd5b61091a8783880161083e565b9350604086013591508082111561093057600080fd5b5061093d8682870161083e565b9150509250925092565b60008060006060848603121561095c57600080fd5b833567ffffffffffffffff8082111561097457600080fd5b6109808783880161083e565b9450602086013591508082111561090e57600080fd5b634e487b7160e01b600052603260045260246000fd5b6000600182016109cc57634e487b7160e01b600052601160045260246000fd5b5060010190565b6000815180845260005b818110156109f9576020818501810151868301820152016109dd565b81811115610a0b576000602083870101525b50601f01601f19169290920160200192915050565b60ff85168152608060208201526000610a3c60808301866109d3565b8460408401528281036060840152610a5481856109d3565b97965050505050505056fea2646970667358221220d671802dabff462e3654d9f2bdf1ee96bba5008a801af38cce2e27575fac56d264736f6c634300080d0033",
}

// DidABI is the input ABI used to generate the binding from.
// Deprecated: Use DidMetaData.ABI instead.
var DidABI = DidMetaData.ABI

// DidBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DidMetaData.Bin instead.
var DidBin = DidMetaData.Bin

// DeployDid deploys a new Ethereum contract, binding an instance of Did to it.
func DeployDid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Did, error) {
	parsed, err := DidMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DidBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Did{DidCaller: DidCaller{contract: contract}, DidTransactor: DidTransactor{contract: contract}, DidFilterer: DidFilterer{contract: contract}}, nil
}

// Did is an auto generated Go binding around an Ethereum contract.
type Did struct {
	DidCaller     // Read-only binding to the contract
	DidTransactor // Write-only binding to the contract
	DidFilterer   // Log filterer for contract events
}

// DidCaller is an auto generated read-only Go binding around an Ethereum contract.
type DidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DidSession struct {
	Contract     *Did              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DidCallerSession struct {
	Contract *DidCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DidTransactorSession struct {
	Contract     *DidTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidRaw is an auto generated low-level Go binding around an Ethereum contract.
type DidRaw struct {
	Contract *Did // Generic contract binding to access the raw methods on
}

// DidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DidCallerRaw struct {
	Contract *DidCaller // Generic read-only contract binding to access the raw methods on
}

// DidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DidTransactorRaw struct {
	Contract *DidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDid creates a new instance of Did, bound to a specific deployed contract.
func NewDid(address common.Address, backend bind.ContractBackend) (*Did, error) {
	contract, err := bindDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Did{DidCaller: DidCaller{contract: contract}, DidTransactor: DidTransactor{contract: contract}, DidFilterer: DidFilterer{contract: contract}}, nil
}

// NewDidCaller creates a new read-only instance of Did, bound to a specific deployed contract.
func NewDidCaller(address common.Address, caller bind.ContractCaller) (*DidCaller, error) {
	contract, err := bindDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DidCaller{contract: contract}, nil
}

// NewDidTransactor creates a new write-only instance of Did, bound to a specific deployed contract.
func NewDidTransactor(address common.Address, transactor bind.ContractTransactor) (*DidTransactor, error) {
	contract, err := bindDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DidTransactor{contract: contract}, nil
}

// NewDidFilterer creates a new log filterer instance of Did, bound to a specific deployed contract.
func NewDidFilterer(address common.Address, filterer bind.ContractFilterer) (*DidFilterer, error) {
	contract, err := bindDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DidFilterer{contract: contract}, nil
}

// bindDid binds a generic wrapper to an already deployed contract.
func bindDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Did.Contract.DidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Did.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.contract.Transact(opts, method, params...)
}

// GetLatestBlock is a free data retrieval call binding the contract method 0xe4e9f5c0.
//
// Solidity: function getLatestBlock(address identify) view returns(uint256)
func (_Did *DidCaller) GetLatestBlock(opts *bind.CallOpts, identify common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Did.contract.Call(opts, &out, "getLatestBlock", identify)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestBlock is a free data retrieval call binding the contract method 0xe4e9f5c0.
//
// Solidity: function getLatestBlock(address identify) view returns(uint256)
func (_Did *DidSession) GetLatestBlock(identify common.Address) (*big.Int, error) {
	return _Did.Contract.GetLatestBlock(&_Did.CallOpts, identify)
}

// GetLatestBlock is a free data retrieval call binding the contract method 0xe4e9f5c0.
//
// Solidity: function getLatestBlock(address identify) view returns(uint256)
func (_Did *DidCallerSession) GetLatestBlock(identify common.Address) (*big.Int, error) {
	return _Did.Contract.GetLatestBlock(&_Did.CallOpts, identify)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address identify) view returns(int8)
func (_Did *DidCaller) GetStatus(opts *bind.CallOpts, identify common.Address) (int8, error) {
	var out []interface{}
	err := _Did.contract.Call(opts, &out, "getStatus", identify)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address identify) view returns(int8)
func (_Did *DidSession) GetStatus(identify common.Address) (int8, error) {
	return _Did.Contract.GetStatus(&_Did.CallOpts, identify)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address identify) view returns(int8)
func (_Did *DidCallerSession) GetStatus(identify common.Address) (int8, error) {
	return _Did.Contract.GetStatus(&_Did.CallOpts, identify)
}

// IsIdentityExist is a free data retrieval call binding the contract method 0x170abf9c.
//
// Solidity: function isIdentityExist(address identify) view returns(bool success)
func (_Did *DidCaller) IsIdentityExist(opts *bind.CallOpts, identify common.Address) (bool, error) {
	var out []interface{}
	err := _Did.contract.Call(opts, &out, "isIdentityExist", identify)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityExist is a free data retrieval call binding the contract method 0x170abf9c.
//
// Solidity: function isIdentityExist(address identify) view returns(bool success)
func (_Did *DidSession) IsIdentityExist(identify common.Address) (bool, error) {
	return _Did.Contract.IsIdentityExist(&_Did.CallOpts, identify)
}

// IsIdentityExist is a free data retrieval call binding the contract method 0x170abf9c.
//
// Solidity: function isIdentityExist(address identify) view returns(bool success)
func (_Did *DidCallerSession) IsIdentityExist(identify common.Address) (bool, error) {
	return _Did.Contract.IsIdentityExist(&_Did.CallOpts, identify)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Did *DidCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Did.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Did *DidSession) Owner() (common.Address, error) {
	return _Did.Contract.Owner(&_Did.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Did *DidCallerSession) Owner() (common.Address, error) {
	return _Did.Contract.Owner(&_Did.CallOpts)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x01dbe858.
//
// Solidity: function changeStatus(int8 status) returns()
func (_Did *DidTransactor) ChangeStatus(opts *bind.TransactOpts, status int8) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeStatus", status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x01dbe858.
//
// Solidity: function changeStatus(int8 status) returns()
func (_Did *DidSession) ChangeStatus(status int8) (*types.Transaction, error) {
	return _Did.Contract.ChangeStatus(&_Did.TransactOpts, status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x01dbe858.
//
// Solidity: function changeStatus(int8 status) returns()
func (_Did *DidTransactorSession) ChangeStatus(status int8) (*types.Transaction, error) {
	return _Did.Contract.ChangeStatus(&_Did.TransactOpts, status)
}

// CreatePid is a paid mutator transaction binding the contract method 0x96d6b634.
//
// Solidity: function createPid(string createTime, string publicKey, string updateTime) returns(bool success)
func (_Did *DidTransactor) CreatePid(opts *bind.TransactOpts, createTime string, publicKey string, updateTime string) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "createPid", createTime, publicKey, updateTime)
}

// CreatePid is a paid mutator transaction binding the contract method 0x96d6b634.
//
// Solidity: function createPid(string createTime, string publicKey, string updateTime) returns(bool success)
func (_Did *DidSession) CreatePid(createTime string, publicKey string, updateTime string) (*types.Transaction, error) {
	return _Did.Contract.CreatePid(&_Did.TransactOpts, createTime, publicKey, updateTime)
}

// CreatePid is a paid mutator transaction binding the contract method 0x96d6b634.
//
// Solidity: function createPid(string createTime, string publicKey, string updateTime) returns(bool success)
func (_Did *DidTransactorSession) CreatePid(createTime string, publicKey string, updateTime string) (*types.Transaction, error) {
	return _Did.Contract.CreatePid(&_Did.TransactOpts, createTime, publicKey, updateTime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Did *DidTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Did *DidSession) RenounceOwnership() (*types.Transaction, error) {
	return _Did.Contract.RenounceOwnership(&_Did.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Did *DidTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Did.Contract.RenounceOwnership(&_Did.TransactOpts)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x3bcc8a87.
//
// Solidity: function setAttribute(uint8 fieldKey, string fieldValue, string updateTime) returns(bool success)
func (_Did *DidTransactor) SetAttribute(opts *bind.TransactOpts, fieldKey uint8, fieldValue string, updateTime string) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttribute", fieldKey, fieldValue, updateTime)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x3bcc8a87.
//
// Solidity: function setAttribute(uint8 fieldKey, string fieldValue, string updateTime) returns(bool success)
func (_Did *DidSession) SetAttribute(fieldKey uint8, fieldValue string, updateTime string) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, fieldKey, fieldValue, updateTime)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x3bcc8a87.
//
// Solidity: function setAttribute(uint8 fieldKey, string fieldValue, string updateTime) returns(bool success)
func (_Did *DidTransactorSession) SetAttribute(fieldKey uint8, fieldValue string, updateTime string) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, fieldKey, fieldValue, updateTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Did *DidTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Did *DidSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.TransferOwnership(&_Did.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Did *DidTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.TransferOwnership(&_Did.TransactOpts, newOwner)
}

// DidInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Did contract.
type DidInitializedIterator struct {
	Event *DidInitialized // Event containing the contract specifics and raw log

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
func (it *DidInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidInitialized)
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
		it.Event = new(DidInitialized)
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
func (it *DidInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidInitialized represents a Initialized event raised by the Did contract.
type DidInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Did *DidFilterer) FilterInitialized(opts *bind.FilterOpts) (*DidInitializedIterator, error) {

	logs, sub, err := _Did.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DidInitializedIterator{contract: _Did.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Did *DidFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DidInitialized) (event.Subscription, error) {

	logs, sub, err := _Did.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidInitialized)
				if err := _Did.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Did *DidFilterer) ParseInitialized(log types.Log) (*DidInitialized, error) {
	event := new(DidInitialized)
	if err := _Did.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DidOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Did contract.
type DidOwnershipTransferredIterator struct {
	Event *DidOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DidOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidOwnershipTransferred)
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
		it.Event = new(DidOwnershipTransferred)
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
func (it *DidOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidOwnershipTransferred represents a OwnershipTransferred event raised by the Did contract.
type DidOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Did *DidFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DidOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DidOwnershipTransferredIterator{contract: _Did.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Did *DidFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DidOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidOwnershipTransferred)
				if err := _Did.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Did *DidFilterer) ParseOwnershipTransferred(log types.Log) (*DidOwnershipTransferred, error) {
	event := new(DidOwnershipTransferred)
	if err := _Did.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DidPIDAttributeChangeIterator is returned from FilterPIDAttributeChange and is used to iterate over the raw logs and unpacked data for PIDAttributeChange events raised by the Did contract.
type DidPIDAttributeChangeIterator struct {
	Event *DidPIDAttributeChange // Event containing the contract specifics and raw log

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
func (it *DidPIDAttributeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidPIDAttributeChange)
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
		it.Event = new(DidPIDAttributeChange)
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
func (it *DidPIDAttributeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidPIDAttributeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidPIDAttributeChange represents a PIDAttributeChange event raised by the Did contract.
type DidPIDAttributeChange struct {
	Identity    common.Address
	FieldKey    uint8
	FieldValue  string
	BlockNumber *big.Int
	UpdateTime  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPIDAttributeChange is a free log retrieval operation binding the contract event 0x78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd.
//
// Solidity: event PIDAttributeChange(address indexed identity, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Did *DidFilterer) FilterPIDAttributeChange(opts *bind.FilterOpts, identity []common.Address) (*DidPIDAttributeChangeIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "PIDAttributeChange", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidPIDAttributeChangeIterator{contract: _Did.contract, event: "PIDAttributeChange", logs: logs, sub: sub}, nil
}

// WatchPIDAttributeChange is a free log subscription operation binding the contract event 0x78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd.
//
// Solidity: event PIDAttributeChange(address indexed identity, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Did *DidFilterer) WatchPIDAttributeChange(opts *bind.WatchOpts, sink chan<- *DidPIDAttributeChange, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "PIDAttributeChange", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidPIDAttributeChange)
				if err := _Did.contract.UnpackLog(event, "PIDAttributeChange", log); err != nil {
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

// ParsePIDAttributeChange is a log parse operation binding the contract event 0x78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd.
//
// Solidity: event PIDAttributeChange(address indexed identity, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Did *DidFilterer) ParsePIDAttributeChange(log types.Log) (*DidPIDAttributeChange, error) {
	event := new(DidPIDAttributeChange)
	if err := _Did.contract.UnpackLog(event, "PIDAttributeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
