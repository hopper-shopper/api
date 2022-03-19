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

// AdventureRiverMetaData contains all meta data concerning the AdventureRiver contract.
var AdventureRiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vefly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hopper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NoHopperStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnfitHopper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongTokenID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"}],\"name\":\"UpdatedBallot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionRate\",\"type\":\"uint256\"}],\"name\":\"UpdatedEmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOPPER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEVEL_GAUGE_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"baseSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEmissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"}],\"name\":\"canEnter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"changeHopperName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"emergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flyLevelCapRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"forceUnvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"generatedPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHopperAndGauge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"hopperGauge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gaugeLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLevel\",\"type\":\"uint256\"}],\"name\":\"getLevelUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBonusShares\",\"type\":\"uint256\"}],\"name\":\"getUserBonusGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBaseShares\",\"type\":\"uint256\"}],\"name\":\"getUserGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBonusUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"levelUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"}],\"name\":\"setBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"setBonusEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emissionRate\",\"type\":\"uint256\"}],\"name\":\"setEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flyLevelCapRatio\",\"type\":\"uint256\"}],\"name\":\"setFlyLevelCapRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenCapFilledPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBonusRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMaxFlyGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veFlyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdventureRiverABI is the input ABI used to generate the binding from.
// Deprecated: Use AdventureRiverMetaData.ABI instead.
var AdventureRiverABI = AdventureRiverMetaData.ABI

// AdventureRiver is an auto generated Go binding around an Ethereum contract.
type AdventureRiver struct {
	AdventureRiverCaller     // Read-only binding to the contract
	AdventureRiverTransactor // Write-only binding to the contract
	AdventureRiverFilterer   // Log filterer for contract events
}

// AdventureRiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdventureRiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureRiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdventureRiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureRiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdventureRiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureRiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdventureRiverSession struct {
	Contract     *AdventureRiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdventureRiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdventureRiverCallerSession struct {
	Contract *AdventureRiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AdventureRiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdventureRiverTransactorSession struct {
	Contract     *AdventureRiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AdventureRiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdventureRiverRaw struct {
	Contract *AdventureRiver // Generic contract binding to access the raw methods on
}

// AdventureRiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdventureRiverCallerRaw struct {
	Contract *AdventureRiverCaller // Generic read-only contract binding to access the raw methods on
}

// AdventureRiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdventureRiverTransactorRaw struct {
	Contract *AdventureRiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdventureRiver creates a new instance of AdventureRiver, bound to a specific deployed contract.
func NewAdventureRiver(address common.Address, backend bind.ContractBackend) (*AdventureRiver, error) {
	contract, err := bindAdventureRiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdventureRiver{AdventureRiverCaller: AdventureRiverCaller{contract: contract}, AdventureRiverTransactor: AdventureRiverTransactor{contract: contract}, AdventureRiverFilterer: AdventureRiverFilterer{contract: contract}}, nil
}

// NewAdventureRiverCaller creates a new read-only instance of AdventureRiver, bound to a specific deployed contract.
func NewAdventureRiverCaller(address common.Address, caller bind.ContractCaller) (*AdventureRiverCaller, error) {
	contract, err := bindAdventureRiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureRiverCaller{contract: contract}, nil
}

// NewAdventureRiverTransactor creates a new write-only instance of AdventureRiver, bound to a specific deployed contract.
func NewAdventureRiverTransactor(address common.Address, transactor bind.ContractTransactor) (*AdventureRiverTransactor, error) {
	contract, err := bindAdventureRiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureRiverTransactor{contract: contract}, nil
}

// NewAdventureRiverFilterer creates a new log filterer instance of AdventureRiver, bound to a specific deployed contract.
func NewAdventureRiverFilterer(address common.Address, filterer bind.ContractFilterer) (*AdventureRiverFilterer, error) {
	contract, err := bindAdventureRiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdventureRiverFilterer{contract: contract}, nil
}

// bindAdventureRiver binds a generic wrapper to an already deployed contract.
func bindAdventureRiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdventureRiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureRiver *AdventureRiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureRiver.Contract.AdventureRiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureRiver *AdventureRiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureRiver.Contract.AdventureRiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureRiver *AdventureRiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureRiver.Contract.AdventureRiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureRiver *AdventureRiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureRiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureRiver *AdventureRiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureRiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureRiver *AdventureRiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureRiver.Contract.contract.Transact(opts, method, params...)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureRiver *AdventureRiverCaller) FLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureRiver *AdventureRiverSession) FLY() (common.Address, error) {
	return _AdventureRiver.Contract.FLY(&_AdventureRiver.CallOpts)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureRiver *AdventureRiverCallerSession) FLY() (common.Address, error) {
	return _AdventureRiver.Contract.FLY(&_AdventureRiver.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureRiver *AdventureRiverCaller) HOPPER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "HOPPER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureRiver *AdventureRiverSession) HOPPER() (common.Address, error) {
	return _AdventureRiver.Contract.HOPPER(&_AdventureRiver.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureRiver *AdventureRiverCallerSession) HOPPER() (common.Address, error) {
	return _AdventureRiver.Contract.HOPPER(&_AdventureRiver.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureRiver *AdventureRiverCaller) LEVELGAUGEKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "LEVEL_GAUGE_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureRiver *AdventureRiverSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureRiver.Contract.LEVELGAUGEKEY(&_AdventureRiver.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureRiver *AdventureRiverCallerSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureRiver.Contract.LEVELGAUGEKEY(&_AdventureRiver.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureRiver *AdventureRiverCaller) VEFLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "VE_FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureRiver *AdventureRiverSession) VEFLY() (common.Address, error) {
	return _AdventureRiver.Contract.VEFLY(&_AdventureRiver.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureRiver *AdventureRiverCallerSession) VEFLY() (common.Address, error) {
	return _AdventureRiver.Contract.VEFLY(&_AdventureRiver.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureRiver *AdventureRiverCaller) Ballot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "ballot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureRiver *AdventureRiverSession) Ballot() (common.Address, error) {
	return _AdventureRiver.Contract.Ballot(&_AdventureRiver.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureRiver *AdventureRiverCallerSession) Ballot() (common.Address, error) {
	return _AdventureRiver.Contract.Ballot(&_AdventureRiver.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) BaseRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "baseRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureRiver.Contract.BaseRewardPerShare(&_AdventureRiver.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureRiver.Contract.BaseRewardPerShare(&_AdventureRiver.CallOpts)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) BaseSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "baseSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.BaseSharesBalance(&_AdventureRiver.CallOpts, arg0)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.BaseSharesBalance(&_AdventureRiver.CallOpts, arg0)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) BonusEmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "bonusEmissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureRiver.Contract.BonusEmissionRate(&_AdventureRiver.CallOpts)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureRiver.Contract.BonusEmissionRate(&_AdventureRiver.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) BonusRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "bonusRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureRiver.Contract.BonusRewardPerShare(&_AdventureRiver.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureRiver.Contract.BonusRewardPerShare(&_AdventureRiver.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) BonusRewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "bonusRewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureRiver.Contract.BonusRewardPerShareStored(&_AdventureRiver.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureRiver.Contract.BonusRewardPerShareStored(&_AdventureRiver.CallOpts)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureRiver *AdventureRiverCaller) CanEnter(opts *bind.CallOpts, hopper HopperNFTHopper) (bool, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "canEnter", hopper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureRiver *AdventureRiverSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureRiver.Contract.CanEnter(&_AdventureRiver.CallOpts, hopper)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureRiver *AdventureRiverCallerSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureRiver.Contract.CanEnter(&_AdventureRiver.CallOpts, hopper)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.Claimable(&_AdventureRiver.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.Claimable(&_AdventureRiver.CallOpts, _account)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureRiver *AdventureRiverCaller) Emergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "emergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureRiver *AdventureRiverSession) Emergency() (bool, error) {
	return _AdventureRiver.Contract.Emergency(&_AdventureRiver.CallOpts)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureRiver *AdventureRiverCallerSession) Emergency() (bool, error) {
	return _AdventureRiver.Contract.Emergency(&_AdventureRiver.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) EmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "emissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) EmissionRate() (*big.Int, error) {
	return _AdventureRiver.Contract.EmissionRate(&_AdventureRiver.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) EmissionRate() (*big.Int, error) {
	return _AdventureRiver.Contract.EmissionRate(&_AdventureRiver.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) FlyLevelCapRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "flyLevelCapRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureRiver.Contract.FlyLevelCapRatio(&_AdventureRiver.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureRiver.Contract.FlyLevelCapRatio(&_AdventureRiver.CallOpts)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) GeneratedPerShareStored(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "generatedPerShareStored", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.GeneratedPerShareStored(&_AdventureRiver.CallOpts, arg0)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.GeneratedPerShareStored(&_AdventureRiver.CallOpts, arg0)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureRiver *AdventureRiverCaller) GetHopperAndGauge(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "getHopperAndGauge", tokenId)

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
func (_AdventureRiver *AdventureRiverSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureRiver.Contract.GetHopperAndGauge(&_AdventureRiver.CallOpts, tokenId)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureRiver *AdventureRiverCallerSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureRiver.Contract.GetHopperAndGauge(&_AdventureRiver.CallOpts, tokenId)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) GetLevelUpCost(opts *bind.CallOpts, currentLevel *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "getLevelUpCost", currentLevel)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureRiver *AdventureRiverSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureRiver.Contract.GetLevelUpCost(&_AdventureRiver.CallOpts, currentLevel)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureRiver.Contract.GetLevelUpCost(&_AdventureRiver.CallOpts, currentLevel)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureRiver *AdventureRiverCaller) GetUserBonusGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "getUserBonusGeneratedFly", account, _totalUserBonusShares)

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
func (_AdventureRiver *AdventureRiverSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureRiver.Contract.GetUserBonusGeneratedFly(&_AdventureRiver.CallOpts, account, _totalUserBonusShares)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureRiver *AdventureRiverCallerSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureRiver.Contract.GetUserBonusGeneratedFly(&_AdventureRiver.CallOpts, account, _totalUserBonusShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureRiver *AdventureRiverCaller) GetUserGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "getUserGeneratedFly", account, _totalUserBaseShares)

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
func (_AdventureRiver *AdventureRiverSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureRiver.Contract.GetUserGeneratedFly(&_AdventureRiver.CallOpts, account, _totalUserBaseShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureRiver *AdventureRiverCallerSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureRiver.Contract.GetUserGeneratedFly(&_AdventureRiver.CallOpts, account, _totalUserBaseShares)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) HopperBaseShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "hopperBaseShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureRiver.Contract.HopperBaseShare(&_AdventureRiver.CallOpts, arg0)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureRiver.Contract.HopperBaseShare(&_AdventureRiver.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureRiver *AdventureRiverCaller) HopperOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "hopperOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureRiver *AdventureRiverSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureRiver.Contract.HopperOwners(&_AdventureRiver.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureRiver *AdventureRiverCallerSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureRiver.Contract.HopperOwners(&_AdventureRiver.CallOpts, arg0)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) LastBonusUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "lastBonusUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureRiver.Contract.LastBonusUpdatedTime(&_AdventureRiver.CallOpts)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureRiver.Contract.LastBonusUpdatedTime(&_AdventureRiver.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) LastUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "lastUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureRiver.Contract.LastUpdatedTime(&_AdventureRiver.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureRiver.Contract.LastUpdatedTime(&_AdventureRiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureRiver *AdventureRiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureRiver *AdventureRiverSession) Owner() (common.Address, error) {
	return _AdventureRiver.Contract.Owner(&_AdventureRiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureRiver *AdventureRiverCallerSession) Owner() (common.Address, error) {
	return _AdventureRiver.Contract.Owner(&_AdventureRiver.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) RewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "rewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureRiver.Contract.RewardPerShareStored(&_AdventureRiver.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureRiver.Contract.RewardPerShareStored(&_AdventureRiver.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.Rewards(&_AdventureRiver.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.Rewards(&_AdventureRiver.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) TokenCapFilledPerShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "tokenCapFilledPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureRiver.Contract.TokenCapFilledPerShare(&_AdventureRiver.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureRiver.Contract.TokenCapFilledPerShare(&_AdventureRiver.CallOpts, arg0)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) TotalBaseShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "totalBaseShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureRiver.Contract.TotalBaseShare(&_AdventureRiver.CallOpts)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureRiver.Contract.TotalBaseShare(&_AdventureRiver.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) TotalVeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "totalVeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) TotalVeShare() (*big.Int, error) {
	return _AdventureRiver.Contract.TotalVeShare(&_AdventureRiver.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) TotalVeShare() (*big.Int, error) {
	return _AdventureRiver.Contract.TotalVeShare(&_AdventureRiver.CallOpts)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) UserBonusRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "userBonusRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.UserBonusRewardPerSharePaid(&_AdventureRiver.CallOpts, arg0)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.UserBonusRewardPerSharePaid(&_AdventureRiver.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) UserMaxFlyGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "userMaxFlyGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.UserMaxFlyGeneration(&_AdventureRiver.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.UserMaxFlyGeneration(&_AdventureRiver.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) UserRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "userRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.UserRewardPerSharePaid(&_AdventureRiver.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.UserRewardPerSharePaid(&_AdventureRiver.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) VeFlyBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "veFlyBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.VeFlyBalance(&_AdventureRiver.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.VeFlyBalance(&_AdventureRiver.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCaller) VeSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureRiver.contract.Call(opts, &out, "veSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.VeSharesBalance(&_AdventureRiver.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureRiver *AdventureRiverCallerSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureRiver.Contract.VeSharesBalance(&_AdventureRiver.CallOpts, arg0)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureRiver *AdventureRiverTransactor) ChangeHopperName(opts *bind.TransactOpts, tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "changeHopperName", tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureRiver *AdventureRiverSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.ChangeHopperName(&_AdventureRiver.TransactOpts, tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.ChangeHopperName(&_AdventureRiver.TransactOpts, tokenId, name, useOwnRewards)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureRiver *AdventureRiverTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureRiver *AdventureRiverSession) Claim() (*types.Transaction, error) {
	return _AdventureRiver.Contract.Claim(&_AdventureRiver.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureRiver *AdventureRiverTransactorSession) Claim() (*types.Transaction, error) {
	return _AdventureRiver.Contract.Claim(&_AdventureRiver.TransactOpts)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureRiver *AdventureRiverTransactor) EmergencyExit(opts *bind.TransactOpts, tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "emergencyExit", tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureRiver *AdventureRiverSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.EmergencyExit(&_AdventureRiver.TransactOpts, tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.EmergencyExit(&_AdventureRiver.TransactOpts, tokenIds, user)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureRiver *AdventureRiverTransactor) EnableEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "enableEmergency")
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureRiver *AdventureRiverSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureRiver.Contract.EnableEmergency(&_AdventureRiver.TransactOpts)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureRiver *AdventureRiverTransactorSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureRiver.Contract.EnableEmergency(&_AdventureRiver.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureRiver *AdventureRiverTransactor) Enter(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "enter", tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureRiver *AdventureRiverSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Enter(&_AdventureRiver.TransactOpts, tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Enter(&_AdventureRiver.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureRiver *AdventureRiverTransactor) Exit(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "exit", tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureRiver *AdventureRiverSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Exit(&_AdventureRiver.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Exit(&_AdventureRiver.TransactOpts, tokenIds)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureRiver *AdventureRiverTransactor) ForceUnvote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "forceUnvote", user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureRiver *AdventureRiverSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.ForceUnvote(&_AdventureRiver.TransactOpts, user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.ForceUnvote(&_AdventureRiver.TransactOpts, user)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureRiver *AdventureRiverTransactor) LevelUp(opts *bind.TransactOpts, tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "levelUp", tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureRiver *AdventureRiverSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.LevelUp(&_AdventureRiver.TransactOpts, tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.LevelUp(&_AdventureRiver.TransactOpts, tokenId, useOwnRewards)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureRiver *AdventureRiverTransactor) SetBallot(opts *bind.TransactOpts, _ballot common.Address) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "setBallot", _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureRiver *AdventureRiverSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetBallot(&_AdventureRiver.TransactOpts, _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetBallot(&_AdventureRiver.TransactOpts, _ballot)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureRiver *AdventureRiverTransactor) SetBonusEmissionRate(opts *bind.TransactOpts, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "setBonusEmissionRate", _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureRiver *AdventureRiverSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetBonusEmissionRate(&_AdventureRiver.TransactOpts, _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetBonusEmissionRate(&_AdventureRiver.TransactOpts, _bonusEmissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureRiver *AdventureRiverTransactor) SetEmissionRate(opts *bind.TransactOpts, _emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "setEmissionRate", _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureRiver *AdventureRiverSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetEmissionRate(&_AdventureRiver.TransactOpts, _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetEmissionRate(&_AdventureRiver.TransactOpts, _emissionRate)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureRiver *AdventureRiverTransactor) SetFlyLevelCapRatio(opts *bind.TransactOpts, _flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "setFlyLevelCapRatio", _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureRiver *AdventureRiverSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetFlyLevelCapRatio(&_AdventureRiver.TransactOpts, _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetFlyLevelCapRatio(&_AdventureRiver.TransactOpts, _flyLevelCapRatio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureRiver *AdventureRiverTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureRiver *AdventureRiverSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetOwner(&_AdventureRiver.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureRiver.Contract.SetOwner(&_AdventureRiver.TransactOpts, _owner)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureRiver *AdventureRiverTransactor) Unvote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "unvote", veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureRiver *AdventureRiverSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Unvote(&_AdventureRiver.TransactOpts, veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Unvote(&_AdventureRiver.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureRiver *AdventureRiverTransactor) Vote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureRiver.contract.Transact(opts, "vote", veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureRiver *AdventureRiverSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Vote(&_AdventureRiver.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureRiver *AdventureRiverTransactorSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureRiver.Contract.Vote(&_AdventureRiver.TransactOpts, veFlyAmount, recount)
}

// AdventureRiverUpdatedBallotIterator is returned from FilterUpdatedBallot and is used to iterate over the raw logs and unpacked data for UpdatedBallot events raised by the AdventureRiver contract.
type AdventureRiverUpdatedBallotIterator struct {
	Event *AdventureRiverUpdatedBallot // Event containing the contract specifics and raw log

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
func (it *AdventureRiverUpdatedBallotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureRiverUpdatedBallot)
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
		it.Event = new(AdventureRiverUpdatedBallot)
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
func (it *AdventureRiverUpdatedBallotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureRiverUpdatedBallotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureRiverUpdatedBallot represents a UpdatedBallot event raised by the AdventureRiver contract.
type AdventureRiverUpdatedBallot struct {
	Ballot common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBallot is a free log retrieval operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureRiver *AdventureRiverFilterer) FilterUpdatedBallot(opts *bind.FilterOpts, ballot []common.Address) (*AdventureRiverUpdatedBallotIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureRiver.contract.FilterLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return &AdventureRiverUpdatedBallotIterator{contract: _AdventureRiver.contract, event: "UpdatedBallot", logs: logs, sub: sub}, nil
}

// WatchUpdatedBallot is a free log subscription operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureRiver *AdventureRiverFilterer) WatchUpdatedBallot(opts *bind.WatchOpts, sink chan<- *AdventureRiverUpdatedBallot, ballot []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureRiver.contract.WatchLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureRiverUpdatedBallot)
				if err := _AdventureRiver.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
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
func (_AdventureRiver *AdventureRiverFilterer) ParseUpdatedBallot(log types.Log) (*AdventureRiverUpdatedBallot, error) {
	event := new(AdventureRiverUpdatedBallot)
	if err := _AdventureRiver.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureRiverUpdatedEmissionIterator is returned from FilterUpdatedEmission and is used to iterate over the raw logs and unpacked data for UpdatedEmission events raised by the AdventureRiver contract.
type AdventureRiverUpdatedEmissionIterator struct {
	Event *AdventureRiverUpdatedEmission // Event containing the contract specifics and raw log

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
func (it *AdventureRiverUpdatedEmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureRiverUpdatedEmission)
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
		it.Event = new(AdventureRiverUpdatedEmission)
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
func (it *AdventureRiverUpdatedEmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureRiverUpdatedEmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureRiverUpdatedEmission represents a UpdatedEmission event raised by the AdventureRiver contract.
type AdventureRiverUpdatedEmission struct {
	EmissionRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEmission is a free log retrieval operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureRiver *AdventureRiverFilterer) FilterUpdatedEmission(opts *bind.FilterOpts) (*AdventureRiverUpdatedEmissionIterator, error) {

	logs, sub, err := _AdventureRiver.contract.FilterLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return &AdventureRiverUpdatedEmissionIterator{contract: _AdventureRiver.contract, event: "UpdatedEmission", logs: logs, sub: sub}, nil
}

// WatchUpdatedEmission is a free log subscription operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureRiver *AdventureRiverFilterer) WatchUpdatedEmission(opts *bind.WatchOpts, sink chan<- *AdventureRiverUpdatedEmission) (event.Subscription, error) {

	logs, sub, err := _AdventureRiver.contract.WatchLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureRiverUpdatedEmission)
				if err := _AdventureRiver.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
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
func (_AdventureRiver *AdventureRiverFilterer) ParseUpdatedEmission(log types.Log) (*AdventureRiverUpdatedEmission, error) {
	event := new(AdventureRiverUpdatedEmission)
	if err := _AdventureRiver.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureRiverUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the AdventureRiver contract.
type AdventureRiverUpdatedOwnerIterator struct {
	Event *AdventureRiverUpdatedOwner // Event containing the contract specifics and raw log

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
func (it *AdventureRiverUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureRiverUpdatedOwner)
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
		it.Event = new(AdventureRiverUpdatedOwner)
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
func (it *AdventureRiverUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureRiverUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureRiverUpdatedOwner represents a UpdatedOwner event raised by the AdventureRiver contract.
type AdventureRiverUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureRiver *AdventureRiverFilterer) FilterUpdatedOwner(opts *bind.FilterOpts, owner []common.Address) (*AdventureRiverUpdatedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureRiver.contract.FilterLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AdventureRiverUpdatedOwnerIterator{contract: _AdventureRiver.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureRiver *AdventureRiverFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *AdventureRiverUpdatedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureRiver.contract.WatchLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureRiverUpdatedOwner)
				if err := _AdventureRiver.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
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
func (_AdventureRiver *AdventureRiverFilterer) ParseUpdatedOwner(log types.Log) (*AdventureRiverUpdatedOwner, error) {
	event := new(AdventureRiverUpdatedOwner)
	if err := _AdventureRiver.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
