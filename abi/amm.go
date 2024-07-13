// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.ConvertType
)

// AMMMetaData contains all meta data concerning the AMM contract.
var AMMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subscriptionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_win_optional_probability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_optional_down_bound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_optional_up_bound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_market_swap_rate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_basicToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"OnlyOwnerOrCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"}],\"name\":\"CoordinatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"sellType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"basicToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetToken\",\"type\":\"uint256\"}],\"name\":\"GetRandoms\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"InTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preOutTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalOutTokenAmount\",\"type\":\"uint256\"}],\"name\":\"SwapSuceesfully\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BalanceK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"persentage\",\"type\":\"uint256\"}],\"name\":\"Settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basicTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_swapType\",\"type\":\"uint8\"}],\"name\":\"Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_basicTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_sellType\",\"type\":\"uint8\"}],\"name\":\"SwapResultOfTargetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_win_optional_probability\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_optional_down_bound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_optional_up_bound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_market_swap_rate\",\"type\":\"uint256\"}],\"name\":\"UpdateConf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basicToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basicTokenReserve\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callbackGasLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"k\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market_swap_rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numWords\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optional_down_bound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optional_up_bound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"randoms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestConfirmations\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_subscriptionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_vrfCoordinator\",\"outputs\":[{\"internalType\":\"contractIVRFCoordinatorV2Plus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"settleResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetTokenReserve\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfCoordinator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"win_optional_probability\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AMMABI is the input ABI used to generate the binding from.
// Deprecated: Use AMMMetaData.ABI instead.
var AMMABI = AMMMetaData.ABI

// AMM is an auto generated Go binding around an Ethereum contract.
type AMM struct {
	AMMCaller     // Read-only binding to the contract
	AMMTransactor // Write-only binding to the contract
	AMMFilterer   // Log filterer for contract events
}

// AMMCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMMSession struct {
	Contract     *AMM              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMMCallerSession struct {
	Contract *AMMCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AMMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMMTransactorSession struct {
	Contract     *AMMTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMMRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMMRaw struct {
	Contract *AMM // Generic contract binding to access the raw methods on
}

// AMMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMMCallerRaw struct {
	Contract *AMMCaller // Generic read-only contract binding to access the raw methods on
}

// AMMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMMTransactorRaw struct {
	Contract *AMMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAMM creates a new instance of AMM, bound to a specific deployed contract.
func NewAMM(address common.Address, backend bind.ContractBackend) (*AMM, error) {
	contract, err := bindAMM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AMM{AMMCaller: AMMCaller{contract: contract}, AMMTransactor: AMMTransactor{contract: contract}, AMMFilterer: AMMFilterer{contract: contract}}, nil
}

// NewAMMCaller creates a new read-only instance of AMM, bound to a specific deployed contract.
func NewAMMCaller(address common.Address, caller bind.ContractCaller) (*AMMCaller, error) {
	contract, err := bindAMM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMMCaller{contract: contract}, nil
}

// NewAMMTransactor creates a new write-only instance of AMM, bound to a specific deployed contract.
func NewAMMTransactor(address common.Address, transactor bind.ContractTransactor) (*AMMTransactor, error) {
	contract, err := bindAMM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMMTransactor{contract: contract}, nil
}

// NewAMMFilterer creates a new log filterer instance of AMM, bound to a specific deployed contract.
func NewAMMFilterer(address common.Address, filterer bind.ContractFilterer) (*AMMFilterer, error) {
	contract, err := bindAMM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMMFilterer{contract: contract}, nil
}

// bindAMM binds a generic wrapper to an already deployed contract.
func bindAMM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AMMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMM *AMMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMM.Contract.AMMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMM *AMMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMM.Contract.AMMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMM *AMMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMM.Contract.AMMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMM *AMMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMM *AMMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMM *AMMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMM.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AMM *AMMCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AMM *AMMSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AMM.Contract.DEFAULTADMINROLE(&_AMM.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AMM *AMMCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AMM.Contract.DEFAULTADMINROLE(&_AMM.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_AMM *AMMCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_AMM *AMMSession) OPERATORROLE() ([32]byte, error) {
	return _AMM.Contract.OPERATORROLE(&_AMM.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_AMM *AMMCallerSession) OPERATORROLE() ([32]byte, error) {
	return _AMM.Contract.OPERATORROLE(&_AMM.CallOpts)
}

// SwapResultOfTargetAmount is a free data retrieval call binding the contract method 0xdb9d1e3c.
//
// Solidity: function SwapResultOfTargetAmount(uint256 _targetTokenAmount, uint256 _basicTokenAmount, uint8 _sellType) view returns(uint256, uint256, uint256)
func (_AMM *AMMCaller) SwapResultOfTargetAmount(opts *bind.CallOpts, _targetTokenAmount *big.Int, _basicTokenAmount *big.Int, _sellType uint8) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "SwapResultOfTargetAmount", _targetTokenAmount, _basicTokenAmount, _sellType)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// SwapResultOfTargetAmount is a free data retrieval call binding the contract method 0xdb9d1e3c.
//
// Solidity: function SwapResultOfTargetAmount(uint256 _targetTokenAmount, uint256 _basicTokenAmount, uint8 _sellType) view returns(uint256, uint256, uint256)
func (_AMM *AMMSession) SwapResultOfTargetAmount(_targetTokenAmount *big.Int, _basicTokenAmount *big.Int, _sellType uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _AMM.Contract.SwapResultOfTargetAmount(&_AMM.CallOpts, _targetTokenAmount, _basicTokenAmount, _sellType)
}

// SwapResultOfTargetAmount is a free data retrieval call binding the contract method 0xdb9d1e3c.
//
// Solidity: function SwapResultOfTargetAmount(uint256 _targetTokenAmount, uint256 _basicTokenAmount, uint8 _sellType) view returns(uint256, uint256, uint256)
func (_AMM *AMMCallerSession) SwapResultOfTargetAmount(_targetTokenAmount *big.Int, _basicTokenAmount *big.Int, _sellType uint8) (*big.Int, *big.Int, *big.Int, error) {
	return _AMM.Contract.SwapResultOfTargetAmount(&_AMM.CallOpts, _targetTokenAmount, _basicTokenAmount, _sellType)
}

// BasicToken is a free data retrieval call binding the contract method 0x190b8dc5.
//
// Solidity: function basicToken() view returns(address)
func (_AMM *AMMCaller) BasicToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "basicToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasicToken is a free data retrieval call binding the contract method 0x190b8dc5.
//
// Solidity: function basicToken() view returns(address)
func (_AMM *AMMSession) BasicToken() (common.Address, error) {
	return _AMM.Contract.BasicToken(&_AMM.CallOpts)
}

// BasicToken is a free data retrieval call binding the contract method 0x190b8dc5.
//
// Solidity: function basicToken() view returns(address)
func (_AMM *AMMCallerSession) BasicToken() (common.Address, error) {
	return _AMM.Contract.BasicToken(&_AMM.CallOpts)
}

// BasicTokenReserve is a free data retrieval call binding the contract method 0x1a4ff275.
//
// Solidity: function basicTokenReserve() view returns(uint112)
func (_AMM *AMMCaller) BasicTokenReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "basicTokenReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasicTokenReserve is a free data retrieval call binding the contract method 0x1a4ff275.
//
// Solidity: function basicTokenReserve() view returns(uint112)
func (_AMM *AMMSession) BasicTokenReserve() (*big.Int, error) {
	return _AMM.Contract.BasicTokenReserve(&_AMM.CallOpts)
}

// BasicTokenReserve is a free data retrieval call binding the contract method 0x1a4ff275.
//
// Solidity: function basicTokenReserve() view returns(uint112)
func (_AMM *AMMCallerSession) BasicTokenReserve() (*big.Int, error) {
	return _AMM.Contract.BasicTokenReserve(&_AMM.CallOpts)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint32)
func (_AMM *AMMCaller) CallbackGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "callbackGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint32)
func (_AMM *AMMSession) CallbackGasLimit() (uint32, error) {
	return _AMM.Contract.CallbackGasLimit(&_AMM.CallOpts)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint32)
func (_AMM *AMMCallerSession) CallbackGasLimit() (uint32, error) {
	return _AMM.Contract.CallbackGasLimit(&_AMM.CallOpts)
}

// GetConfInfo is a free data retrieval call binding the contract method 0xb54f681e.
//
// Solidity: function getConfInfo() view returns(uint256, uint256, uint256, uint256)
func (_AMM *AMMCaller) GetConfInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "getConfInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetConfInfo is a free data retrieval call binding the contract method 0xb54f681e.
//
// Solidity: function getConfInfo() view returns(uint256, uint256, uint256, uint256)
func (_AMM *AMMSession) GetConfInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AMM.Contract.GetConfInfo(&_AMM.CallOpts)
}

// GetConfInfo is a free data retrieval call binding the contract method 0xb54f681e.
//
// Solidity: function getConfInfo() view returns(uint256, uint256, uint256, uint256)
func (_AMM *AMMCallerSession) GetConfInfo() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AMM.Contract.GetConfInfo(&_AMM.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AMM *AMMCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AMM *AMMSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AMM.Contract.GetRoleAdmin(&_AMM.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AMM *AMMCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AMM.Contract.GetRoleAdmin(&_AMM.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AMM *AMMCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AMM *AMMSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AMM.Contract.HasRole(&_AMM.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AMM *AMMCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AMM.Contract.HasRole(&_AMM.CallOpts, role, account)
}

// K is a free data retrieval call binding the contract method 0xb4f40c61.
//
// Solidity: function k() view returns(uint256)
func (_AMM *AMMCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "k")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xb4f40c61.
//
// Solidity: function k() view returns(uint256)
func (_AMM *AMMSession) K() (*big.Int, error) {
	return _AMM.Contract.K(&_AMM.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xb4f40c61.
//
// Solidity: function k() view returns(uint256)
func (_AMM *AMMCallerSession) K() (*big.Int, error) {
	return _AMM.Contract.K(&_AMM.CallOpts)
}

// MarketSwapRate is a free data retrieval call binding the contract method 0x49f3a6f0.
//
// Solidity: function market_swap_rate() view returns(uint256)
func (_AMM *AMMCaller) MarketSwapRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "market_swap_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketSwapRate is a free data retrieval call binding the contract method 0x49f3a6f0.
//
// Solidity: function market_swap_rate() view returns(uint256)
func (_AMM *AMMSession) MarketSwapRate() (*big.Int, error) {
	return _AMM.Contract.MarketSwapRate(&_AMM.CallOpts)
}

// MarketSwapRate is a free data retrieval call binding the contract method 0x49f3a6f0.
//
// Solidity: function market_swap_rate() view returns(uint256)
func (_AMM *AMMCallerSession) MarketSwapRate() (*big.Int, error) {
	return _AMM.Contract.MarketSwapRate(&_AMM.CallOpts)
}

// NumWords is a free data retrieval call binding the contract method 0x7ccfd7fc.
//
// Solidity: function numWords() view returns(uint32)
func (_AMM *AMMCaller) NumWords(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "numWords")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumWords is a free data retrieval call binding the contract method 0x7ccfd7fc.
//
// Solidity: function numWords() view returns(uint32)
func (_AMM *AMMSession) NumWords() (uint32, error) {
	return _AMM.Contract.NumWords(&_AMM.CallOpts)
}

// NumWords is a free data retrieval call binding the contract method 0x7ccfd7fc.
//
// Solidity: function numWords() view returns(uint32)
func (_AMM *AMMCallerSession) NumWords() (uint32, error) {
	return _AMM.Contract.NumWords(&_AMM.CallOpts)
}

// OptionalDownBound is a free data retrieval call binding the contract method 0x6acab104.
//
// Solidity: function optional_down_bound() view returns(uint256)
func (_AMM *AMMCaller) OptionalDownBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "optional_down_bound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OptionalDownBound is a free data retrieval call binding the contract method 0x6acab104.
//
// Solidity: function optional_down_bound() view returns(uint256)
func (_AMM *AMMSession) OptionalDownBound() (*big.Int, error) {
	return _AMM.Contract.OptionalDownBound(&_AMM.CallOpts)
}

// OptionalDownBound is a free data retrieval call binding the contract method 0x6acab104.
//
// Solidity: function optional_down_bound() view returns(uint256)
func (_AMM *AMMCallerSession) OptionalDownBound() (*big.Int, error) {
	return _AMM.Contract.OptionalDownBound(&_AMM.CallOpts)
}

// OptionalUpBound is a free data retrieval call binding the contract method 0xc0452021.
//
// Solidity: function optional_up_bound() view returns(uint256)
func (_AMM *AMMCaller) OptionalUpBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "optional_up_bound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OptionalUpBound is a free data retrieval call binding the contract method 0xc0452021.
//
// Solidity: function optional_up_bound() view returns(uint256)
func (_AMM *AMMSession) OptionalUpBound() (*big.Int, error) {
	return _AMM.Contract.OptionalUpBound(&_AMM.CallOpts)
}

// OptionalUpBound is a free data retrieval call binding the contract method 0xc0452021.
//
// Solidity: function optional_up_bound() view returns(uint256)
func (_AMM *AMMCallerSession) OptionalUpBound() (*big.Int, error) {
	return _AMM.Contract.OptionalUpBound(&_AMM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AMM *AMMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AMM *AMMSession) Owner() (common.Address, error) {
	return _AMM.Contract.Owner(&_AMM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AMM *AMMCallerSession) Owner() (common.Address, error) {
	return _AMM.Contract.Owner(&_AMM.CallOpts)
}

// Randoms is a free data retrieval call binding the contract method 0xc0d0ed7c.
//
// Solidity: function randoms(uint256 ) view returns(uint256)
func (_AMM *AMMCaller) Randoms(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "randoms", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Randoms is a free data retrieval call binding the contract method 0xc0d0ed7c.
//
// Solidity: function randoms(uint256 ) view returns(uint256)
func (_AMM *AMMSession) Randoms(arg0 *big.Int) (*big.Int, error) {
	return _AMM.Contract.Randoms(&_AMM.CallOpts, arg0)
}

// Randoms is a free data retrieval call binding the contract method 0xc0d0ed7c.
//
// Solidity: function randoms(uint256 ) view returns(uint256)
func (_AMM *AMMCallerSession) Randoms(arg0 *big.Int) (*big.Int, error) {
	return _AMM.Contract.Randoms(&_AMM.CallOpts, arg0)
}

// RequestConfirmations is a free data retrieval call binding the contract method 0xb0fb162f.
//
// Solidity: function requestConfirmations() view returns(uint16)
func (_AMM *AMMCaller) RequestConfirmations(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "requestConfirmations")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RequestConfirmations is a free data retrieval call binding the contract method 0xb0fb162f.
//
// Solidity: function requestConfirmations() view returns(uint16)
func (_AMM *AMMSession) RequestConfirmations() (uint16, error) {
	return _AMM.Contract.RequestConfirmations(&_AMM.CallOpts)
}

// RequestConfirmations is a free data retrieval call binding the contract method 0xb0fb162f.
//
// Solidity: function requestConfirmations() view returns(uint16)
func (_AMM *AMMCallerSession) RequestConfirmations() (uint16, error) {
	return _AMM.Contract.RequestConfirmations(&_AMM.CallOpts)
}

// SKeyHash is a free data retrieval call binding the contract method 0x45bb327b.
//
// Solidity: function s_keyHash() view returns(bytes32)
func (_AMM *AMMCaller) SKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "s_keyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SKeyHash is a free data retrieval call binding the contract method 0x45bb327b.
//
// Solidity: function s_keyHash() view returns(bytes32)
func (_AMM *AMMSession) SKeyHash() ([32]byte, error) {
	return _AMM.Contract.SKeyHash(&_AMM.CallOpts)
}

// SKeyHash is a free data retrieval call binding the contract method 0x45bb327b.
//
// Solidity: function s_keyHash() view returns(bytes32)
func (_AMM *AMMCallerSession) SKeyHash() ([32]byte, error) {
	return _AMM.Contract.SKeyHash(&_AMM.CallOpts)
}

// SSubscriptionId is a free data retrieval call binding the contract method 0x8ac00021.
//
// Solidity: function s_subscriptionId() view returns(uint256)
func (_AMM *AMMCaller) SSubscriptionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "s_subscriptionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SSubscriptionId is a free data retrieval call binding the contract method 0x8ac00021.
//
// Solidity: function s_subscriptionId() view returns(uint256)
func (_AMM *AMMSession) SSubscriptionId() (*big.Int, error) {
	return _AMM.Contract.SSubscriptionId(&_AMM.CallOpts)
}

// SSubscriptionId is a free data retrieval call binding the contract method 0x8ac00021.
//
// Solidity: function s_subscriptionId() view returns(uint256)
func (_AMM *AMMCallerSession) SSubscriptionId() (*big.Int, error) {
	return _AMM.Contract.SSubscriptionId(&_AMM.CallOpts)
}

// SVrfCoordinator is a free data retrieval call binding the contract method 0x9eccacf6.
//
// Solidity: function s_vrfCoordinator() view returns(address)
func (_AMM *AMMCaller) SVrfCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "s_vrfCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SVrfCoordinator is a free data retrieval call binding the contract method 0x9eccacf6.
//
// Solidity: function s_vrfCoordinator() view returns(address)
func (_AMM *AMMSession) SVrfCoordinator() (common.Address, error) {
	return _AMM.Contract.SVrfCoordinator(&_AMM.CallOpts)
}

// SVrfCoordinator is a free data retrieval call binding the contract method 0x9eccacf6.
//
// Solidity: function s_vrfCoordinator() view returns(address)
func (_AMM *AMMCallerSession) SVrfCoordinator() (common.Address, error) {
	return _AMM.Contract.SVrfCoordinator(&_AMM.CallOpts)
}

// SettleResult is a free data retrieval call binding the contract method 0xd3cd5d64.
//
// Solidity: function settleResult(uint256 ) view returns(uint256)
func (_AMM *AMMCaller) SettleResult(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "settleResult", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettleResult is a free data retrieval call binding the contract method 0xd3cd5d64.
//
// Solidity: function settleResult(uint256 ) view returns(uint256)
func (_AMM *AMMSession) SettleResult(arg0 *big.Int) (*big.Int, error) {
	return _AMM.Contract.SettleResult(&_AMM.CallOpts, arg0)
}

// SettleResult is a free data retrieval call binding the contract method 0xd3cd5d64.
//
// Solidity: function settleResult(uint256 ) view returns(uint256)
func (_AMM *AMMCallerSession) SettleResult(arg0 *big.Int) (*big.Int, error) {
	return _AMM.Contract.SettleResult(&_AMM.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AMM *AMMCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AMM *AMMSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AMM.Contract.SupportsInterface(&_AMM.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AMM *AMMCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AMM.Contract.SupportsInterface(&_AMM.CallOpts, interfaceId)
}

// TargetToken is a free data retrieval call binding the contract method 0x327107f7.
//
// Solidity: function targetToken() view returns(address)
func (_AMM *AMMCaller) TargetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "targetToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetToken is a free data retrieval call binding the contract method 0x327107f7.
//
// Solidity: function targetToken() view returns(address)
func (_AMM *AMMSession) TargetToken() (common.Address, error) {
	return _AMM.Contract.TargetToken(&_AMM.CallOpts)
}

// TargetToken is a free data retrieval call binding the contract method 0x327107f7.
//
// Solidity: function targetToken() view returns(address)
func (_AMM *AMMCallerSession) TargetToken() (common.Address, error) {
	return _AMM.Contract.TargetToken(&_AMM.CallOpts)
}

// TargetTokenReserve is a free data retrieval call binding the contract method 0x04ff431d.
//
// Solidity: function targetTokenReserve() view returns(uint112)
func (_AMM *AMMCaller) TargetTokenReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "targetTokenReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetTokenReserve is a free data retrieval call binding the contract method 0x04ff431d.
//
// Solidity: function targetTokenReserve() view returns(uint112)
func (_AMM *AMMSession) TargetTokenReserve() (*big.Int, error) {
	return _AMM.Contract.TargetTokenReserve(&_AMM.CallOpts)
}

// TargetTokenReserve is a free data retrieval call binding the contract method 0x04ff431d.
//
// Solidity: function targetTokenReserve() view returns(uint112)
func (_AMM *AMMCallerSession) TargetTokenReserve() (*big.Int, error) {
	return _AMM.Contract.TargetTokenReserve(&_AMM.CallOpts)
}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_AMM *AMMCaller) VrfCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "vrfCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_AMM *AMMSession) VrfCoordinator() (common.Address, error) {
	return _AMM.Contract.VrfCoordinator(&_AMM.CallOpts)
}

// VrfCoordinator is a free data retrieval call binding the contract method 0xa3e56fa8.
//
// Solidity: function vrfCoordinator() view returns(address)
func (_AMM *AMMCallerSession) VrfCoordinator() (common.Address, error) {
	return _AMM.Contract.VrfCoordinator(&_AMM.CallOpts)
}

// WinOptionalProbability is a free data retrieval call binding the contract method 0xcca02e16.
//
// Solidity: function win_optional_probability() view returns(uint256)
func (_AMM *AMMCaller) WinOptionalProbability(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AMM.contract.Call(opts, &out, "win_optional_probability")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinOptionalProbability is a free data retrieval call binding the contract method 0xcca02e16.
//
// Solidity: function win_optional_probability() view returns(uint256)
func (_AMM *AMMSession) WinOptionalProbability() (*big.Int, error) {
	return _AMM.Contract.WinOptionalProbability(&_AMM.CallOpts)
}

// WinOptionalProbability is a free data retrieval call binding the contract method 0xcca02e16.
//
// Solidity: function win_optional_probability() view returns(uint256)
func (_AMM *AMMCallerSession) WinOptionalProbability() (*big.Int, error) {
	return _AMM.Contract.WinOptionalProbability(&_AMM.CallOpts)
}

// BalanceK is a paid mutator transaction binding the contract method 0x29a97b4f.
//
// Solidity: function BalanceK() returns()
func (_AMM *AMMTransactor) BalanceK(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "BalanceK")
}

// BalanceK is a paid mutator transaction binding the contract method 0x29a97b4f.
//
// Solidity: function BalanceK() returns()
func (_AMM *AMMSession) BalanceK() (*types.Transaction, error) {
	return _AMM.Contract.BalanceK(&_AMM.TransactOpts)
}

// BalanceK is a paid mutator transaction binding the contract method 0x29a97b4f.
//
// Solidity: function BalanceK() returns()
func (_AMM *AMMTransactorSession) BalanceK() (*types.Transaction, error) {
	return _AMM.Contract.BalanceK(&_AMM.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x88a84ea6.
//
// Solidity: function Settle(uint256 requestId, uint256 persentage) returns()
func (_AMM *AMMTransactor) Settle(opts *bind.TransactOpts, requestId *big.Int, persentage *big.Int) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "Settle", requestId, persentage)
}

// Settle is a paid mutator transaction binding the contract method 0x88a84ea6.
//
// Solidity: function Settle(uint256 requestId, uint256 persentage) returns()
func (_AMM *AMMSession) Settle(requestId *big.Int, persentage *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.Settle(&_AMM.TransactOpts, requestId, persentage)
}

// Settle is a paid mutator transaction binding the contract method 0x88a84ea6.
//
// Solidity: function Settle(uint256 requestId, uint256 persentage) returns()
func (_AMM *AMMTransactorSession) Settle(requestId *big.Int, persentage *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.Settle(&_AMM.TransactOpts, requestId, persentage)
}

// Swap is a paid mutator transaction binding the contract method 0x503e0988.
//
// Solidity: function Swap(uint256 _targetTokenAmount, uint256 _basicTokenAmount, uint8 _swapType) returns(uint256 requestId)
func (_AMM *AMMTransactor) Swap(opts *bind.TransactOpts, _targetTokenAmount *big.Int, _basicTokenAmount *big.Int, _swapType uint8) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "Swap", _targetTokenAmount, _basicTokenAmount, _swapType)
}

// Swap is a paid mutator transaction binding the contract method 0x503e0988.
//
// Solidity: function Swap(uint256 _targetTokenAmount, uint256 _basicTokenAmount, uint8 _swapType) returns(uint256 requestId)
func (_AMM *AMMSession) Swap(_targetTokenAmount *big.Int, _basicTokenAmount *big.Int, _swapType uint8) (*types.Transaction, error) {
	return _AMM.Contract.Swap(&_AMM.TransactOpts, _targetTokenAmount, _basicTokenAmount, _swapType)
}

// Swap is a paid mutator transaction binding the contract method 0x503e0988.
//
// Solidity: function Swap(uint256 _targetTokenAmount, uint256 _basicTokenAmount, uint8 _swapType) returns(uint256 requestId)
func (_AMM *AMMTransactorSession) Swap(_targetTokenAmount *big.Int, _basicTokenAmount *big.Int, _swapType uint8) (*types.Transaction, error) {
	return _AMM.Contract.Swap(&_AMM.TransactOpts, _targetTokenAmount, _basicTokenAmount, _swapType)
}

// UpdateConf is a paid mutator transaction binding the contract method 0xf0cba149.
//
// Solidity: function UpdateConf(uint256 _win_optional_probability, uint256 _optional_down_bound, uint256 _optional_up_bound, uint256 _market_swap_rate) returns()
func (_AMM *AMMTransactor) UpdateConf(opts *bind.TransactOpts, _win_optional_probability *big.Int, _optional_down_bound *big.Int, _optional_up_bound *big.Int, _market_swap_rate *big.Int) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "UpdateConf", _win_optional_probability, _optional_down_bound, _optional_up_bound, _market_swap_rate)
}

// UpdateConf is a paid mutator transaction binding the contract method 0xf0cba149.
//
// Solidity: function UpdateConf(uint256 _win_optional_probability, uint256 _optional_down_bound, uint256 _optional_up_bound, uint256 _market_swap_rate) returns()
func (_AMM *AMMSession) UpdateConf(_win_optional_probability *big.Int, _optional_down_bound *big.Int, _optional_up_bound *big.Int, _market_swap_rate *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.UpdateConf(&_AMM.TransactOpts, _win_optional_probability, _optional_down_bound, _optional_up_bound, _market_swap_rate)
}

// UpdateConf is a paid mutator transaction binding the contract method 0xf0cba149.
//
// Solidity: function UpdateConf(uint256 _win_optional_probability, uint256 _optional_down_bound, uint256 _optional_up_bound, uint256 _market_swap_rate) returns()
func (_AMM *AMMTransactorSession) UpdateConf(_win_optional_probability *big.Int, _optional_down_bound *big.Int, _optional_up_bound *big.Int, _market_swap_rate *big.Int) (*types.Transaction, error) {
	return _AMM.Contract.UpdateConf(&_AMM.TransactOpts, _win_optional_probability, _optional_down_bound, _optional_up_bound, _market_swap_rate)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AMM *AMMTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AMM *AMMSession) AcceptOwnership() (*types.Transaction, error) {
	return _AMM.Contract.AcceptOwnership(&_AMM.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AMM *AMMTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AMM.Contract.AcceptOwnership(&_AMM.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AMM *AMMTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AMM *AMMSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.Contract.GrantRole(&_AMM.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AMM *AMMTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.Contract.GrantRole(&_AMM.TransactOpts, role, account)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_AMM *AMMTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_AMM *AMMSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _AMM.Contract.RawFulfillRandomWords(&_AMM.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_AMM *AMMTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _AMM.Contract.RawFulfillRandomWords(&_AMM.TransactOpts, requestId, randomWords)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AMM *AMMTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AMM *AMMSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.Contract.RenounceRole(&_AMM.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AMM *AMMTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.Contract.RenounceRole(&_AMM.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AMM *AMMTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AMM *AMMSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.Contract.RevokeRole(&_AMM.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AMM *AMMTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AMM.Contract.RevokeRole(&_AMM.TransactOpts, role, account)
}

// SetCoordinator is a paid mutator transaction binding the contract method 0x8ea98117.
//
// Solidity: function setCoordinator(address _vrfCoordinator) returns()
func (_AMM *AMMTransactor) SetCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "setCoordinator", _vrfCoordinator)
}

// SetCoordinator is a paid mutator transaction binding the contract method 0x8ea98117.
//
// Solidity: function setCoordinator(address _vrfCoordinator) returns()
func (_AMM *AMMSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _AMM.Contract.SetCoordinator(&_AMM.TransactOpts, _vrfCoordinator)
}

// SetCoordinator is a paid mutator transaction binding the contract method 0x8ea98117.
//
// Solidity: function setCoordinator(address _vrfCoordinator) returns()
func (_AMM *AMMTransactorSession) SetCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _AMM.Contract.SetCoordinator(&_AMM.TransactOpts, _vrfCoordinator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AMM *AMMTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AMM.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AMM *AMMSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AMM.Contract.TransferOwnership(&_AMM.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AMM *AMMTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AMM.Contract.TransferOwnership(&_AMM.TransactOpts, to)
}

// AMMCoordinatorSetIterator is returned from FilterCoordinatorSet and is used to iterate over the raw logs and unpacked data for CoordinatorSet events raised by the AMM contract.
type AMMCoordinatorSetIterator struct {
	Event *AMMCoordinatorSet // Event containing the contract specifics and raw log

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
func (it *AMMCoordinatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMCoordinatorSet)
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
		it.Event = new(AMMCoordinatorSet)
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
func (it *AMMCoordinatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMCoordinatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMCoordinatorSet represents a CoordinatorSet event raised by the AMM contract.
type AMMCoordinatorSet struct {
	VrfCoordinator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCoordinatorSet is a free log retrieval operation binding the contract event 0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6.
//
// Solidity: event CoordinatorSet(address vrfCoordinator)
func (_AMM *AMMFilterer) FilterCoordinatorSet(opts *bind.FilterOpts) (*AMMCoordinatorSetIterator, error) {

	logs, sub, err := _AMM.contract.FilterLogs(opts, "CoordinatorSet")
	if err != nil {
		return nil, err
	}
	return &AMMCoordinatorSetIterator{contract: _AMM.contract, event: "CoordinatorSet", logs: logs, sub: sub}, nil
}

// WatchCoordinatorSet is a free log subscription operation binding the contract event 0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6.
//
// Solidity: event CoordinatorSet(address vrfCoordinator)
func (_AMM *AMMFilterer) WatchCoordinatorSet(opts *bind.WatchOpts, sink chan<- *AMMCoordinatorSet) (event.Subscription, error) {

	logs, sub, err := _AMM.contract.WatchLogs(opts, "CoordinatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMCoordinatorSet)
				if err := _AMM.contract.UnpackLog(event, "CoordinatorSet", log); err != nil {
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

// ParseCoordinatorSet is a log parse operation binding the contract event 0xd1a6a14209a385a964d036e404cb5cfb71f4000cdb03c9366292430787261be6.
//
// Solidity: event CoordinatorSet(address vrfCoordinator)
func (_AMM *AMMFilterer) ParseCoordinatorSet(log types.Log) (*AMMCoordinatorSet, error) {
	event := new(AMMCoordinatorSet)
	if err := _AMM.contract.UnpackLog(event, "CoordinatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMGetRandomsIterator is returned from FilterGetRandoms and is used to iterate over the raw logs and unpacked data for GetRandoms events raised by the AMM contract.
type AMMGetRandomsIterator struct {
	Event *AMMGetRandoms // Event containing the contract specifics and raw log

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
func (it *AMMGetRandomsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMGetRandoms)
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
		it.Event = new(AMMGetRandoms)
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
func (it *AMMGetRandomsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMGetRandomsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMGetRandoms represents a GetRandoms event raised by the AMM contract.
type AMMGetRandoms struct {
	RequestId   *big.Int
	Randomness  *big.Int
	User        common.Address
	SellType    uint8
	BasicToken  *big.Int
	TargetToken *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGetRandoms is a free log retrieval operation binding the contract event 0xa4d91522ca7b1c3b07cc2a7e606812f0953daed38ec14f873c47e11aa255f89f.
//
// Solidity: event GetRandoms(uint256 indexed requestId, uint256 indexed randomness, address user, uint8 sellType, uint256 basicToken, uint256 targetToken)
func (_AMM *AMMFilterer) FilterGetRandoms(opts *bind.FilterOpts, requestId []*big.Int, randomness []*big.Int) (*AMMGetRandomsIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var randomnessRule []interface{}
	for _, randomnessItem := range randomness {
		randomnessRule = append(randomnessRule, randomnessItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "GetRandoms", requestIdRule, randomnessRule)
	if err != nil {
		return nil, err
	}
	return &AMMGetRandomsIterator{contract: _AMM.contract, event: "GetRandoms", logs: logs, sub: sub}, nil
}

// WatchGetRandoms is a free log subscription operation binding the contract event 0xa4d91522ca7b1c3b07cc2a7e606812f0953daed38ec14f873c47e11aa255f89f.
//
// Solidity: event GetRandoms(uint256 indexed requestId, uint256 indexed randomness, address user, uint8 sellType, uint256 basicToken, uint256 targetToken)
func (_AMM *AMMFilterer) WatchGetRandoms(opts *bind.WatchOpts, sink chan<- *AMMGetRandoms, requestId []*big.Int, randomness []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var randomnessRule []interface{}
	for _, randomnessItem := range randomness {
		randomnessRule = append(randomnessRule, randomnessItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "GetRandoms", requestIdRule, randomnessRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMGetRandoms)
				if err := _AMM.contract.UnpackLog(event, "GetRandoms", log); err != nil {
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

// ParseGetRandoms is a log parse operation binding the contract event 0xa4d91522ca7b1c3b07cc2a7e606812f0953daed38ec14f873c47e11aa255f89f.
//
// Solidity: event GetRandoms(uint256 indexed requestId, uint256 indexed randomness, address user, uint8 sellType, uint256 basicToken, uint256 targetToken)
func (_AMM *AMMFilterer) ParseGetRandoms(log types.Log) (*AMMGetRandoms, error) {
	event := new(AMMGetRandoms)
	if err := _AMM.contract.UnpackLog(event, "GetRandoms", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the AMM contract.
type AMMOwnershipTransferRequestedIterator struct {
	Event *AMMOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *AMMOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMOwnershipTransferRequested)
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
		it.Event = new(AMMOwnershipTransferRequested)
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
func (it *AMMOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the AMM contract.
type AMMOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AMM *AMMFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AMMOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AMMOwnershipTransferRequestedIterator{contract: _AMM.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AMM *AMMFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AMMOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMOwnershipTransferRequested)
				if err := _AMM.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AMM *AMMFilterer) ParseOwnershipTransferRequested(log types.Log) (*AMMOwnershipTransferRequested, error) {
	event := new(AMMOwnershipTransferRequested)
	if err := _AMM.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AMM contract.
type AMMOwnershipTransferredIterator struct {
	Event *AMMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AMMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMOwnershipTransferred)
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
		it.Event = new(AMMOwnershipTransferred)
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
func (it *AMMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMOwnershipTransferred represents a OwnershipTransferred event raised by the AMM contract.
type AMMOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AMM *AMMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AMMOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AMMOwnershipTransferredIterator{contract: _AMM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AMM *AMMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AMMOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMOwnershipTransferred)
				if err := _AMM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AMM *AMMFilterer) ParseOwnershipTransferred(log types.Log) (*AMMOwnershipTransferred, error) {
	event := new(AMMOwnershipTransferred)
	if err := _AMM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AMM contract.
type AMMRoleAdminChangedIterator struct {
	Event *AMMRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AMMRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMRoleAdminChanged)
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
		it.Event = new(AMMRoleAdminChanged)
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
func (it *AMMRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMRoleAdminChanged represents a RoleAdminChanged event raised by the AMM contract.
type AMMRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AMM *AMMFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AMMRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AMMRoleAdminChangedIterator{contract: _AMM.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AMM *AMMFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AMMRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMRoleAdminChanged)
				if err := _AMM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AMM *AMMFilterer) ParseRoleAdminChanged(log types.Log) (*AMMRoleAdminChanged, error) {
	event := new(AMMRoleAdminChanged)
	if err := _AMM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AMM contract.
type AMMRoleGrantedIterator struct {
	Event *AMMRoleGranted // Event containing the contract specifics and raw log

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
func (it *AMMRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMRoleGranted)
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
		it.Event = new(AMMRoleGranted)
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
func (it *AMMRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMRoleGranted represents a RoleGranted event raised by the AMM contract.
type AMMRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AMM *AMMFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AMMRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AMMRoleGrantedIterator{contract: _AMM.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AMM *AMMFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AMMRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMRoleGranted)
				if err := _AMM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AMM *AMMFilterer) ParseRoleGranted(log types.Log) (*AMMRoleGranted, error) {
	event := new(AMMRoleGranted)
	if err := _AMM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AMM contract.
type AMMRoleRevokedIterator struct {
	Event *AMMRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AMMRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMRoleRevoked)
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
		it.Event = new(AMMRoleRevoked)
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
func (it *AMMRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMRoleRevoked represents a RoleRevoked event raised by the AMM contract.
type AMMRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AMM *AMMFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AMMRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AMMRoleRevokedIterator{contract: _AMM.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AMM *AMMFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AMMRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMRoleRevoked)
				if err := _AMM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AMM *AMMFilterer) ParseRoleRevoked(log types.Log) (*AMMRoleRevoked, error) {
	event := new(AMMRoleRevoked)
	if err := _AMM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AMMSwapSuceesfullyIterator is returned from FilterSwapSuceesfully and is used to iterate over the raw logs and unpacked data for SwapSuceesfully events raised by the AMM contract.
type AMMSwapSuceesfullyIterator struct {
	Event *AMMSwapSuceesfully // Event containing the contract specifics and raw log

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
func (it *AMMSwapSuceesfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AMMSwapSuceesfully)
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
		it.Event = new(AMMSwapSuceesfully)
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
func (it *AMMSwapSuceesfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AMMSwapSuceesfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AMMSwapSuceesfully represents a SwapSuceesfully event raised by the AMM contract.
type AMMSwapSuceesfully struct {
	RequestId           *big.Int
	User                common.Address
	InTokenAmount       *big.Int
	SwapType            uint8
	PreOutTokenAmount   *big.Int
	FinalOutTokenAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSwapSuceesfully is a free log retrieval operation binding the contract event 0xb2395e47ade8f8160853da862b40fe88f5e09e5e7b88f2736ffc25f1e5905083.
//
// Solidity: event SwapSuceesfully(uint256 indexed requestId, address indexed user, uint256 InTokenAmount, uint8 swapType, uint256 preOutTokenAmount, uint256 finalOutTokenAmount)
func (_AMM *AMMFilterer) FilterSwapSuceesfully(opts *bind.FilterOpts, requestId []*big.Int, user []common.Address) (*AMMSwapSuceesfullyIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AMM.contract.FilterLogs(opts, "SwapSuceesfully", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AMMSwapSuceesfullyIterator{contract: _AMM.contract, event: "SwapSuceesfully", logs: logs, sub: sub}, nil
}

// WatchSwapSuceesfully is a free log subscription operation binding the contract event 0xb2395e47ade8f8160853da862b40fe88f5e09e5e7b88f2736ffc25f1e5905083.
//
// Solidity: event SwapSuceesfully(uint256 indexed requestId, address indexed user, uint256 InTokenAmount, uint8 swapType, uint256 preOutTokenAmount, uint256 finalOutTokenAmount)
func (_AMM *AMMFilterer) WatchSwapSuceesfully(opts *bind.WatchOpts, sink chan<- *AMMSwapSuceesfully, requestId []*big.Int, user []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AMM.contract.WatchLogs(opts, "SwapSuceesfully", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AMMSwapSuceesfully)
				if err := _AMM.contract.UnpackLog(event, "SwapSuceesfully", log); err != nil {
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

// ParseSwapSuceesfully is a log parse operation binding the contract event 0xb2395e47ade8f8160853da862b40fe88f5e09e5e7b88f2736ffc25f1e5905083.
//
// Solidity: event SwapSuceesfully(uint256 indexed requestId, address indexed user, uint256 InTokenAmount, uint8 swapType, uint256 preOutTokenAmount, uint256 finalOutTokenAmount)
func (_AMM *AMMFilterer) ParseSwapSuceesfully(log types.Log) (*AMMSwapSuceesfully, error) {
	event := new(AMMSwapSuceesfully)
	if err := _AMM.contract.UnpackLog(event, "SwapSuceesfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
