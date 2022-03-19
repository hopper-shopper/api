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

// AdventureForestMetaData contains all meta data concerning the AdventureForest contract.
var AdventureForestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vefly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hopper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NoHopperStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnfitHopper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongTokenID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"}],\"name\":\"UpdatedBallot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionRate\",\"type\":\"uint256\"}],\"name\":\"UpdatedEmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOPPER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEVEL_GAUGE_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"baseSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEmissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"}],\"name\":\"canEnter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"changeHopperName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"emergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flyLevelCapRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"forceUnvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"generatedPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHopperAndGauge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"hopperGauge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gaugeLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLevel\",\"type\":\"uint256\"}],\"name\":\"getLevelUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBonusShares\",\"type\":\"uint256\"}],\"name\":\"getUserBonusGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBaseShares\",\"type\":\"uint256\"}],\"name\":\"getUserGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBonusUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"levelUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"}],\"name\":\"setBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"setBonusEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emissionRate\",\"type\":\"uint256\"}],\"name\":\"setEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flyLevelCapRatio\",\"type\":\"uint256\"}],\"name\":\"setFlyLevelCapRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenCapFilledPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBonusRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMaxFlyGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veFlyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdventureForestABI is the input ABI used to generate the binding from.
// Deprecated: Use AdventureForestMetaData.ABI instead.
var AdventureForestABI = AdventureForestMetaData.ABI

// AdventureForest is an auto generated Go binding around an Ethereum contract.
type AdventureForest struct {
	AdventureForestCaller     // Read-only binding to the contract
	AdventureForestTransactor // Write-only binding to the contract
	AdventureForestFilterer   // Log filterer for contract events
}

// AdventureForestCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdventureForestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureForestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdventureForestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureForestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdventureForestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureForestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdventureForestSession struct {
	Contract     *AdventureForest  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdventureForestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdventureForestCallerSession struct {
	Contract *AdventureForestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AdventureForestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdventureForestTransactorSession struct {
	Contract     *AdventureForestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AdventureForestRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdventureForestRaw struct {
	Contract *AdventureForest // Generic contract binding to access the raw methods on
}

// AdventureForestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdventureForestCallerRaw struct {
	Contract *AdventureForestCaller // Generic read-only contract binding to access the raw methods on
}

// AdventureForestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdventureForestTransactorRaw struct {
	Contract *AdventureForestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdventureForest creates a new instance of AdventureForest, bound to a specific deployed contract.
func NewAdventureForest(address common.Address, backend bind.ContractBackend) (*AdventureForest, error) {
	contract, err := bindAdventureForest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdventureForest{AdventureForestCaller: AdventureForestCaller{contract: contract}, AdventureForestTransactor: AdventureForestTransactor{contract: contract}, AdventureForestFilterer: AdventureForestFilterer{contract: contract}}, nil
}

// NewAdventureForestCaller creates a new read-only instance of AdventureForest, bound to a specific deployed contract.
func NewAdventureForestCaller(address common.Address, caller bind.ContractCaller) (*AdventureForestCaller, error) {
	contract, err := bindAdventureForest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureForestCaller{contract: contract}, nil
}

// NewAdventureForestTransactor creates a new write-only instance of AdventureForest, bound to a specific deployed contract.
func NewAdventureForestTransactor(address common.Address, transactor bind.ContractTransactor) (*AdventureForestTransactor, error) {
	contract, err := bindAdventureForest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureForestTransactor{contract: contract}, nil
}

// NewAdventureForestFilterer creates a new log filterer instance of AdventureForest, bound to a specific deployed contract.
func NewAdventureForestFilterer(address common.Address, filterer bind.ContractFilterer) (*AdventureForestFilterer, error) {
	contract, err := bindAdventureForest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdventureForestFilterer{contract: contract}, nil
}

// bindAdventureForest binds a generic wrapper to an already deployed contract.
func bindAdventureForest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdventureForestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureForest *AdventureForestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureForest.Contract.AdventureForestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureForest *AdventureForestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureForest.Contract.AdventureForestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureForest *AdventureForestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureForest.Contract.AdventureForestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureForest *AdventureForestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureForest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureForest *AdventureForestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureForest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureForest *AdventureForestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureForest.Contract.contract.Transact(opts, method, params...)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureForest *AdventureForestCaller) FLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureForest *AdventureForestSession) FLY() (common.Address, error) {
	return _AdventureForest.Contract.FLY(&_AdventureForest.CallOpts)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureForest *AdventureForestCallerSession) FLY() (common.Address, error) {
	return _AdventureForest.Contract.FLY(&_AdventureForest.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureForest *AdventureForestCaller) HOPPER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "HOPPER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureForest *AdventureForestSession) HOPPER() (common.Address, error) {
	return _AdventureForest.Contract.HOPPER(&_AdventureForest.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureForest *AdventureForestCallerSession) HOPPER() (common.Address, error) {
	return _AdventureForest.Contract.HOPPER(&_AdventureForest.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureForest *AdventureForestCaller) LEVELGAUGEKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "LEVEL_GAUGE_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureForest *AdventureForestSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureForest.Contract.LEVELGAUGEKEY(&_AdventureForest.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureForest *AdventureForestCallerSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureForest.Contract.LEVELGAUGEKEY(&_AdventureForest.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureForest *AdventureForestCaller) VEFLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "VE_FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureForest *AdventureForestSession) VEFLY() (common.Address, error) {
	return _AdventureForest.Contract.VEFLY(&_AdventureForest.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureForest *AdventureForestCallerSession) VEFLY() (common.Address, error) {
	return _AdventureForest.Contract.VEFLY(&_AdventureForest.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureForest *AdventureForestCaller) Ballot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "ballot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureForest *AdventureForestSession) Ballot() (common.Address, error) {
	return _AdventureForest.Contract.Ballot(&_AdventureForest.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureForest *AdventureForestCallerSession) Ballot() (common.Address, error) {
	return _AdventureForest.Contract.Ballot(&_AdventureForest.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) BaseRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "baseRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureForest *AdventureForestSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureForest.Contract.BaseRewardPerShare(&_AdventureForest.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureForest.Contract.BaseRewardPerShare(&_AdventureForest.CallOpts)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) BaseSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "baseSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.BaseSharesBalance(&_AdventureForest.CallOpts, arg0)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.BaseSharesBalance(&_AdventureForest.CallOpts, arg0)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) BonusEmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "bonusEmissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureForest *AdventureForestSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureForest.Contract.BonusEmissionRate(&_AdventureForest.CallOpts)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureForest.Contract.BonusEmissionRate(&_AdventureForest.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) BonusRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "bonusRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureForest *AdventureForestSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureForest.Contract.BonusRewardPerShare(&_AdventureForest.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureForest.Contract.BonusRewardPerShare(&_AdventureForest.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) BonusRewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "bonusRewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureForest *AdventureForestSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureForest.Contract.BonusRewardPerShareStored(&_AdventureForest.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureForest.Contract.BonusRewardPerShareStored(&_AdventureForest.CallOpts)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureForest *AdventureForestCaller) CanEnter(opts *bind.CallOpts, hopper HopperNFTHopper) (bool, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "canEnter", hopper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureForest *AdventureForestSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureForest.Contract.CanEnter(&_AdventureForest.CallOpts, hopper)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureForest *AdventureForestCallerSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureForest.Contract.CanEnter(&_AdventureForest.CallOpts, hopper)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureForest *AdventureForestSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.Claimable(&_AdventureForest.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.Claimable(&_AdventureForest.CallOpts, _account)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureForest *AdventureForestCaller) Emergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "emergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureForest *AdventureForestSession) Emergency() (bool, error) {
	return _AdventureForest.Contract.Emergency(&_AdventureForest.CallOpts)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureForest *AdventureForestCallerSession) Emergency() (bool, error) {
	return _AdventureForest.Contract.Emergency(&_AdventureForest.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) EmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "emissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureForest *AdventureForestSession) EmissionRate() (*big.Int, error) {
	return _AdventureForest.Contract.EmissionRate(&_AdventureForest.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) EmissionRate() (*big.Int, error) {
	return _AdventureForest.Contract.EmissionRate(&_AdventureForest.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) FlyLevelCapRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "flyLevelCapRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureForest *AdventureForestSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureForest.Contract.FlyLevelCapRatio(&_AdventureForest.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureForest.Contract.FlyLevelCapRatio(&_AdventureForest.CallOpts)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) GeneratedPerShareStored(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "generatedPerShareStored", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.GeneratedPerShareStored(&_AdventureForest.CallOpts, arg0)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.GeneratedPerShareStored(&_AdventureForest.CallOpts, arg0)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureForest *AdventureForestCaller) GetHopperAndGauge(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "getHopperAndGauge", tokenId)

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
func (_AdventureForest *AdventureForestSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureForest.Contract.GetHopperAndGauge(&_AdventureForest.CallOpts, tokenId)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureForest *AdventureForestCallerSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureForest.Contract.GetHopperAndGauge(&_AdventureForest.CallOpts, tokenId)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureForest *AdventureForestCaller) GetLevelUpCost(opts *bind.CallOpts, currentLevel *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "getLevelUpCost", currentLevel)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureForest *AdventureForestSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureForest.Contract.GetLevelUpCost(&_AdventureForest.CallOpts, currentLevel)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureForest.Contract.GetLevelUpCost(&_AdventureForest.CallOpts, currentLevel)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureForest *AdventureForestCaller) GetUserBonusGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "getUserBonusGeneratedFly", account, _totalUserBonusShares)

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
func (_AdventureForest *AdventureForestSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureForest.Contract.GetUserBonusGeneratedFly(&_AdventureForest.CallOpts, account, _totalUserBonusShares)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureForest *AdventureForestCallerSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureForest.Contract.GetUserBonusGeneratedFly(&_AdventureForest.CallOpts, account, _totalUserBonusShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureForest *AdventureForestCaller) GetUserGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "getUserGeneratedFly", account, _totalUserBaseShares)

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
func (_AdventureForest *AdventureForestSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureForest.Contract.GetUserGeneratedFly(&_AdventureForest.CallOpts, account, _totalUserBaseShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureForest *AdventureForestCallerSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureForest.Contract.GetUserGeneratedFly(&_AdventureForest.CallOpts, account, _totalUserBaseShares)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) HopperBaseShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "hopperBaseShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureForest.Contract.HopperBaseShare(&_AdventureForest.CallOpts, arg0)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureForest.Contract.HopperBaseShare(&_AdventureForest.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureForest *AdventureForestCaller) HopperOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "hopperOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureForest *AdventureForestSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureForest.Contract.HopperOwners(&_AdventureForest.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureForest *AdventureForestCallerSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureForest.Contract.HopperOwners(&_AdventureForest.CallOpts, arg0)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) LastBonusUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "lastBonusUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureForest *AdventureForestSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureForest.Contract.LastBonusUpdatedTime(&_AdventureForest.CallOpts)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureForest.Contract.LastBonusUpdatedTime(&_AdventureForest.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) LastUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "lastUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureForest *AdventureForestSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureForest.Contract.LastUpdatedTime(&_AdventureForest.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureForest.Contract.LastUpdatedTime(&_AdventureForest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureForest *AdventureForestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureForest *AdventureForestSession) Owner() (common.Address, error) {
	return _AdventureForest.Contract.Owner(&_AdventureForest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureForest *AdventureForestCallerSession) Owner() (common.Address, error) {
	return _AdventureForest.Contract.Owner(&_AdventureForest.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) RewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "rewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureForest *AdventureForestSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureForest.Contract.RewardPerShareStored(&_AdventureForest.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureForest.Contract.RewardPerShareStored(&_AdventureForest.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.Rewards(&_AdventureForest.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.Rewards(&_AdventureForest.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) TokenCapFilledPerShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "tokenCapFilledPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureForest.Contract.TokenCapFilledPerShare(&_AdventureForest.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureForest.Contract.TokenCapFilledPerShare(&_AdventureForest.CallOpts, arg0)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) TotalBaseShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "totalBaseShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureForest *AdventureForestSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureForest.Contract.TotalBaseShare(&_AdventureForest.CallOpts)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureForest.Contract.TotalBaseShare(&_AdventureForest.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureForest *AdventureForestCaller) TotalVeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "totalVeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureForest *AdventureForestSession) TotalVeShare() (*big.Int, error) {
	return _AdventureForest.Contract.TotalVeShare(&_AdventureForest.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) TotalVeShare() (*big.Int, error) {
	return _AdventureForest.Contract.TotalVeShare(&_AdventureForest.CallOpts)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) UserBonusRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "userBonusRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.UserBonusRewardPerSharePaid(&_AdventureForest.CallOpts, arg0)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.UserBonusRewardPerSharePaid(&_AdventureForest.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) UserMaxFlyGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "userMaxFlyGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.UserMaxFlyGeneration(&_AdventureForest.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.UserMaxFlyGeneration(&_AdventureForest.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) UserRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "userRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.UserRewardPerSharePaid(&_AdventureForest.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.UserRewardPerSharePaid(&_AdventureForest.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) VeFlyBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "veFlyBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.VeFlyBalance(&_AdventureForest.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.VeFlyBalance(&_AdventureForest.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCaller) VeSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureForest.contract.Call(opts, &out, "veSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.VeSharesBalance(&_AdventureForest.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureForest *AdventureForestCallerSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureForest.Contract.VeSharesBalance(&_AdventureForest.CallOpts, arg0)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureForest *AdventureForestTransactor) ChangeHopperName(opts *bind.TransactOpts, tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "changeHopperName", tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureForest *AdventureForestSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.ChangeHopperName(&_AdventureForest.TransactOpts, tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureForest *AdventureForestTransactorSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.ChangeHopperName(&_AdventureForest.TransactOpts, tokenId, name, useOwnRewards)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureForest *AdventureForestTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureForest *AdventureForestSession) Claim() (*types.Transaction, error) {
	return _AdventureForest.Contract.Claim(&_AdventureForest.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureForest *AdventureForestTransactorSession) Claim() (*types.Transaction, error) {
	return _AdventureForest.Contract.Claim(&_AdventureForest.TransactOpts)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureForest *AdventureForestTransactor) EmergencyExit(opts *bind.TransactOpts, tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "emergencyExit", tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureForest *AdventureForestSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.EmergencyExit(&_AdventureForest.TransactOpts, tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureForest *AdventureForestTransactorSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.EmergencyExit(&_AdventureForest.TransactOpts, tokenIds, user)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureForest *AdventureForestTransactor) EnableEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "enableEmergency")
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureForest *AdventureForestSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureForest.Contract.EnableEmergency(&_AdventureForest.TransactOpts)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureForest *AdventureForestTransactorSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureForest.Contract.EnableEmergency(&_AdventureForest.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureForest *AdventureForestTransactor) Enter(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "enter", tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureForest *AdventureForestSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.Enter(&_AdventureForest.TransactOpts, tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureForest *AdventureForestTransactorSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.Enter(&_AdventureForest.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureForest *AdventureForestTransactor) Exit(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "exit", tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureForest *AdventureForestSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.Exit(&_AdventureForest.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureForest *AdventureForestTransactorSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.Exit(&_AdventureForest.TransactOpts, tokenIds)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureForest *AdventureForestTransactor) ForceUnvote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "forceUnvote", user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureForest *AdventureForestSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.ForceUnvote(&_AdventureForest.TransactOpts, user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureForest *AdventureForestTransactorSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.ForceUnvote(&_AdventureForest.TransactOpts, user)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureForest *AdventureForestTransactor) LevelUp(opts *bind.TransactOpts, tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "levelUp", tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureForest *AdventureForestSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.LevelUp(&_AdventureForest.TransactOpts, tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureForest *AdventureForestTransactorSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.LevelUp(&_AdventureForest.TransactOpts, tokenId, useOwnRewards)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureForest *AdventureForestTransactor) SetBallot(opts *bind.TransactOpts, _ballot common.Address) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "setBallot", _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureForest *AdventureForestSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetBallot(&_AdventureForest.TransactOpts, _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureForest *AdventureForestTransactorSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetBallot(&_AdventureForest.TransactOpts, _ballot)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureForest *AdventureForestTransactor) SetBonusEmissionRate(opts *bind.TransactOpts, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "setBonusEmissionRate", _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureForest *AdventureForestSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetBonusEmissionRate(&_AdventureForest.TransactOpts, _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureForest *AdventureForestTransactorSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetBonusEmissionRate(&_AdventureForest.TransactOpts, _bonusEmissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureForest *AdventureForestTransactor) SetEmissionRate(opts *bind.TransactOpts, _emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "setEmissionRate", _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureForest *AdventureForestSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetEmissionRate(&_AdventureForest.TransactOpts, _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureForest *AdventureForestTransactorSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetEmissionRate(&_AdventureForest.TransactOpts, _emissionRate)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureForest *AdventureForestTransactor) SetFlyLevelCapRatio(opts *bind.TransactOpts, _flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "setFlyLevelCapRatio", _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureForest *AdventureForestSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetFlyLevelCapRatio(&_AdventureForest.TransactOpts, _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureForest *AdventureForestTransactorSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetFlyLevelCapRatio(&_AdventureForest.TransactOpts, _flyLevelCapRatio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureForest *AdventureForestTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureForest *AdventureForestSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetOwner(&_AdventureForest.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureForest *AdventureForestTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureForest.Contract.SetOwner(&_AdventureForest.TransactOpts, _owner)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureForest *AdventureForestTransactor) Unvote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "unvote", veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureForest *AdventureForestSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.Unvote(&_AdventureForest.TransactOpts, veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureForest *AdventureForestTransactorSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.Unvote(&_AdventureForest.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureForest *AdventureForestTransactor) Vote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureForest.contract.Transact(opts, "vote", veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureForest *AdventureForestSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.Vote(&_AdventureForest.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureForest *AdventureForestTransactorSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureForest.Contract.Vote(&_AdventureForest.TransactOpts, veFlyAmount, recount)
}

// AdventureForestUpdatedBallotIterator is returned from FilterUpdatedBallot and is used to iterate over the raw logs and unpacked data for UpdatedBallot events raised by the AdventureForest contract.
type AdventureForestUpdatedBallotIterator struct {
	Event *AdventureForestUpdatedBallot // Event containing the contract specifics and raw log

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
func (it *AdventureForestUpdatedBallotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureForestUpdatedBallot)
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
		it.Event = new(AdventureForestUpdatedBallot)
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
func (it *AdventureForestUpdatedBallotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureForestUpdatedBallotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureForestUpdatedBallot represents a UpdatedBallot event raised by the AdventureForest contract.
type AdventureForestUpdatedBallot struct {
	Ballot common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBallot is a free log retrieval operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureForest *AdventureForestFilterer) FilterUpdatedBallot(opts *bind.FilterOpts, ballot []common.Address) (*AdventureForestUpdatedBallotIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureForest.contract.FilterLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return &AdventureForestUpdatedBallotIterator{contract: _AdventureForest.contract, event: "UpdatedBallot", logs: logs, sub: sub}, nil
}

// WatchUpdatedBallot is a free log subscription operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureForest *AdventureForestFilterer) WatchUpdatedBallot(opts *bind.WatchOpts, sink chan<- *AdventureForestUpdatedBallot, ballot []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureForest.contract.WatchLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureForestUpdatedBallot)
				if err := _AdventureForest.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
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
func (_AdventureForest *AdventureForestFilterer) ParseUpdatedBallot(log types.Log) (*AdventureForestUpdatedBallot, error) {
	event := new(AdventureForestUpdatedBallot)
	if err := _AdventureForest.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureForestUpdatedEmissionIterator is returned from FilterUpdatedEmission and is used to iterate over the raw logs and unpacked data for UpdatedEmission events raised by the AdventureForest contract.
type AdventureForestUpdatedEmissionIterator struct {
	Event *AdventureForestUpdatedEmission // Event containing the contract specifics and raw log

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
func (it *AdventureForestUpdatedEmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureForestUpdatedEmission)
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
		it.Event = new(AdventureForestUpdatedEmission)
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
func (it *AdventureForestUpdatedEmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureForestUpdatedEmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureForestUpdatedEmission represents a UpdatedEmission event raised by the AdventureForest contract.
type AdventureForestUpdatedEmission struct {
	EmissionRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEmission is a free log retrieval operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureForest *AdventureForestFilterer) FilterUpdatedEmission(opts *bind.FilterOpts) (*AdventureForestUpdatedEmissionIterator, error) {

	logs, sub, err := _AdventureForest.contract.FilterLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return &AdventureForestUpdatedEmissionIterator{contract: _AdventureForest.contract, event: "UpdatedEmission", logs: logs, sub: sub}, nil
}

// WatchUpdatedEmission is a free log subscription operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureForest *AdventureForestFilterer) WatchUpdatedEmission(opts *bind.WatchOpts, sink chan<- *AdventureForestUpdatedEmission) (event.Subscription, error) {

	logs, sub, err := _AdventureForest.contract.WatchLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureForestUpdatedEmission)
				if err := _AdventureForest.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
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
func (_AdventureForest *AdventureForestFilterer) ParseUpdatedEmission(log types.Log) (*AdventureForestUpdatedEmission, error) {
	event := new(AdventureForestUpdatedEmission)
	if err := _AdventureForest.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureForestUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the AdventureForest contract.
type AdventureForestUpdatedOwnerIterator struct {
	Event *AdventureForestUpdatedOwner // Event containing the contract specifics and raw log

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
func (it *AdventureForestUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureForestUpdatedOwner)
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
		it.Event = new(AdventureForestUpdatedOwner)
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
func (it *AdventureForestUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureForestUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureForestUpdatedOwner represents a UpdatedOwner event raised by the AdventureForest contract.
type AdventureForestUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureForest *AdventureForestFilterer) FilterUpdatedOwner(opts *bind.FilterOpts, owner []common.Address) (*AdventureForestUpdatedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureForest.contract.FilterLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AdventureForestUpdatedOwnerIterator{contract: _AdventureForest.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureForest *AdventureForestFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *AdventureForestUpdatedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureForest.contract.WatchLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureForestUpdatedOwner)
				if err := _AdventureForest.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
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
func (_AdventureForest *AdventureForestFilterer) ParseUpdatedOwner(log types.Log) (*AdventureForestUpdatedOwner, error) {
	event := new(AdventureForestUpdatedOwner)
	if err := _AdventureForest.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
