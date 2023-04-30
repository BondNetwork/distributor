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

// MerkledistributorMetaData contains all meta data concerning the Merkledistributor contract.
var MerkledistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregatorProxy_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"projectId_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundId\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200159d3803806200159d83398181016040528101906200003791906200016d565b826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600281905550505050620001c9565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000fa82620000cd565b9050919050565b6200010c81620000ed565b81146200011857600080fd5b50565b6000815190506200012c8162000101565b92915050565b6000819050919050565b620001478162000132565b81146200015357600080fd5b50565b60008151905062000167816200013c565b92915050565b600080600060608486031215620001895762000188620000c8565b5b600062000199868287016200011b565b9350506020620001ac868287016200011b565b9250506040620001bf8682870162000156565b9150509250925092565b6113c480620001d96000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80637519ab50116100715780637519ab50146101565780638cd221c914610174578063a4e2d63414610192578063f21f537d146101b0578063f364c90c146101ce578063fc0c546a146101fe576100a9565b80632eb4a7ab146100ae5780633fafa127146100cc5780635ca1e165146100ea5780635d4df3bf14610108578063733db88f14610138575b600080fd5b6100b661021c565b6040516100c39190610b31565b60405180910390f35b6100d4610222565b6040516100e19190610b65565b60405180910390f35b6100f2610228565b6040516100ff9190610b9b565b60405180910390f35b610122600480360381019061011d9190610caf565b61032f565b60405161012f9190610b9b565b60405180910390f35b6101406104a2565b60405161014d9190610b65565b60405180910390f35b61015e6104a8565b60405161016b9190610b65565b60405180910390f35b61017c6104ae565b6040516101899190610d6e565b60405180910390f35b61019a6104ca565b6040516101a79190610b9b565b60405180910390f35b6101b8610561565b6040516101c59190610b65565b60405180910390f35b6101e860048036038101906101e39190610d89565b610567565b6040516101f59190610b9b565b60405180910390f35b610206610658565b6040516102139190610dd8565b60405180910390f35b60035481565b60025481565b60006102326104ca565b15610240576000905061032c565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c00675d66040518163ffffffff1660e01b815260040160a060405180830381865afa1580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190610e60565b60056000600460006003600060066000600760008a919050558991905055889190505587919050558691906101000a81548169ffffffffffffffffffff021916908369ffffffffffffffffffff1602179055505050505050600190505b90565b60006103396104ca565b156103475760009050610498565b60045487146103595760009050610498565b61036560045487610567565b156103735760009050610498565b600086868660405160200161038a93929190610f44565b6040516020818303038152906040528051906020012090506103f0848480806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506003548361067e565b6103fe576000915050610498565b6104088787610695565b6104558686600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166107669092919063ffffffff16565b7fb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba8888888860405161048a9493929190610f81565b60405180910390a160019150505b9695505050505050565b60045481565b60075481565b600560009054906101000a900469ffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a4e2d6346040518163ffffffff1660e01b8152600401602060405180830381865afa158015610538573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061055c9190610ff2565b905090565b60065481565b6000826004541461057b5760019050610652565b60006008600085815260200190815260200160002060405180608001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff161515151581525050905082816020015114801561063d57506001151581606001511515145b1561064c576001915050610652565b60009150505b92915050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008261068b85846107ec565b1490509392505050565b604051806080016040528060045481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1681526020016001151581525060086000848152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff0219169083151502179055509050505050565b6107e78363a9059cbb60e01b848460405160240161078592919061101f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610842565b505050565b60008082905060005b8451811015610837576108228286838151811061081557610814611048565b5b6020026020010151610909565b9150808061082f906110a6565b9150506107f5565b508091505092915050565b60006108a4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166109349092919063ffffffff16565b905060008151111561090457808060200190518101906108c49190610ff2565b610903576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fa90611171565b60405180910390fd5b5b505050565b60008183106109215761091c828461094c565b61092c565b61092b838361094c565b5b905092915050565b60606109438484600085610963565b90509392505050565b600082600052816020526040600020905092915050565b6060824710156109a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099f90611203565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516109d19190611294565b60006040518083038185875af1925050503d8060008114610a0e576040519150601f19603f3d011682016040523d82523d6000602084013e610a13565b606091505b5091509150610a2487838387610a30565b92505050949350505050565b60608315610a92576000835103610a8a57610a4a85610aa5565b610a89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a80906112f7565b60405180910390fd5b5b829050610a9d565b610a9c8383610ac8565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600082511115610adb5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0f919061136c565b60405180910390fd5b6000819050919050565b610b2b81610b18565b82525050565b6000602082019050610b466000830184610b22565b92915050565b6000819050919050565b610b5f81610b4c565b82525050565b6000602082019050610b7a6000830184610b56565b92915050565b60008115159050919050565b610b9581610b80565b82525050565b6000602082019050610bb06000830184610b8c565b92915050565b600080fd5b600080fd5b610bc981610b4c565b8114610bd457600080fd5b50565b600081359050610be681610bc0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c1782610bec565b9050919050565b610c2781610c0c565b8114610c3257600080fd5b50565b600081359050610c4481610c1e565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610c6f57610c6e610c4a565b5b8235905067ffffffffffffffff811115610c8c57610c8b610c4f565b5b602083019150836020820283011115610ca857610ca7610c54565b5b9250929050565b60008060008060008060a08789031215610ccc57610ccb610bb6565b5b6000610cda89828a01610bd7565b9650506020610ceb89828a01610bd7565b9550506040610cfc89828a01610c35565b9450506060610d0d89828a01610bd7565b935050608087013567ffffffffffffffff811115610d2e57610d2d610bbb565b5b610d3a89828a01610c59565b92509250509295509295509295565b600069ffffffffffffffffffff82169050919050565b610d6881610d49565b82525050565b6000602082019050610d836000830184610d5f565b92915050565b60008060408385031215610da057610d9f610bb6565b5b6000610dae85828601610bd7565b9250506020610dbf85828601610bd7565b9150509250929050565b610dd281610c0c565b82525050565b6000602082019050610ded6000830184610dc9565b92915050565b610dfc81610d49565b8114610e0757600080fd5b50565b600081519050610e1981610df3565b92915050565b600081519050610e2e81610bc0565b92915050565b610e3d81610b18565b8114610e4857600080fd5b50565b600081519050610e5a81610e34565b92915050565b600080600080600060a08688031215610e7c57610e7b610bb6565b5b6000610e8a88828901610e0a565b9550506020610e9b88828901610e1f565b9450506040610eac88828901610e4b565b9350506060610ebd88828901610e1f565b9250506080610ece88828901610e1f565b9150509295509295909350565b6000819050919050565b610ef6610ef182610b4c565b610edb565b82525050565b60008160601b9050919050565b6000610f1482610efc565b9050919050565b6000610f2682610f09565b9050919050565b610f3e610f3982610c0c565b610f1b565b82525050565b6000610f508286610ee5565b602082019150610f608285610f2d565b601482019150610f708284610ee5565b602082019150819050949350505050565b6000608082019050610f966000830187610b56565b610fa36020830186610b56565b610fb06040830185610dc9565b610fbd6060830184610b56565b95945050505050565b610fcf81610b80565b8114610fda57600080fd5b50565b600081519050610fec81610fc6565b92915050565b60006020828403121561100857611007610bb6565b5b600061101684828501610fdd565b91505092915050565b60006040820190506110346000830185610dc9565b6110416020830184610b56565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006110b182610b4c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110e3576110e2611077565b5b600182019050919050565b600082825260208201905092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b600061115b602a836110ee565b9150611166826110ff565b604082019050919050565b6000602082019050818103600083015261118a8161114e565b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b60006111ed6026836110ee565b91506111f882611191565b604082019050919050565b6000602082019050818103600083015261121c816111e0565b9050919050565b600081519050919050565b600081905092915050565b60005b8381101561125757808201518184015260208101905061123c565b60008484015250505050565b600061126e82611223565b611278818561122e565b9350611288818560208601611239565b80840191505092915050565b60006112a08284611263565b915081905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b60006112e1601d836110ee565b91506112ec826112ab565b602082019050919050565b60006020820190508181036000830152611310816112d4565b9050919050565b600081519050919050565b6000601f19601f8301169050919050565b600061133e82611317565b61134881856110ee565b9350611358818560208601611239565b61136181611322565b840191505092915050565b600060208201905081810360008301526113868184611333565b90509291505056fea2646970667358221220665f1d942bd4c6a116ad1682cafc32843081c24eec73b2bfe738e85f50d797ff64736f6c63430008130033",
}

// MerkledistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkledistributorMetaData.ABI instead.
var MerkledistributorABI = MerkledistributorMetaData.ABI

// MerkledistributorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkledistributorMetaData.Bin instead.
var MerkledistributorBin = MerkledistributorMetaData.Bin

// DeployMerkledistributor deploys a new Ethereum contract, binding an instance of Merkledistributor to it.
func DeployMerkledistributor(auth *bind.TransactOpts, backend bind.ContractBackend, aggregatorProxy_ common.Address, token_ common.Address, projectId_ *big.Int) (common.Address, *types.Transaction, *Merkledistributor, error) {
	parsed, err := MerkledistributorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkledistributorBin), backend, aggregatorProxy_, token_, projectId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Merkledistributor{MerkledistributorCaller: MerkledistributorCaller{contract: contract}, MerkledistributorTransactor: MerkledistributorTransactor{contract: contract}, MerkledistributorFilterer: MerkledistributorFilterer{contract: contract}}, nil
}

// Merkledistributor is an auto generated Go binding around an Ethereum contract.
type Merkledistributor struct {
	MerkledistributorCaller     // Read-only binding to the contract
	MerkledistributorTransactor // Write-only binding to the contract
	MerkledistributorFilterer   // Log filterer for contract events
}

// MerkledistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkledistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkledistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkledistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkledistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkledistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkledistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkledistributorSession struct {
	Contract     *Merkledistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MerkledistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkledistributorCallerSession struct {
	Contract *MerkledistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MerkledistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkledistributorTransactorSession struct {
	Contract     *MerkledistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MerkledistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkledistributorRaw struct {
	Contract *Merkledistributor // Generic contract binding to access the raw methods on
}

// MerkledistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkledistributorCallerRaw struct {
	Contract *MerkledistributorCaller // Generic read-only contract binding to access the raw methods on
}

// MerkledistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkledistributorTransactorRaw struct {
	Contract *MerkledistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkledistributor creates a new instance of Merkledistributor, bound to a specific deployed contract.
func NewMerkledistributor(address common.Address, backend bind.ContractBackend) (*Merkledistributor, error) {
	contract, err := bindMerkledistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkledistributor{MerkledistributorCaller: MerkledistributorCaller{contract: contract}, MerkledistributorTransactor: MerkledistributorTransactor{contract: contract}, MerkledistributorFilterer: MerkledistributorFilterer{contract: contract}}, nil
}

// NewMerkledistributorCaller creates a new read-only instance of Merkledistributor, bound to a specific deployed contract.
func NewMerkledistributorCaller(address common.Address, caller bind.ContractCaller) (*MerkledistributorCaller, error) {
	contract, err := bindMerkledistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkledistributorCaller{contract: contract}, nil
}

// NewMerkledistributorTransactor creates a new write-only instance of Merkledistributor, bound to a specific deployed contract.
func NewMerkledistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkledistributorTransactor, error) {
	contract, err := bindMerkledistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkledistributorTransactor{contract: contract}, nil
}

