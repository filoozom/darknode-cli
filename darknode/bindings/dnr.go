// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DarknodeRegistryABI is the input ABI used to generate the binding from.
const DarknodeRegistryABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_VERSION\",\"type\":\"string\"},{\"internalType\":\"contractRenToken\",\"name\":\"_renAddress\",\"type\":\"address\"},{\"internalType\":\"contractDarknodeRegistryStore\",\"name\":\"_storeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumPodSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumEpochIntervalSeconds\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"LogDarknodeDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"LogDarknodeOwnerRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIDarknodePayment\",\"name\":\"_previousDarknodePayment\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIDarknodePayment\",\"name\":\"_nextDarknodePayment\",\"type\":\"address\"}],\"name\":\"LogDarknodePaymentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bond\",\"type\":\"uint256\"}],\"name\":\"LogDarknodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"LogDarknodeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousMinimumBond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextMinimumBond\",\"type\":\"uint256\"}],\"name\":\"LogMinimumBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousMinimumEpochInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextMinimumEpochInterval\",\"type\":\"uint256\"}],\"name\":\"LogMinimumEpochIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousMinimumPodSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextMinimumPodSize\",\"type\":\"uint256\"}],\"name\":\"LogMinimumPodSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochhash\",\"type\":\"uint256\"}],\"name\":\"LogNewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_previousSlasher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nextSlasher\",\"type\":\"address\"}],\"name\":\"LogSlasherUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimStoreOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochhash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"darknodePayment\",\"outputs\":[{\"internalType\":\"contractIDarknodePayment\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"deregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"getDarknodeBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"getDarknodeOwner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"getDarknodePublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getDarknodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getPreviousDarknodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isDeregisterable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isDeregistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isPendingDeregistration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isPendingRegistration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRefundable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRefunded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"isRegisteredInPreviousEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumEpochInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumPodSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumEpochInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextMinimumPodSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextSlasher\",\"outputs\":[{\"internalType\":\"contractIDarknodeSlasher\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodesNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numDarknodesPreviousEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochhash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocktime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_darknodeID\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ren\",\"outputs\":[{\"internalType\":\"contractRenToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guilty\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractIDarknodeSlasher\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"internalType\":\"contractDarknodeRegistryStore\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDarknodeRegistry\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferStoreOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIDarknodePayment\",\"name\":\"_darknodePayment\",\"type\":\"address\"}],\"name\":\"updateDarknodePayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextMinimumBond\",\"type\":\"uint256\"}],\"name\":\"updateMinimumBond\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextMinimumEpochInterval\",\"type\":\"uint256\"}],\"name\":\"updateMinimumEpochInterval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextMinimumPodSize\",\"type\":\"uint256\"}],\"name\":\"updateMinimumPodSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIDarknodeSlasher\",\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"updateSlasher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DarknodeRegistry is an auto generated Go binding around an Ethereum contract.
type DarknodeRegistry struct {
	DarknodeRegistryCaller     // Read-only binding to the contract
	DarknodeRegistryTransactor // Write-only binding to the contract
	DarknodeRegistryFilterer   // Log filterer for contract events
}

// DarknodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DarknodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DarknodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DarknodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DarknodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DarknodeRegistrySession struct {
	Contract     *DarknodeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DarknodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DarknodeRegistryCallerSession struct {
	Contract *DarknodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DarknodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DarknodeRegistryTransactorSession struct {
	Contract     *DarknodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DarknodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DarknodeRegistryRaw struct {
	Contract *DarknodeRegistry // Generic contract binding to access the raw methods on
}

// DarknodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DarknodeRegistryCallerRaw struct {
	Contract *DarknodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DarknodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DarknodeRegistryTransactorRaw struct {
	Contract *DarknodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDarknodeRegistry creates a new instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistry(address common.Address, backend bind.ContractBackend) (*DarknodeRegistry, error) {
	contract, err := bindDarknodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistry{DarknodeRegistryCaller: DarknodeRegistryCaller{contract: contract}, DarknodeRegistryTransactor: DarknodeRegistryTransactor{contract: contract}, DarknodeRegistryFilterer: DarknodeRegistryFilterer{contract: contract}}, nil
}

// NewDarknodeRegistryCaller creates a new read-only instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*DarknodeRegistryCaller, error) {
	contract, err := bindDarknodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryCaller{contract: contract}, nil
}

// NewDarknodeRegistryTransactor creates a new write-only instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DarknodeRegistryTransactor, error) {
	contract, err := bindDarknodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryTransactor{contract: contract}, nil
}

// NewDarknodeRegistryFilterer creates a new log filterer instance of DarknodeRegistry, bound to a specific deployed contract.
func NewDarknodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DarknodeRegistryFilterer, error) {
	contract, err := bindDarknodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryFilterer{contract: contract}, nil
}

