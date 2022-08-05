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

// PctMetaData contains all meta data concerning the Pct contract.
var PctMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pctId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"jsonSchema\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"name\":\"RegisterPct\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getNextPctId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pctId\",\"type\":\"uint256\"}],\"name\":\"getPctInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonSchema\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"extra\",\"type\":\"bytes\"}],\"name\":\"registerPct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pctId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d88806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638f795e221161005b5780638f795e22146100cd57806395220d00146100d5578063c4d66de8146100f7578063f2fde38b1461010a57600080fd5b80630f47e21d14610082578063715018a6146100a85780638da5cb5b146100b2575b600080fd5b6100956100903660046108b8565b61011d565b6040519081526020015b60405180910390f35b6100b06102c3565b005b6033546040516001600160a01b03909116815260200161009f565b606654610095565b6100e86100e3366004610947565b6102d7565b60405161009f939291906109bc565b6100b0610105366004610a07565b610427565b6100b0610118366004610a07565b610552565b6065546000908190610137906001600160a01b03166105cb565b8051909150600090815b8181101561019657336001600160a01b031684828151811061016557610165610a2b565b60200260200101516001600160a01b0316036101845760019250610196565b8061018e81610a57565b915050610141565b50816101de5760405162461bcd60e51b815260206004820152601260248201527134b73b30b634b21036b9b39739b2b73232b960711b60448201526064015b60405180910390fd5b60408051606081018252338152602080820189815282840189905260665460009081526067835293909320825181546001600160a01b0319166001600160a01b039091161781559251805192939261023c9260018501920190610772565b5060408201518051610258916002840191602090910190610772565b50506066546040513392507f455459674d2e900484971e2223af5549736bd88a214c9d38a17a0c2957569e9390610292908a908a90610a70565b60405180910390a36066546102a8906001610a9e565b60668190556102b990600190610ab6565b9695505050505050565b6102cb6106c6565b6102d56000610720565b565b6000818152606760205260408120805460018201805460609384936001600160a01b03169291600290910190829061030e90610acd565b80601f016020809104026020016040519081016040528092919081815260200182805461033a90610acd565b80156103875780601f1061035c57610100808354040283529160200191610387565b820191906000526020600020905b81548152906001019060200180831161036a57829003601f168201915b5050505050915080805461039a90610acd565b80601f01602080910402602001604051908101604052809291908181526020018280546103c690610acd565b80156104135780601f106103e857610100808354040283529160200191610413565b820191906000526020600020905b8154815290600101906020018083116103f657829003601f168201915b505050505090509250925092509193909250565b600054610100900460ff16158080156104475750600054600160ff909116105b806104615750303b158015610461575060005460ff166001145b6104c45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016101d5565b6000805460ff1916600117905580156104e7576000805461ff0019166101001790555b606580546001600160a01b0319166001600160a01b0384161790556103e8606655801561054e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b61055a6106c6565b6001600160a01b0381166105bf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101d5565b6105c881610720565b50565b60408051600481526024810182526020810180516001600160e01b03166311ddc2b160e31b179052905160609160009182916001600160a01b038616916106129190610b07565b600060405180830381855afa9150503d806000811461064d576040519150601f19603f3d011682016040523d82523d6000602084013e610652565b606091505b5091509150816106a45760405162461bcd60e51b815260206004820152601b60248201527f73746174696363616c6c20616c6c6f77616e6365206661696c6564000000000060448201526064016101d5565b6000818060200190518101906106ba9190610c69565b50909695505050505050565b6033546001600160a01b031633146102d55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101d5565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b82805461077e90610acd565b90600052602060002090601f0160209004810192826107a057600085556107e6565b82601f106107b957805160ff19168380011785556107e6565b828001600101855582156107e6579182015b828111156107e65782518255916020019190600101906107cb565b506107f29291506107f6565b5090565b5b808211156107f257600081556001016107f7565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561084a5761084a61080b565b604052919050565b600067ffffffffffffffff82111561086c5761086c61080b565b50601f01601f191660200190565b600061088d61088884610852565b610821565b90508281528383830111156108a157600080fd5b828260208301376000602084830101529392505050565b600080604083850312156108cb57600080fd5b823567ffffffffffffffff808211156108e357600080fd5b818501915085601f8301126108f757600080fd5b6109068683356020850161087a565b9350602085013591508082111561091c57600080fd5b508301601f8101851361092e57600080fd5b61093d8582356020840161087a565b9150509250929050565b60006020828403121561095957600080fd5b5035919050565b60005b8381101561097b578181015183820152602001610963565b8381111561098a576000848401525b50505050565b600081518084526109a8816020860160208601610960565b601f01601f19169290920160200192915050565b6001600160a01b03841681526060602082018190526000906109e090830185610990565b82810360408401526102b98185610990565b6001600160a01b03811681146105c857600080fd5b600060208284031215610a1957600080fd5b8135610a24816109f2565b9392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610a6957610a69610a41565b5060010190565b604081526000610a836040830185610990565b8281036020840152610a958185610990565b95945050505050565b60008219821115610ab157610ab1610a41565b500190565b600082821015610ac857610ac8610a41565b500390565b600181811c90821680610ae157607f821691505b602082108103610b0157634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610b19818460208701610960565b9190910192915050565b600067ffffffffffffffff821115610b3d57610b3d61080b565b5060051b60200190565b600082601f830112610b5857600080fd5b81516020610b6861088883610b23565b82815260059290921b84018101918181019086841115610b8757600080fd5b8286015b84811015610c0357805167ffffffffffffffff811115610bab5760008081fd5b8701603f81018913610bbd5760008081fd5b848101516040610bcf61088883610852565b8281528b82848601011115610be45760008081fd5b610bf383898301848701610960565b8652505050918301918301610b8b565b509695505050505050565b600082601f830112610c1f57600080fd5b81516020610c2f61088883610b23565b82815260059290921b84018101918181019086841115610c4e57600080fd5b8286015b84811015610c035780518352918301918301610c52565b600080600060608486031215610c7e57600080fd5b835167ffffffffffffffff80821115610c9657600080fd5b818601915086601f830112610caa57600080fd5b81516020610cba61088883610b23565b82815260059290921b8401810191818101908a841115610cd957600080fd5b948201945b83861015610d00578551610cf1816109f2565b82529482019490820190610cde565b91890151919750909350505080821115610d1957600080fd5b610d2587838801610b47565b93506040860151915080821115610d3b57600080fd5b50610d4886828701610c0e565b915050925092509256fea2646970667358221220fff074923486a4d3a303c10099eae109212b3058a7c24ecb5b2b81dcbc6732b764736f6c634300080d0033",
}

