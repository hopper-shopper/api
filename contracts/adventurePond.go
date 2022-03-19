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

// HopperNFTHopper is an auto generated low-level Go binding around an user-defined struct.
type HopperNFTHopper struct {
	Level        *big.Int
	Rebirths     uint16
	Strength     uint8
	Agility      uint8
	Vitality     uint8
	Intelligence uint8
	Fertility    uint8
}

// AdventurePondMetaData contains all meta data concerning the AdventurePond contract.
var AdventurePondMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vefly\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hopper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NoHopperStaked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnfitHopper\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongTokenID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ballot\",\"type\":\"address\"}],\"name\":\"UpdatedBallot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionRate\",\"type\":\"uint256\"}],\"name\":\"UpdatedEmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UpdatedOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOPPER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEVEL_GAUGE_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_FLY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"baseSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEmissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusRewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"}],\"name\":\"canEnter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"changeHopperName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"emergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flyLevelCapRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"forceUnvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"generatedPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getHopperAndGauge\",\"outputs\":[{\"components\":[{\"internalType\":\"uint200\",\"name\":\"level\",\"type\":\"uint200\"},{\"internalType\":\"uint16\",\"name\":\"rebirths\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"strength\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"agility\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"vitality\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"intelligence\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"fertility\",\"type\":\"uint8\"}],\"internalType\":\"structHopperNFT.Hopper\",\"name\":\"hopper\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"hopperGauge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gaugeLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentLevel\",\"type\":\"uint256\"}],\"name\":\"getLevelUpCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBonusShares\",\"type\":\"uint256\"}],\"name\":\"getUserBonusGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalUserBaseShares\",\"type\":\"uint256\"}],\"name\":\"getUserGeneratedFly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hopperOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBonusUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useOwnRewards\",\"type\":\"bool\"}],\"name\":\"levelUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerShareStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"}],\"name\":\"setBallot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bonusEmissionRate\",\"type\":\"uint256\"}],\"name\":\"setBonusEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emissionRate\",\"type\":\"uint256\"}],\"name\":\"setEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flyLevelCapRatio\",\"type\":\"uint256\"}],\"name\":\"setFlyLevelCapRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenCapFilledPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBaseShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"unvote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBonusRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMaxFlyGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerSharePaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veFlyBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"veSharesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"veFlyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"recount\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdventurePondABI is the input ABI used to generate the binding from.
// Deprecated: Use AdventurePondMetaData.ABI instead.
var AdventurePondABI = AdventurePondMetaData.ABI

// AdventurePond is an auto generated Go binding around an Ethereum contract.
type AdventurePond struct {
	AdventurePondCaller     // Read-only binding to the contract
	AdventurePondTransactor // Write-only binding to the contract
	AdventurePondFilterer   // Log filterer for contract events
}

// AdventurePondCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdventurePondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventurePondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdventurePondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventurePondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdventurePondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdventurePondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdventurePondSession struct {
	Contract     *AdventurePond    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdventurePondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdventurePondCallerSession struct {
	Contract *AdventurePondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdventurePondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdventurePondTransactorSession struct {
	Contract     *AdventurePondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdventurePondRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdventurePondRaw struct {
	Contract *AdventurePond // Generic contract binding to access the raw methods on
}

// AdventurePondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdventurePondCallerRaw struct {
	Contract *AdventurePondCaller // Generic read-only contract binding to access the raw methods on
}

// AdventurePondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdventurePondTransactorRaw struct {
	Contract *AdventurePondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdventurePond creates a new instance of AdventurePond, bound to a specific deployed contract.
func NewAdventurePond(address common.Address, backend bind.ContractBackend) (*AdventurePond, error) {
	contract, err := bindAdventurePond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdventurePond{AdventurePondCaller: AdventurePondCaller{contract: contract}, AdventurePondTransactor: AdventurePondTransactor{contract: contract}, AdventurePondFilterer: AdventurePondFilterer{contract: contract}}, nil
}

// NewAdventurePondCaller creates a new read-only instance of AdventurePond, bound to a specific deployed contract.
func NewAdventurePondCaller(address common.Address, caller bind.ContractCaller) (*AdventurePondCaller, error) {
	contract, err := bindAdventurePond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdventurePondCaller{contract: contract}, nil
}

// NewAdventurePondTransactor creates a new write-only instance of AdventurePond, bound to a specific deployed contract.
func NewAdventurePondTransactor(address common.Address, transactor bind.ContractTransactor) (*AdventurePondTransactor, error) {
	contract, err := bindAdventurePond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdventurePondTransactor{contract: contract}, nil
}

// NewAdventurePondFilterer creates a new log filterer instance of AdventurePond, bound to a specific deployed contract.
func NewAdventurePondFilterer(address common.Address, filterer bind.ContractFilterer) (*AdventurePondFilterer, error) {
	contract, err := bindAdventurePond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdventurePondFilterer{contract: contract}, nil
}

// bindAdventurePond binds a generic wrapper to an already deployed contract.
func bindAdventurePond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdventurePondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventurePond *AdventurePondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventurePond.Contract.AdventurePondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventurePond *AdventurePondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventurePond.Contract.AdventurePondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventurePond *AdventurePondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventurePond.Contract.AdventurePondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdventurePond *AdventurePondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdventurePond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdventurePond *AdventurePondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventurePond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdventurePond *AdventurePondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdventurePond.Contract.contract.Transact(opts, method, params...)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventurePond *AdventurePondCaller) FLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventurePond *AdventurePondSession) FLY() (common.Address, error) {
	return _AdventurePond.Contract.FLY(&_AdventurePond.CallOpts)
}