// bindDarknodeRegistry binds a generic wrapper to an already deployed contract.
func bindDarknodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DarknodeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistry *DarknodeRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistry.Contract.DarknodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistry *DarknodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.DarknodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistry *DarknodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.DarknodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DarknodeRegistry *DarknodeRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DarknodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DarknodeRegistry *DarknodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DarknodeRegistry *DarknodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_DarknodeRegistry *DarknodeRegistryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_DarknodeRegistry *DarknodeRegistrySession) VERSION() (string, error) {
	return _DarknodeRegistry.Contract.VERSION(&_DarknodeRegistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) VERSION() (string, error) {
	return _DarknodeRegistry.Contract.VERSION(&_DarknodeRegistry.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocktime uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) CurrentEpoch(opts *bind.CallOpts) (struct {
	Epochhash *big.Int
	Blocktime *big.Int
}, error) {
	ret := new(struct {
		Epochhash *big.Int
		Blocktime *big.Int
	})
	out := ret
	err := _DarknodeRegistry.contract.Call(opts, out, "currentEpoch")
	return *ret, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocktime uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) CurrentEpoch() (struct {
	Epochhash *big.Int
	Blocktime *big.Int
}, error) {
	return _DarknodeRegistry.Contract.CurrentEpoch(&_DarknodeRegistry.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() constant returns(epochhash uint256, blocktime uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) CurrentEpoch() (struct {
	Epochhash *big.Int
	Blocktime *big.Int
}, error) {
	return _DarknodeRegistry.Contract.CurrentEpoch(&_DarknodeRegistry.CallOpts)
}

// DarknodePayment is a free data retrieval call binding the contract method 0xb6b34c67.
//
// Solidity: function darknodePayment() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) DarknodePayment(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "darknodePayment")
	return *ret0, err
}

// DarknodePayment is a free data retrieval call binding the contract method 0xb6b34c67.
//
// Solidity: function darknodePayment() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) DarknodePayment() (common.Address, error) {
	return _DarknodeRegistry.Contract.DarknodePayment(&_DarknodeRegistry.CallOpts)
}

// DarknodePayment is a free data retrieval call binding the contract method 0xb6b34c67.
//
// Solidity: function darknodePayment() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) DarknodePayment() (common.Address, error) {
	return _DarknodeRegistry.Contract.DarknodePayment(&_DarknodeRegistry.CallOpts)
}

// GetDarknodeBond is a free data retrieval call binding the contract method 0xba0f5b20.
//
// Solidity: function getDarknodeBond(_darknodeID address) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodeBond(opts *bind.CallOpts, _darknodeID common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodeBond", _darknodeID)
	return *ret0, err
}

// GetDarknodeBond is a free data retrieval call binding the contract method 0xba0f5b20.
//
// Solidity: function getDarknodeBond(_darknodeID address) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodeBond(_darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistry.Contract.GetDarknodeBond(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeBond is a free data retrieval call binding the contract method 0xba0f5b20.
//
// Solidity: function getDarknodeBond(_darknodeID address) constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodeBond(_darknodeID common.Address) (*big.Int, error) {
	return _DarknodeRegistry.Contract.GetDarknodeBond(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x1cedf8a3.
//
// Solidity: function getDarknodeOwner(_darknodeID address) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodeOwner(opts *bind.CallOpts, _darknodeID common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodeOwner", _darknodeID)
	return *ret0, err
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x1cedf8a3.
//
// Solidity: function getDarknodeOwner(_darknodeID address) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodeOwner(_darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodeOwner(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodeOwner is a free data retrieval call binding the contract method 0x1cedf8a3.
//
// Solidity: function getDarknodeOwner(_darknodeID address) constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodeOwner(_darknodeID common.Address) (common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodeOwner(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodePublicKey is a free data retrieval call binding the contract method 0x84d2688c.
//
// Solidity: function getDarknodePublicKey(_darknodeID address) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodePublicKey(opts *bind.CallOpts, _darknodeID common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodePublicKey", _darknodeID)
	return *ret0, err
}

// GetDarknodePublicKey is a free data retrieval call binding the contract method 0x84d2688c.
//
// Solidity: function getDarknodePublicKey(_darknodeID address) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodePublicKey(_darknodeID common.Address) ([]byte, error) {
	return _DarknodeRegistry.Contract.GetDarknodePublicKey(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodePublicKey is a free data retrieval call binding the contract method 0x84d2688c.
//
// Solidity: function getDarknodePublicKey(_darknodeID address) constant returns(bytes)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodePublicKey(_darknodeID common.Address) ([]byte, error) {
	return _DarknodeRegistry.Contract.GetDarknodePublicKey(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// GetDarknodes is a free data retrieval call binding the contract method 0xec5325c1.
//
// Solidity: function getDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCaller) GetDarknodes(opts *bind.CallOpts, _start common.Address, _count *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getDarknodes", _start, _count)
	return *ret0, err
}

// GetDarknodes is a free data retrieval call binding the contract method 0xec5325c1.
//
// Solidity: function getDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistrySession) GetDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// GetDarknodes is a free data retrieval call binding the contract method 0xec5325c1.
//
// Solidity: function getDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// GetPreviousDarknodes is a free data retrieval call binding the contract method 0x43846074.
//
// Solidity: function getPreviousDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCaller) GetPreviousDarknodes(opts *bind.CallOpts, _start common.Address, _count *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "getPreviousDarknodes", _start, _count)
	return *ret0, err
}

// GetPreviousDarknodes is a free data retrieval call binding the contract method 0x43846074.
//
// Solidity: function getPreviousDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistrySession) GetPreviousDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetPreviousDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// GetPreviousDarknodes is a free data retrieval call binding the contract method 0x43846074.
//
// Solidity: function getPreviousDarknodes(_start address, _count uint256) constant returns(address[])
func (_DarknodeRegistry *DarknodeRegistryCallerSession) GetPreviousDarknodes(_start common.Address, _count *big.Int) ([]common.Address, error) {
	return _DarknodeRegistry.Contract.GetPreviousDarknodes(&_DarknodeRegistry.CallOpts, _start, _count)
}

// IsDeregisterable is a free data retrieval call binding the contract method 0xe1878925.
//
// Solidity: function isDeregisterable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsDeregisterable(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isDeregisterable", _darknodeID)
	return *ret0, err
}

// IsDeregisterable is a free data retrieval call binding the contract method 0xe1878925.
//
// Solidity: function isDeregisterable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsDeregisterable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregisterable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregisterable is a free data retrieval call binding the contract method 0xe1878925.
//
// Solidity: function isDeregisterable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsDeregisterable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregisterable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregistered is a free data retrieval call binding the contract method 0x8020fc1f.
//
// Solidity: function isDeregistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsDeregistered(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isDeregistered", _darknodeID)
	return *ret0, err
}

