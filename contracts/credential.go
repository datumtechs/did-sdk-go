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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fieldKey\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fieldValue\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"CredentialAttributeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"int8\",\"name\":\"status\",\"type\":\"int8\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"signerPublicKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signatureData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"updateTime\",\"type\":\"string\"}],\"name\":\"createCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"}],\"name\":\"getLatestBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"}],\"name\":\"isHashExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e7f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b14610103578063b2de79fd1461011e578063bfae430214610131578063c4d66de814610152578063f2fde38b1461016557600080fd5b80633ac87a5f146100985780635de28ae0146100ad578063715018a6146100d85780638a2d099c146100e0575b600080fd5b6100ab6100a6366004610947565b610178565b005b6100c06100bb36600461097d565b6102d3565b60405160009190910b81526020015b60405180910390f35b6100ab610348565b6100f36100ee36600461097d565b61035c565b60405190151581526020016100cf565b6033546040516001600160a01b0390911681526020016100cf565b6100f361012c366004610a5b565b6103ad565b61014461013f36600461097d565b61058b565b6040519081526020016100cf565b6100ab610160366004610b02565b610602565b6100ab610173366004610b02565b610727565b60665460009081805b828110156101cb57856066828154811061019d5761019d610b26565b9060005260206000200154036101b957600193508091506101cb565b806101c381610b3c565b915050610181565b508261021e5760405162461bcd60e51b815260206004820152601760248201527f646f63756d656e7420646f6573206e6f7420657869737400000000000000000060448201526064015b60405180910390fd5b6000858152606560205260409020600101546001600160a01b031633146102a05760405162461bcd60e51b815260206004820152603060248201527f4f6e6c7920746865206973737565722063616e206368616e676520746865206360448201526f726564656e7469616c2073746174757360801b6064820152608401610215565b505050600091825260656020526040909120600101805460ff909216600160a01b0260ff60a01b19909216919091179055565b60665460009060001990825b8181101561033f5784606682815481106102fb576102fb610b26565b90600052602060002001540361032d57600085815260656020526040812060010154600160a01b9004900b925061033f565b8061033781610b3c565b9150506102df565b50909392505050565b6103506107a0565b61035a60006107fa565b565b6066546000908190815b8181101561033f57846066828154811061038257610382610b26565b90600052602060002001540361039b576001925061033f565b806103a581610b3c565b915050610366565b60675460009081906103c7906001600160a01b031661084c565b8051909150600090815b8181101561042657336001600160a01b03168482815181106103f5576103f5610b26565b60200260200101516001600160a01b0316036104145760019250610426565b8061041e81610b3c565b9150506103d1565b50816104695760405162461bcd60e51b815260206004820152601260248201527134b73b30b634b21036b9b39739b2b73232b960711b6044820152606401610215565b877fbda6aff1adc27399496f953e769dd5eaea248b63011f5b641aae2d9531bbd3eb6000896000896040516104a19493929190610bbf565b60405180910390a2877fbda6aff1adc27399496f953e769dd5eaea248b63011f5b641aae2d9531bbd3eb6001886000896040516104e19493929190610bbf565b60405180910390a250506040805160608101825243815233602080830191825260008385018181528b8252606590925293842092518355905160019283018054925160ff16600160a01b026001600160a81b03199093166001600160a01b03929092169190911791909117905560668054808301825592527f46501879b8ca8525e8c2fd519e2fbfcfa2ebea26501294aa02cbfcfb12e94354909101879055915050949350505050565b6066546000908190815b818110156105dc5784606682815481106105b1576105b1610b26565b9060005260206000200154036105ca57600192506105dc565b806105d481610b3c565b915050610595565b5081156105f85750505060009081526065602052604090205490565b5060009392505050565b600054610100900460ff16158080156106225750600054600160ff909116105b8061063c5750303b15801561063c575060005460ff166001145b61069f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610215565b6000805460ff1916600117905580156106c2576000805461ff0019166101001790555b606780546001600160a01b0319166001600160a01b0384161790558015610723576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b61072f6107a0565b6001600160a01b0381166107945760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610215565b61079d816107fa565b50565b6033546001600160a01b0316331461035a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610215565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60408051600481526024810182526020810180516001600160e01b03166311ddc2b160e31b179052905160609160009182916001600160a01b038616916108939190610bfe565b600060405180830381855afa9150503d80600081146108ce576040519150601f19603f3d011682016040523d82523d6000602084013e6108d3565b606091505b5091509150816109255760405162461bcd60e51b815260206004820152601b60248201527f73746174696363616c6c20616c6c6f77616e6365206661696c656400000000006044820152606401610215565b60008180602001905181019061093b9190610d60565b50909695505050505050565b6000806040838503121561095a57600080fd5b8235915060208301358060000b811461097257600080fd5b809150509250929050565b60006020828403121561098f57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156109d5576109d5610996565b604052919050565b600067ffffffffffffffff8211156109f7576109f7610996565b50601f01601f191660200190565b600082601f830112610a1657600080fd5b8135610a29610a24826109dd565b6109ac565b818152846020838601011115610a3e57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215610a7157600080fd5b84359350602085013567ffffffffffffffff80821115610a9057600080fd5b610a9c88838901610a05565b94506040870135915080821115610ab257600080fd5b610abe88838901610a05565b93506060870135915080821115610ad457600080fd5b50610ae187828801610a05565b91505092959194509250565b6001600160a01b038116811461079d57600080fd5b600060208284031215610b1457600080fd5b8135610b1f81610aed565b9392505050565b634e487b7160e01b600052603260045260246000fd5b600060018201610b5c57634e487b7160e01b600052601160045260246000fd5b5060010190565b60005b83811015610b7e578181015183820152602001610b66565b83811115610b8d576000848401525b50505050565b60008151808452610bab816020860160208601610b63565b601f01601f19169290920160200192915050565b60ff85168152608060208201526000610bdb6080830186610b93565b8460408401528281036060840152610bf38185610b93565b979650505050505050565b60008251610c10818460208701610b63565b9190910192915050565b600067ffffffffffffffff821115610c3457610c34610996565b5060051b60200190565b600082601f830112610c4f57600080fd5b81516020610c5f610a2483610c1a565b82815260059290921b84018101918181019086841115610c7e57600080fd5b8286015b84811015610cfa57805167ffffffffffffffff811115610ca25760008081fd5b8701603f81018913610cb45760008081fd5b848101516040610cc6610a24836109dd565b8281528b82848601011115610cdb5760008081fd5b610cea83898301848701610b63565b8652505050918301918301610c82565b509695505050505050565b600082601f830112610d1657600080fd5b81516020610d26610a2483610c1a565b82815260059290921b84018101918181019086841115610d4557600080fd5b8286015b84811015610cfa5780518352918301918301610d49565b600080600060608486031215610d7557600080fd5b835167ffffffffffffffff80821115610d8d57600080fd5b818601915086601f830112610da157600080fd5b81516020610db1610a2483610c1a565b82815260059290921b8401810191818101908a841115610dd057600080fd5b948201945b83861015610df7578551610de881610aed565b82529482019490820190610dd5565b91890151919750909350505080821115610e1057600080fd5b610e1c87838801610c3e565b93506040860151915080821115610e3257600080fd5b50610e3f86828701610d05565b915050925092509256fea26469706673582212202c847dfe5fb349ba73093aed82d725ff2e3ce4b704d9668d5deb95ec3c6e4f2564736f6c634300080d0033",
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

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 credentialHash) view returns(int8)
func (_Credential *CredentialCaller) GetStatus(opts *bind.CallOpts, credentialHash [32]byte) (int8, error) {
	var out []interface{}
	err := _Credential.contract.Call(opts, &out, "getStatus", credentialHash)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 credentialHash) view returns(int8)
