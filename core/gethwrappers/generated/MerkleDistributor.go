// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merkledistributor

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200ba2c6ff950bc2af74cc1aaa06e4e3fd9098072335b41fcd00cadd595d48699b64736f6c63430008110033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AggregatorMerkleInterfaceMetaData contains all meta data concerning the AggregatorMerkleInterface contract.
var AggregatorMerkleInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestMerkleRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkelAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"}],\"name\":\"setLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AggregatorMerkleInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorMerkleInterfaceMetaData.ABI instead.
var AggregatorMerkleInterfaceABI = AggregatorMerkleInterfaceMetaData.ABI

// AggregatorMerkleInterface is an auto generated Go binding around an Ethereum contract.
type AggregatorMerkleInterface struct {
	AggregatorMerkleInterfaceCaller     // Read-only binding to the contract
	AggregatorMerkleInterfaceTransactor // Write-only binding to the contract
	AggregatorMerkleInterfaceFilterer   // Log filterer for contract events
}

// AggregatorMerkleInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorMerkleInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorMerkleInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorMerkleInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorMerkleInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorMerkleInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorMerkleInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorMerkleInterfaceSession struct {
	Contract     *AggregatorMerkleInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AggregatorMerkleInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorMerkleInterfaceCallerSession struct {
	Contract *AggregatorMerkleInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// AggregatorMerkleInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorMerkleInterfaceTransactorSession struct {
	Contract     *AggregatorMerkleInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// AggregatorMerkleInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorMerkleInterfaceRaw struct {
	Contract *AggregatorMerkleInterface // Generic contract binding to access the raw methods on
}

// AggregatorMerkleInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorMerkleInterfaceCallerRaw struct {
	Contract *AggregatorMerkleInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorMerkleInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorMerkleInterfaceTransactorRaw struct {
	Contract *AggregatorMerkleInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorMerkleInterface creates a new instance of AggregatorMerkleInterface, bound to a specific deployed contract.
func NewAggregatorMerkleInterface(address common.Address, backend bind.ContractBackend) (*AggregatorMerkleInterface, error) {
	contract, err := bindAggregatorMerkleInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorMerkleInterface{AggregatorMerkleInterfaceCaller: AggregatorMerkleInterfaceCaller{contract: contract}, AggregatorMerkleInterfaceTransactor: AggregatorMerkleInterfaceTransactor{contract: contract}, AggregatorMerkleInterfaceFilterer: AggregatorMerkleInterfaceFilterer{contract: contract}}, nil
}

// NewAggregatorMerkleInterfaceCaller creates a new read-only instance of AggregatorMerkleInterface, bound to a specific deployed contract.
func NewAggregatorMerkleInterfaceCaller(address common.Address, caller bind.ContractCaller) (*AggregatorMerkleInterfaceCaller, error) {
	contract, err := bindAggregatorMerkleInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorMerkleInterfaceCaller{contract: contract}, nil
}

// NewAggregatorMerkleInterfaceTransactor creates a new write-only instance of AggregatorMerkleInterface, bound to a specific deployed contract.
func NewAggregatorMerkleInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorMerkleInterfaceTransactor, error) {
	contract, err := bindAggregatorMerkleInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorMerkleInterfaceTransactor{contract: contract}, nil
}

// NewAggregatorMerkleInterfaceFilterer creates a new log filterer instance of AggregatorMerkleInterface, bound to a specific deployed contract.
func NewAggregatorMerkleInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorMerkleInterfaceFilterer, error) {
	contract, err := bindAggregatorMerkleInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorMerkleInterfaceFilterer{contract: contract}, nil
}