// FLY is a free data retrieval call binding the contract method 0xe1a61a0c.
//
// Solidity: function FLY() view returns(address)
func (_AdventurePond *AdventurePondCallerSession) FLY() (common.Address, error) {
	return _AdventurePond.Contract.FLY(&_AdventurePond.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventurePond *AdventurePondCaller) HOPPER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "HOPPER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventurePond *AdventurePondSession) HOPPER() (common.Address, error) {
	return _AdventurePond.Contract.HOPPER(&_AdventurePond.CallOpts)
}

// HOPPER is a free data retrieval call binding the contract method 0x4b905e3f.
//
// Solidity: function HOPPER() view returns(address)
func (_AdventurePond *AdventurePondCallerSession) HOPPER() (common.Address, error) {
	return _AdventurePond.Contract.HOPPER(&_AdventurePond.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventurePond *AdventurePondCaller) LEVELGAUGEKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "LEVEL_GAUGE_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventurePond *AdventurePondSession) LEVELGAUGEKEY() (string, error) {
	return _AdventurePond.Contract.LEVELGAUGEKEY(&_AdventurePond.CallOpts)
}

// LEVELGAUGEKEY is a free data retrieval call binding the contract method 0xa49a3e1b.
//
// Solidity: function LEVEL_GAUGE_KEY() view returns(string)
func (_AdventurePond *AdventurePondCallerSession) LEVELGAUGEKEY() (string, error) {
	return _AdventurePond.Contract.LEVELGAUGEKEY(&_AdventurePond.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventurePond *AdventurePondCaller) VEFLY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "VE_FLY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventurePond *AdventurePondSession) VEFLY() (common.Address, error) {
	return _AdventurePond.Contract.VEFLY(&_AdventurePond.CallOpts)
}

// VEFLY is a free data retrieval call binding the contract method 0x2677d07f.
//
// Solidity: function VE_FLY() view returns(address)
func (_AdventurePond *AdventurePondCallerSession) VEFLY() (common.Address, error) {
	return _AdventurePond.Contract.VEFLY(&_AdventurePond.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventurePond *AdventurePondCaller) Ballot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "ballot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventurePond *AdventurePondSession) Ballot() (common.Address, error) {
	return _AdventurePond.Contract.Ballot(&_AdventurePond.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_AdventurePond *AdventurePondCallerSession) Ballot() (common.Address, error) {
	return _AdventurePond.Contract.Ballot(&_AdventurePond.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) BaseRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "baseRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventurePond *AdventurePondSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventurePond.Contract.BaseRewardPerShare(&_AdventurePond.CallOpts)
}

// BaseRewardPerShare is a free data retrieval call binding the contract method 0x0b242163.
//
// Solidity: function baseRewardPerShare() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) BaseRewardPerShare() (*big.Int, error) {
	return _AdventurePond.Contract.BaseRewardPerShare(&_AdventurePond.CallOpts)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) BaseSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "baseSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.BaseSharesBalance(&_AdventurePond.CallOpts, arg0)
}

// BaseSharesBalance is a free data retrieval call binding the contract method 0x215e0f4d.
//
// Solidity: function baseSharesBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) BaseSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.BaseSharesBalance(&_AdventurePond.CallOpts, arg0)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) BonusEmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "bonusEmissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventurePond *AdventurePondSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventurePond.Contract.BonusEmissionRate(&_AdventurePond.CallOpts)
}

