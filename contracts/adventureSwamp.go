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

// AdventureSwampMetaData contains all meta data concerning the AdventureSwamp contract.
var AdventureSwampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vefly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hopper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NoHopperStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnfitHopper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongTokenID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"}],\"name\":\"UpdatedBallot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionRate\",\"type\":\"uint256\"}],\"name\":\"UpdatedEmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOPPER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEVEL_GAUGE_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"baseSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEmissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"}],\"name\":\"canEnter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"changeHopperName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"emergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flyLevelCapRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"forceUnvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"generatedPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHopperAndGauge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"hopperGauge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gaugeLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLevel\",\"type\":\"uint256\"}],\"name\":\"getLevelUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBonusShares\",\"type\":\"uint256\"}],\"name\":\"getUserBonusGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBaseShares\",\"type\":\"uint256\"}],\"name\":\"getUserGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBonusUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"levelUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"}],\"name\":\"setBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"setBonusEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emissionRate\",\"type\":\"uint256\"}],\"name\":\"setEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flyLevelCapRatio\",\"type\":\"uint256\"}],\"name\":\"setFlyLevelCapRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenCapFilledPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBonusRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMaxFlyGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veFlyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdventureSwampABI is the input ABI used to generate the binding from.
// Deprecated: Use AdventureSwampMetaData.ABI instead.
var AdventureSwampABI = AdventureSwampMetaData.ABI

// AdventureSwamp is an auto generated Go binding around an Ethereum contract.
type AdventureSwamp struct {
	AdventureSwampCaller     // Read-only binding to the contract
	AdventureSwampTransactor // Write-only binding to the contract
	AdventureSwampFilterer   // Log filterer for contract events
}

// AdventureSwampCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdventureSwampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureSwampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdventureSwampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureSwampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdventureSwampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureSwampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdventureSwampSession struct {
	Contract     *AdventureSwamp   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdventureSwampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdventureSwampCallerSession struct {
	Contract *AdventureSwampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AdventureSwampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdventureSwampTransactorSession struct {
	Contract     *AdventureSwampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AdventureSwampRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdventureSwampRaw struct {
	Contract *AdventureSwamp // Generic contract binding to access the raw methods on
}

// AdventureSwampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdventureSwampCallerRaw struct {
	Contract *AdventureSwampCaller // Generic read-only contract binding to access the raw methods on
}

// AdventureSwampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdventureSwampTransactorRaw struct {
	Contract *AdventureSwampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdventureSwamp creates a new instance of AdventureSwamp, bound to a specific deployed contract.
func NewAdventureSwamp(address common.Address, backend bind.ContractBackend) (*AdventureSwamp, error) {
	contract, err := bindAdventureSwamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdventureSwamp{AdventureSwampCaller: AdventureSwampCaller{contract: contract}, AdventureSwampTransactor: AdventureSwampTransactor{contract: contract}, AdventureSwampFilterer: AdventureSwampFilterer{contract: contract}}, nil
}

// NewAdventureSwampCaller creates a new read-only instance of AdventureSwamp, bound to a specific deployed contract.
func NewAdventureSwampCaller(address common.Address, caller bind.ContractCaller) (*AdventureSwampCaller, error) {
	contract, err := bindAdventureSwamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureSwampCaller{contract: contract}, nil
}

// NewAdventureSwampTransactor creates a new write-only instance of AdventureSwamp, bound to a specific deployed contract.
func NewAdventureSwampTransactor(address common.Address, transactor bind.ContractTransactor) (*AdventureSwampTransactor, error) {
	contract, err := bindAdventureSwamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureSwampTransactor{contract: contract}, nil
}

// NewAdventureSwampFilterer creates a new log filterer instance of AdventureSwamp, bound to a specific deployed contract.
func NewAdventureSwampFilterer(address common.Address, filterer bind.ContractFilterer) (*AdventureSwampFilterer, error) {
	contract, err := bindAdventureSwamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdventureSwampFilterer{contract: contract}, nil
}

// bindAdventureSwamp binds a generic wrapper to an already deployed contract.
func bindAdventureSwamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdventureSwampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureSwamp *AdventureSwampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureSwamp.Contract.AdventureSwampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureSwamp *AdventureSwampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.AdventureSwampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureSwamp *AdventureSwampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.AdventureSwampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureSwamp *AdventureSwampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureSwamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureSwamp *AdventureSwampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureSwamp *AdventureSwampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.contract.Transact(opts, method, params...)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureSwamp *AdventureSwampCaller) FLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureSwamp *AdventureSwampSession) FLY() (common.Address, error) {
	return _AdventureSwamp.Contract.FLY(&_AdventureSwamp.CallOpts)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureSwamp *AdventureSwampCallerSession) FLY() (common.Address, error) {
	return _AdventureSwamp.Contract.FLY(&_AdventureSwamp.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureSwamp *AdventureSwampCaller) HOPPER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "HOPPER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureSwamp *AdventureSwampSession) HOPPER() (common.Address, error) {
	return _AdventureSwamp.Contract.HOPPER(&_AdventureSwamp.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureSwamp *AdventureSwampCallerSession) HOPPER() (common.Address, error) {
	return _AdventureSwamp.Contract.HOPPER(&_AdventureSwamp.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureSwamp *AdventureSwampCaller) LEVELGAUGEKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "LEVEL_GAUGE_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureSwamp *AdventureSwampSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureSwamp.Contract.LEVELGAUGEKEY(&_AdventureSwamp.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureSwamp *AdventureSwampCallerSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureSwamp.Contract.LEVELGAUGEKEY(&_AdventureSwamp.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureSwamp *AdventureSwampCaller) VEFLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "VE_FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureSwamp *AdventureSwampSession) VEFLY() (common.Address, error) {
	return _AdventureSwamp.Contract.VEFLY(&_AdventureSwamp.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureSwamp *AdventureSwampCallerSession) VEFLY() (common.Address, error) {
	return _AdventureSwamp.Contract.VEFLY(&_AdventureSwamp.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureSwamp *AdventureSwampCaller) Ballot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "ballot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureSwamp *AdventureSwampSession) Ballot() (common.Address, error) {
	return _AdventureSwamp.Contract.Ballot(&_AdventureSwamp.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureSwamp *AdventureSwampCallerSession) Ballot() (common.Address, error) {
	return _AdventureSwamp.Contract.Ballot(&_AdventureSwamp.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) BaseRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "baseRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.BaseRewardPerShare(&_AdventureSwamp.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.BaseRewardPerShare(&_AdventureSwamp.CallOpts)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) BaseSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "baseSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.BaseSharesBalance(&_AdventureSwamp.CallOpts, arg0)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.BaseSharesBalance(&_AdventureSwamp.CallOpts, arg0)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) BonusEmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "bonusEmissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureSwamp.Contract.BonusEmissionRate(&_AdventureSwamp.CallOpts)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureSwamp.Contract.BonusEmissionRate(&_AdventureSwamp.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) BonusRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "bonusRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.BonusRewardPerShare(&_AdventureSwamp.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.BonusRewardPerShare(&_AdventureSwamp.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) BonusRewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "bonusRewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureSwamp.Contract.BonusRewardPerShareStored(&_AdventureSwamp.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureSwamp.Contract.BonusRewardPerShareStored(&_AdventureSwamp.CallOpts)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureSwamp *AdventureSwampCaller) CanEnter(opts *bind.CallOpts, hopper HopperNFTHopper) (bool, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "canEnter", hopper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureSwamp *AdventureSwampSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureSwamp.Contract.CanEnter(&_AdventureSwamp.CallOpts, hopper)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureSwamp *AdventureSwampCallerSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureSwamp.Contract.CanEnter(&_AdventureSwamp.CallOpts, hopper)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.Claimable(&_AdventureSwamp.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.Claimable(&_AdventureSwamp.CallOpts, _account)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureSwamp *AdventureSwampCaller) Emergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "emergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureSwamp *AdventureSwampSession) Emergency() (bool, error) {
	return _AdventureSwamp.Contract.Emergency(&_AdventureSwamp.CallOpts)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureSwamp *AdventureSwampCallerSession) Emergency() (bool, error) {
	return _AdventureSwamp.Contract.Emergency(&_AdventureSwamp.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) EmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "emissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) EmissionRate() (*big.Int, error) {
	return _AdventureSwamp.Contract.EmissionRate(&_AdventureSwamp.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) EmissionRate() (*big.Int, error) {
	return _AdventureSwamp.Contract.EmissionRate(&_AdventureSwamp.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) FlyLevelCapRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "flyLevelCapRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureSwamp.Contract.FlyLevelCapRatio(&_AdventureSwamp.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureSwamp.Contract.FlyLevelCapRatio(&_AdventureSwamp.CallOpts)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) GeneratedPerShareStored(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "generatedPerShareStored", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.GeneratedPerShareStored(&_AdventureSwamp.CallOpts, arg0)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.GeneratedPerShareStored(&_AdventureSwamp.CallOpts, arg0)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureSwamp *AdventureSwampCaller) GetHopperAndGauge(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "getHopperAndGauge", tokenId)

	outstruct := new(struct {
		Hopper      HopperNFTHopper
		HopperGauge *big.Int
		GaugeLimit  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hopper = *abi.ConvertType(out[0], new(HopperNFTHopper)).(*HopperNFTHopper)
	outstruct.HopperGauge = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GaugeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureSwamp *AdventureSwampSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureSwamp.Contract.GetHopperAndGauge(&_AdventureSwamp.CallOpts, tokenId)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureSwamp *AdventureSwampCallerSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureSwamp.Contract.GetHopperAndGauge(&_AdventureSwamp.CallOpts, tokenId)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) GetLevelUpCost(opts *bind.CallOpts, currentLevel *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "getLevelUpCost", currentLevel)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureSwamp.Contract.GetLevelUpCost(&_AdventureSwamp.CallOpts, currentLevel)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureSwamp.Contract.GetLevelUpCost(&_AdventureSwamp.CallOpts, currentLevel)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureSwamp *AdventureSwampCaller) GetUserBonusGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "getUserBonusGeneratedFly", account, _totalUserBonusShares)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureSwamp *AdventureSwampSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureSwamp.Contract.GetUserBonusGeneratedFly(&_AdventureSwamp.CallOpts, account, _totalUserBonusShares)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureSwamp.Contract.GetUserBonusGeneratedFly(&_AdventureSwamp.CallOpts, account, _totalUserBonusShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureSwamp *AdventureSwampCaller) GetUserGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "getUserGeneratedFly", account, _totalUserBaseShares)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureSwamp *AdventureSwampSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureSwamp.Contract.GetUserGeneratedFly(&_AdventureSwamp.CallOpts, account, _totalUserBaseShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureSwamp.Contract.GetUserGeneratedFly(&_AdventureSwamp.CallOpts, account, _totalUserBaseShares)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) HopperBaseShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "hopperBaseShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureSwamp.Contract.HopperBaseShare(&_AdventureSwamp.CallOpts, arg0)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureSwamp.Contract.HopperBaseShare(&_AdventureSwamp.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureSwamp *AdventureSwampCaller) HopperOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "hopperOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureSwamp *AdventureSwampSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureSwamp.Contract.HopperOwners(&_AdventureSwamp.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureSwamp *AdventureSwampCallerSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureSwamp.Contract.HopperOwners(&_AdventureSwamp.CallOpts, arg0)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) LastBonusUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "lastBonusUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureSwamp.Contract.LastBonusUpdatedTime(&_AdventureSwamp.CallOpts)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureSwamp.Contract.LastBonusUpdatedTime(&_AdventureSwamp.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) LastUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "lastUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureSwamp.Contract.LastUpdatedTime(&_AdventureSwamp.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureSwamp.Contract.LastUpdatedTime(&_AdventureSwamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureSwamp *AdventureSwampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureSwamp *AdventureSwampSession) Owner() (common.Address, error) {
	return _AdventureSwamp.Contract.Owner(&_AdventureSwamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureSwamp *AdventureSwampCallerSession) Owner() (common.Address, error) {
	return _AdventureSwamp.Contract.Owner(&_AdventureSwamp.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) RewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "rewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureSwamp.Contract.RewardPerShareStored(&_AdventureSwamp.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureSwamp.Contract.RewardPerShareStored(&_AdventureSwamp.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.Rewards(&_AdventureSwamp.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.Rewards(&_AdventureSwamp.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) TokenCapFilledPerShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "tokenCapFilledPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureSwamp.Contract.TokenCapFilledPerShare(&_AdventureSwamp.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureSwamp.Contract.TokenCapFilledPerShare(&_AdventureSwamp.CallOpts, arg0)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) TotalBaseShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "totalBaseShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.TotalBaseShare(&_AdventureSwamp.CallOpts)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.TotalBaseShare(&_AdventureSwamp.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) TotalVeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "totalVeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) TotalVeShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.TotalVeShare(&_AdventureSwamp.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) TotalVeShare() (*big.Int, error) {
	return _AdventureSwamp.Contract.TotalVeShare(&_AdventureSwamp.CallOpts)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) UserBonusRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "userBonusRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.UserBonusRewardPerSharePaid(&_AdventureSwamp.CallOpts, arg0)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.UserBonusRewardPerSharePaid(&_AdventureSwamp.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) UserMaxFlyGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "userMaxFlyGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.UserMaxFlyGeneration(&_AdventureSwamp.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.UserMaxFlyGeneration(&_AdventureSwamp.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) UserRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "userRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.UserRewardPerSharePaid(&_AdventureSwamp.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.UserRewardPerSharePaid(&_AdventureSwamp.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) VeFlyBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "veFlyBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.VeFlyBalance(&_AdventureSwamp.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.VeFlyBalance(&_AdventureSwamp.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCaller) VeSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureSwamp.contract.Call(opts, &out, "veSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.VeSharesBalance(&_AdventureSwamp.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureSwamp *AdventureSwampCallerSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureSwamp.Contract.VeSharesBalance(&_AdventureSwamp.CallOpts, arg0)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureSwamp *AdventureSwampTransactor) ChangeHopperName(opts *bind.TransactOpts, tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "changeHopperName", tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureSwamp *AdventureSwampSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.ChangeHopperName(&_AdventureSwamp.TransactOpts, tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.ChangeHopperName(&_AdventureSwamp.TransactOpts, tokenId, name, useOwnRewards)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureSwamp *AdventureSwampTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureSwamp *AdventureSwampSession) Claim() (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Claim(&_AdventureSwamp.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) Claim() (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Claim(&_AdventureSwamp.TransactOpts)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureSwamp *AdventureSwampTransactor) EmergencyExit(opts *bind.TransactOpts, tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "emergencyExit", tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureSwamp *AdventureSwampSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.EmergencyExit(&_AdventureSwamp.TransactOpts, tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.EmergencyExit(&_AdventureSwamp.TransactOpts, tokenIds, user)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureSwamp *AdventureSwampTransactor) EnableEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "enableEmergency")
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureSwamp *AdventureSwampSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureSwamp.Contract.EnableEmergency(&_AdventureSwamp.TransactOpts)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureSwamp.Contract.EnableEmergency(&_AdventureSwamp.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureSwamp *AdventureSwampTransactor) Enter(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "enter", tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureSwamp *AdventureSwampSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Enter(&_AdventureSwamp.TransactOpts, tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Enter(&_AdventureSwamp.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureSwamp *AdventureSwampTransactor) Exit(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "exit", tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureSwamp *AdventureSwampSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Exit(&_AdventureSwamp.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Exit(&_AdventureSwamp.TransactOpts, tokenIds)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureSwamp *AdventureSwampTransactor) ForceUnvote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "forceUnvote", user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureSwamp *AdventureSwampSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.ForceUnvote(&_AdventureSwamp.TransactOpts, user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.ForceUnvote(&_AdventureSwamp.TransactOpts, user)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureSwamp *AdventureSwampTransactor) LevelUp(opts *bind.TransactOpts, tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "levelUp", tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureSwamp *AdventureSwampSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.LevelUp(&_AdventureSwamp.TransactOpts, tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.LevelUp(&_AdventureSwamp.TransactOpts, tokenId, useOwnRewards)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureSwamp *AdventureSwampTransactor) SetBallot(opts *bind.TransactOpts, _ballot common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "setBallot", _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureSwamp *AdventureSwampSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetBallot(&_AdventureSwamp.TransactOpts, _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetBallot(&_AdventureSwamp.TransactOpts, _ballot)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureSwamp *AdventureSwampTransactor) SetBonusEmissionRate(opts *bind.TransactOpts, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "setBonusEmissionRate", _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureSwamp *AdventureSwampSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetBonusEmissionRate(&_AdventureSwamp.TransactOpts, _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetBonusEmissionRate(&_AdventureSwamp.TransactOpts, _bonusEmissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureSwamp *AdventureSwampTransactor) SetEmissionRate(opts *bind.TransactOpts, _emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "setEmissionRate", _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureSwamp *AdventureSwampSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetEmissionRate(&_AdventureSwamp.TransactOpts, _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetEmissionRate(&_AdventureSwamp.TransactOpts, _emissionRate)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureSwamp *AdventureSwampTransactor) SetFlyLevelCapRatio(opts *bind.TransactOpts, _flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "setFlyLevelCapRatio", _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureSwamp *AdventureSwampSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetFlyLevelCapRatio(&_AdventureSwamp.TransactOpts, _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetFlyLevelCapRatio(&_AdventureSwamp.TransactOpts, _flyLevelCapRatio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureSwamp *AdventureSwampTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureSwamp *AdventureSwampSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetOwner(&_AdventureSwamp.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.SetOwner(&_AdventureSwamp.TransactOpts, _owner)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureSwamp *AdventureSwampTransactor) Unvote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "unvote", veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureSwamp *AdventureSwampSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Unvote(&_AdventureSwamp.TransactOpts, veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Unvote(&_AdventureSwamp.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureSwamp *AdventureSwampTransactor) Vote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureSwamp.contract.Transact(opts, "vote", veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureSwamp *AdventureSwampSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Vote(&_AdventureSwamp.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureSwamp *AdventureSwampTransactorSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureSwamp.Contract.Vote(&_AdventureSwamp.TransactOpts, veFlyAmount, recount)
}

// AdventureSwampUpdatedBallotIterator is returned from FilterUpdatedBallot and is used to iterate over the raw logs and unpacked data for UpdatedBallot events raised by the AdventureSwamp contract.
type AdventureSwampUpdatedBallotIterator struct {
	Event *AdventureSwampUpdatedBallot // Event containing the contract specifics and raw log

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
func (it *AdventureSwampUpdatedBallotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureSwampUpdatedBallot)
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
		it.Event = new(AdventureSwampUpdatedBallot)
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
func (it *AdventureSwampUpdatedBallotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureSwampUpdatedBallotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureSwampUpdatedBallot represents a UpdatedBallot event raised by the AdventureSwamp contract.
type AdventureSwampUpdatedBallot struct {
	Ballot common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBallot is a free log retrieval operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureSwamp *AdventureSwampFilterer) FilterUpdatedBallot(opts *bind.FilterOpts, ballot []common.Address) (*AdventureSwampUpdatedBallotIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureSwamp.contract.FilterLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return &AdventureSwampUpdatedBallotIterator{contract: _AdventureSwamp.contract, event: "UpdatedBallot", logs: logs, sub: sub}, nil
}

// WatchUpdatedBallot is a free log subscription operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureSwamp *AdventureSwampFilterer) WatchUpdatedBallot(opts *bind.WatchOpts, sink chan<- *AdventureSwampUpdatedBallot, ballot []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureSwamp.contract.WatchLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureSwampUpdatedBallot)
				if err := _AdventureSwamp.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
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

// ParseUpdatedBallot is a log parse operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureSwamp *AdventureSwampFilterer) ParseUpdatedBallot(log types.Log) (*AdventureSwampUpdatedBallot, error) {
	event := new(AdventureSwampUpdatedBallot)
	if err := _AdventureSwamp.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureSwampUpdatedEmissionIterator is returned from FilterUpdatedEmission and is used to iterate over the raw logs and unpacked data for UpdatedEmission events raised by the AdventureSwamp contract.
type AdventureSwampUpdatedEmissionIterator struct {
	Event *AdventureSwampUpdatedEmission // Event containing the contract specifics and raw log

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
func (it *AdventureSwampUpdatedEmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureSwampUpdatedEmission)
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
		it.Event = new(AdventureSwampUpdatedEmission)
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
func (it *AdventureSwampUpdatedEmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureSwampUpdatedEmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureSwampUpdatedEmission represents a UpdatedEmission event raised by the AdventureSwamp contract.
type AdventureSwampUpdatedEmission struct {
	EmissionRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEmission is a free log retrieval operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureSwamp *AdventureSwampFilterer) FilterUpdatedEmission(opts *bind.FilterOpts) (*AdventureSwampUpdatedEmissionIterator, error) {

	logs, sub, err := _AdventureSwamp.contract.FilterLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return &AdventureSwampUpdatedEmissionIterator{contract: _AdventureSwamp.contract, event: "UpdatedEmission", logs: logs, sub: sub}, nil
}

// WatchUpdatedEmission is a free log subscription operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureSwamp *AdventureSwampFilterer) WatchUpdatedEmission(opts *bind.WatchOpts, sink chan<- *AdventureSwampUpdatedEmission) (event.Subscription, error) {

	logs, sub, err := _AdventureSwamp.contract.WatchLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureSwampUpdatedEmission)
				if err := _AdventureSwamp.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
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

// ParseUpdatedEmission is a log parse operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureSwamp *AdventureSwampFilterer) ParseUpdatedEmission(log types.Log) (*AdventureSwampUpdatedEmission, error) {
	event := new(AdventureSwampUpdatedEmission)
	if err := _AdventureSwamp.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureSwampUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the AdventureSwamp contract.
type AdventureSwampUpdatedOwnerIterator struct {
	Event *AdventureSwampUpdatedOwner // Event containing the contract specifics and raw log

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
func (it *AdventureSwampUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureSwampUpdatedOwner)
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
		it.Event = new(AdventureSwampUpdatedOwner)
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
func (it *AdventureSwampUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureSwampUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureSwampUpdatedOwner represents a UpdatedOwner event raised by the AdventureSwamp contract.
type AdventureSwampUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureSwamp *AdventureSwampFilterer) FilterUpdatedOwner(opts *bind.FilterOpts, owner []common.Address) (*AdventureSwampUpdatedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureSwamp.contract.FilterLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AdventureSwampUpdatedOwnerIterator{contract: _AdventureSwamp.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureSwamp *AdventureSwampFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *AdventureSwampUpdatedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureSwamp.contract.WatchLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureSwampUpdatedOwner)
				if err := _AdventureSwamp.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
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

// ParseUpdatedOwner is a log parse operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureSwamp *AdventureSwampFilterer) ParseUpdatedOwner(log types.Log) (*AdventureSwampUpdatedOwner, error) {
	event := new(AdventureSwampUpdatedOwner)
	if err := _AdventureSwamp.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