// IsDeregistered is a free data retrieval call binding the contract method 0x8020fc1f.
//
// Solidity: function isDeregistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsDeregistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsDeregistered is a free data retrieval call binding the contract method 0x8020fc1f.
//
// Solidity: function isDeregistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsDeregistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsDeregistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsOwner() (bool, error) {
	return _DarknodeRegistry.Contract.IsOwner(&_DarknodeRegistry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsOwner() (bool, error) {
	return _DarknodeRegistry.Contract.IsOwner(&_DarknodeRegistry.CallOpts)
}

// IsPendingDeregistration is a free data retrieval call binding the contract method 0x303ee989.
//
// Solidity: function isPendingDeregistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsPendingDeregistration(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isPendingDeregistration", _darknodeID)
	return *ret0, err
}

// IsPendingDeregistration is a free data retrieval call binding the contract method 0x303ee989.
//
// Solidity: function isPendingDeregistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsPendingDeregistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingDeregistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingDeregistration is a free data retrieval call binding the contract method 0x303ee989.
//
// Solidity: function isPendingDeregistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsPendingDeregistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingDeregistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingRegistration is a free data retrieval call binding the contract method 0x040fa051.
//
// Solidity: function isPendingRegistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsPendingRegistration(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isPendingRegistration", _darknodeID)
	return *ret0, err
}

// IsPendingRegistration is a free data retrieval call binding the contract method 0x040fa051.
//
// Solidity: function isPendingRegistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsPendingRegistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingRegistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsPendingRegistration is a free data retrieval call binding the contract method 0x040fa051.
//
// Solidity: function isPendingRegistration(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsPendingRegistration(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsPendingRegistration(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefundable is a free data retrieval call binding the contract method 0x5aebd1cb.
//
// Solidity: function isRefundable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRefundable(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRefundable", _darknodeID)
	return *ret0, err
}

// IsRefundable is a free data retrieval call binding the contract method 0x5aebd1cb.
//
// Solidity: function isRefundable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRefundable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefundable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefundable is a free data retrieval call binding the contract method 0x5aebd1cb.
//
// Solidity: function isRefundable(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRefundable(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefundable(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRefunded(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRefunded", _darknodeID)
	return *ret0, err
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRefunded(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefunded(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRefunded(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRefunded(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRegistered(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRegistered", _darknodeID)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRegistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRegistered(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegistered(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegisteredInPreviousEpoch is a free data retrieval call binding the contract method 0x7be266da.
//
// Solidity: function isRegisteredInPreviousEpoch(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCaller) IsRegisteredInPreviousEpoch(opts *bind.CallOpts, _darknodeID common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "isRegisteredInPreviousEpoch", _darknodeID)
	return *ret0, err
}

// IsRegisteredInPreviousEpoch is a free data retrieval call binding the contract method 0x7be266da.
//
// Solidity: function isRegisteredInPreviousEpoch(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistrySession) IsRegisteredInPreviousEpoch(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegisteredInPreviousEpoch(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// IsRegisteredInPreviousEpoch is a free data retrieval call binding the contract method 0x7be266da.
//
// Solidity: function isRegisteredInPreviousEpoch(_darknodeID address) constant returns(bool)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) IsRegisteredInPreviousEpoch(_darknodeID common.Address) (bool, error) {
	return _DarknodeRegistry.Contract.IsRegisteredInPreviousEpoch(&_DarknodeRegistry.CallOpts, _darknodeID)
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumBond")
	return *ret0, err
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumBond(&_DarknodeRegistry.CallOpts)
}

// MinimumBond is a free data retrieval call binding the contract method 0xaa7517e1.
//
// Solidity: function minimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumBond(&_DarknodeRegistry.CallOpts)
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumEpochInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumEpochInterval")
	return *ret0, err
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// MinimumEpochInterval is a free data retrieval call binding the contract method 0x55cacda5.
//
// Solidity: function minimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) MinimumPodSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "minimumPodSize")
	return *ret0, err
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) MinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// MinimumPodSize is a free data retrieval call binding the contract method 0xc7dbc2be.
//
// Solidity: function minimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) MinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.MinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumBond")
	return *ret0, err
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumBond(&_DarknodeRegistry.CallOpts)
}

// NextMinimumBond is a free data retrieval call binding the contract method 0x60a22fe4.
//
// Solidity: function nextMinimumBond() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumBond() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumBond(&_DarknodeRegistry.CallOpts)
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumEpochInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumEpochInterval")
	return *ret0, err
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// NextMinimumEpochInterval is a free data retrieval call binding the contract method 0x455dc46d.
//
// Solidity: function nextMinimumEpochInterval() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumEpochInterval() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumEpochInterval(&_DarknodeRegistry.CallOpts)
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextMinimumPodSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextMinimumPodSize")
	return *ret0, err
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NextMinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextMinimumPodSize is a free data retrieval call binding the contract method 0x702c25ee.
//
// Solidity: function nextMinimumPodSize() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextMinimumPodSize() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NextMinimumPodSize(&_DarknodeRegistry.CallOpts)
}

// NextSlasher is a free data retrieval call binding the contract method 0x21a2ad3a.
//
// Solidity: function nextSlasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) NextSlasher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "nextSlasher")
	return *ret0, err
}

// NextSlasher is a free data retrieval call binding the contract method 0x21a2ad3a.
//
// Solidity: function nextSlasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) NextSlasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.NextSlasher(&_DarknodeRegistry.CallOpts)
}

// NextSlasher is a free data retrieval call binding the contract method 0x21a2ad3a.
//
// Solidity: function nextSlasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NextSlasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.NextSlasher(&_DarknodeRegistry.CallOpts)
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodes")
	return *ret0, err
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodes() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodes(&_DarknodeRegistry.CallOpts)
}

// NumDarknodes is a free data retrieval call binding the contract method 0x1460e603.
//
// Solidity: function numDarknodes() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodes() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodes(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodesNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodesNextEpoch")
	return *ret0, err
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodesNextEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesNextEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesNextEpoch is a free data retrieval call binding the contract method 0x0847e9fa.
//
// Solidity: function numDarknodesNextEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodesNextEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesNextEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesPreviousEpoch is a free data retrieval call binding the contract method 0x71740d16.
//
// Solidity: function numDarknodesPreviousEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) NumDarknodesPreviousEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "numDarknodesPreviousEpoch")
	return *ret0, err
}

