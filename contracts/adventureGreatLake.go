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

// AdventureGreatLakeMetaData contains all meta data concerning the AdventureGreatLake contract.
var AdventureGreatLakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vefly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hopper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NoHopperStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnfitHopper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongTokenID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"}],\"name\":\"UpdatedBallot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionRate\",\"type\":\"uint256\"}],\"name\":\"UpdatedEmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOPPER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEVEL_GAUGE_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"baseSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEmissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"}],\"name\":\"canEnter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"changeHopperName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"emergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flyLevelCapRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"forceUnvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"generatedPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHopperAndGauge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"hopperGauge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gaugeLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLevel\",\"type\":\"uint256\"}],\"name\":\"getLevelUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBonusShares\",\"type\":\"uint256\"}],\"name\":\"getUserBonusGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBaseShares\",\"type\":\"uint256\"}],\"name\":\"getUserGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBonusUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"levelUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"}],\"name\":\"setBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"setBonusEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emissionRate\",\"type\":\"uint256\"}],\"name\":\"setEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flyLevelCapRatio\",\"type\":\"uint256\"}],\"name\":\"setFlyLevelCapRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenCapFilledPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBonusRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMaxFlyGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veFlyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdventureGreatLakeABI is the input ABI used to generate the binding from.
// Deprecated: Use AdventureGreatLakeMetaData.ABI instead.
var AdventureGreatLakeABI = AdventureGreatLakeMetaData.ABI

// AdventureGreatLake is an auto generated Go binding around an Ethereum contract.
type AdventureGreatLake struct {
	AdventureGreatLakeCaller     // Read-only binding to the contract
	AdventureGreatLakeTransactor // Write-only binding to the contract
	AdventureGreatLakeFilterer   // Log filterer for contract events
}

// AdventureGreatLakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdventureGreatLakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureGreatLakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdventureGreatLakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureGreatLakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdventureGreatLakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventureGreatLakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdventureGreatLakeSession struct {
	Contract     *AdventureGreatLake // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AdventureGreatLakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdventureGreatLakeCallerSession struct {
	Contract *AdventureGreatLakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AdventureGreatLakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdventureGreatLakeTransactorSession struct {
	Contract     *AdventureGreatLakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AdventureGreatLakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdventureGreatLakeRaw struct {
	Contract *AdventureGreatLake // Generic contract binding to access the raw methods on
}

// AdventureGreatLakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdventureGreatLakeCallerRaw struct {
	Contract *AdventureGreatLakeCaller // Generic read-only contract binding to access the raw methods on
}

// AdventureGreatLakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdventureGreatLakeTransactorRaw struct {
	Contract *AdventureGreatLakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdventureGreatLake creates a new instance of AdventureGreatLake, bound to a specific deployed contract.
func NewAdventureGreatLake(address common.Address, backend bind.ContractBackend) (*AdventureGreatLake, error) {
	contract, err := bindAdventureGreatLake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLake{AdventureGreatLakeCaller: AdventureGreatLakeCaller{contract: contract}, AdventureGreatLakeTransactor: AdventureGreatLakeTransactor{contract: contract}, AdventureGreatLakeFilterer: AdventureGreatLakeFilterer{contract: contract}}, nil
}

// NewAdventureGreatLakeCaller creates a new read-only instance of AdventureGreatLake, bound to a specific deployed contract.
func NewAdventureGreatLakeCaller(address common.Address, caller bind.ContractCaller) (*AdventureGreatLakeCaller, error) {
	contract, err := bindAdventureGreatLake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLakeCaller{contract: contract}, nil
}

// NewAdventureGreatLakeTransactor creates a new write-only instance of AdventureGreatLake, bound to a specific deployed contract.
func NewAdventureGreatLakeTransactor(address common.Address, transactor bind.ContractTransactor) (*AdventureGreatLakeTransactor, error) {
	contract, err := bindAdventureGreatLake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLakeTransactor{contract: contract}, nil
}

// NewAdventureGreatLakeFilterer creates a new log filterer instance of AdventureGreatLake, bound to a specific deployed contract.
func NewAdventureGreatLakeFilterer(address common.Address, filterer bind.ContractFilterer) (*AdventureGreatLakeFilterer, error) {
	contract, err := bindAdventureGreatLake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLakeFilterer{contract: contract}, nil
}

