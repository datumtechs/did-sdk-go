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

// CredentialMetaData contains all meta data concerning the Credential contract.
var CredentialMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fieldKey\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fieldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"CredentialAttributeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"signer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatureData\",\"type\":\"string\"}],\"name\":\"createCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"}],\"name\":\"getLatestBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"}],\"name\":\"isHashExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610b57806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063abf66aed1161005b578063abf66aed146100cf578063bfae4302146100e2578063c4d66de814610103578063f2fde38b1461011657600080fd5b8063715018a6146100825780638a2d099c1461008c5780638da5cb5b146100b4575b600080fd5b61008a610129565b005b61009f61009a366004610807565b610194565b60405190151581526020015b60405180910390f35b6033546040516001600160a01b0390911681526020016100ab565b61009f6100dd3660046108d7565b6101ee565b6100f56100f0366004610807565b6103b8565b6040519081526020016100ab565b61008a610111366004610959565b61042f565b61008a610124366004610959565b610504565b6033546001600160a01b031633146101885760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b61019260006105cf565b565b6066546000908190815b818110156101e55784606682815481106101ba576101ba61097d565b9060005260206000200154036101d357600192506101e5565b806101dd81610993565b91505061019e565b50909392505050565b6067546000908190610208906001600160a01b0316610621565b606754909150600090610223906001600160a01b0316610719565b8251909150600090815b8181101561028257336001600160a01b03168582815181106102515761025161097d565b60200260200101516001600160a01b0316036102705760019250610282565b8061027a81610993565b91505061022d565b5081806102975750336001600160a01b038416145b6102d85760405162461bcd60e51b815260206004820152601260248201527134b73b30b634b21036b9b39739b2b73232b960711b604482015260640161017f565b7fad1f4fc02828882c8dc0931fe80a0ffbe0bf083316996ccf319e8b4200b9b0cd88600089600060405161030f94939291906109ea565b60405180910390a17fad1f4fc02828882c8dc0931fe80a0ffbe0bf083316996ccf319e8b4200b9b0cd88600188600060405161034e94939291906109ea565b60405180910390a15050604080516020808201835243825260008981526065909152918220905190556066805460018181018355919092527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e94354909101879055925050509392505050565b6066546000908190815b818110156104095784606682815481106103de576103de61097d565b9060005260206000200154036103f75760019250610409565b8061040181610993565b9150506103c2565b5081156104255750505060009081526065602052604090205490565b5060009392505050565b600054610100900460ff1661044a5760005460ff161561044e565b303b155b6104b15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161017f565b600054610100900460ff161580156104d3576000805461ffff19166101011790555b606780546001600160a01b0319166001600160a01b0384161790558015610500576000805461ff00191690555b5050565b6033546001600160a01b0316331461055e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161017f565b6001600160a01b0381166105c35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161017f565b6105cc816105cf565b50565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60408051600481526024810182526020810180516001600160e01b03166311ddc2b160e31b179052905160609160009182916001600160a01b038616916106689190610a36565b600060405180830381855afa9150503d80600081146106a3576040519150601f19603f3d011682016040523d82523d6000602084013e6106a8565b606091505b5091509150816106fa5760405162461bcd60e51b815260206004820152601b60248201527f73746174696363616c6c20616c6c6f77616e6365206661696c65640000000000604482015260640161017f565b6000818060200190518101906107109190610a52565b95945050505050565b60408051600481526024810182526020810180516001600160e01b0316636e9960c360e01b1790529051600091829182916001600160a01b0386169161075f9190610a36565b600060405180830381855afa9150503d806000811461079a576040519150601f19603f3d011682016040523d82523d6000602084013e61079f565b606091505b5091509150816107f15760405162461bcd60e51b815260206004820152601b60248201527f73746174696363616c6c20616c6c6f77616e6365206661696c65640000000000604482015260640161017f565b6000818060200190518101906107109190610b04565b60006020828403121561081957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561085f5761085f610820565b604052919050565b600082601f83011261087857600080fd5b813567ffffffffffffffff81111561089257610892610820565b6108a5601f8201601f1916602001610836565b8181528460208386010111156108ba57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156108ec57600080fd5b83359250602084013567ffffffffffffffff8082111561090b57600080fd5b61091787838801610867565b9350604086013591508082111561092d57600080fd5b5061093a86828701610867565b9150509250925092565b6001600160a01b03811681146105cc57600080fd5b60006020828403121561096b57600080fd5b813561097681610944565b9392505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016109b357634e487b7160e01b600052601160045260246000fd5b5060010190565b60005b838110156109d55781810151838201526020016109bd565b838111156109e4576000848401525b50505050565b84815260ff841660208201526080604082015260008351806080840152610a188160a08501602088016109ba565b606083019390935250601f91909101601f19160160a0019392505050565b60008251610a488184602087016109ba565b9190910192915050565b60006020808385031215610a6557600080fd5b825167ffffffffffffffff80821115610a7d57600080fd5b818501915085601f830112610a9157600080fd5b815181811115610aa357610aa3610820565b8060051b9150610ab4848301610836565b8181529183018401918481019088841115610ace57600080fd5b938501935b83851015610af85784519250610ae883610944565b8282529385019390850190610ad3565b98975050505050505050565b600060208284031215610b1657600080fd5b81516109768161094456fea2646970667358221220a7dce3ef873d12719777efae6084d5ad41fca0cdd78d98212fbd3f6b86a7dd8f64736f6c634300080d0033",
}