// NewMerkledistributorFilterer creates a new log filterer instance of Merkledistributor, bound to a specific deployed contract.
func NewMerkledistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkledistributorFilterer, error) {
	contract, err := bindMerkledistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkledistributorFilterer{contract: contract}, nil
}

// bindMerkledistributor binds a generic wrapper to an already deployed contract.
func bindMerkledistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkledistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkledistributor *MerkledistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merkledistributor.Contract.MerkledistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkledistributor *MerkledistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkledistributor.Contract.MerkledistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkledistributor *MerkledistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkledistributor.Contract.MerkledistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkledistributor *MerkledistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Merkledistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkledistributor *MerkledistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkledistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkledistributor *MerkledistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkledistributor.Contract.contract.Transact(opts, method, params...)
}

// CurBatch is a free data retrieval call binding the contract method 0x733db88f.
//
// Solidity: function curBatch() view returns(uint256)
func (_Merkledistributor *MerkledistributorCaller) CurBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "curBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurBatch is a free data retrieval call binding the contract method 0x733db88f.
//
// Solidity: function curBatch() view returns(uint256)
func (_Merkledistributor *MerkledistributorSession) CurBatch() (*big.Int, error) {
	return _Merkledistributor.Contract.CurBatch(&_Merkledistributor.CallOpts)
}