// bindAggregatorMerkleInterface binds a generic wrapper to an already deployed contract.
func bindAggregatorMerkleInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregatorMerkleInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorMerkleInterface.Contract.AggregatorMerkleInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.Contract.AggregatorMerkleInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.Contract.AggregatorMerkleInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorMerkleInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.Contract.contract.Transact(opts, method, params...)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AggregatorMerkleInterface.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceSession) IsLocked() (bool, error) {
	return _AggregatorMerkleInterface.Contract.IsLocked(&_AggregatorMerkleInterface.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceCallerSession) IsLocked() (bool, error) {
	return _AggregatorMerkleInterface.Contract.IsLocked(&_AggregatorMerkleInterface.CallOpts)
}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xc00675d6.
//
// Solidity: function latestMerkleRoundData() view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceCaller) LatestMerkleRoundData(opts *bind.CallOpts) (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorMerkleInterface.contract.Call(opts, &out, "latestMerkleRoundData")

	outstruct := new(struct {
		RoundId      *big.Int
		BatchId      *big.Int
		MerkelAnswer [32]byte
		StartedAt    *big.Int
		UpdatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MerkelAnswer = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StartedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xc00675d6.
//
// Solidity: function latestMerkleRoundData() view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceSession) LatestMerkleRoundData() (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _AggregatorMerkleInterface.Contract.LatestMerkleRoundData(&_AggregatorMerkleInterface.CallOpts)
}

// LatestMerkleRoundData is a free data retrieval call binding the contract method 0xc00675d6.
//
// Solidity: function latestMerkleRoundData() view returns(uint80 roundId, uint256 batchId, bytes32 merkelAnswer, uint256 startedAt, uint256 updatedAt)
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceCallerSession) LatestMerkleRoundData() (struct {
	RoundId      *big.Int
	BatchId      *big.Int
	MerkelAnswer [32]byte
	StartedAt    *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _AggregatorMerkleInterface.Contract.LatestMerkleRoundData(&_AggregatorMerkleInterface.CallOpts)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceTransactor) SetLock(opts *bind.TransactOpts, isLock bool) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.contract.Transact(opts, "setLock", isLock)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceSession) SetLock(isLock bool) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.Contract.SetLock(&_AggregatorMerkleInterface.TransactOpts, isLock)
}