// NumDarknodesPreviousEpoch is a free data retrieval call binding the contract method 0x71740d16.
//
// Solidity: function numDarknodesPreviousEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) NumDarknodesPreviousEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesPreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// NumDarknodesPreviousEpoch is a free data retrieval call binding the contract method 0x71740d16.
//
// Solidity: function numDarknodesPreviousEpoch() constant returns(uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) NumDarknodesPreviousEpoch() (*big.Int, error) {
	return _DarknodeRegistry.Contract.NumDarknodesPreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Owner() (common.Address, error) {
	return _DarknodeRegistry.Contract.Owner(&_DarknodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Owner() (common.Address, error) {
	return _DarknodeRegistry.Contract.Owner(&_DarknodeRegistry.CallOpts)
}

// PreviousEpoch is a free data retrieval call binding the contract method 0x5cdaab48.
//
// Solidity: function previousEpoch() constant returns(epochhash uint256, blocktime uint256)
func (_DarknodeRegistry *DarknodeRegistryCaller) PreviousEpoch(opts *bind.CallOpts) (struct {
	Epochhash *big.Int
	Blocktime *big.Int
}, error) {
	ret := new(struct {
		Epochhash *big.Int
		Blocktime *big.Int
	})
	out := ret
	err := _DarknodeRegistry.contract.Call(opts, out, "previousEpoch")
	return *ret, err
}

// PreviousEpoch is a free data retrieval call binding the contract method 0x5cdaab48.
//
// Solidity: function previousEpoch() constant returns(epochhash uint256, blocktime uint256)
func (_DarknodeRegistry *DarknodeRegistrySession) PreviousEpoch() (struct {
	Epochhash *big.Int
	Blocktime *big.Int
}, error) {
	return _DarknodeRegistry.Contract.PreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// PreviousEpoch is a free data retrieval call binding the contract method 0x5cdaab48.
//
// Solidity: function previousEpoch() constant returns(epochhash uint256, blocktime uint256)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) PreviousEpoch() (struct {
	Epochhash *big.Int
	Blocktime *big.Int
}, error) {
	return _DarknodeRegistry.Contract.PreviousEpoch(&_DarknodeRegistry.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Ren(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "ren")
	return *ret0, err
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Ren() (common.Address, error) {
	return _DarknodeRegistry.Contract.Ren(&_DarknodeRegistry.CallOpts)
}

// Ren is a free data retrieval call binding the contract method 0x8a9b4067.
//
// Solidity: function ren() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Ren() (common.Address, error) {
	return _DarknodeRegistry.Contract.Ren(&_DarknodeRegistry.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "slasher")
	return *ret0, err
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Slasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.Slasher(&_DarknodeRegistry.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Slasher() (common.Address, error) {
	return _DarknodeRegistry.Contract.Slasher(&_DarknodeRegistry.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DarknodeRegistry.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistrySession) Store() (common.Address, error) {
	return _DarknodeRegistry.Contract.Store(&_DarknodeRegistry.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_DarknodeRegistry *DarknodeRegistryCallerSession) Store() (common.Address, error) {
	return _DarknodeRegistry.Contract.Store(&_DarknodeRegistry.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) ClaimOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.ClaimOwnership(&_DarknodeRegistry.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.ClaimOwnership(&_DarknodeRegistry.TransactOpts)
}

// ClaimStoreOwnership is a paid mutator transaction binding the contract method 0x6fd689e8.
//
// Solidity: function claimStoreOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) ClaimStoreOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "claimStoreOwnership")
}

// ClaimStoreOwnership is a paid mutator transaction binding the contract method 0x6fd689e8.
//
// Solidity: function claimStoreOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) ClaimStoreOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.ClaimStoreOwnership(&_DarknodeRegistry.TransactOpts)
}

// ClaimStoreOwnership is a paid mutator transaction binding the contract method 0x6fd689e8.
//
// Solidity: function claimStoreOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) ClaimStoreOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.ClaimStoreOwnership(&_DarknodeRegistry.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Deregister(opts *bind.TransactOpts, _darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "deregister", _darknodeID)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Deregister(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Deregister(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Deregister is a paid mutator transaction binding the contract method 0x84ac33ec.
//
// Solidity: function deregister(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Deregister(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Deregister(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Epoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "epoch")
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Epoch() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Epoch(&_DarknodeRegistry.TransactOpts)
}

// Epoch is a paid mutator transaction binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Epoch() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Epoch(&_DarknodeRegistry.TransactOpts)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(_token address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) RecoverTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "recoverTokens", _token)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(_token address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) RecoverTokens(_token common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RecoverTokens(&_DarknodeRegistry.TransactOpts, _token)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(_token address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) RecoverTokens(_token common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RecoverTokens(&_DarknodeRegistry.TransactOpts, _token)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Refund(opts *bind.TransactOpts, _darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "refund", _darknodeID)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Refund(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Refund(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_darknodeID address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Refund(_darknodeID common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Refund(&_DarknodeRegistry.TransactOpts, _darknodeID)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(_darknodeID address, _publicKey bytes) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Register(opts *bind.TransactOpts, _darknodeID common.Address, _publicKey []byte) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "register", _darknodeID, _publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(_darknodeID address, _publicKey bytes) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Register(_darknodeID common.Address, _publicKey []byte) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Register(&_DarknodeRegistry.TransactOpts, _darknodeID, _publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(_darknodeID address, _publicKey bytes) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Register(_darknodeID common.Address, _publicKey []byte) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Register(&_DarknodeRegistry.TransactOpts, _darknodeID, _publicKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RenounceOwnership(&_DarknodeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.RenounceOwnership(&_DarknodeRegistry.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0xe74f8239.
//
// Solidity: function slash(_guilty address, _challenger address, _percentage uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) Slash(opts *bind.TransactOpts, _guilty common.Address, _challenger common.Address, _percentage *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "slash", _guilty, _challenger, _percentage)
}

// Slash is a paid mutator transaction binding the contract method 0xe74f8239.
//
// Solidity: function slash(_guilty address, _challenger address, _percentage uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) Slash(_guilty common.Address, _challenger common.Address, _percentage *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Slash(&_DarknodeRegistry.TransactOpts, _guilty, _challenger, _percentage)
}

// Slash is a paid mutator transaction binding the contract method 0xe74f8239.
//
// Solidity: function slash(_guilty address, _challenger address, _percentage uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) Slash(_guilty common.Address, _challenger common.Address, _percentage *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.Slash(&_DarknodeRegistry.TransactOpts, _guilty, _challenger, _percentage)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferOwnership(&_DarknodeRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferOwnership(&_DarknodeRegistry.TransactOpts, newOwner)
}

// TransferStoreOwnership is a paid mutator transaction binding the contract method 0xc2250a99.
//
// Solidity: function transferStoreOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) TransferStoreOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "transferStoreOwnership", _newOwner)
}