// CurBatch is a free data retrieval call binding the contract method 0x733db88f.
//
// Solidity: function curBatch() view returns(uint256)
func (_Merkledistributor *MerkledistributorCallerSession) CurBatch() (*big.Int, error) {
	return _Merkledistributor.Contract.CurBatch(&_Merkledistributor.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_Merkledistributor *MerkledistributorCaller) IsClaimed(opts *bind.CallOpts, batch *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "isClaimed", batch, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_Merkledistributor *MerkledistributorSession) IsClaimed(batch *big.Int, index *big.Int) (bool, error) {
	return _Merkledistributor.Contract.IsClaimed(&_Merkledistributor.CallOpts, batch, index)
}

// IsClaimed is a free data retrieval call binding the contract method 0xf364c90c.
//
// Solidity: function isClaimed(uint256 batch, uint256 index) view returns(bool)
func (_Merkledistributor *MerkledistributorCallerSession) IsClaimed(batch *big.Int, index *big.Int) (bool, error) {
	return _Merkledistributor.Contract.IsClaimed(&_Merkledistributor.CallOpts, batch, index)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Merkledistributor *MerkledistributorCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Merkledistributor *MerkledistributorSession) MerkleRoot() ([32]byte, error) {
	return _Merkledistributor.Contract.MerkleRoot(&_Merkledistributor.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Merkledistributor *MerkledistributorCallerSession) MerkleRoot() ([32]byte, error) {
	return _Merkledistributor.Contract.MerkleRoot(&_Merkledistributor.CallOpts)
}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_Merkledistributor *MerkledistributorCaller) ProjectId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "projectId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_Merkledistributor *MerkledistributorSession) ProjectId() (*big.Int, error) {
	return _Merkledistributor.Contract.ProjectId(&_Merkledistributor.CallOpts)
}

// ProjectId is a free data retrieval call binding the contract method 0x3fafa127.
//
// Solidity: function projectId() view returns(uint256)
func (_Merkledistributor *MerkledistributorCallerSession) ProjectId() (*big.Int, error) {
	return _Merkledistributor.Contract.ProjectId(&_Merkledistributor.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_Merkledistributor *MerkledistributorCaller) RoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "roundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_Merkledistributor *MerkledistributorSession) RoundId() (*big.Int, error) {
	return _Merkledistributor.Contract.RoundId(&_Merkledistributor.CallOpts)
}

// RoundId is a free data retrieval call binding the contract method 0x8cd221c9.
//
// Solidity: function roundId() view returns(uint80)
func (_Merkledistributor *MerkledistributorCallerSession) RoundId() (*big.Int, error) {
	return _Merkledistributor.Contract.RoundId(&_Merkledistributor.CallOpts)
}

// StartedAt is a free data retrieval call binding the contract method 0xf21f537d.
//
// Solidity: function startedAt() view returns(uint256)
func (_Merkledistributor *MerkledistributorCaller) StartedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "startedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartedAt is a free data retrieval call binding the contract method 0xf21f537d.
//
// Solidity: function startedAt() view returns(uint256)
func (_Merkledistributor *MerkledistributorSession) StartedAt() (*big.Int, error) {
	return _Merkledistributor.Contract.StartedAt(&_Merkledistributor.CallOpts)
}