// BonusEmissionRate is a free data retrieval call binding the contract method 0x43aba89d.
//
// Solidity: function bonusEmissionRate() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) BonusEmissionRate() (*big.Int, error) {
	return _AdventurePond.Contract.BonusEmissionRate(&_AdventurePond.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) BonusRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "bonusRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventurePond *AdventurePondSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventurePond.Contract.BonusRewardPerShare(&_AdventurePond.CallOpts)
}

// BonusRewardPerShare is a free data retrieval call binding the contract method 0xe532a154.
//
// Solidity: function bonusRewardPerShare() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) BonusRewardPerShare() (*big.Int, error) {
	return _AdventurePond.Contract.BonusRewardPerShare(&_AdventurePond.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) BonusRewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "bonusRewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventurePond *AdventurePondSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventurePond.Contract.BonusRewardPerShareStored(&_AdventurePond.CallOpts)
}

// BonusRewardPerShareStored is a free data retrieval call binding the contract method 0x1c1502c4.
//
// Solidity: function bonusRewardPerShareStored() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) BonusRewardPerShareStored() (*big.Int, error) {
	return _AdventurePond.Contract.BonusRewardPerShareStored(&_AdventurePond.CallOpts)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventurePond *AdventurePondCaller) CanEnter(opts *bind.CallOpts, hopper HopperNFTHopper) (bool, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "canEnter", hopper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventurePond *AdventurePondSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventurePond.Contract.CanEnter(&_AdventurePond.CallOpts, hopper)
}

// CanEnter is a free data retrieval call binding the contract method 0x769022a3.
//
// Solidity: function canEnter((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper) pure returns(bool)
func (_AdventurePond *AdventurePondCallerSession) CanEnter(hopper HopperNFTHopper) (bool, error) {
	return _AdventurePond.Contract.CanEnter(&_AdventurePond.CallOpts, hopper)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventurePond *AdventurePondSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.Claimable(&_AdventurePond.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.Claimable(&_AdventurePond.CallOpts, _account)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventurePond *AdventurePondCaller) Emergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "emergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventurePond *AdventurePondSession) Emergency() (bool, error) {
	return _AdventurePond.Contract.Emergency(&_AdventurePond.CallOpts)
}

// Emergency is a free data retrieval call binding the contract method 0xcaa6fea4.
//
// Solidity: function emergency() view returns(bool)
func (_AdventurePond *AdventurePondCallerSession) Emergency() (bool, error) {
	return _AdventurePond.Contract.Emergency(&_AdventurePond.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) EmissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "emissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventurePond *AdventurePondSession) EmissionRate() (*big.Int, error) {
	return _AdventurePond.Contract.EmissionRate(&_AdventurePond.CallOpts)
}

// EmissionRate is a free data retrieval call binding the contract method 0x96afc450.
//
// Solidity: function emissionRate() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) EmissionRate() (*big.Int, error) {
	return _AdventurePond.Contract.EmissionRate(&_AdventurePond.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) FlyLevelCapRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "flyLevelCapRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventurePond *AdventurePondSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventurePond.Contract.FlyLevelCapRatio(&_AdventurePond.CallOpts)
}

// FlyLevelCapRatio is a free data retrieval call binding the contract method 0x5d35b5d5.
//
// Solidity: function flyLevelCapRatio() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) FlyLevelCapRatio() (*big.Int, error) {
	return _AdventurePond.Contract.FlyLevelCapRatio(&_AdventurePond.CallOpts)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) GeneratedPerShareStored(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "generatedPerShareStored", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.GeneratedPerShareStored(&_AdventurePond.CallOpts, arg0)
}

// GeneratedPerShareStored is a free data retrieval call binding the contract method 0x3cfa989e.
//
// Solidity: function generatedPerShareStored(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) GeneratedPerShareStored(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.GeneratedPerShareStored(&_AdventurePond.CallOpts, arg0)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventurePond *AdventurePondCaller) GetHopperAndGauge(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "getHopperAndGauge", tokenId)

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
func (_AdventurePond *AdventurePondSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventurePond.Contract.GetHopperAndGauge(&_AdventurePond.CallOpts, tokenId)
}

// GetHopperAndGauge is a free data retrieval call binding the contract method 0xad914e2e.
//
// Solidity: function getHopperAndGauge(uint256 tokenId) view returns((uint200,uint16,uint8,uint8,uint8,uint8,uint8) hopper, uint256 hopperGauge, uint256 gaugeLimit)
func (_AdventurePond *AdventurePondCallerSession) GetHopperAndGauge(tokenId *big.Int) (struct {
	Hopper      HopperNFTHopper
	HopperGauge *big.Int
	GaugeLimit  *big.Int
}, error) {
	return _AdventurePond.Contract.GetHopperAndGauge(&_AdventurePond.CallOpts, tokenId)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventurePond *AdventurePondCaller) GetLevelUpCost(opts *bind.CallOpts, currentLevel *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "getLevelUpCost", currentLevel)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventurePond *AdventurePondSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventurePond.Contract.GetLevelUpCost(&_AdventurePond.CallOpts, currentLevel)
}