// bindAdventureGreatLake binds a generic wrapper to an already deployed contract.
func bindAdventureGreatLake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdventureGreatLakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureGreatLake *AdventureGreatLakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureGreatLake.Contract.AdventureGreatLakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureGreatLake *AdventureGreatLakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.AdventureGreatLakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureGreatLake *AdventureGreatLakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.AdventureGreatLakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventureGreatLake *AdventureGreatLakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventureGreatLake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventureGreatLake *AdventureGreatLakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventureGreatLake *AdventureGreatLakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.contract.Transact(opts, method, params...)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCaller) FLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeSession) FLY() (common.Address, error) {
	return _AdventureGreatLake.Contract.FLY(&_AdventureGreatLake.CallOpts)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) FLY() (common.Address, error) {
	return _AdventureGreatLake.Contract.FLY(&_AdventureGreatLake.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCaller) HOPPER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "HOPPER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeSession) HOPPER() (common.Address, error) {
	return _AdventureGreatLake.Contract.HOPPER(&_AdventureGreatLake.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) HOPPER() (common.Address, error) {
	return _AdventureGreatLake.Contract.HOPPER(&_AdventureGreatLake.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureGreatLake *AdventureGreatLakeCaller) LEVELGAUGEKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "LEVEL_GAUGE_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureGreatLake *AdventureGreatLakeSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureGreatLake.Contract.LEVELGAUGEKEY(&_AdventureGreatLake.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) LEVELGAUGEKEY() (string, error) {
	return _AdventureGreatLake.Contract.LEVELGAUGEKEY(&_AdventureGreatLake.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCaller) VEFLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "VE_FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeSession) VEFLY() (common.Address, error) {
	return _AdventureGreatLake.Contract.VEFLY(&_AdventureGreatLake.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) VEFLY() (common.Address, error) {
	return _AdventureGreatLake.Contract.VEFLY(&_AdventureGreatLake.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCaller) Ballot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "ballot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeSession) Ballot() (common.Address, error) {
	return _AdventureGreatLake.Contract.Ballot(&_AdventureGreatLake.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) Ballot() (common.Address, error) {
	return _AdventureGreatLake.Contract.Ballot(&_AdventureGreatLake.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) BaseRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "baseRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BaseRewardPerShare(&_AdventureGreatLake.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BaseRewardPerShare(&_AdventureGreatLake.CallOpts)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) BaseSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "baseSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.BaseSharesBalance(&_AdventureGreatLake.CallOpts, arg0)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.BaseSharesBalance(&_AdventureGreatLake.CallOpts, arg0)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) BonusEmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "bonusEmissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BonusEmissionRate(&_AdventureGreatLake.CallOpts)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BonusEmissionRate(&_AdventureGreatLake.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) BonusRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "bonusRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BonusRewardPerShare(&_AdventureGreatLake.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BonusRewardPerShare(&_AdventureGreatLake.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) BonusRewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "bonusRewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BonusRewardPerShareStored(&_AdventureGreatLake.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventureGreatLake.Contract.BonusRewardPerShareStored(&_AdventureGreatLake.CallOpts)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureGreatLake *AdventureGreatLakeCaller) CanEnter(opts *bind.CallOpts, hopper HopperNFTHopper) (bool, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "canEnter", hopper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureGreatLake *AdventureGreatLakeSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureGreatLake.Contract.CanEnter(&_AdventureGreatLake.CallOpts, hopper)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventureGreatLake.Contract.CanEnter(&_AdventureGreatLake.CallOpts, hopper)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.Claimable(&_AdventureGreatLake.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.Claimable(&_AdventureGreatLake.CallOpts, _account)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureGreatLake *AdventureGreatLakeCaller) Emergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "emergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureGreatLake *AdventureGreatLakeSession) Emergency() (bool, error) {
	return _AdventureGreatLake.Contract.Emergency(&_AdventureGreatLake.CallOpts)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) Emergency() (bool, error) {
	return _AdventureGreatLake.Contract.Emergency(&_AdventureGreatLake.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) EmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "emissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) EmissionRate() (*big.Int, error) {
	return _AdventureGreatLake.Contract.EmissionRate(&_AdventureGreatLake.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) EmissionRate() (*big.Int, error) {
	return _AdventureGreatLake.Contract.EmissionRate(&_AdventureGreatLake.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) FlyLevelCapRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "flyLevelCapRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureGreatLake.Contract.FlyLevelCapRatio(&_AdventureGreatLake.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventureGreatLake.Contract.FlyLevelCapRatio(&_AdventureGreatLake.CallOpts)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) GeneratedPerShareStored(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "generatedPerShareStored", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.GeneratedPerShareStored(&_AdventureGreatLake.CallOpts, arg0)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.GeneratedPerShareStored(&_AdventureGreatLake.CallOpts, arg0)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureGreatLake *AdventureGreatLakeCaller) GetHopperAndGauge(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "getHopperAndGauge", tokenId)

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
func (_AdventureGreatLake *AdventureGreatLakeSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureGreatLake.Contract.GetHopperAndGauge(&_AdventureGreatLake.CallOpts, tokenId)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventureGreatLake.Contract.GetHopperAndGauge(&_AdventureGreatLake.CallOpts, tokenId)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) GetLevelUpCost(opts *bind.CallOpts, currentLevel *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "getLevelUpCost", currentLevel)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureGreatLake.Contract.GetLevelUpCost(&_AdventureGreatLake.CallOpts, currentLevel)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventureGreatLake.Contract.GetLevelUpCost(&_AdventureGreatLake.CallOpts, currentLevel)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) GetUserBonusGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "getUserBonusGeneratedFly", account, _totalUserBonusShares)

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
func (_AdventureGreatLake *AdventureGreatLakeSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureGreatLake.Contract.GetUserBonusGeneratedFly(&_AdventureGreatLake.CallOpts, account, _totalUserBonusShares)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureGreatLake.Contract.GetUserBonusGeneratedFly(&_AdventureGreatLake.CallOpts, account, _totalUserBonusShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) GetUserGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "getUserGeneratedFly", account, _totalUserBaseShares)

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
func (_AdventureGreatLake *AdventureGreatLakeSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureGreatLake.Contract.GetUserGeneratedFly(&_AdventureGreatLake.CallOpts, account, _totalUserBaseShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventureGreatLake.Contract.GetUserGeneratedFly(&_AdventureGreatLake.CallOpts, account, _totalUserBaseShares)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) HopperBaseShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "hopperBaseShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureGreatLake.Contract.HopperBaseShare(&_AdventureGreatLake.CallOpts, arg0)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureGreatLake.Contract.HopperBaseShare(&_AdventureGreatLake.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCaller) HopperOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "hopperOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureGreatLake.Contract.HopperOwners(&_AdventureGreatLake.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventureGreatLake.Contract.HopperOwners(&_AdventureGreatLake.CallOpts, arg0)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) LastBonusUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "lastBonusUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureGreatLake.Contract.LastBonusUpdatedTime(&_AdventureGreatLake.CallOpts)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventureGreatLake.Contract.LastBonusUpdatedTime(&_AdventureGreatLake.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) LastUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "lastUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureGreatLake.Contract.LastUpdatedTime(&_AdventureGreatLake.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventureGreatLake.Contract.LastUpdatedTime(&_AdventureGreatLake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeSession) Owner() (common.Address, error) {
	return _AdventureGreatLake.Contract.Owner(&_AdventureGreatLake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) Owner() (common.Address, error) {
	return _AdventureGreatLake.Contract.Owner(&_AdventureGreatLake.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) RewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "rewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureGreatLake.Contract.RewardPerShareStored(&_AdventureGreatLake.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventureGreatLake.Contract.RewardPerShareStored(&_AdventureGreatLake.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.Rewards(&_AdventureGreatLake.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.Rewards(&_AdventureGreatLake.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) TokenCapFilledPerShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "tokenCapFilledPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureGreatLake.Contract.TokenCapFilledPerShare(&_AdventureGreatLake.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventureGreatLake.Contract.TokenCapFilledPerShare(&_AdventureGreatLake.CallOpts, arg0)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) TotalBaseShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "totalBaseShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.TotalBaseShare(&_AdventureGreatLake.CallOpts)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) TotalBaseShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.TotalBaseShare(&_AdventureGreatLake.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) TotalVeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "totalVeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) TotalVeShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.TotalVeShare(&_AdventureGreatLake.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) TotalVeShare() (*big.Int, error) {
	return _AdventureGreatLake.Contract.TotalVeShare(&_AdventureGreatLake.CallOpts)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) UserBonusRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "userBonusRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.UserBonusRewardPerSharePaid(&_AdventureGreatLake.CallOpts, arg0)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.UserBonusRewardPerSharePaid(&_AdventureGreatLake.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) UserMaxFlyGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "userMaxFlyGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.UserMaxFlyGeneration(&_AdventureGreatLake.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.UserMaxFlyGeneration(&_AdventureGreatLake.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) UserRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "userRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.UserRewardPerSharePaid(&_AdventureGreatLake.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.UserRewardPerSharePaid(&_AdventureGreatLake.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) VeFlyBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "veFlyBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.VeFlyBalance(&_AdventureGreatLake.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.VeFlyBalance(&_AdventureGreatLake.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCaller) VeSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventureGreatLake.contract.Call(opts, &out, "veSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.VeSharesBalance(&_AdventureGreatLake.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventureGreatLake *AdventureGreatLakeCallerSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventureGreatLake.Contract.VeSharesBalance(&_AdventureGreatLake.CallOpts, arg0)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) ChangeHopperName(opts *bind.TransactOpts, tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "changeHopperName", tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.ChangeHopperName(&_AdventureGreatLake.TransactOpts, tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.ChangeHopperName(&_AdventureGreatLake.TransactOpts, tokenId, name, useOwnRewards)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) Claim() (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Claim(&_AdventureGreatLake.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) Claim() (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Claim(&_AdventureGreatLake.TransactOpts)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) EmergencyExit(opts *bind.TransactOpts, tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "emergencyExit", tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.EmergencyExit(&_AdventureGreatLake.TransactOpts, tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.EmergencyExit(&_AdventureGreatLake.TransactOpts, tokenIds, user)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) EnableEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "enableEmergency")
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.EnableEmergency(&_AdventureGreatLake.TransactOpts)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.EnableEmergency(&_AdventureGreatLake.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) Enter(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "enter", tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Enter(&_AdventureGreatLake.TransactOpts, tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Enter(&_AdventureGreatLake.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) Exit(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "exit", tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Exit(&_AdventureGreatLake.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Exit(&_AdventureGreatLake.TransactOpts, tokenIds)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) ForceUnvote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "forceUnvote", user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.ForceUnvote(&_AdventureGreatLake.TransactOpts, user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.ForceUnvote(&_AdventureGreatLake.TransactOpts, user)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) LevelUp(opts *bind.TransactOpts, tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "levelUp", tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.LevelUp(&_AdventureGreatLake.TransactOpts, tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.LevelUp(&_AdventureGreatLake.TransactOpts, tokenId, useOwnRewards)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) SetBallot(opts *bind.TransactOpts, _ballot common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "setBallot", _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetBallot(&_AdventureGreatLake.TransactOpts, _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetBallot(&_AdventureGreatLake.TransactOpts, _ballot)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) SetBonusEmissionRate(opts *bind.TransactOpts, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "setBonusEmissionRate", _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetBonusEmissionRate(&_AdventureGreatLake.TransactOpts, _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetBonusEmissionRate(&_AdventureGreatLake.TransactOpts, _bonusEmissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) SetEmissionRate(opts *bind.TransactOpts, _emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "setEmissionRate", _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetEmissionRate(&_AdventureGreatLake.TransactOpts, _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetEmissionRate(&_AdventureGreatLake.TransactOpts, _emissionRate)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) SetFlyLevelCapRatio(opts *bind.TransactOpts, _flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "setFlyLevelCapRatio", _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetFlyLevelCapRatio(&_AdventureGreatLake.TransactOpts, _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetFlyLevelCapRatio(&_AdventureGreatLake.TransactOpts, _flyLevelCapRatio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetOwner(&_AdventureGreatLake.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.SetOwner(&_AdventureGreatLake.TransactOpts, _owner)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) Unvote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "unvote", veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Unvote(&_AdventureGreatLake.TransactOpts, veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Unvote(&_AdventureGreatLake.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactor) Vote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureGreatLake.contract.Transact(opts, "vote", veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureGreatLake *AdventureGreatLakeSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Vote(&_AdventureGreatLake.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventureGreatLake *AdventureGreatLakeTransactorSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventureGreatLake.Contract.Vote(&_AdventureGreatLake.TransactOpts, veFlyAmount, recount)
}

// AdventureGreatLakeUpdatedBallotIterator is returned from FilterUpdatedBallot and is used to iterate over the raw logs and unpacked data for UpdatedBallot events raised by the AdventureGreatLake contract.
type AdventureGreatLakeUpdatedBallotIterator struct {
	Event *AdventureGreatLakeUpdatedBallot // Event containing the contract specifics and raw log

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
func (it *AdventureGreatLakeUpdatedBallotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureGreatLakeUpdatedBallot)
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
		it.Event = new(AdventureGreatLakeUpdatedBallot)
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
func (it *AdventureGreatLakeUpdatedBallotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureGreatLakeUpdatedBallotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureGreatLakeUpdatedBallot represents a UpdatedBallot event raised by the AdventureGreatLake contract.
type AdventureGreatLakeUpdatedBallot struct {
	Ballot common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBallot is a free log retrieval operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureGreatLake *AdventureGreatLakeFilterer) FilterUpdatedBallot(opts *bind.FilterOpts, ballot []common.Address) (*AdventureGreatLakeUpdatedBallotIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureGreatLake.contract.FilterLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLakeUpdatedBallotIterator{contract: _AdventureGreatLake.contract, event: "UpdatedBallot", logs: logs, sub: sub}, nil
}

// WatchUpdatedBallot is a free log subscription operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventureGreatLake *AdventureGreatLakeFilterer) WatchUpdatedBallot(opts *bind.WatchOpts, sink chan<- *AdventureGreatLakeUpdatedBallot, ballot []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventureGreatLake.contract.WatchLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureGreatLakeUpdatedBallot)
				if err := _AdventureGreatLake.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
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
func (_AdventureGreatLake *AdventureGreatLakeFilterer) ParseUpdatedBallot(log types.Log) (*AdventureGreatLakeUpdatedBallot, error) {
	event := new(AdventureGreatLakeUpdatedBallot)
	if err := _AdventureGreatLake.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureGreatLakeUpdatedEmissionIterator is returned from FilterUpdatedEmission and is used to iterate over the raw logs and unpacked data for UpdatedEmission events raised by the AdventureGreatLake contract.
type AdventureGreatLakeUpdatedEmissionIterator struct {
	Event *AdventureGreatLakeUpdatedEmission // Event containing the contract specifics and raw log

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
func (it *AdventureGreatLakeUpdatedEmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureGreatLakeUpdatedEmission)
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
		it.Event = new(AdventureGreatLakeUpdatedEmission)
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
func (it *AdventureGreatLakeUpdatedEmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureGreatLakeUpdatedEmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureGreatLakeUpdatedEmission represents a UpdatedEmission event raised by the AdventureGreatLake contract.
type AdventureGreatLakeUpdatedEmission struct {
	EmissionRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEmission is a free log retrieval operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureGreatLake *AdventureGreatLakeFilterer) FilterUpdatedEmission(opts *bind.FilterOpts) (*AdventureGreatLakeUpdatedEmissionIterator, error) {

	logs, sub, err := _AdventureGreatLake.contract.FilterLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLakeUpdatedEmissionIterator{contract: _AdventureGreatLake.contract, event: "UpdatedEmission", logs: logs, sub: sub}, nil
}

// WatchUpdatedEmission is a free log subscription operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventureGreatLake *AdventureGreatLakeFilterer) WatchUpdatedEmission(opts *bind.WatchOpts, sink chan<- *AdventureGreatLakeUpdatedEmission) (event.Subscription, error) {

	logs, sub, err := _AdventureGreatLake.contract.WatchLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureGreatLakeUpdatedEmission)
				if err := _AdventureGreatLake.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
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
func (_AdventureGreatLake *AdventureGreatLakeFilterer) ParseUpdatedEmission(log types.Log) (*AdventureGreatLakeUpdatedEmission, error) {
	event := new(AdventureGreatLakeUpdatedEmission)
	if err := _AdventureGreatLake.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventureGreatLakeUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the AdventureGreatLake contract.
type AdventureGreatLakeUpdatedOwnerIterator struct {
	Event *AdventureGreatLakeUpdatedOwner // Event containing the contract specifics and raw log

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
func (it *AdventureGreatLakeUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventureGreatLakeUpdatedOwner)
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
		it.Event = new(AdventureGreatLakeUpdatedOwner)
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
func (it *AdventureGreatLakeUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventureGreatLakeUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventureGreatLakeUpdatedOwner represents a UpdatedOwner event raised by the AdventureGreatLake contract.
type AdventureGreatLakeUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureGreatLake *AdventureGreatLakeFilterer) FilterUpdatedOwner(opts *bind.FilterOpts, owner []common.Address) (*AdventureGreatLakeUpdatedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureGreatLake.contract.FilterLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AdventureGreatLakeUpdatedOwnerIterator{contract: _AdventureGreatLake.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventureGreatLake *AdventureGreatLakeFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *AdventureGreatLakeUpdatedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventureGreatLake.contract.WatchLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventureGreatLakeUpdatedOwner)
				if err := _AdventureGreatLake.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
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
func (_AdventureGreatLake *AdventureGreatLakeFilterer) ParseUpdatedOwner(log types.Log) (*AdventureGreatLakeUpdatedOwner, error) {
	event := new(AdventureGreatLakeUpdatedOwner)
	if err := _AdventureGreatLake.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