// PctABI is the input ABI used to generate the binding from.
// Deprecated: Use PctMetaData.ABI instead.
var PctABI = PctMetaData.ABI

// PctBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PctMetaData.Bin instead.
var PctBin = PctMetaData.Bin

// DeployPct deploys a new Ethereum contract, binding an instance of Pct to it.
func DeployPct(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pct, error) {
	parsed, err := PctMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PctBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pct{PctCaller: PctCaller{contract: contract}, PctTransactor: PctTransactor{contract: contract}, PctFilterer: PctFilterer{contract: contract}}, nil
}

// Pct is an auto generated Go binding around an Ethereum contract.
type Pct struct {
	PctCaller     // Read-only binding to the contract
	PctTransactor // Write-only binding to the contract
	PctFilterer   // Log filterer for contract events
}

// PctCaller is an auto generated read-only Go binding around an Ethereum contract.
type PctCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PctTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PctTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PctFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PctFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PctSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PctSession struct {
	Contract     *Pct              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PctCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PctCallerSession struct {
	Contract *PctCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PctTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PctTransactorSession struct {
	Contract     *PctTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PctRaw is an auto generated low-level Go binding around an Ethereum contract.
type PctRaw struct {
	Contract *Pct // Generic contract binding to access the raw methods on
}

// PctCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PctCallerRaw struct {
	Contract *PctCaller // Generic read-only contract binding to access the raw methods on
}

// PctTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PctTransactorRaw struct {
	Contract *PctTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPct creates a new instance of Pct, bound to a specific deployed contract.
func NewPct(address common.Address, backend bind.ContractBackend) (*Pct, error) {
	contract, err := bindPct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pct{PctCaller: PctCaller{contract: contract}, PctTransactor: PctTransactor{contract: contract}, PctFilterer: PctFilterer{contract: contract}}, nil
}

// NewPctCaller creates a new read-only instance of Pct, bound to a specific deployed contract.
func NewPctCaller(address common.Address, caller bind.ContractCaller) (*PctCaller, error) {
	contract, err := bindPct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PctCaller{contract: contract}, nil
}

// NewPctTransactor creates a new write-only instance of Pct, bound to a specific deployed contract.
func NewPctTransactor(address common.Address, transactor bind.ContractTransactor) (*PctTransactor, error) {
	contract, err := bindPct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PctTransactor{contract: contract}, nil
}

// NewPctFilterer creates a new log filterer instance of Pct, bound to a specific deployed contract.
func NewPctFilterer(address common.Address, filterer bind.ContractFilterer) (*PctFilterer, error) {
	contract, err := bindPct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PctFilterer{contract: contract}, nil
}

// bindPct binds a generic wrapper to an already deployed contract.
func bindPct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PctABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pct *PctRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pct.Contract.PctCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pct *PctRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pct.Contract.PctTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pct *PctRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pct.Contract.PctTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pct *PctCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pct *PctTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pct *PctTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pct.Contract.contract.Transact(opts, method, params...)
}

// GetNextPctId is a free data retrieval call binding the contract method 0x8f795e22.
//
// Solidity: function getNextPctId() view returns(uint256)
func (_Pct *PctCaller) GetNextPctId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pct.contract.Call(opts, &out, "getNextPctId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextPctId is a free data retrieval call binding the contract method 0x8f795e22.
//
// Solidity: function getNextPctId() view returns(uint256)
func (_Pct *PctSession) GetNextPctId() (*big.Int, error) {
	return _Pct.Contract.GetNextPctId(&_Pct.CallOpts)
}

// GetNextPctId is a free data retrieval call binding the contract method 0x8f795e22.
//
// Solidity: function getNextPctId() view returns(uint256)
func (_Pct *PctCallerSession) GetNextPctId() (*big.Int, error) {
	return _Pct.Contract.GetNextPctId(&_Pct.CallOpts)
}

// GetPctInfo is a free data retrieval call binding the contract method 0x95220d00.
//
// Solidity: function getPctInfo(uint256 pctId) view returns(address, string, bytes)
func (_Pct *PctCaller) GetPctInfo(opts *bind.CallOpts, pctId *big.Int) (common.Address, string, []byte, error) {
	var out []interface{}
	err := _Pct.contract.Call(opts, &out, "getPctInfo", pctId)

	if err != nil {
		return *new(common.Address), *new(string), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// GetPctInfo is a free data retrieval call binding the contract method 0x95220d00.
//
// Solidity: function getPctInfo(uint256 pctId) view returns(address, string, bytes)
func (_Pct *PctSession) GetPctInfo(pctId *big.Int) (common.Address, string, []byte, error) {
	return _Pct.Contract.GetPctInfo(&_Pct.CallOpts, pctId)
}

// GetPctInfo is a free data retrieval call binding the contract method 0x95220d00.
//
// Solidity: function getPctInfo(uint256 pctId) view returns(address, string, bytes)
func (_Pct *PctCallerSession) GetPctInfo(pctId *big.Int) (common.Address, string, []byte, error) {
	return _Pct.Contract.GetPctInfo(&_Pct.CallOpts, pctId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pct *PctCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pct.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pct *PctSession) Owner() (common.Address, error) {
	return _Pct.Contract.Owner(&_Pct.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pct *PctCallerSession) Owner() (common.Address, error) {
	return _Pct.Contract.Owner(&_Pct.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address voteAddress) returns()
func (_Pct *PctTransactor) Initialize(opts *bind.TransactOpts, voteAddress common.Address) (*types.Transaction, error) {
	return _Pct.contract.Transact(opts, "initialize", voteAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address voteAddress) returns()
func (_Pct *PctSession) Initialize(voteAddress common.Address) (*types.Transaction, error) {
	return _Pct.Contract.Initialize(&_Pct.TransactOpts, voteAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address voteAddress) returns()
func (_Pct *PctTransactorSession) Initialize(voteAddress common.Address) (*types.Transaction, error) {
	return _Pct.Contract.Initialize(&_Pct.TransactOpts, voteAddress)
}

// RegisterPct is a paid mutator transaction binding the contract method 0x0f47e21d.
//
// Solidity: function registerPct(string jsonSchema, bytes extra) returns(uint256 pctId)
func (_Pct *PctTransactor) RegisterPct(opts *bind.TransactOpts, jsonSchema string, extra []byte) (*types.Transaction, error) {
	return _Pct.contract.Transact(opts, "registerPct", jsonSchema, extra)
}

// RegisterPct is a paid mutator transaction binding the contract method 0x0f47e21d.
//
// Solidity: function registerPct(string jsonSchema, bytes extra) returns(uint256 pctId)
func (_Pct *PctSession) RegisterPct(jsonSchema string, extra []byte) (*types.Transaction, error) {
	return _Pct.Contract.RegisterPct(&_Pct.TransactOpts, jsonSchema, extra)
}

// RegisterPct is a paid mutator transaction binding the contract method 0x0f47e21d.
//
// Solidity: function registerPct(string jsonSchema, bytes extra) returns(uint256 pctId)
func (_Pct *PctTransactorSession) RegisterPct(jsonSchema string, extra []byte) (*types.Transaction, error) {
	return _Pct.Contract.RegisterPct(&_Pct.TransactOpts, jsonSchema, extra)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pct *PctTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pct.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pct *PctSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pct.Contract.RenounceOwnership(&_Pct.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pct *PctTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pct.Contract.RenounceOwnership(&_Pct.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pct *PctTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pct.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pct *PctSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pct.Contract.TransferOwnership(&_Pct.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pct *PctTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pct.Contract.TransferOwnership(&_Pct.TransactOpts, newOwner)
}

// PctInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pct contract.
type PctInitializedIterator struct {
	Event *PctInitialized // Event containing the contract specifics and raw log

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
func (it *PctInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PctInitialized)
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
		it.Event = new(PctInitialized)
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
func (it *PctInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PctInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PctInitialized represents a Initialized event raised by the Pct contract.
type PctInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pct *PctFilterer) FilterInitialized(opts *bind.FilterOpts) (*PctInitializedIterator, error) {

	logs, sub, err := _Pct.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PctInitializedIterator{contract: _Pct.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pct *PctFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PctInitialized) (event.Subscription, error) {

	logs, sub, err := _Pct.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PctInitialized)
				if err := _Pct.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pct *PctFilterer) ParseInitialized(log types.Log) (*PctInitialized, error) {
	event := new(PctInitialized)
	if err := _Pct.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PctOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pct contract.
type PctOwnershipTransferredIterator struct {
	Event *PctOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PctOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PctOwnershipTransferred)
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
		it.Event = new(PctOwnershipTransferred)
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
func (it *PctOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PctOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PctOwnershipTransferred represents a OwnershipTransferred event raised by the Pct contract.
type PctOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pct *PctFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PctOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pct.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PctOwnershipTransferredIterator{contract: _Pct.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pct *PctFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PctOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pct.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PctOwnershipTransferred)
				if err := _Pct.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pct *PctFilterer) ParseOwnershipTransferred(log types.Log) (*PctOwnershipTransferred, error) {
	event := new(PctOwnershipTransferred)
	if err := _Pct.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PctRegisterPctIterator is returned from FilterRegisterPct and is used to iterate over the raw logs and unpacked data for RegisterPct events raised by the Pct contract.
type PctRegisterPctIterator struct {
	Event *PctRegisterPct // Event containing the contract specifics and raw log

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
func (it *PctRegisterPctIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PctRegisterPct)
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
		it.Event = new(PctRegisterPct)
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
func (it *PctRegisterPctIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PctRegisterPctIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PctRegisterPct represents a RegisterPct event raised by the Pct contract.
type PctRegisterPct struct {
	PctId      *big.Int
	Issuer     common.Address
	JsonSchema string
	Extra      []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegisterPct is a free log retrieval operation binding the contract event 0x455459674d2e900484971e2223af5549736bd88a214c9d38a17a0c2957569e93.
//
// Solidity: event RegisterPct(uint256 indexed pctId, address indexed issuer, string jsonSchema, bytes extra)
func (_Pct *PctFilterer) FilterRegisterPct(opts *bind.FilterOpts, pctId []*big.Int, issuer []common.Address) (*PctRegisterPctIterator, error) {

	var pctIdRule []interface{}
	for _, pctIdItem := range pctId {
		pctIdRule = append(pctIdRule, pctIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Pct.contract.FilterLogs(opts, "RegisterPct", pctIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &PctRegisterPctIterator{contract: _Pct.contract, event: "RegisterPct", logs: logs, sub: sub}, nil
}

// WatchRegisterPct is a free log subscription operation binding the contract event 0x455459674d2e900484971e2223af5549736bd88a214c9d38a17a0c2957569e93.
//
// Solidity: event RegisterPct(uint256 indexed pctId, address indexed issuer, string jsonSchema, bytes extra)
func (_Pct *PctFilterer) WatchRegisterPct(opts *bind.WatchOpts, sink chan<- *PctRegisterPct, pctId []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var pctIdRule []interface{}
	for _, pctIdItem := range pctId {
		pctIdRule = append(pctIdRule, pctIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Pct.contract.WatchLogs(opts, "RegisterPct", pctIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PctRegisterPct)
				if err := _Pct.contract.UnpackLog(event, "RegisterPct", log); err != nil {
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

// ParseRegisterPct is a log parse operation binding the contract event 0x455459674d2e900484971e2223af5549736bd88a214c9d38a17a0c2957569e93.
//
// Solidity: event RegisterPct(uint256 indexed pctId, address indexed issuer, string jsonSchema, bytes extra)
func (_Pct *PctFilterer) ParseRegisterPct(log types.Log) (*PctRegisterPct, error) {
	event := new(PctRegisterPct)
	if err := _Pct.contract.UnpackLog(event, "RegisterPct", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