// GetLevelUpCost is a free data retrieval call binding the contract method 0xbf9bc0b2.
//
// Solidity: function getLevelUpCost(uint256 currentLevel) pure returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) GetLevelUpCost(currentLevel *big.Int) (*big.Int, error) {
	return _AdventurePond.Contract.GetLevelUpCost(&_AdventurePond.CallOpts, currentLevel)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventurePond *AdventurePondCaller) GetUserBonusGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "getUserBonusGeneratedFly", account, _totalUserBonusShares)

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
func (_AdventurePond *AdventurePondSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventurePond.Contract.GetUserBonusGeneratedFly(&_AdventurePond.CallOpts, account, _totalUserBonusShares)
}

// GetUserBonusGeneratedFly is a free data retrieval call binding the contract method 0xab20355c.
//
// Solidity: function getUserBonusGeneratedFly(address account, uint256 _totalUserBonusShares) view returns(uint256, uint256)
func (_AdventurePond *AdventurePondCallerSession) GetUserBonusGeneratedFly(account common.Address, _totalUserBonusShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventurePond.Contract.GetUserBonusGeneratedFly(&_AdventurePond.CallOpts, account, _totalUserBonusShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventurePond *AdventurePondCaller) GetUserGeneratedFly(opts *bind.CallOpts, account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "getUserGeneratedFly", account, _totalUserBaseShares)

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
func (_AdventurePond *AdventurePondSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventurePond.Contract.GetUserGeneratedFly(&_AdventurePond.CallOpts, account, _totalUserBaseShares)
}

// GetUserGeneratedFly is a free data retrieval call binding the contract method 0x79f409e0.
//
// Solidity: function getUserGeneratedFly(address account, uint256 _totalUserBaseShares) view returns(uint256, uint256)
func (_AdventurePond *AdventurePondCallerSession) GetUserGeneratedFly(account common.Address, _totalUserBaseShares *big.Int) (*big.Int, *big.Int, error) {
	return _AdventurePond.Contract.GetUserGeneratedFly(&_AdventurePond.CallOpts, account, _totalUserBaseShares)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) HopperBaseShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "hopperBaseShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventurePond.Contract.HopperBaseShare(&_AdventurePond.CallOpts, arg0)
}

// HopperBaseShare is a free data retrieval call binding the contract method 0x968c603d.
//
// Solidity: function hopperBaseShare(uint256 ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) HopperBaseShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventurePond.Contract.HopperBaseShare(&_AdventurePond.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventurePond *AdventurePondCaller) HopperOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "hopperOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventurePond *AdventurePondSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventurePond.Contract.HopperOwners(&_AdventurePond.CallOpts, arg0)
}

// HopperOwners is a free data retrieval call binding the contract method 0x201e29a0.
//
// Solidity: function hopperOwners(uint256 ) view returns(address)
func (_AdventurePond *AdventurePondCallerSession) HopperOwners(arg0 *big.Int) (common.Address, error) {
	return _AdventurePond.Contract.HopperOwners(&_AdventurePond.CallOpts, arg0)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) LastBonusUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "lastBonusUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventurePond *AdventurePondSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventurePond.Contract.LastBonusUpdatedTime(&_AdventurePond.CallOpts)
}

// LastBonusUpdatedTime is a free data retrieval call binding the contract method 0x5834c36d.
//
// Solidity: function lastBonusUpdatedTime() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) LastBonusUpdatedTime() (*big.Int, error) {
	return _AdventurePond.Contract.LastBonusUpdatedTime(&_AdventurePond.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) LastUpdatedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "lastUpdatedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventurePond *AdventurePondSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventurePond.Contract.LastUpdatedTime(&_AdventurePond.CallOpts)
}