// TransferStoreOwnership is a paid mutator transaction binding the contract method 0xc2250a99.
//
// Solidity: function transferStoreOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) TransferStoreOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferStoreOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// TransferStoreOwnership is a paid mutator transaction binding the contract method 0xc2250a99.
//
// Solidity: function transferStoreOwnership(_newOwner address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) TransferStoreOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.TransferStoreOwnership(&_DarknodeRegistry.TransactOpts, _newOwner)
}

// UpdateDarknodePayment is a paid mutator transaction binding the contract method 0x71b4133c.
//
// Solidity: function updateDarknodePayment(_darknodePayment address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateDarknodePayment(opts *bind.TransactOpts, _darknodePayment common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateDarknodePayment", _darknodePayment)
}

// UpdateDarknodePayment is a paid mutator transaction binding the contract method 0x71b4133c.
//
// Solidity: function updateDarknodePayment(_darknodePayment address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateDarknodePayment(_darknodePayment common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateDarknodePayment(&_DarknodeRegistry.TransactOpts, _darknodePayment)
}

// UpdateDarknodePayment is a paid mutator transaction binding the contract method 0x71b4133c.
//
// Solidity: function updateDarknodePayment(_darknodePayment address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateDarknodePayment(_darknodePayment common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateDarknodePayment(&_DarknodeRegistry.TransactOpts, _darknodePayment)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumBond(opts *bind.TransactOpts, _nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumBond", _nextMinimumBond)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumBond(_nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumBond(&_DarknodeRegistry.TransactOpts, _nextMinimumBond)
}

// UpdateMinimumBond is a paid mutator transaction binding the contract method 0x0ff9aafe.
//
// Solidity: function updateMinimumBond(_nextMinimumBond uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumBond(_nextMinimumBond *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumBond(&_DarknodeRegistry.TransactOpts, _nextMinimumBond)
}

// UpdateMinimumEpochInterval is a paid mutator transaction binding the contract method 0x63b851b9.
//
// Solidity: function updateMinimumEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumEpochInterval(opts *bind.TransactOpts, _nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumEpochInterval", _nextMinimumEpochInterval)
}

// UpdateMinimumEpochInterval is a paid mutator transaction binding the contract method 0x63b851b9.
//
// Solidity: function updateMinimumEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumEpochInterval(_nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumEpochInterval(&_DarknodeRegistry.TransactOpts, _nextMinimumEpochInterval)
}

// UpdateMinimumEpochInterval is a paid mutator transaction binding the contract method 0x63b851b9.
//
// Solidity: function updateMinimumEpochInterval(_nextMinimumEpochInterval uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumEpochInterval(_nextMinimumEpochInterval *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumEpochInterval(&_DarknodeRegistry.TransactOpts, _nextMinimumEpochInterval)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateMinimumPodSize(opts *bind.TransactOpts, _nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateMinimumPodSize", _nextMinimumPodSize)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateMinimumPodSize(_nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumPodSize(&_DarknodeRegistry.TransactOpts, _nextMinimumPodSize)
}

// UpdateMinimumPodSize is a paid mutator transaction binding the contract method 0x80a0c461.
//
// Solidity: function updateMinimumPodSize(_nextMinimumPodSize uint256) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateMinimumPodSize(_nextMinimumPodSize *big.Int) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateMinimumPodSize(&_DarknodeRegistry.TransactOpts, _nextMinimumPodSize)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_slasher address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactor) UpdateSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.contract.Transact(opts, "updateSlasher", _slasher)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_slasher address) returns()
func (_DarknodeRegistry *DarknodeRegistrySession) UpdateSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateSlasher(&_DarknodeRegistry.TransactOpts, _slasher)
}

// UpdateSlasher is a paid mutator transaction binding the contract method 0xb3139d38.
//
// Solidity: function updateSlasher(_slasher address) returns()
func (_DarknodeRegistry *DarknodeRegistryTransactorSession) UpdateSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _DarknodeRegistry.Contract.UpdateSlasher(&_DarknodeRegistry.TransactOpts, _slasher)
}