// CredentialABI is the input ABI used to generate the binding from.
// Deprecated: Use CredentialMetaData.ABI instead.
var CredentialABI = CredentialMetaData.ABI

// CredentialBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CredentialMetaData.Bin instead.
var CredentialBin = CredentialMetaData.Bin

// DeployCredential deploys a new Ethereum contract, binding an instance of Credential to it.
func DeployCredential(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Credential, error) {
	parsed, err := CredentialMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CredentialBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Credential{CredentialCaller: CredentialCaller{contract: contract}, CredentialTransactor: CredentialTransactor{contract: contract}, CredentialFilterer: CredentialFilterer{contract: contract}}, nil
}

// Credential is an auto generated Go binding around an Ethereum contract.
type Credential struct {
	CredentialCaller     // Read-only binding to the contract
	CredentialTransactor // Write-only binding to the contract
	CredentialFilterer   // Log filterer for contract events
}

// CredentialCaller is an auto generated read-only Go binding around an Ethereum contract.
type CredentialCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CredentialTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CredentialFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CredentialSession struct {
	Contract     *Credential       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CredentialCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CredentialCallerSession struct {
	Contract *CredentialCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CredentialTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CredentialTransactorSession struct {
	Contract     *CredentialTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CredentialRaw is an auto generated low-level Go binding around an Ethereum contract.
type CredentialRaw struct {
	Contract *Credential // Generic contract binding to access the raw methods on
}

// CredentialCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CredentialCallerRaw struct {
	Contract *CredentialCaller // Generic read-only contract binding to access the raw methods on
}

// CredentialTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CredentialTransactorRaw struct {
	Contract *CredentialTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredential creates a new instance of Credential, bound to a specific deployed contract.
func NewCredential(address common.Address, backend bind.ContractBackend) (*Credential, error) {
	contract, err := bindCredential(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Credential{CredentialCaller: CredentialCaller{contract: contract}, CredentialTransactor: CredentialTransactor{contract: contract}, CredentialFilterer: CredentialFilterer{contract: contract}}, nil
}

// NewCredentialCaller creates a new read-only instance of Credential, bound to a specific deployed contract.
func NewCredentialCaller(address common.Address, caller bind.ContractCaller) (*CredentialCaller, error) {
	contract, err := bindCredential(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CredentialCaller{contract: contract}, nil
}

// NewCredentialTransactor creates a new write-only instance of Credential, bound to a specific deployed contract.
func NewCredentialTransactor(address common.Address, transactor bind.ContractTransactor) (*CredentialTransactor, error) {
	contract, err := bindCredential(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CredentialTransactor{contract: contract}, nil
}

// NewCredentialFilterer creates a new log filterer instance of Credential, bound to a specific deployed contract.
func NewCredentialFilterer(address common.Address, filterer bind.ContractFilterer) (*CredentialFilterer, error) {
	contract, err := bindCredential(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CredentialFilterer{contract: contract}, nil
}

// bindCredential binds a generic wrapper to an already deployed contract.
func bindCredential(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CredentialABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credential *CredentialRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credential.Contract.CredentialCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credential *CredentialRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credential.Contract.CredentialTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credential *CredentialRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credential.Contract.CredentialTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credential *CredentialCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credential.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credential *CredentialTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credential.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credential *CredentialTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credential.Contract.contract.Transact(opts, method, params...)
}

// GetLatestBlock is a free data retrieval call binding the contract method 0xbfae4302.
//
// Solidity: function getLatestBlock(bytes32 credentialHash) view returns(uint256)
func (_Credential *CredentialCaller) GetLatestBlock(opts *bind.CallOpts, credentialHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Credential.contract.Call(opts, &out, "getLatestBlock", credentialHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestBlock is a free data retrieval call binding the contract method 0xbfae4302.
//
// Solidity: function getLatestBlock(bytes32 credentialHash) view returns(uint256)
func (_Credential *CredentialSession) GetLatestBlock(credentialHash [32]byte) (*big.Int, error) {
	return _Credential.Contract.GetLatestBlock(&_Credential.CallOpts, credentialHash)
}

// GetLatestBlock is a free data retrieval call binding the contract method 0xbfae4302.
//
// Solidity: function getLatestBlock(bytes32 credentialHash) view returns(uint256)
func (_Credential *CredentialCallerSession) GetLatestBlock(credentialHash [32]byte) (*big.Int, error) {
	return _Credential.Contract.GetLatestBlock(&_Credential.CallOpts, credentialHash)
}

// IsHashExist is a free data retrieval call binding the contract method 0x8a2d099c.
//
// Solidity: function isHashExist(bytes32 credentialHash) view returns(bool success)
func (_Credential *CredentialCaller) IsHashExist(opts *bind.CallOpts, credentialHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Credential.contract.Call(opts, &out, "isHashExist", credentialHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHashExist is a free data retrieval call binding the contract method 0x8a2d099c.
//
// Solidity: function isHashExist(bytes32 credentialHash) view returns(bool success)
func (_Credential *CredentialSession) IsHashExist(credentialHash [32]byte) (bool, error) {
	return _Credential.Contract.IsHashExist(&_Credential.CallOpts, credentialHash)
}

// IsHashExist is a free data retrieval call binding the contract method 0x8a2d099c.
//
// Solidity: function isHashExist(bytes32 credentialHash) view returns(bool success)
func (_Credential *CredentialCallerSession) IsHashExist(credentialHash [32]byte) (bool, error) {
	return _Credential.Contract.IsHashExist(&_Credential.CallOpts, credentialHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credential *CredentialCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Credential.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credential *CredentialSession) Owner() (common.Address, error) {
	return _Credential.Contract.Owner(&_Credential.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credential *CredentialCallerSession) Owner() (common.Address, error) {
	return _Credential.Contract.Owner(&_Credential.CallOpts)
}

// CreateCredential is a paid mutator transaction binding the contract method 0xabf66aed.
//
// Solidity: function createCredential(bytes32 credentialHash, string signer, string signatureData) returns(bool success)
func (_Credential *CredentialTransactor) CreateCredential(opts *bind.TransactOpts, credentialHash [32]byte, signer string, signatureData string) (*types.Transaction, error) {
	return _Credential.contract.Transact(opts, "createCredential", credentialHash, signer, signatureData)
}

// CreateCredential is a paid mutator transaction binding the contract method 0xabf66aed.
//
// Solidity: function createCredential(bytes32 credentialHash, string signer, string signatureData) returns(bool success)
func (_Credential *CredentialSession) CreateCredential(credentialHash [32]byte, signer string, signatureData string) (*types.Transaction, error) {
	return _Credential.Contract.CreateCredential(&_Credential.TransactOpts, credentialHash, signer, signatureData)
}

// CreateCredential is a paid mutator transaction binding the contract method 0xabf66aed.
//
// Solidity: function createCredential(bytes32 credentialHash, string signer, string signatureData) returns(bool success)
func (_Credential *CredentialTransactorSession) CreateCredential(credentialHash [32]byte, signer string, signatureData string) (*types.Transaction, error) {
	return _Credential.Contract.CreateCredential(&_Credential.TransactOpts, credentialHash, signer, signatureData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address voteAddress) returns()
func (_Credential *CredentialTransactor) Initialize(opts *bind.TransactOpts, voteAddress common.Address) (*types.Transaction, error) {
	return _Credential.contract.Transact(opts, "initialize", voteAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address voteAddress) returns()
func (_Credential *CredentialSession) Initialize(voteAddress common.Address) (*types.Transaction, error) {
	return _Credential.Contract.Initialize(&_Credential.TransactOpts, voteAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address voteAddress) returns()
func (_Credential *CredentialTransactorSession) Initialize(voteAddress common.Address) (*types.Transaction, error) {
	return _Credential.Contract.Initialize(&_Credential.TransactOpts, voteAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credential *CredentialTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credential.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credential *CredentialSession) RenounceOwnership() (*types.Transaction, error) {
	return _Credential.Contract.RenounceOwnership(&_Credential.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credential *CredentialTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Credential.Contract.RenounceOwnership(&_Credential.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credential *CredentialTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Credential.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credential *CredentialSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Credential.Contract.TransferOwnership(&_Credential.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credential *CredentialTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Credential.Contract.TransferOwnership(&_Credential.TransactOpts, newOwner)
}

// CredentialCredentialAttributeChangeIterator is returned from FilterCredentialAttributeChange and is used to iterate over the raw logs and unpacked data for CredentialAttributeChange events raised by the Credential contract.
type CredentialCredentialAttributeChangeIterator struct {
	Event *CredentialCredentialAttributeChange // Event containing the contract specifics and raw log

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
func (it *CredentialCredentialAttributeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CredentialCredentialAttributeChange)
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
		it.Event = new(CredentialCredentialAttributeChange)
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
func (it *CredentialCredentialAttributeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CredentialCredentialAttributeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CredentialCredentialAttributeChange represents a CredentialAttributeChange event raised by the Credential contract.
type CredentialCredentialAttributeChange struct {
	CredentialHash [32]byte
	FieldKey       uint8
	FieldValue     string
	BlockNumber    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCredentialAttributeChange is a free log retrieval operation binding the contract event 0xad1f4fc02828882c8dc0931fe80a0ffbe0bf083316996ccf319e8b4200b9b0cd.
//
// Solidity: event CredentialAttributeChange(bytes32 credentialHash, uint8 fieldKey, string fieldValue, uint256 blockNumber)
func (_Credential *CredentialFilterer) FilterCredentialAttributeChange(opts *bind.FilterOpts) (*CredentialCredentialAttributeChangeIterator, error) {

	logs, sub, err := _Credential.contract.FilterLogs(opts, "CredentialAttributeChange")
	if err != nil {
		return nil, err
	}
	return &CredentialCredentialAttributeChangeIterator{contract: _Credential.contract, event: "CredentialAttributeChange", logs: logs, sub: sub}, nil
}

// WatchCredentialAttributeChange is a free log subscription operation binding the contract event 0xad1f4fc02828882c8dc0931fe80a0ffbe0bf083316996ccf319e8b4200b9b0cd.
//
// Solidity: event CredentialAttributeChange(bytes32 credentialHash, uint8 fieldKey, string fieldValue, uint256 blockNumber)
func (_Credential *CredentialFilterer) WatchCredentialAttributeChange(opts *bind.WatchOpts, sink chan<- *CredentialCredentialAttributeChange) (event.Subscription, error) {

	logs, sub, err := _Credential.contract.WatchLogs(opts, "CredentialAttributeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CredentialCredentialAttributeChange)
				if err := _Credential.contract.UnpackLog(event, "CredentialAttributeChange", log); err != nil {
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

// ParseCredentialAttributeChange is a log parse operation binding the contract event 0xad1f4fc02828882c8dc0931fe80a0ffbe0bf083316996ccf319e8b4200b9b0cd.
//
// Solidity: event CredentialAttributeChange(bytes32 credentialHash, uint8 fieldKey, string fieldValue, uint256 blockNumber)
func (_Credential *CredentialFilterer) ParseCredentialAttributeChange(log types.Log) (*CredentialCredentialAttributeChange, error) {
	event := new(CredentialCredentialAttributeChange)
	if err := _Credential.contract.UnpackLog(event, "CredentialAttributeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CredentialOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Credential contract.
type CredentialOwnershipTransferredIterator struct {
	Event *CredentialOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CredentialOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CredentialOwnershipTransferred)
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
		it.Event = new(CredentialOwnershipTransferred)
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
func (it *CredentialOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CredentialOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CredentialOwnershipTransferred represents a OwnershipTransferred event raised by the Credential contract.
type CredentialOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Credential *CredentialFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CredentialOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Credential.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CredentialOwnershipTransferredIterator{contract: _Credential.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Credential *CredentialFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CredentialOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Credential.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CredentialOwnershipTransferred)
				if err := _Credential.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Credential *CredentialFilterer) ParseOwnershipTransferred(log types.Log) (*CredentialOwnershipTransferred, error) {
	event := new(CredentialOwnershipTransferred)
	if err := _Credential.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