// LastUpdatedTime is a free data retrieval call binding the contract method 0xbf856895.
//
// Solidity: function lastUpdatedTime() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) LastUpdatedTime() (*big.Int, error) {
	return _AdventurePond.Contract.LastUpdatedTime(&_AdventurePond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventurePond *AdventurePondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventurePond *AdventurePondSession) Owner() (common.Address, error) {
	return _AdventurePond.Contract.Owner(&_AdventurePond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdventurePond *AdventurePondCallerSession) Owner() (common.Address, error) {
	return _AdventurePond.Contract.Owner(&_AdventurePond.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) RewardPerShareStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "rewardPerShareStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventurePond *AdventurePondSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventurePond.Contract.RewardPerShareStored(&_AdventurePond.CallOpts)
}

// RewardPerShareStored is a free data retrieval call binding the contract method 0x0ab747f0.
//
// Solidity: function rewardPerShareStored() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) RewardPerShareStored() (*big.Int, error) {
	return _AdventurePond.Contract.RewardPerShareStored(&_AdventurePond.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.Rewards(&_AdventurePond.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.Rewards(&_AdventurePond.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) TokenCapFilledPerShare(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "tokenCapFilledPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventurePond.Contract.TokenCapFilledPerShare(&_AdventurePond.CallOpts, arg0)
}

// TokenCapFilledPerShare is a free data retrieval call binding the contract method 0xdee93568.
//
// Solidity: function tokenCapFilledPerShare(uint256 ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) TokenCapFilledPerShare(arg0 *big.Int) (*big.Int, error) {
	return _AdventurePond.Contract.TokenCapFilledPerShare(&_AdventurePond.CallOpts, arg0)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) TotalBaseShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "totalBaseShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventurePond *AdventurePondSession) TotalBaseShare() (*big.Int, error) {
	return _AdventurePond.Contract.TotalBaseShare(&_AdventurePond.CallOpts)
}

// TotalBaseShare is a free data retrieval call binding the contract method 0x3bad0f71.
//
// Solidity: function totalBaseShare() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) TotalBaseShare() (*big.Int, error) {
	return _AdventurePond.Contract.TotalBaseShare(&_AdventurePond.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventurePond *AdventurePondCaller) TotalVeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "totalVeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventurePond *AdventurePondSession) TotalVeShare() (*big.Int, error) {
	return _AdventurePond.Contract.TotalVeShare(&_AdventurePond.CallOpts)
}

// TotalVeShare is a free data retrieval call binding the contract method 0x4b9ca431.
//
// Solidity: function totalVeShare() view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) TotalVeShare() (*big.Int, error) {
	return _AdventurePond.Contract.TotalVeShare(&_AdventurePond.CallOpts)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) UserBonusRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "userBonusRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.UserBonusRewardPerSharePaid(&_AdventurePond.CallOpts, arg0)
}

// UserBonusRewardPerSharePaid is a free data retrieval call binding the contract method 0xee569ba9.
//
// Solidity: function userBonusRewardPerSharePaid(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) UserBonusRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.UserBonusRewardPerSharePaid(&_AdventurePond.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) UserMaxFlyGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "userMaxFlyGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.UserMaxFlyGeneration(&_AdventurePond.CallOpts, arg0)
}

// UserMaxFlyGeneration is a free data retrieval call binding the contract method 0x5c1ab8c0.
//
// Solidity: function userMaxFlyGeneration(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) UserMaxFlyGeneration(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.UserMaxFlyGeneration(&_AdventurePond.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) UserRewardPerSharePaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "userRewardPerSharePaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.UserRewardPerSharePaid(&_AdventurePond.CallOpts, arg0)
}

// UserRewardPerSharePaid is a free data retrieval call binding the contract method 0x818ae1ce.
//
// Solidity: function userRewardPerSharePaid(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) UserRewardPerSharePaid(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.UserRewardPerSharePaid(&_AdventurePond.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) VeFlyBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "veFlyBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.VeFlyBalance(&_AdventurePond.CallOpts, arg0)
}

// VeFlyBalance is a free data retrieval call binding the contract method 0x7e2d9af8.
//
// Solidity: function veFlyBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) VeFlyBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.VeFlyBalance(&_AdventurePond.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCaller) VeSharesBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdventurePond.contract.Call(opts, &out, "veSharesBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.VeSharesBalance(&_AdventurePond.CallOpts, arg0)
}

// VeSharesBalance is a free data retrieval call binding the contract method 0x2df97550.
//
// Solidity: function veSharesBalance(address ) view returns(uint256)
func (_AdventurePond *AdventurePondCallerSession) VeSharesBalance(arg0 common.Address) (*big.Int, error) {
	return _AdventurePond.Contract.VeSharesBalance(&_AdventurePond.CallOpts, arg0)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventurePond *AdventurePondTransactor) ChangeHopperName(opts *bind.TransactOpts, tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "changeHopperName", tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventurePond *AdventurePondSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.ChangeHopperName(&_AdventurePond.TransactOpts, tokenId, name, useOwnRewards)
}