func (_Credential *CredentialSession) GetStatus(credentialHash [32]byte) (int8, error) {
	return _Credential.Contract.GetStatus(&_Credential.CallOpts, credentialHash)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 credentialHash) view returns(int8)
func (_Credential *CredentialCallerSession) GetStatus(credentialHash [32]byte) (int8, error) {
	return _Credential.Contract.GetStatus(&_Credential.CallOpts, credentialHash)
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

// ChangeStatus is a paid mutator transaction binding the contract method 0x3ac87a5f.
//
// Solidity: function changeStatus(bytes32 credentialHash, int8 status) returns()
func (_Credential *CredentialTransactor) ChangeStatus(opts *bind.TransactOpts, credentialHash [32]byte, status int8) (*types.Transaction, error) {
	return _Credential.contract.Transact(opts, "changeStatus", credentialHash, status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x3ac87a5f.
//
// Solidity: function changeStatus(bytes32 credentialHash, int8 status) returns()
func (_Credential *CredentialSession) ChangeStatus(credentialHash [32]byte, status int8) (*types.Transaction, error) {
	return _Credential.Contract.ChangeStatus(&_Credential.TransactOpts, credentialHash, status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0x3ac87a5f.
//
// Solidity: function changeStatus(bytes32 credentialHash, int8 status) returns()
func (_Credential *CredentialTransactorSession) ChangeStatus(credentialHash [32]byte, status int8) (*types.Transaction, error) {
	return _Credential.Contract.ChangeStatus(&_Credential.TransactOpts, credentialHash, status)
}

// CreateCredential is a paid mutator transaction binding the contract method 0xb2de79fd.
//
// Solidity: function createCredential(bytes32 credentialHash, string signerPublicKey, string signatureData, string updateTime) returns(bool success)
func (_Credential *CredentialTransactor) CreateCredential(opts *bind.TransactOpts, credentialHash [32]byte, signerPublicKey string, signatureData string, updateTime string) (*types.Transaction, error) {
	return _Credential.contract.Transact(opts, "createCredential", credentialHash, signerPublicKey, signatureData, updateTime)
}

// CreateCredential is a paid mutator transaction binding the contract method 0xb2de79fd.
//
// Solidity: function createCredential(bytes32 credentialHash, string signerPublicKey, string signatureData, string updateTime) returns(bool success)
func (_Credential *CredentialSession) CreateCredential(credentialHash [32]byte, signerPublicKey string, signatureData string, updateTime string) (*types.Transaction, error) {
	return _Credential.Contract.CreateCredential(&_Credential.TransactOpts, credentialHash, signerPublicKey, signatureData, updateTime)
}

// CreateCredential is a paid mutator transaction binding the contract method 0xb2de79fd.
//
// Solidity: function createCredential(bytes32 credentialHash, string signerPublicKey, string signatureData, string updateTime) returns(bool success)
func (_Credential *CredentialTransactorSession) CreateCredential(credentialHash [32]byte, signerPublicKey string, signatureData string, updateTime string) (*types.Transaction, error) {
	return _Credential.Contract.CreateCredential(&_Credential.TransactOpts, credentialHash, signerPublicKey, signatureData, updateTime)
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
	UpdateTime     string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCredentialAttributeChange is a free log retrieval operation binding the contract event 0xbda6aff1adc27399496f953e769dd5eaea248b63011f5b641aae2d9531bbd3eb.
//
// Solidity: event CredentialAttributeChange(bytes32 indexed credentialHash, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Credential *CredentialFilterer) FilterCredentialAttributeChange(opts *bind.FilterOpts, credentialHash [][32]byte) (*CredentialCredentialAttributeChangeIterator, error) {

	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _Credential.contract.FilterLogs(opts, "CredentialAttributeChange", credentialHashRule)
	if err != nil {
		return nil, err
	}
	return &CredentialCredentialAttributeChangeIterator{contract: _Credential.contract, event: "CredentialAttributeChange", logs: logs, sub: sub}, nil
}

// WatchCredentialAttributeChange is a free log subscription operation binding the contract event 0xbda6aff1adc27399496f953e769dd5eaea248b63011f5b641aae2d9531bbd3eb.
//
// Solidity: event CredentialAttributeChange(bytes32 indexed credentialHash, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Credential *CredentialFilterer) WatchCredentialAttributeChange(opts *bind.WatchOpts, sink chan<- *CredentialCredentialAttributeChange, credentialHash [][32]byte) (event.Subscription, error) {

	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _Credential.contract.WatchLogs(opts, "CredentialAttributeChange", credentialHashRule)
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

// ParseCredentialAttributeChange is a log parse operation binding the contract event 0xbda6aff1adc27399496f953e769dd5eaea248b63011f5b641aae2d9531bbd3eb.
//
// Solidity: event CredentialAttributeChange(bytes32 indexed credentialHash, uint8 fieldKey, string fieldValue, uint256 blockNumber, string updateTime)
func (_Credential *CredentialFilterer) ParseCredentialAttributeChange(log types.Log) (*CredentialCredentialAttributeChange, error) {
	event := new(CredentialCredentialAttributeChange)
	if err := _Credential.contract.UnpackLog(event, "CredentialAttributeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CredentialInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Credential contract.
type CredentialInitializedIterator struct {
	Event *CredentialInitialized // Event containing the contract specifics and raw log

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
func (it *CredentialInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CredentialInitialized)
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
		it.Event = new(CredentialInitialized)
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
func (it *CredentialInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CredentialInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CredentialInitialized represents a Initialized event raised by the Credential contract.
type CredentialInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Credential *CredentialFilterer) FilterInitialized(opts *bind.FilterOpts) (*CredentialInitializedIterator, error) {

	logs, sub, err := _Credential.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CredentialInitializedIterator{contract: _Credential.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Credential *CredentialFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CredentialInitialized) (event.Subscription, error) {

	logs, sub, err := _Credential.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CredentialInitialized)
				if err := _Credential.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Credential *CredentialFilterer) ParseInitialized(log types.Log) (*CredentialInitialized, error) {
	event := new(CredentialInitialized)
	if err := _Credential.contract.UnpackLog(event, "Initialized", log); err != nil {
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
