// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// DidABI is the input ABI used to generate the binding from.
const DidABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fieldKey\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fieldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"PIDAttributeChange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"changeStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"createTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"authentication\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"createPid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identify\",\"type\":\"address\"}],\"name\":\"getLatestBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"identify\",\"type\":\"address\"}],\"name\":\"isIdentityExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"fieldKey\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fieldValue\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"setAttribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DidBin is the compiled bytecode used for deploying new contracts.
var DidBin = "0x608060405234801561001057600080fd5b50610a1a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b146100e557806395bc953814610100578063e4e9f5c014610113578063f2fde38b1461013457600080fd5b8063170abf9c1461008d5780633bcc8a87146100b55780636596dc56146100c8578063715018a6146100db575b600080fd5b6100a061009b3660046106c4565b610147565b60405190151581526020015b60405180910390f35b6100a06100c33660046107ad565b6101b4565b6100a06100d6366004610821565b6102c0565b6100e36103d5565b005b6033546040516001600160a01b0390911681526020016100ac565b6100a061010e3660046108ce565b61043b565b6101266101213660046106c4565b610514565b6040519081526020016100ac565b6100e36101423660046106c4565b6105a7565b6066546000908190815b818110156101ab57846001600160a01b031660668281548110610176576101766108e9565b6000918252602090912001546001600160a01b03160361019957600192506101ab565b806101a3816108ff565b915050610151565b50909392505050565b6066546000908190815b8181101561021857336001600160a01b0316606682815481106101e3576101e36108e9565b6000918252602090912001546001600160a01b0316036102065760019250610218565b80610210816108ff565b9150506101be565b50816102655760405162461bcd60e51b8152602060048201526017602482015276191bd8dd5b595b9d08191bd95cc81b9bdd08195e1a5cdd604a1b60448201526064015b60405180910390fd5b33600081815260656020526040908190205490516000805160206109c58339815191529261029a9290918a918a918a90610973565b60405180910390a150503360009081526065602052604090204390555060019392505050565b60006000805160206109c5833981519152336000876000866040516102e9959493929190610973565b60405180910390a16000805160206109c583398151915233600186600086604051610318959493929190610973565b60405180910390a16000805160206109c583398151915233600285600086604051610347959493929190610973565b60405180910390a15050604080518082018252438152600060208083018281523380845260659092529382209251835592516001928301805460ff191660ff90921691909117905560668054808401825591527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e943540180546001600160a01b0319169092179091559392505050565b6033546001600160a01b0316331461042f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161025c565b6104396000610672565b565b6066546000908190815b8181101561049f57336001600160a01b03166066828154811061046a5761046a6108e9565b6000918252602090912001546001600160a01b03160361048d576001925061049f565b80610497816108ff565b915050610445565b50816104e75760405162461bcd60e51b8152602060048201526017602482015276191bd8dd5b595b9d08191bd95cc81b9bdd08195e1a5cdd604a1b604482015260640161025c565b50503360009081526065602052604090206001908101805460ff191660ff94909416939093179092555090565b6066546000908190815b8181101561057857846001600160a01b031660668281548110610543576105436108e9565b6000918252602090912001546001600160a01b0316036105665760019250610578565b80610570816108ff565b91505061051e565b50811561059d575050506001600160a01b031660009081526065602052604090205490565b5060009392505050565b6033546001600160a01b031633146106015760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161025c565b6001600160a01b0381166106665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161025c565b61066f81610672565b50565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000602082840312156106d657600080fd5b81356001600160a01b03811681146106ed57600080fd5b9392505050565b803560ff8116811461070557600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261073157600080fd5b813567ffffffffffffffff8082111561074c5761074c61070a565b604051601f8301601f19908116603f011681019082821181831017156107745761077461070a565b8160405283815286602085880101111561078d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156107c257600080fd5b6107cb846106f4565b9250602084013567ffffffffffffffff808211156107e857600080fd5b6107f487838801610720565b9350604086013591508082111561080a57600080fd5b5061081786828701610720565b9150509250925092565b6000806000806080858703121561083757600080fd5b843567ffffffffffffffff8082111561084f57600080fd5b61085b88838901610720565b9550602087013591508082111561087157600080fd5b61087d88838901610720565b9450604087013591508082111561089357600080fd5b61089f88838901610720565b935060608701359150808211156108b557600080fd5b506108c287828801610720565b91505092959194509250565b6000602082840312156108e057600080fd5b6106ed826106f4565b634e487b7160e01b600052603260045260246000fd5b60006001820161091f57634e487b7160e01b600052601160045260246000fd5b5060010190565b6000815180845260005b8181101561094c57602081850181015186830182015201610930565b8181111561095e576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b038616815260ff8516602082015260a0604082018190526000906109a090830186610926565b84606084015282810360808401526109b88185610926565b9897505050505050505056fe78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadda2646970667358221220bbfc2a145a4c49d1cd73f4382344d5be5c52cb67dd5e564016e5ce286f5eb71364736f6c634300080d0033"