// DarknodeRegistryLogDarknodeDeregisteredIterator is returned from FilterLogDarknodeDeregistered and is used to iterate over the raw logs and unpacked data for LogDarknodeDeregistered events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeDeregisteredIterator struct {
	Event *DarknodeRegistryLogDarknodeDeregistered // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeDeregistered)
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
		it.Event = new(DarknodeRegistryLogDarknodeDeregistered)
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
func (it *DarknodeRegistryLogDarknodeDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeDeregistered represents a LogDarknodeDeregistered event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeDeregistered struct {
	Operator   common.Address
	DarknodeID common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeDeregistered is a free log retrieval operation binding the contract event 0xf73268ea792d9dbf3e21a95ec9711f0b535c5f6c99f6b4f54f6766838086b842.
//
// Solidity: event LogDarknodeDeregistered(_operator indexed address, _darknodeID indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeDeregistered(opts *bind.FilterOpts, _operator []common.Address, _darknodeID []common.Address) (*DarknodeRegistryLogDarknodeDeregisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _darknodeIDRule []interface{}
	for _, _darknodeIDItem := range _darknodeID {
		_darknodeIDRule = append(_darknodeIDRule, _darknodeIDItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeDeregistered", _operatorRule, _darknodeIDRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeDeregisteredIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeDeregistered", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeDeregistered is a free log subscription operation binding the contract event 0xf73268ea792d9dbf3e21a95ec9711f0b535c5f6c99f6b4f54f6766838086b842.
//
// Solidity: event LogDarknodeDeregistered(_operator indexed address, _darknodeID indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeDeregistered(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeDeregistered, _operator []common.Address, _darknodeID []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _darknodeIDRule []interface{}
	for _, _darknodeIDItem := range _darknodeID {
		_darknodeIDRule = append(_darknodeIDRule, _darknodeIDItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeDeregistered", _operatorRule, _darknodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeDeregistered)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeDeregistered", log); err != nil {
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

// DarknodeRegistryLogDarknodeOwnerRefundedIterator is returned from FilterLogDarknodeOwnerRefunded and is used to iterate over the raw logs and unpacked data for LogDarknodeOwnerRefunded events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeOwnerRefundedIterator struct {
	Event *DarknodeRegistryLogDarknodeOwnerRefunded // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeOwnerRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeOwnerRefunded)
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
		it.Event = new(DarknodeRegistryLogDarknodeOwnerRefunded)
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
func (it *DarknodeRegistryLogDarknodeOwnerRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeOwnerRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeOwnerRefunded represents a LogDarknodeOwnerRefunded event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeOwnerRefunded struct {
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeOwnerRefunded is a free log retrieval operation binding the contract event 0x96ab9e56c79eee4a72db6e2879cbfbecdba5c65b411f4861824e66b89df19764.
//
// Solidity: event LogDarknodeOwnerRefunded(_operator indexed address, _amount uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeOwnerRefunded(opts *bind.FilterOpts, _operator []common.Address) (*DarknodeRegistryLogDarknodeOwnerRefundedIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeOwnerRefunded", _operatorRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeOwnerRefundedIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeOwnerRefunded", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeOwnerRefunded is a free log subscription operation binding the contract event 0x96ab9e56c79eee4a72db6e2879cbfbecdba5c65b411f4861824e66b89df19764.
//
// Solidity: event LogDarknodeOwnerRefunded(_operator indexed address, _amount uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeOwnerRefunded(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeOwnerRefunded, _operator []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeOwnerRefunded", _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeOwnerRefunded)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeOwnerRefunded", log); err != nil {
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

// DarknodeRegistryLogDarknodePaymentUpdatedIterator is returned from FilterLogDarknodePaymentUpdated and is used to iterate over the raw logs and unpacked data for LogDarknodePaymentUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodePaymentUpdatedIterator struct {
	Event *DarknodeRegistryLogDarknodePaymentUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodePaymentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodePaymentUpdated)
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
		it.Event = new(DarknodeRegistryLogDarknodePaymentUpdated)
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
func (it *DarknodeRegistryLogDarknodePaymentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodePaymentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodePaymentUpdated represents a LogDarknodePaymentUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodePaymentUpdated struct {
	PreviousDarknodePayment common.Address
	NextDarknodePayment     common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodePaymentUpdated is a free log retrieval operation binding the contract event 0xe3e25a79a5ba7c894fcc55794b2712e225537e89f777b9b9df307cc5504ba0e9.
//
// Solidity: event LogDarknodePaymentUpdated(_previousDarknodePayment address, _nextDarknodePayment address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodePaymentUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogDarknodePaymentUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodePaymentUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodePaymentUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodePaymentUpdated", logs: logs, sub: sub}, nil
}

// WatchLogDarknodePaymentUpdated is a free log subscription operation binding the contract event 0xe3e25a79a5ba7c894fcc55794b2712e225537e89f777b9b9df307cc5504ba0e9.
//
// Solidity: event LogDarknodePaymentUpdated(_previousDarknodePayment address, _nextDarknodePayment address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodePaymentUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodePaymentUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodePaymentUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodePaymentUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodePaymentUpdated", log); err != nil {
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

// DarknodeRegistryLogDarknodeRegisteredIterator is returned from FilterLogDarknodeRegistered and is used to iterate over the raw logs and unpacked data for LogDarknodeRegistered events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeRegisteredIterator struct {
	Event *DarknodeRegistryLogDarknodeRegistered // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeRegistered)
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
		it.Event = new(DarknodeRegistryLogDarknodeRegistered)
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
func (it *DarknodeRegistryLogDarknodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeRegistered represents a LogDarknodeRegistered event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeRegistered struct {
	Operator   common.Address
	DarknodeID common.Address
	Bond       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeRegistered is a free log retrieval operation binding the contract event 0x7c56cb7f63b6922d24414bf7c2b2c40c7ea1ea637c3f400efa766a85ecf2f093.
//
// Solidity: event LogDarknodeRegistered(_operator indexed address, _darknodeID indexed address, _bond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeRegistered(opts *bind.FilterOpts, _operator []common.Address, _darknodeID []common.Address) (*DarknodeRegistryLogDarknodeRegisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _darknodeIDRule []interface{}
	for _, _darknodeIDItem := range _darknodeID {
		_darknodeIDRule = append(_darknodeIDRule, _darknodeIDItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeRegistered", _operatorRule, _darknodeIDRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeRegisteredIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeRegistered", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeRegistered is a free log subscription operation binding the contract event 0x7c56cb7f63b6922d24414bf7c2b2c40c7ea1ea637c3f400efa766a85ecf2f093.
//
// Solidity: event LogDarknodeRegistered(_operator indexed address, _darknodeID indexed address, _bond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeRegistered(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeRegistered, _operator []common.Address, _darknodeID []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _darknodeIDRule []interface{}
	for _, _darknodeIDItem := range _darknodeID {
		_darknodeIDRule = append(_darknodeIDRule, _darknodeIDItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeRegistered", _operatorRule, _darknodeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeRegistered)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeRegistered", log); err != nil {
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

// DarknodeRegistryLogDarknodeSlashedIterator is returned from FilterLogDarknodeSlashed and is used to iterate over the raw logs and unpacked data for LogDarknodeSlashed events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeSlashedIterator struct {
	Event *DarknodeRegistryLogDarknodeSlashed // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogDarknodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogDarknodeSlashed)
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
		it.Event = new(DarknodeRegistryLogDarknodeSlashed)
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
func (it *DarknodeRegistryLogDarknodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogDarknodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogDarknodeSlashed represents a LogDarknodeSlashed event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogDarknodeSlashed struct {
	Operator   common.Address
	DarknodeID common.Address
	Challenger common.Address
	Percentage *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogDarknodeSlashed is a free log retrieval operation binding the contract event 0xb43e0cc88b4d6ae901c6c99d1b58769cb8c9ded8e6f20a0d3712d09bf9e1ea77.
//
// Solidity: event LogDarknodeSlashed(_operator indexed address, _darknodeID indexed address, _challenger indexed address, _percentage uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogDarknodeSlashed(opts *bind.FilterOpts, _operator []common.Address, _darknodeID []common.Address, _challenger []common.Address) (*DarknodeRegistryLogDarknodeSlashedIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _darknodeIDRule []interface{}
	for _, _darknodeIDItem := range _darknodeID {
		_darknodeIDRule = append(_darknodeIDRule, _darknodeIDItem)
	}
	var _challengerRule []interface{}
	for _, _challengerItem := range _challenger {
		_challengerRule = append(_challengerRule, _challengerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogDarknodeSlashed", _operatorRule, _darknodeIDRule, _challengerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogDarknodeSlashedIterator{contract: _DarknodeRegistry.contract, event: "LogDarknodeSlashed", logs: logs, sub: sub}, nil
}

// WatchLogDarknodeSlashed is a free log subscription operation binding the contract event 0xb43e0cc88b4d6ae901c6c99d1b58769cb8c9ded8e6f20a0d3712d09bf9e1ea77.
//
// Solidity: event LogDarknodeSlashed(_operator indexed address, _darknodeID indexed address, _challenger indexed address, _percentage uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogDarknodeSlashed(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogDarknodeSlashed, _operator []common.Address, _darknodeID []common.Address, _challenger []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _darknodeIDRule []interface{}
	for _, _darknodeIDItem := range _darknodeID {
		_darknodeIDRule = append(_darknodeIDRule, _darknodeIDItem)
	}
	var _challengerRule []interface{}
	for _, _challengerItem := range _challenger {
		_challengerRule = append(_challengerRule, _challengerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogDarknodeSlashed", _operatorRule, _darknodeIDRule, _challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogDarknodeSlashed)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogDarknodeSlashed", log); err != nil {
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

// DarknodeRegistryLogMinimumBondUpdatedIterator is returned from FilterLogMinimumBondUpdated and is used to iterate over the raw logs and unpacked data for LogMinimumBondUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumBondUpdatedIterator struct {
	Event *DarknodeRegistryLogMinimumBondUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogMinimumBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogMinimumBondUpdated)
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
		it.Event = new(DarknodeRegistryLogMinimumBondUpdated)
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
func (it *DarknodeRegistryLogMinimumBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogMinimumBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogMinimumBondUpdated represents a LogMinimumBondUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumBondUpdated struct {
	PreviousMinimumBond *big.Int
	NextMinimumBond     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogMinimumBondUpdated is a free log retrieval operation binding the contract event 0x7c77c94944e9e4e5b0d46f1297127d060020792687cd743401d782346c68f655.
//
// Solidity: event LogMinimumBondUpdated(_previousMinimumBond uint256, _nextMinimumBond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogMinimumBondUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogMinimumBondUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogMinimumBondUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogMinimumBondUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogMinimumBondUpdated", logs: logs, sub: sub}, nil
}

// WatchLogMinimumBondUpdated is a free log subscription operation binding the contract event 0x7c77c94944e9e4e5b0d46f1297127d060020792687cd743401d782346c68f655.
//
// Solidity: event LogMinimumBondUpdated(_previousMinimumBond uint256, _nextMinimumBond uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogMinimumBondUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogMinimumBondUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogMinimumBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogMinimumBondUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogMinimumBondUpdated", log); err != nil {
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

// DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator is returned from FilterLogMinimumEpochIntervalUpdated and is used to iterate over the raw logs and unpacked data for LogMinimumEpochIntervalUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator struct {
	Event *DarknodeRegistryLogMinimumEpochIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogMinimumEpochIntervalUpdated)
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
		it.Event = new(DarknodeRegistryLogMinimumEpochIntervalUpdated)
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
func (it *DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogMinimumEpochIntervalUpdated represents a LogMinimumEpochIntervalUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumEpochIntervalUpdated struct {
	PreviousMinimumEpochInterval *big.Int
	NextMinimumEpochInterval     *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterLogMinimumEpochIntervalUpdated is a free log retrieval operation binding the contract event 0xb218cde2730b79a0667ddf869466ee66a12ef56fe65fa4986a590f8a7108c9de.
//
// Solidity: event LogMinimumEpochIntervalUpdated(_previousMinimumEpochInterval uint256, _nextMinimumEpochInterval uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogMinimumEpochIntervalUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogMinimumEpochIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogMinimumEpochIntervalUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogMinimumEpochIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchLogMinimumEpochIntervalUpdated is a free log subscription operation binding the contract event 0xb218cde2730b79a0667ddf869466ee66a12ef56fe65fa4986a590f8a7108c9de.
//
// Solidity: event LogMinimumEpochIntervalUpdated(_previousMinimumEpochInterval uint256, _nextMinimumEpochInterval uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogMinimumEpochIntervalUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogMinimumEpochIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogMinimumEpochIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogMinimumEpochIntervalUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogMinimumEpochIntervalUpdated", log); err != nil {
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

// DarknodeRegistryLogMinimumPodSizeUpdatedIterator is returned from FilterLogMinimumPodSizeUpdated and is used to iterate over the raw logs and unpacked data for LogMinimumPodSizeUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumPodSizeUpdatedIterator struct {
	Event *DarknodeRegistryLogMinimumPodSizeUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogMinimumPodSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogMinimumPodSizeUpdated)
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
		it.Event = new(DarknodeRegistryLogMinimumPodSizeUpdated)
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
func (it *DarknodeRegistryLogMinimumPodSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogMinimumPodSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogMinimumPodSizeUpdated represents a LogMinimumPodSizeUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogMinimumPodSizeUpdated struct {
	PreviousMinimumPodSize *big.Int
	NextMinimumPodSize     *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLogMinimumPodSizeUpdated is a free log retrieval operation binding the contract event 0x6d520e46e5714982ddf8cb6216bcb3e1c1d5b79d337afc305335f819394f5d6a.
//
// Solidity: event LogMinimumPodSizeUpdated(_previousMinimumPodSize uint256, _nextMinimumPodSize uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogMinimumPodSizeUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogMinimumPodSizeUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogMinimumPodSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogMinimumPodSizeUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogMinimumPodSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchLogMinimumPodSizeUpdated is a free log subscription operation binding the contract event 0x6d520e46e5714982ddf8cb6216bcb3e1c1d5b79d337afc305335f819394f5d6a.
//
// Solidity: event LogMinimumPodSizeUpdated(_previousMinimumPodSize uint256, _nextMinimumPodSize uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogMinimumPodSizeUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogMinimumPodSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogMinimumPodSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogMinimumPodSizeUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogMinimumPodSizeUpdated", log); err != nil {
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

// DarknodeRegistryLogNewEpochIterator is returned from FilterLogNewEpoch and is used to iterate over the raw logs and unpacked data for LogNewEpoch events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogNewEpochIterator struct {
	Event *DarknodeRegistryLogNewEpoch // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogNewEpoch)
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
		it.Event = new(DarknodeRegistryLogNewEpoch)
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
func (it *DarknodeRegistryLogNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogNewEpoch represents a LogNewEpoch event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogNewEpoch struct {
	Epochhash *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNewEpoch is a free log retrieval operation binding the contract event 0xaf2fc4796f2932ce294c3684deffe5098d3ef65dc2dd64efa80ef94eed88b01e.
//
// Solidity: event LogNewEpoch(epochhash indexed uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogNewEpoch(opts *bind.FilterOpts, epochhash []*big.Int) (*DarknodeRegistryLogNewEpochIterator, error) {

	var epochhashRule []interface{}
	for _, epochhashItem := range epochhash {
		epochhashRule = append(epochhashRule, epochhashItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogNewEpoch", epochhashRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogNewEpochIterator{contract: _DarknodeRegistry.contract, event: "LogNewEpoch", logs: logs, sub: sub}, nil
}

// WatchLogNewEpoch is a free log subscription operation binding the contract event 0xaf2fc4796f2932ce294c3684deffe5098d3ef65dc2dd64efa80ef94eed88b01e.
//
// Solidity: event LogNewEpoch(epochhash indexed uint256)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogNewEpoch(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogNewEpoch, epochhash []*big.Int) (event.Subscription, error) {

	var epochhashRule []interface{}
	for _, epochhashItem := range epochhash {
		epochhashRule = append(epochhashRule, epochhashItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogNewEpoch", epochhashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogNewEpoch)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogNewEpoch", log); err != nil {
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

// DarknodeRegistryLogSlasherUpdatedIterator is returned from FilterLogSlasherUpdated and is used to iterate over the raw logs and unpacked data for LogSlasherUpdated events raised by the DarknodeRegistry contract.
type DarknodeRegistryLogSlasherUpdatedIterator struct {
	Event *DarknodeRegistryLogSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryLogSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryLogSlasherUpdated)
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
		it.Event = new(DarknodeRegistryLogSlasherUpdated)
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
func (it *DarknodeRegistryLogSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryLogSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryLogSlasherUpdated represents a LogSlasherUpdated event raised by the DarknodeRegistry contract.
type DarknodeRegistryLogSlasherUpdated struct {
	PreviousSlasher common.Address
	NextSlasher     common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSlasherUpdated is a free log retrieval operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: event LogSlasherUpdated(_previousSlasher address, _nextSlasher address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterLogSlasherUpdated(opts *bind.FilterOpts) (*DarknodeRegistryLogSlasherUpdatedIterator, error) {

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryLogSlasherUpdatedIterator{contract: _DarknodeRegistry.contract, event: "LogSlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchLogSlasherUpdated is a free log subscription operation binding the contract event 0x933228a1c3ba8fadd3ce47a9db5b898be647f89af99ba7c1b9a655f59ea306c8.
//
// Solidity: event LogSlasherUpdated(_previousSlasher address, _nextSlasher address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchLogSlasherUpdated(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryLogSlasherUpdated) (event.Subscription, error) {

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "LogSlasherUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryLogSlasherUpdated)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "LogSlasherUpdated", log); err != nil {
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

// DarknodeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipTransferredIterator struct {
	Event *DarknodeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DarknodeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DarknodeRegistryOwnershipTransferred)
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
		it.Event = new(DarknodeRegistryOwnershipTransferred)
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
func (it *DarknodeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DarknodeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DarknodeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the DarknodeRegistry contract.
type DarknodeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DarknodeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DarknodeRegistryOwnershipTransferredIterator{contract: _DarknodeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DarknodeRegistry *DarknodeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DarknodeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DarknodeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DarknodeRegistryOwnershipTransferred)
				if err := _DarknodeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