// SetLock is a paid mutator transaction binding the contract method 0x619d5194.
//
// Solidity: function setLock(bool isLock) returns()
func (_AggregatorMerkleInterface *AggregatorMerkleInterfaceTransactorSession) SetLock(isLock bool) (*types.Transaction, error) {
	return _AggregatorMerkleInterface.Contract.SetLock(&_AggregatorMerkleInterface.TransactOpts, isLock)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20PermitMetaData contains all meta data concerning the IERC20Permit contract.
var IERC20PermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitMetaData.ABI instead.
var IERC20PermitABI = IERC20PermitMetaData.ABI

// IERC20Permit is an auto generated Go binding around an Ethereum contract.
type IERC20Permit struct {
	IERC20PermitCaller     // Read-only binding to the contract
	IERC20PermitTransactor // Write-only binding to the contract
	IERC20PermitFilterer   // Log filterer for contract events
}

// IERC20PermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitSession struct {
	Contract     *IERC20Permit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20PermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitCallerSession struct {
	Contract *IERC20PermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IERC20PermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitTransactorSession struct {
	Contract     *IERC20PermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IERC20PermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitRaw struct {
	Contract *IERC20Permit // Generic contract binding to access the raw methods on
}

// IERC20PermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitCallerRaw struct {
	Contract *IERC20PermitCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitTransactorRaw struct {
	Contract *IERC20PermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Permit creates a new instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20Permit(address common.Address, backend bind.ContractBackend) (*IERC20Permit, error) {
	contract, err := bindIERC20Permit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Permit{IERC20PermitCaller: IERC20PermitCaller{contract: contract}, IERC20PermitTransactor: IERC20PermitTransactor{contract: contract}, IERC20PermitFilterer: IERC20PermitFilterer{contract: contract}}, nil
}

// NewIERC20PermitCaller creates a new read-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitCaller, error) {
	contract, err := bindIERC20Permit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitCaller{contract: contract}, nil
}

// NewIERC20PermitTransactor creates a new write-only instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitTransactor, error) {
	contract, err := bindIERC20Permit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitTransactor{contract: contract}, nil
}

// NewIERC20PermitFilterer creates a new log filterer instance of IERC20Permit, bound to a specific deployed contract.
func NewIERC20PermitFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitFilterer, error) {
	contract, err := bindIERC20Permit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitFilterer{contract: contract}, nil
}

// bindIERC20Permit binds a generic wrapper to an already deployed contract.
func bindIERC20Permit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20PermitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.IERC20PermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.IERC20PermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Permit *IERC20PermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Permit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Permit *IERC20PermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Permit.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20Permit *IERC20PermitCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20Permit.Contract.DOMAINSEPARATOR(&_IERC20Permit.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Permit.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20Permit *IERC20PermitCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20Permit.Contract.Nonces(&_IERC20Permit.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20Permit *IERC20PermitTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20Permit.Contract.Permit(&_IERC20Permit.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// IMerkleDistributorMetaData contains all meta data concerning the IMerkleDistributor contract.
var IMerkleDistributorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IMerkleDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use IMerkleDistributorMetaData.ABI instead.
var IMerkleDistributorABI = IMerkleDistributorMetaData.ABI

// IMerkleDistributor is an auto generated Go binding around an Ethereum contract.
type IMerkleDistributor struct {
	IMerkleDistributorCaller     // Read-only binding to the contract
	IMerkleDistributorTransactor // Write-only binding to the contract
	IMerkleDistributorFilterer   // Log filterer for contract events
}

// IMerkleDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMerkleDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMerkleDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMerkleDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMerkleDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMerkleDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMerkleDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMerkleDistributorSession struct {
	Contract     *IMerkleDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IMerkleDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMerkleDistributorCallerSession struct {
	Contract *IMerkleDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IMerkleDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMerkleDistributorTransactorSession struct {
	Contract     *IMerkleDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IMerkleDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMerkleDistributorRaw struct {
	Contract *IMerkleDistributor // Generic contract binding to access the raw methods on
}

// IMerkleDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMerkleDistributorCallerRaw struct {
	Contract *IMerkleDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// IMerkleDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMerkleDistributorTransactorRaw struct {
	Contract *IMerkleDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMerkleDistributor creates a new instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributor(address common.Address, backend bind.ContractBackend) (*IMerkleDistributor, error) {
	contract, err := bindIMerkleDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributor{IMerkleDistributorCaller: IMerkleDistributorCaller{contract: contract}, IMerkleDistributorTransactor: IMerkleDistributorTransactor{contract: contract}, IMerkleDistributorFilterer: IMerkleDistributorFilterer{contract: contract}}, nil
}

// NewIMerkleDistributorCaller creates a new read-only instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributorCaller(address common.Address, caller bind.ContractCaller) (*IMerkleDistributorCaller, error) {
	contract, err := bindIMerkleDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorCaller{contract: contract}, nil
}

// NewIMerkleDistributorTransactor creates a new write-only instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*IMerkleDistributorTransactor, error) {
	contract, err := bindIMerkleDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorTransactor{contract: contract}, nil
}

// NewIMerkleDistributorFilterer creates a new log filterer instance of IMerkleDistributor, bound to a specific deployed contract.
func NewIMerkleDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*IMerkleDistributorFilterer, error) {
	contract, err := bindIMerkleDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorFilterer{contract: contract}, nil
}

// bindIMerkleDistributor binds a generic wrapper to an already deployed contract.
func bindIMerkleDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMerkleDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMerkleDistributor *IMerkleDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMerkleDistributor.Contract.IMerkleDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMerkleDistributor *IMerkleDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.IMerkleDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMerkleDistributor *IMerkleDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.IMerkleDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMerkleDistributor *IMerkleDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMerkleDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMerkleDistributor *IMerkleDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMerkleDistributor *IMerkleDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.contract.Transact(opts, method, params...)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_IMerkleDistributor *IMerkleDistributorCaller) IsClaimed(opts *bind.CallOpts, batch *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _IMerkleDistributor.contract.Call(opts, &out, "isClaimed", batch, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_IMerkleDistributor *IMerkleDistributorSession) IsClaimed(batch *big.Int, index *big.Int) (bool, error) {
	return _IMerkleDistributor.Contract.IsClaimed(&_IMerkleDistributor.CallOpts, batch, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_IMerkleDistributor *IMerkleDistributorCallerSession) IsClaimed(batch *big.Int, index *big.Int) (bool, error) {
	return _IMerkleDistributor.Contract.IsClaimed(&_IMerkleDistributor.CallOpts, batch, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_IMerkleDistributor *IMerkleDistributorCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IMerkleDistributor.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_IMerkleDistributor *IMerkleDistributorSession) MerkleRoot() ([32]byte, error) {
	return _IMerkleDistributor.Contract.MerkleRoot(&_IMerkleDistributor.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_IMerkleDistributor *IMerkleDistributorCallerSession) MerkleRoot() ([32]byte, error) {
	return _IMerkleDistributor.Contract.MerkleRoot(&_IMerkleDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IMerkleDistributor *IMerkleDistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMerkleDistributor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IMerkleDistributor *IMerkleDistributorSession) Token() (common.Address, error) {
	return _IMerkleDistributor.Contract.Token(&_IMerkleDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IMerkleDistributor *IMerkleDistributorCallerSession) Token() (common.Address, error) {
	return _IMerkleDistributor.Contract.Token(&_IMerkleDistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_IMerkleDistributor *IMerkleDistributorTransactor) Claim(opts *bind.TransactOpts, batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IMerkleDistributor.contract.Transact(opts, "claim", batch, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_IMerkleDistributor *IMerkleDistributorSession) Claim(batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.Claim(&_IMerkleDistributor.TransactOpts, batch, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_IMerkleDistributor *IMerkleDistributorTransactorSession) Claim(batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _IMerkleDistributor.Contract.Claim(&_IMerkleDistributor.TransactOpts, batch, index, account, amount, merkleProof)
}

// IMerkleDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the IMerkleDistributor contract.
type IMerkleDistributorClaimedIterator struct {
	Event *IMerkleDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *IMerkleDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMerkleDistributorClaimed)
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
		it.Event = new(IMerkleDistributorClaimed)
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
func (it *IMerkleDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMerkleDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMerkleDistributorClaimed represents a Claimed event raised by the IMerkleDistributor contract.
type IMerkleDistributorClaimed struct {
	Batch   *big.Int
	Index   *big.Int
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_IMerkleDistributor *IMerkleDistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*IMerkleDistributorClaimedIterator, error) {

	logs, sub, err := _IMerkleDistributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &IMerkleDistributorClaimedIterator{contract: _IMerkleDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_IMerkleDistributor *IMerkleDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *IMerkleDistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _IMerkleDistributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMerkleDistributorClaimed)
				if err := _IMerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_IMerkleDistributor *IMerkleDistributorFilterer) ParseClaimed(log types.Log) (*IMerkleDistributorClaimed, error) {
	event := new(IMerkleDistributorClaimed)
	if err := _IMerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleDistributorMetaData contains all meta data concerning the MerkleDistributor contract.
var MerkleDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregatorProxy_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"projectId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620015e1380380620015e183398181016040528101906200003791906200016d565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600281905550505050620001c9565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000fa82620000cd565b9050919050565b6200010c81620000ed565b81146200011857600080fd5b50565b6000815190506200012c8162000101565b92915050565b6000819050919050565b620001478162000132565b81146200015357600080fd5b50565b60008151905062000167816200013c565b92915050565b600080600060608486031215620001895762000188620000c8565b5b600062000199868287016200011b565b9350506020620001ac868287016200011b565b9250506040620001bf8682870162000156565b9150509250925092565b61140880620001d96000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638cd221c9116100665780638cd221c91461014b578063b1de162614610169578063f21f537d14610187578063f364c90c146101a5578063fc0c546a146101d55761009e565b80632eb4a7ab146100a35780633fafa127146100c15780635d4df3bf146100df578063733db88f1461010f5780637519ab501461012d575b600080fd5b6100ab6101f3565b6040516100b89190610b46565b60405180910390f35b6100c96101f9565b6040516100d69190610b7a565b60405180910390f35b6100f960048036038101906100f49190610c8e565b6101ff565b6040516101069190610d43565b60405180910390f35b610117610370565b6040516101249190610b7a565b60405180910390f35b610135610376565b6040516101429190610b7a565b60405180910390f35b61015361037c565b6040516101609190610d83565b60405180910390f35b610171610398565b60405161017e9190610d43565b60405180910390f35b61018f6105d7565b60405161019c9190610b7a565b60405180910390f35b6101bf60048036038101906101ba9190610d9e565b6105dd565b6040516101cc9190610d43565b60405180910390f35b6101dd61061e565b6040516101ea9190610ded565b60405180910390f35b60035481565b60025481565b6000610209610398565b6102165760009050610366565b60045487146102285760009050610366565b610234600454876105dd565b156102425760009050610366565b600086868660405160200161025993929190610e71565b6040516020818303038152906040528051906020012090506102bf848480806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505060035483610644565b6102cd576000915050610366565b6102d68761065b565b6103238686600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661077b9092919063ffffffff16565b7fb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba888888886040516103589493929190610eae565b60405180910390a160019150505b9695505050505050565b60045481565b60075481565b600560009054906101000a900469ffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a4e2d6346040518163ffffffff1660e01b8152600401602060405180830381865afa158015610406573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042a9190610f1f565b1561043857600090506105d4565b60008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c00675d66040518163ffffffff1660e01b815260040160a060405180830381865afa1580156104ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d09190610fb9565b809550819650829850839950849750505050505060045485146105ca576000600980549050905060005b8181101561056c57600860006009838154811061051a57610519611034565b5b90600052602060002090600202016000015481526020019081526020016000206000808201600090556001820160006101000a81549060ff02191690555050808061056490611092565b9150506104fa565b505b600060098054905011156105c857600980548061058e5761058d6110da565b5b60019003818190600052602060002090600202016000808201600090556001820160006101000a81549060ff02191690555050905561056e565b505b6001955050505050505b90565b60065481565b600082600454146105f15760019050610618565b6008600083815260200190815260200160002060010160009054906101000a900460ff1690505b92915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000826106518584610801565b1490509392505050565b600060086000838152602001908152602001600020604051806040016040529081600082015481526020016001820160009054906101000a900460ff1615151515815250509050600015158160200151151503610777576000600980549050905060016008600085815260200190815260200160002060010160006101000a81548160ff021916908315150217905550806008600085815260200190815260200160002060000181905550600960405180604001604052808581526020016000151581525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010160006101000a81548160ff0219169083151502179055505050505b5050565b6107fc8363a9059cbb60e01b848460405160240161079a929190611109565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610857565b505050565b60008082905060005b845181101561084c576108378286838151811061082a57610829611034565b5b602002602001015161091e565b9150808061084490611092565b91505061080a565b508091505092915050565b60006108b9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166109499092919063ffffffff16565b905060008151111561091957808060200190518101906108d99190610f1f565b610918576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090f906111b5565b60405180910390fd5b5b505050565b6000818310610936576109318284610961565b610941565b6109408383610961565b5b905092915050565b60606109588484600085610978565b90509392505050565b600082600052816020526040600020905092915050565b6060824710156109bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b490611247565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516109e691906112d8565b60006040518083038185875af1925050503d8060008114610a23576040519150601f19603f3d011682016040523d82523d6000602084013e610a28565b606091505b5091509150610a3987838387610a45565b92505050949350505050565b60608315610aa7576000835103610a9f57610a5f85610aba565b610a9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a959061133b565b60405180910390fd5b5b829050610ab2565b610ab18383610add565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115610af05781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2491906113b0565b60405180910390fd5b6000819050919050565b610b4081610b2d565b82525050565b6000602082019050610b5b6000830184610b37565b92915050565b6000819050919050565b610b7481610b61565b82525050565b6000602082019050610b8f6000830184610b6b565b92915050565b600080fd5b600080fd5b610ba881610b61565b8114610bb357600080fd5b50565b600081359050610bc581610b9f565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bf682610bcb565b9050919050565b610c0681610beb565b8114610c1157600080fd5b50565b600081359050610c2381610bfd565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610c4e57610c4d610c29565b5b8235905067ffffffffffffffff811115610c6b57610c6a610c2e565b5b602083019150836020820283011115610c8757610c86610c33565b5b9250929050565b60008060008060008060a08789031215610cab57610caa610b95565b5b6000610cb989828a01610bb6565b9650506020610cca89828a01610bb6565b9550506040610cdb89828a01610c14565b9450506060610cec89828a01610bb6565b935050608087013567ffffffffffffffff811115610d0d57610d0c610b9a565b5b610d1989828a01610c38565b92509250509295509295509295565b60008115159050919050565b610d3d81610d28565b82525050565b6000602082019050610d586000830184610d34565b92915050565b600069ffffffffffffffffffff82169050919050565b610d7d81610d5e565b82525050565b6000602082019050610d986000830184610d74565b92915050565b60008060408385031215610db557610db4610b95565b5b6000610dc385828601610bb6565b9250506020610dd485828601610bb6565b9150509250929050565b610de781610beb565b82525050565b6000602082019050610e026000830184610dde565b92915050565b6000819050919050565b610e23610e1e82610b61565b610e08565b82525050565b60008160601b9050919050565b6000610e4182610e29565b9050919050565b6000610e5382610e36565b9050919050565b610e6b610e6682610beb565b610e48565b82525050565b6000610e7d8286610e12565b602082019150610e8d8285610e5a565b601482019150610e9d8284610e12565b602082019150819050949350505050565b6000608082019050610ec36000830187610b6b565b610ed06020830186610b6b565b610edd6040830185610dde565b610eea6060830184610b6b565b95945050505050565b610efc81610d28565b8114610f0757600080fd5b50565b600081519050610f1981610ef3565b92915050565b600060208284031215610f3557610f34610b95565b5b6000610f4384828501610f0a565b91505092915050565b610f5581610d5e565b8114610f6057600080fd5b50565b600081519050610f7281610f4c565b92915050565b600081519050610f8781610b9f565b92915050565b610f9681610b2d565b8114610fa157600080fd5b50565b600081519050610fb381610f8d565b92915050565b600080600080600060a08688031215610fd557610fd4610b95565b5b6000610fe388828901610f63565b9550506020610ff488828901610f78565b945050604061100588828901610fa4565b935050606061101688828901610f78565b925050608061102788828901610f78565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061109d82610b61565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110cf576110ce611063565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060408201905061111e6000830185610dde565b61112b6020830184610b6b565b9392505050565b600082825260208201905092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b600061119f602a83611132565b91506111aa82611143565b604082019050919050565b600060208201905081810360008301526111ce81611192565b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b6000611231602683611132565b915061123c826111d5565b604082019050919050565b6000602082019050818103600083015261126081611224565b9050919050565b600081519050919050565b600081905092915050565b60005b8381101561129b578082015181840152602081019050611280565b60008484015250505050565b60006112b282611267565b6112bc8185611272565b93506112cc81856020860161127d565b80840191505092915050565b60006112e482846112a7565b915081905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b6000611325601d83611132565b9150611330826112ef565b602082019050919050565b6000602082019050818103600083015261135481611318565b9050919050565b600081519050919050565b6000601f19601f8301169050919050565b60006113828261135b565b61138c8185611132565b935061139c81856020860161127d565b6113a581611366565b840191505092915050565b600060208201905081810360008301526113ca8184611377565b90509291505056fea26469706673582212205ba18f8cc95fe34b4a35dc8acea2902d7de5cc22621455d90852e23b1ca2b83664736f6c63430008110033",
}

// MerkleDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleDistributorMetaData.ABI instead.
var MerkleDistributorABI = MerkleDistributorMetaData.ABI

// MerkleDistributorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleDistributorMetaData.Bin instead.
var MerkleDistributorBin = MerkleDistributorMetaData.Bin

// DeployMerkleDistributor deploys a new Ethereum contract, binding an instance of MerkleDistributor to it.
func DeployMerkleDistributor(auth *bind.TransactOpts, backend bind.ContractBackend, aggregatorProxy_ common.Address, token_ common.Address, projectId_ *big.Int) (common.Address, *types.Transaction, *MerkleDistributor, error) {
	parsed, err := MerkleDistributorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleDistributorBin), backend, aggregatorProxy_, token_, projectId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleDistributor{MerkleDistributorCaller: MerkleDistributorCaller{contract: contract}, MerkleDistributorTransactor: MerkleDistributorTransactor{contract: contract}, MerkleDistributorFilterer: MerkleDistributorFilterer{contract: contract}}, nil
}

// MerkleDistributor is an auto generated Go binding around an Ethereum contract.
type MerkleDistributor struct {
	MerkleDistributorCaller     // Read-only binding to the contract
	MerkleDistributorTransactor // Write-only binding to the contract
	MerkleDistributorFilterer   // Log filterer for contract events
}

// MerkleDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleDistributorSession struct {
	Contract     *MerkleDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MerkleDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleDistributorCallerSession struct {
	Contract *MerkleDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MerkleDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleDistributorTransactorSession struct {
	Contract     *MerkleDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MerkleDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleDistributorRaw struct {
	Contract *MerkleDistributor // Generic contract binding to access the raw methods on
}

// MerkleDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleDistributorCallerRaw struct {
	Contract *MerkleDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleDistributorTransactorRaw struct {
	Contract *MerkleDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleDistributor creates a new instance of MerkleDistributor, bound to a specific deployed contract.
func NewMerkleDistributor(address common.Address, backend bind.ContractBackend) (*MerkleDistributor, error) {
	contract, err := bindMerkleDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleDistributor{MerkleDistributorCaller: MerkleDistributorCaller{contract: contract}, MerkleDistributorTransactor: MerkleDistributorTransactor{contract: contract}, MerkleDistributorFilterer: MerkleDistributorFilterer{contract: contract}}, nil
}

// NewMerkleDistributorCaller creates a new read-only instance of MerkleDistributor, bound to a specific deployed contract.
func NewMerkleDistributorCaller(address common.Address, caller bind.ContractCaller) (*MerkleDistributorCaller, error) {
	contract, err := bindMerkleDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleDistributorCaller{contract: contract}, nil
}

// NewMerkleDistributorTransactor creates a new write-only instance of MerkleDistributor, bound to a specific deployed contract.
func NewMerkleDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleDistributorTransactor, error) {
	contract, err := bindMerkleDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleDistributorTransactor{contract: contract}, nil
}

// NewMerkleDistributorFilterer creates a new log filterer instance of MerkleDistributor, bound to a specific deployed contract.
func NewMerkleDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleDistributorFilterer, error) {
	contract, err := bindMerkleDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleDistributorFilterer{contract: contract}, nil
}

// bindMerkleDistributor binds a generic wrapper to an already deployed contract.
func bindMerkleDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleDistributor *MerkleDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleDistributor.Contract.MerkleDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleDistributor *MerkleDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleDistributor.Contract.MerkleDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleDistributor *MerkleDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleDistributor.Contract.MerkleDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleDistributor *MerkleDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleDistributor *MerkleDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleDistributor *MerkleDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleDistributor.Contract.contract.Transact(opts, method, params...)
}

// CurBatch is a free data retrieval call binding the contract method 0x733db88f.
//
// Solidity: function curBatch() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCaller) CurBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "curBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurBatch is a free data retrieval call binding the contract method 0x733db88f.
//
// Solidity: function curBatch() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorSession) CurBatch() (*big.Int, error) {
	return _MerkleDistributor.Contract.CurBatch(&_MerkleDistributor.CallOpts)
}

// CurBatch is a free data retrieval call binding the contract method 0x733db88f.
//
// Solidity: function curBatch() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCallerSession) CurBatch() (*big.Int, error) {
	return _MerkleDistributor.Contract.CurBatch(&_MerkleDistributor.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_MerkleDistributor *MerkleDistributorCaller) IsClaimed(opts *bind.CallOpts, batch *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "isClaimed", batch, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_MerkleDistributor *MerkleDistributorSession) IsClaimed(batch *big.Int, index *big.Int) (bool, error) {
	return _MerkleDistributor.Contract.IsClaimed(&_MerkleDistributor.CallOpts, batch, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_MerkleDistributor *MerkleDistributorCallerSession) IsClaimed(batch *big.Int, index *big.Int) (bool, error) {
	return _MerkleDistributor.Contract.IsClaimed(&_MerkleDistributor.CallOpts, batch, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_MerkleDistributor *MerkleDistributorCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_MerkleDistributor *MerkleDistributorSession) MerkleRoot() ([32]byte, error) {
	return _MerkleDistributor.Contract.MerkleRoot(&_MerkleDistributor.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_MerkleDistributor *MerkleDistributorCallerSession) MerkleRoot() ([32]byte, error) {
	return _MerkleDistributor.Contract.MerkleRoot(&_MerkleDistributor.CallOpts)
}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCaller) ProjectId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "projectId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorSession) ProjectId() (*big.Int, error) {
	return _MerkleDistributor.Contract.ProjectId(&_MerkleDistributor.CallOpts)
}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCallerSession) ProjectId() (*big.Int, error) {
	return _MerkleDistributor.Contract.ProjectId(&_MerkleDistributor.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_MerkleDistributor *MerkleDistributorCaller) RoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "roundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_MerkleDistributor *MerkleDistributorSession) RoundId() (*big.Int, error) {
	return _MerkleDistributor.Contract.RoundId(&_MerkleDistributor.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_MerkleDistributor *MerkleDistributorCallerSession) RoundId() (*big.Int, error) {
	return _MerkleDistributor.Contract.RoundId(&_MerkleDistributor.CallOpts)
}

// StartedAt is a free data retrieval call binding the contract method 0xf21f537d.
//
// Solidity: function startedAt() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCaller) StartedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "startedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartedAt is a free data retrieval call binding the contract method 0xf21f537d.
//
// Solidity: function startedAt() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorSession) StartedAt() (*big.Int, error) {
	return _MerkleDistributor.Contract.StartedAt(&_MerkleDistributor.CallOpts)
}

// StartedAt is a free data retrieval call binding the contract method 0xf21f537d.
//
// Solidity: function startedAt() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCallerSession) StartedAt() (*big.Int, error) {
	return _MerkleDistributor.Contract.StartedAt(&_MerkleDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MerkleDistributor *MerkleDistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MerkleDistributor *MerkleDistributorSession) Token() (common.Address, error) {
	return _MerkleDistributor.Contract.Token(&_MerkleDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MerkleDistributor *MerkleDistributorCallerSession) Token() (common.Address, error) {
	return _MerkleDistributor.Contract.Token(&_MerkleDistributor.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCaller) UpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleDistributor.contract.Call(opts, &out, "updatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorSession) UpdatedAt() (*big.Int, error) {
	return _MerkleDistributor.Contract.UpdatedAt(&_MerkleDistributor.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_MerkleDistributor *MerkleDistributorCallerSession) UpdatedAt() (*big.Int, error) {
	return _MerkleDistributor.Contract.UpdatedAt(&_MerkleDistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_MerkleDistributor *MerkleDistributorTransactor) Claim(opts *bind.TransactOpts, batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _MerkleDistributor.contract.Transact(opts, "claim", batch, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_MerkleDistributor *MerkleDistributorSession) Claim(batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _MerkleDistributor.Contract.Claim(&_MerkleDistributor.TransactOpts, batch, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_MerkleDistributor *MerkleDistributorTransactorSession) Claim(batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _MerkleDistributor.Contract.Claim(&_MerkleDistributor.TransactOpts, batch, index, account, amount, merkleProof)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0xb1de1626.
//
// Solidity: function updateRoot() returns(bool)
func (_MerkleDistributor *MerkleDistributorTransactor) UpdateRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleDistributor.contract.Transact(opts, "updateRoot")
}

// UpdateRoot is a paid mutator transaction binding the contract method 0xb1de1626.
//
// Solidity: function updateRoot() returns(bool)
func (_MerkleDistributor *MerkleDistributorSession) UpdateRoot() (*types.Transaction, error) {
	return _MerkleDistributor.Contract.UpdateRoot(&_MerkleDistributor.TransactOpts)
}

// UpdateRoot is a paid mutator transaction binding the contract method 0xb1de1626.
//
// Solidity: function updateRoot() returns(bool)
func (_MerkleDistributor *MerkleDistributorTransactorSession) UpdateRoot() (*types.Transaction, error) {
	return _MerkleDistributor.Contract.UpdateRoot(&_MerkleDistributor.TransactOpts)
}

// MerkleDistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the MerkleDistributor contract.
type MerkleDistributorClaimedIterator struct {
	Event *MerkleDistributorClaimed // Event containing the contract specifics and raw log

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
func (it *MerkleDistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleDistributorClaimed)
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
		it.Event = new(MerkleDistributorClaimed)
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
func (it *MerkleDistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleDistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleDistributorClaimed represents a Claimed event raised by the MerkleDistributor contract.
type MerkleDistributorClaimed struct {
	Batch   *big.Int
	Index   *big.Int
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_MerkleDistributor *MerkleDistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*MerkleDistributorClaimedIterator, error) {

	logs, sub, err := _MerkleDistributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &MerkleDistributorClaimedIterator{contract: _MerkleDistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_MerkleDistributor *MerkleDistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *MerkleDistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _MerkleDistributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleDistributorClaimed)
				if err := _MerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_MerkleDistributor *MerkleDistributorFilterer) ParseClaimed(log types.Log) (*MerkleDistributorClaimed, error) {
	event := new(MerkleDistributorClaimed)
	if err := _MerkleDistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleProofMetaData contains all meta data concerning the MerkleProof contract.
var MerkleProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122027a89ed8031c64db4b491ae520f3e532921ccad232f469abf2ace41bbcd91ab264736f6c63430008110033",
}

// MerkleProofABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleProofMetaData.ABI instead.
var MerkleProofABI = MerkleProofMetaData.ABI

// MerkleProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleProofMetaData.Bin instead.
var MerkleProofBin = MerkleProofMetaData.Bin

// DeployMerkleProof deploys a new Ethereum contract, binding an instance of MerkleProof to it.
func DeployMerkleProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleProof, error) {
	parsed, err := MerkleProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleProof{MerkleProofCaller: MerkleProofCaller{contract: contract}, MerkleProofTransactor: MerkleProofTransactor{contract: contract}, MerkleProofFilterer: MerkleProofFilterer{contract: contract}}, nil
}

// MerkleProof is an auto generated Go binding around an Ethereum contract.
type MerkleProof struct {
	MerkleProofCaller     // Read-only binding to the contract
	MerkleProofTransactor // Write-only binding to the contract
	MerkleProofFilterer   // Log filterer for contract events
}

// MerkleProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleProofSession struct {
	Contract     *MerkleProof      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleProofCallerSession struct {
	Contract *MerkleProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MerkleProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleProofTransactorSession struct {
	Contract     *MerkleProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MerkleProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleProofRaw struct {
	Contract *MerkleProof // Generic contract binding to access the raw methods on
}

// MerkleProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleProofCallerRaw struct {
	Contract *MerkleProofCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleProofTransactorRaw struct {
	Contract *MerkleProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleProof creates a new instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProof(address common.Address, backend bind.ContractBackend) (*MerkleProof, error) {
	contract, err := bindMerkleProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleProof{MerkleProofCaller: MerkleProofCaller{contract: contract}, MerkleProofTransactor: MerkleProofTransactor{contract: contract}, MerkleProofFilterer: MerkleProofFilterer{contract: contract}}, nil
}

// NewMerkleProofCaller creates a new read-only instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofCaller(address common.Address, caller bind.ContractCaller) (*MerkleProofCaller, error) {
	contract, err := bindMerkleProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleProofCaller{contract: contract}, nil
}

// NewMerkleProofTransactor creates a new write-only instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleProofTransactor, error) {
	contract, err := bindMerkleProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleProofTransactor{contract: contract}, nil
}

// NewMerkleProofFilterer creates a new log filterer instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleProofFilterer, error) {
	contract, err := bindMerkleProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleProofFilterer{contract: contract}, nil
}

// bindMerkleProof binds a generic wrapper to an already deployed contract.
func bindMerkleProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleProof *MerkleProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleProof.Contract.MerkleProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleProof *MerkleProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleProof.Contract.MerkleProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleProof *MerkleProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleProof.Contract.MerkleProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleProof *MerkleProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleProof *MerkleProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleProof *MerkleProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleProof.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204d0e33fda8b74337ad3126372d3734595cc17a7a61a918aa070b0aaa83bfe48c64736f6c63430008110033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}