// StartedAt is a free data retrieval call binding the contract method 0xf21f537d.
//
// Solidity: function startedAt() view returns(uint256)
func (_Merkledistributor *MerkledistributorCallerSession) StartedAt() (*big.Int, error) {
	return _Merkledistributor.Contract.StartedAt(&_Merkledistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Merkledistributor *MerkledistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Merkledistributor *MerkledistributorSession) Token() (common.Address, error) {
	return _Merkledistributor.Contract.Token(&_Merkledistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Merkledistributor *MerkledistributorCallerSession) Token() (common.Address, error) {
	return _Merkledistributor.Contract.Token(&_Merkledistributor.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_Merkledistributor *MerkledistributorCaller) UpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Merkledistributor.contract.Call(opts, &out, "updatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_Merkledistributor *MerkledistributorSession) UpdatedAt() (*big.Int, error) {
	return _Merkledistributor.Contract.UpdatedAt(&_Merkledistributor.CallOpts)
}

// UpdatedAt is a free data retrieval call binding the contract method 0x7519ab50.
//
// Solidity: function updatedAt() view returns(uint256)
func (_Merkledistributor *MerkledistributorCallerSession) UpdatedAt() (*big.Int, error) {
	return _Merkledistributor.Contract.UpdatedAt(&_Merkledistributor.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_Merkledistributor *MerkledistributorTransactor) Claim(opts *bind.TransactOpts, batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Merkledistributor.contract.Transact(opts, "claim", batch, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_Merkledistributor *MerkledistributorSession) Claim(batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Merkledistributor.Contract.Claim(&_Merkledistributor.TransactOpts, batch, index, account, amount, merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0x5d4df3bf.
//
// Solidity: function claim(uint256 batch, uint256 index, address account, uint256 amount, bytes32[] merkleProof) returns(bool)
func (_Merkledistributor *MerkledistributorTransactorSession) Claim(batch *big.Int, index *big.Int, account common.Address, amount *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Merkledistributor.Contract.Claim(&_Merkledistributor.TransactOpts, batch, index, account, amount, merkleProof)
}

// GetRoot is a paid mutator transaction binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() returns(bool)
func (_Merkledistributor *MerkledistributorTransactor) GetRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkledistributor.contract.Transact(opts, "getRoot")
}

// GetRoot is a paid mutator transaction binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() returns(bool)
func (_Merkledistributor *MerkledistributorSession) GetRoot() (*types.Transaction, error) {
	return _Merkledistributor.Contract.GetRoot(&_Merkledistributor.TransactOpts)
}

// GetRoot is a paid mutator transaction binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() returns(bool)
func (_Merkledistributor *MerkledistributorTransactorSession) GetRoot() (*types.Transaction, error) {
	return _Merkledistributor.Contract.GetRoot(&_Merkledistributor.TransactOpts)
}

// IsLocked is a paid mutator transaction binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() returns(bool)
func (_Merkledistributor *MerkledistributorTransactor) IsLocked(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkledistributor.contract.Transact(opts, "isLocked")
}

// IsLocked is a paid mutator transaction binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() returns(bool)
func (_Merkledistributor *MerkledistributorSession) IsLocked() (*types.Transaction, error) {
	return _Merkledistributor.Contract.IsLocked(&_Merkledistributor.TransactOpts)
}

// IsLocked is a paid mutator transaction binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() returns(bool)
func (_Merkledistributor *MerkledistributorTransactorSession) IsLocked() (*types.Transaction, error) {
	return _Merkledistributor.Contract.IsLocked(&_Merkledistributor.TransactOpts)
}

// MerkledistributorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Merkledistributor contract.
type MerkledistributorClaimedIterator struct {
	Event *MerkledistributorClaimed // Event containing the contract specifics and raw log

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
func (it *MerkledistributorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkledistributorClaimed)
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
		it.Event = new(MerkledistributorClaimed)
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
func (it *MerkledistributorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkledistributorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkledistributorClaimed represents a Claimed event raised by the Merkledistributor contract.
type MerkledistributorClaimed struct {
	Batch   *big.Int
	Index   *big.Int
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_Merkledistributor *MerkledistributorFilterer) FilterClaimed(opts *bind.FilterOpts) (*MerkledistributorClaimedIterator, error) {

	logs, sub, err := _Merkledistributor.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &MerkledistributorClaimedIterator{contract: _Merkledistributor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xb94bf7f9302edf52a596286915a69b4b0685574cffdedd0712e3c62f2550f0ba.
//
// Solidity: event Claimed(uint256 batch, uint256 index, address account, uint256 amount)
func (_Merkledistributor *MerkledistributorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *MerkledistributorClaimed) (event.Subscription, error) {

	logs, sub, err := _Merkledistributor.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkledistributorClaimed)
				if err := _Merkledistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
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
func (_Merkledistributor *MerkledistributorFilterer) ParseClaimed(log types.Log) (*MerkledistributorClaimed, error) {
	event := new(MerkledistributorClaimed)
	if err := _Merkledistributor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