// ChangeHopperName is a paid mutator transaction binding the contract method 0x8eb3be28.
//
// Solidity: function changeHopperName(uint256 tokenId, string name, bool useOwnRewards) returns()
func (_AdventurePond *AdventurePondTransactorSession) ChangeHopperName(tokenId *big.Int, name string, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.ChangeHopperName(&_AdventurePond.TransactOpts, tokenId, name, useOwnRewards)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventurePond *AdventurePondTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventurePond *AdventurePondSession) Claim() (*types.Transaction, error) {
	return _AdventurePond.Contract.Claim(&_AdventurePond.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdventurePond *AdventurePondTransactorSession) Claim() (*types.Transaction, error) {
	return _AdventurePond.Contract.Claim(&_AdventurePond.TransactOpts)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventurePond *AdventurePondTransactor) EmergencyExit(opts *bind.TransactOpts, tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "emergencyExit", tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventurePond *AdventurePondSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.EmergencyExit(&_AdventurePond.TransactOpts, tokenIds, user)
}

// EmergencyExit is a paid mutator transaction binding the contract method 0xd881e959.
//
// Solidity: function emergencyExit(uint256[] tokenIds, address user) returns()
func (_AdventurePond *AdventurePondTransactorSession) EmergencyExit(tokenIds []*big.Int, user common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.EmergencyExit(&_AdventurePond.TransactOpts, tokenIds, user)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventurePond *AdventurePondTransactor) EnableEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "enableEmergency")
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventurePond *AdventurePondSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventurePond.Contract.EnableEmergency(&_AdventurePond.TransactOpts)
}

// EnableEmergency is a paid mutator transaction binding the contract method 0xe42c4e63.
//
// Solidity: function enableEmergency() returns()
func (_AdventurePond *AdventurePondTransactorSession) EnableEmergency() (*types.Transaction, error) {
	return _AdventurePond.Contract.EnableEmergency(&_AdventurePond.TransactOpts)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventurePond *AdventurePondTransactor) Enter(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "enter", tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventurePond *AdventurePondSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.Enter(&_AdventurePond.TransactOpts, tokenIds)
}

// Enter is a paid mutator transaction binding the contract method 0x7944135d.
//
// Solidity: function enter(uint256[] tokenIds) returns()
func (_AdventurePond *AdventurePondTransactorSession) Enter(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.Enter(&_AdventurePond.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventurePond *AdventurePondTransactor) Exit(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "exit", tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventurePond *AdventurePondSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.Exit(&_AdventurePond.TransactOpts, tokenIds)
}

// Exit is a paid mutator transaction binding the contract method 0x18c08f26.
//
// Solidity: function exit(uint256[] tokenIds) returns()
func (_AdventurePond *AdventurePondTransactorSession) Exit(tokenIds []*big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.Exit(&_AdventurePond.TransactOpts, tokenIds)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventurePond *AdventurePondTransactor) ForceUnvote(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "forceUnvote", user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventurePond *AdventurePondSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.ForceUnvote(&_AdventurePond.TransactOpts, user)
}

// ForceUnvote is a paid mutator transaction binding the contract method 0xf5544eb7.
//
// Solidity: function forceUnvote(address user) returns()
func (_AdventurePond *AdventurePondTransactorSession) ForceUnvote(user common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.ForceUnvote(&_AdventurePond.TransactOpts, user)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventurePond *AdventurePondTransactor) LevelUp(opts *bind.TransactOpts, tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "levelUp", tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventurePond *AdventurePondSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.LevelUp(&_AdventurePond.TransactOpts, tokenId, useOwnRewards)
}

// LevelUp is a paid mutator transaction binding the contract method 0x0c679fa0.
//
// Solidity: function levelUp(uint256 tokenId, bool useOwnRewards) returns()
func (_AdventurePond *AdventurePondTransactorSession) LevelUp(tokenId *big.Int, useOwnRewards bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.LevelUp(&_AdventurePond.TransactOpts, tokenId, useOwnRewards)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventurePond *AdventurePondTransactor) SetBallot(opts *bind.TransactOpts, _ballot common.Address) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "setBallot", _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventurePond *AdventurePondSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetBallot(&_AdventurePond.TransactOpts, _ballot)
}

// SetBallot is a paid mutator transaction binding the contract method 0xd051bfa3.
//
// Solidity: function setBallot(address _ballot) returns()
func (_AdventurePond *AdventurePondTransactorSession) SetBallot(_ballot common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetBallot(&_AdventurePond.TransactOpts, _ballot)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventurePond *AdventurePondTransactor) SetBonusEmissionRate(opts *bind.TransactOpts, _bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "setBonusEmissionRate", _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventurePond *AdventurePondSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetBonusEmissionRate(&_AdventurePond.TransactOpts, _bonusEmissionRate)
}

