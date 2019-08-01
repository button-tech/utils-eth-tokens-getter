// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"getTokensBalancesByAddress\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `0x608060405234801561001057600080fd5b506105f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630d0f9df414610046578063c27be11f14610153578063c489744b14610280575b600080fd5b6100fc6004803603602081101561005c57600080fd5b810190808035906020019064010000000081111561007957600080fd5b82018360208201111561008b57600080fd5b803590602001918460208302840111640100000000831117156100ad57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506102f8565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561013f578082015181840152602081019050610124565b505050509050019250505060405180910390f35b6102296004803603604081101561016957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156101a657600080fd5b8201836020820111156101b857600080fd5b803590602001918460208302840111640100000000831117156101da57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506103a1565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561026c578082015181840152602081019050610251565b505050509050019250505060405180910390f35b6102e26004803603604081101561029657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104f4565b6040518082815260200191505060405180910390f35b606060018251101561030957600080fd5b81516040519080825280602002602001820160405280156103395781602001602082028038833980820191505090505b50905060008090505b825181101561039b5782818151811061035757fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163182828151811061038257fe5b6020026020010181815250508080600101915050610342565b50919050565b60606001825110156103b257600080fd5b81516040519080825280602002602001820160405280156103e25781602001602082028038833980820191505090505b50905060008090505b82518110156104ea57600083828151811061040257fe5b602002602001015190508073ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561048957600080fd5b505afa15801561049d573d6000803e3d6000fd5b505050506040513d60208110156104b357600080fd5b81019080805190602001909291905050508383815181106104d057fe5b6020026020010181815250505080806001019150506103eb565b5080905092915050565b60008060008390508073ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561057957600080fd5b505afa15801561058d573d6000803e3d6000fd5b505050506040513d60208110156105a357600080fd5b8101908080519060200190929190505050915081925050509291505056fea265627a7a72305820162dc2142020227b4e866e769d8ccb3cb21d6195aafb7e0451b24792d54bc1d764736f6c634300050a0032`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x0d0f9df4.
//
// Solidity: function getBalance(address[] _addresses) constant returns(uint256[] balances)
func (_Contract *ContractCaller) GetBalance(opts *bind.CallOpts, _addresses []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getBalance", _addresses)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x0d0f9df4.
//
// Solidity: function getBalance(address[] _addresses) constant returns(uint256[] balances)
func (_Contract *ContractSession) GetBalance(_addresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts, _addresses)
}

// GetBalance is a free data retrieval call binding the contract method 0x0d0f9df4.
//
// Solidity: function getBalance(address[] _addresses) constant returns(uint256[] balances)
func (_Contract *ContractCallerSession) GetBalance(_addresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts, _addresses)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _address, address _tokenAddress) constant returns(uint256 balances)
func (_Contract *ContractCaller) GetTokenBalance(opts *bind.CallOpts, _address common.Address, _tokenAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getTokenBalance", _address, _tokenAddress)
	return *ret0, err
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _address, address _tokenAddress) constant returns(uint256 balances)
func (_Contract *ContractSession) GetTokenBalance(_address common.Address, _tokenAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTokenBalance(&_Contract.CallOpts, _address, _tokenAddress)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0xc489744b.
//
// Solidity: function getTokenBalance(address _address, address _tokenAddress) constant returns(uint256 balances)
func (_Contract *ContractCallerSession) GetTokenBalance(_address common.Address, _tokenAddress common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTokenBalance(&_Contract.CallOpts, _address, _tokenAddress)
}

// GetTokensBalancesByAddress is a free data retrieval call binding the contract method 0xc27be11f.
//
// Solidity: function getTokensBalancesByAddress(address _address, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractCaller) GetTokensBalancesByAddress(opts *bind.CallOpts, _address common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getTokensBalancesByAddress", _address, _tokenAddresses)
	return *ret0, err
}

// GetTokensBalancesByAddress is a free data retrieval call binding the contract method 0xc27be11f.
//
// Solidity: function getTokensBalancesByAddress(address _address, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractSession) GetTokensBalancesByAddress(_address common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetTokensBalancesByAddress(&_Contract.CallOpts, _address, _tokenAddresses)
}

// GetTokensBalancesByAddress is a free data retrieval call binding the contract method 0xc27be11f.
//
// Solidity: function getTokensBalancesByAddress(address _address, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractCallerSession) GetTokensBalancesByAddress(_address common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetTokensBalancesByAddress(&_Contract.CallOpts, _address, _tokenAddresses)
}
