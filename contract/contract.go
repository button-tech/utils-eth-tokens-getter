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
const ContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"},{\"name\":\"_tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_tokenAddresses\",\"type\":\"address[]\"}],\"name\":\"getTokenBalanceByAddress\",\"outputs\":[{\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `0x608060405234801561001057600080fd5b5061087f806100206000396000f3fe608060405260043610610051576000357c0100000000000000000000000000000000000000000000000000000000900480630d0f9df4146100565780633f078cb714610170578063d883c56e1461031e575b600080fd5b34801561006257600080fd5b506101196004803603602081101561007957600080fd5b810190808035906020019064010000000081111561009657600080fd5b8201836020820111156100a857600080fd5b803590602001918460208302840111640100000000831117156100ca57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610458565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561015c578082015181840152602081019050610141565b505050509050019250505060405180910390f35b34801561017c57600080fd5b506102c76004803603604081101561019357600080fd5b81019080803590602001906401000000008111156101b057600080fd5b8201836020820111156101c257600080fd5b803590602001918460208302840111640100000000831117156101e457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561024457600080fd5b82018360208201111561025657600080fd5b8035906020019184602083028401116401000000008311171561027857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505061050b565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561030a5780820151818401526020810190506102ef565b505050509050019250505060405180910390f35b34801561032a57600080fd5b506104016004803603604081101561034157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561037e57600080fd5b82018360208201111561039057600080fd5b803590602001918460208302840111640100000000831117156103b257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506106da565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610444578082015181840152602081019050610429565b505050509050019250505060405180910390f35b6060600182511015151561046b57600080fd5b815160405190808252806020026020018201604052801561049b5781602001602082028038833980820191505090505b50905060008090505b82518110156105055782818151811015156104bb57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff163182828151811015156104ea57fe5b906020019060200201818152505080806001019150506104a4565b50919050565b6060600182511015151561051e57600080fd5b600183511015151561052f57600080fd5b600080905082518451026040519080825280602002602001820160405280156105675781602001602082028038833980820191505090505b50915060008090505b83518110156106cf5760008090505b85518110156106c1576000858381518110151561059857fe5b9060200190602002015190508073ffffffffffffffffffffffffffffffffffffffff166370a0823188848151811015156105ce57fe5b906020019060200201516040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561065457600080fd5b505afa158015610668573d6000803e3d6000fd5b505050506040513d602081101561067e57600080fd5b8101908080519060200190929190505050858581518110151561069d57fe5b9060200190602002018181525050838060010194505050808060010191505061057f565b508080600101915050610570565b508191505092915050565b606060018251101515156106ed57600080fd5b815160405190808252806020026020018201604052801561071d5781602001602082028038833980820191505090505b50905060008090505b8251811015610849576000838281518110151561073f57fe5b9060200190602002015190508073ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156107e457600080fd5b505afa1580156107f8573d6000803e3d6000fd5b505050506040513d602081101561080e57600080fd5b8101908080519060200190929190505050838381518110151561082d57fe5b9060200190602002018181525050508080600101915050610726565b508090509291505056fea165627a7a72305820537f0184ad8a5c630e70f542f10073cc3c72e0f698319b18f2de1d0da459c1b50029`

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

// GetTokenBalance is a free data retrieval call binding the contract method 0x3f078cb7.
//
// Solidity: function getTokenBalance(address[] _addresses, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractCaller) GetTokenBalance(opts *bind.CallOpts, _addresses []common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getTokenBalance", _addresses, _tokenAddresses)
	return *ret0, err
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3f078cb7.
//
// Solidity: function getTokenBalance(address[] _addresses, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractSession) GetTokenBalance(_addresses []common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetTokenBalance(&_Contract.CallOpts, _addresses, _tokenAddresses)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x3f078cb7.
//
// Solidity: function getTokenBalance(address[] _addresses, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractCallerSession) GetTokenBalance(_addresses []common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetTokenBalance(&_Contract.CallOpts, _addresses, _tokenAddresses)
}

// GetTokenBalanceByAddress is a free data retrieval call binding the contract method 0xd883c56e.
//
// Solidity: function getTokenBalanceByAddress(address _address, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractCaller) GetTokenBalanceByAddress(opts *bind.CallOpts, _address common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getTokenBalanceByAddress", _address, _tokenAddresses)
	return *ret0, err
}

// GetTokenBalanceByAddress is a free data retrieval call binding the contract method 0xd883c56e.
//
// Solidity: function getTokenBalanceByAddress(address _address, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractSession) GetTokenBalanceByAddress(_address common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetTokenBalanceByAddress(&_Contract.CallOpts, _address, _tokenAddresses)
}

// GetTokenBalanceByAddress is a free data retrieval call binding the contract method 0xd883c56e.
//
// Solidity: function getTokenBalanceByAddress(address _address, address[] _tokenAddresses) constant returns(uint256[] balances)
func (_Contract *ContractCallerSession) GetTokenBalanceByAddress(_address common.Address, _tokenAddresses []common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetTokenBalanceByAddress(&_Contract.CallOpts, _address, _tokenAddresses)
}