// DeployDid deploys a new Ethereum contract, binding an instance of Did to it.
func DeployDid(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Did, error) {
	parsed, err := abi.JSON(strings.NewReader(DidABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DidBin), backend)
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

// ChangeStatus is a paid mutator transaction binding the contract method 0x95bc9538.
//
// Solidity: function changeStatus(uint8 status) returns(bool success)
func (_Did *DidTransactor) ChangeStatus(opts *bind.TransactOpts, status uint8) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeStatus", status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x95bc9538.
//
// Solidity: function changeStatus(uint8 status) returns(bool success)
func (_Did *DidSession) ChangeStatus(status uint8) (*types.Transaction, error) {
	return _Did.Contract.ChangeStatus(&_Did.TransactOpts, status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x95bc9538.
//
// Solidity: function changeStatus(uint8 status) returns(bool success)
func (_Did *DidTransactorSession) ChangeStatus(status uint8) (*types.Transaction, error) {
	return _Did.Contract.ChangeStatus(&_Did.TransactOpts, status)
}

// CreatePid is a paid mutator transaction binding the contract method 0x6596dc56.
//
// Solidity: function createPid(string createTime, string authentication, string publicKey, string updateTime) returns(bool success)
func (_Did *DidTransactor) CreatePid(opts *bind.TransactOpts, createTime string, authentication string, publicKey string, updateTime string) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "createPid", createTime, authentication, publicKey, updateTime)
}

// CreatePid is a paid mutator transaction binding the contract method 0x6596dc56.
//
// Solidity: function createPid(string createTime, string authentication, string publicKey, string updateTime) returns(bool success)
func (_Did *DidSession) CreatePid(createTime string, authentication string, publicKey string, updateTime string) (*types.Transaction, error) {
	return _Did.Contract.CreatePid(&_Did.TransactOpts, createTime, authentication, publicKey, updateTime)
}

// CreatePid is a paid mutator transaction binding the contract method 0x6596dc56.
//
// Solidity: function createPid(string createTime, string authentication, string publicKey, string updateTime) returns(bool success)
func (_Did *DidTransactorSession) CreatePid(createTime string, authentication string, publicKey string, updateTime string) (*types.Transaction, error) {
	return _Did.Contract.CreatePid(&_Did.TransactOpts, createTime, authentication, publicKey, updateTime)
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
// Solidity: event PIDAttributeChange(address identity, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Did *DidFilterer) FilterPIDAttributeChange(opts *bind.FilterOpts) (*DidPIDAttributeChangeIterator, error) {

	logs, sub, err := _Did.contract.FilterLogs(opts, "PIDAttributeChange")
	if err != nil {
		return nil, err
	}
	return &DidPIDAttributeChangeIterator{contract: _Did.contract, event: "PIDAttributeChange", logs: logs, sub: sub}, nil
}

// WatchPIDAttributeChange is a free log subscription operation binding the contract event 0x78acd80016b76ce5977296fbdd38a0aac1a3ff9531865bd76469cd4cb590dadd.
//
// Solidity: event PIDAttributeChange(address identity, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Did *DidFilterer) WatchPIDAttributeChange(opts *bind.WatchOpts, sink chan<- *DidPIDAttributeChange) (event.Subscription, error) {

	logs, sub, err := _Did.contract.WatchLogs(opts, "PIDAttributeChange")
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
// Solidity: event PIDAttributeChange(address identity, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Did *DidFilterer) ParsePIDAttributeChange(log types.Log) (*DidPIDAttributeChange, error) {
	event := new(DidPIDAttributeChange)
	if err := _Did.contract.UnpackLog(event, "PIDAttributeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