// SetBonusEmissionRate is a paid mutator transaction binding the contract method 0x9d4c93c4.
//
// Solidity: function setBonusEmissionRate(uint256 _bonusEmissionRate) returns()
func (_AdventurePond *AdventurePondTransactorSession) SetBonusEmissionRate(_bonusEmissionRate *big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetBonusEmissionRate(&_AdventurePond.TransactOpts, _bonusEmissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventurePond *AdventurePondTransactor) SetEmissionRate(opts *bind.TransactOpts, _emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "setEmissionRate", _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventurePond *AdventurePondSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetEmissionRate(&_AdventurePond.TransactOpts, _emissionRate)
}

// SetEmissionRate is a paid mutator transaction binding the contract method 0xa1bdb15e.
//
// Solidity: function setEmissionRate(uint256 _emissionRate) returns()
func (_AdventurePond *AdventurePondTransactorSession) SetEmissionRate(_emissionRate *big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetEmissionRate(&_AdventurePond.TransactOpts, _emissionRate)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventurePond *AdventurePondTransactor) SetFlyLevelCapRatio(opts *bind.TransactOpts, _flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "setFlyLevelCapRatio", _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventurePond *AdventurePondSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetFlyLevelCapRatio(&_AdventurePond.TransactOpts, _flyLevelCapRatio)
}

// SetFlyLevelCapRatio is a paid mutator transaction binding the contract method 0x6351520e.
//
// Solidity: function setFlyLevelCapRatio(uint256 _flyLevelCapRatio) returns()
func (_AdventurePond *AdventurePondTransactorSession) SetFlyLevelCapRatio(_flyLevelCapRatio *big.Int) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetFlyLevelCapRatio(&_AdventurePond.TransactOpts, _flyLevelCapRatio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventurePond *AdventurePondTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventurePond *AdventurePondSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetOwner(&_AdventurePond.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AdventurePond *AdventurePondTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AdventurePond.Contract.SetOwner(&_AdventurePond.TransactOpts, _owner)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventurePond *AdventurePondTransactor) Unvote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "unvote", veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventurePond *AdventurePondSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.Unvote(&_AdventurePond.TransactOpts, veFlyAmount, recount)
}

// Unvote is a paid mutator transaction binding the contract method 0x591adce6.
//
// Solidity: function unvote(uint256 veFlyAmount, bool recount) returns()
func (_AdventurePond *AdventurePondTransactorSession) Unvote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.Unvote(&_AdventurePond.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventurePond *AdventurePondTransactor) Vote(opts *bind.TransactOpts, veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventurePond.contract.Transact(opts, "vote", veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventurePond *AdventurePondSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.Vote(&_AdventurePond.TransactOpts, veFlyAmount, recount)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 veFlyAmount, bool recount) returns()
func (_AdventurePond *AdventurePondTransactorSession) Vote(veFlyAmount *big.Int, recount bool) (*types.Transaction, error) {
	return _AdventurePond.Contract.Vote(&_AdventurePond.TransactOpts, veFlyAmount, recount)
}

// AdventurePondUpdatedBallotIterator is returned from FilterUpdatedBallot and is used to iterate over the raw logs and unpacked data for UpdatedBallot events raised by the AdventurePond contract.
type AdventurePondUpdatedBallotIterator struct {
	Event *AdventurePondUpdatedBallot // Event containing the contract specifics and raw log

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
func (it *AdventurePondUpdatedBallotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventurePondUpdatedBallot)
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
		it.Event = new(AdventurePondUpdatedBallot)
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
func (it *AdventurePondUpdatedBallotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventurePondUpdatedBallotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventurePondUpdatedBallot represents a UpdatedBallot event raised by the AdventurePond contract.
type AdventurePondUpdatedBallot struct {
	Ballot common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedBallot is a free log retrieval operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventurePond *AdventurePondFilterer) FilterUpdatedBallot(opts *bind.FilterOpts, ballot []common.Address) (*AdventurePondUpdatedBallotIterator, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventurePond.contract.FilterLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return &AdventurePondUpdatedBallotIterator{contract: _AdventurePond.contract, event: "UpdatedBallot", logs: logs, sub: sub}, nil
}

// WatchUpdatedBallot is a free log subscription operation binding the contract event 0x605aa74980dc92c245d8959d233f8c2d7062d874e49326a42e7418279cc8d1f8.
//
// Solidity: event UpdatedBallot(address indexed ballot)
func (_AdventurePond *AdventurePondFilterer) WatchUpdatedBallot(opts *bind.WatchOpts, sink chan<- *AdventurePondUpdatedBallot, ballot []common.Address) (event.Subscription, error) {

	var ballotRule []interface{}
	for _, ballotItem := range ballot {
		ballotRule = append(ballotRule, ballotItem)
	}

	logs, sub, err := _AdventurePond.contract.WatchLogs(opts, "UpdatedBallot", ballotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventurePondUpdatedBallot)
				if err := _AdventurePond.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
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
func (_AdventurePond *AdventurePondFilterer) ParseUpdatedBallot(log types.Log) (*AdventurePondUpdatedBallot, error) {
	event := new(AdventurePondUpdatedBallot)
	if err := _AdventurePond.contract.UnpackLog(event, "UpdatedBallot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventurePondUpdatedEmissionIterator is returned from FilterUpdatedEmission and is used to iterate over the raw logs and unpacked data for UpdatedEmission events raised by the AdventurePond contract.
type AdventurePondUpdatedEmissionIterator struct {
	Event *AdventurePondUpdatedEmission // Event containing the contract specifics and raw log

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
func (it *AdventurePondUpdatedEmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventurePondUpdatedEmission)
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
		it.Event = new(AdventurePondUpdatedEmission)
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
func (it *AdventurePondUpdatedEmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventurePondUpdatedEmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventurePondUpdatedEmission represents a UpdatedEmission event raised by the AdventurePond contract.
type AdventurePondUpdatedEmission struct {
	EmissionRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEmission is a free log retrieval operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventurePond *AdventurePondFilterer) FilterUpdatedEmission(opts *bind.FilterOpts) (*AdventurePondUpdatedEmissionIterator, error) {

	logs, sub, err := _AdventurePond.contract.FilterLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return &AdventurePondUpdatedEmissionIterator{contract: _AdventurePond.contract, event: "UpdatedEmission", logs: logs, sub: sub}, nil
}

// WatchUpdatedEmission is a free log subscription operation binding the contract event 0x90e51e492b1841b8ed3d459463d02d171ced752c7fcd84a9dd79f90098166aec.
//
// Solidity: event UpdatedEmission(uint256 emissionRate)
func (_AdventurePond *AdventurePondFilterer) WatchUpdatedEmission(opts *bind.WatchOpts, sink chan<- *AdventurePondUpdatedEmission) (event.Subscription, error) {

	logs, sub, err := _AdventurePond.contract.WatchLogs(opts, "UpdatedEmission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventurePondUpdatedEmission)
				if err := _AdventurePond.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
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
func (_AdventurePond *AdventurePondFilterer) ParseUpdatedEmission(log types.Log) (*AdventurePondUpdatedEmission, error) {
	event := new(AdventurePondUpdatedEmission)
	if err := _AdventurePond.contract.UnpackLog(event, "UpdatedEmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdventurePondUpdatedOwnerIterator is returned from FilterUpdatedOwner and is used to iterate over the raw logs and unpacked data for UpdatedOwner events raised by the AdventurePond contract.
type AdventurePondUpdatedOwnerIterator struct {
	Event *AdventurePondUpdatedOwner // Event containing the contract specifics and raw log

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
func (it *AdventurePondUpdatedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdventurePondUpdatedOwner)
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
		it.Event = new(AdventurePondUpdatedOwner)
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
func (it *AdventurePondUpdatedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdventurePondUpdatedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdventurePondUpdatedOwner represents a UpdatedOwner event raised by the AdventurePond contract.
type AdventurePondUpdatedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOwner is a free log retrieval operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventurePond *AdventurePondFilterer) FilterUpdatedOwner(opts *bind.FilterOpts, owner []common.Address) (*AdventurePondUpdatedOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventurePond.contract.FilterLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AdventurePondUpdatedOwnerIterator{contract: _AdventurePond.contract, event: "UpdatedOwner", logs: logs, sub: sub}, nil
}

// WatchUpdatedOwner is a free log subscription operation binding the contract event 0x957090e72c0a1b3ebf83c682eb8c1f88c2a18cd0578b91a819efb28859f0f3a3.
//
// Solidity: event UpdatedOwner(address indexed owner)
func (_AdventurePond *AdventurePondFilterer) WatchUpdatedOwner(opts *bind.WatchOpts, sink chan<- *AdventurePondUpdatedOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _AdventurePond.contract.WatchLogs(opts, "UpdatedOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdventurePondUpdatedOwner)
				if err := _AdventurePond.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
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
func (_AdventurePond *AdventurePondFilterer) ParseUpdatedOwner(log types.Log) (*AdventurePondUpdatedOwner, error) {
	event := new(AdventurePondUpdatedOwner)
	if err := _AdventurePond.contract.UnpackLog(event, "UpdatedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
