package contracts

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/steschwa/hopper-analytics-api/constants"
)

type (
	OnChainClient struct {
		Connection *ethclient.Client
	}

	ZoneContract interface {
		TotalBaseShare(opts *bind.CallOpts) (*big.Int, error)
	}
)

func NewOnChainClient() (*OnChainClient, error) {
	client, err := ethclient.Dial(constants.AVAX_RPC)
	if err != nil {
		return &OnChainClient{}, err
	}

	return &OnChainClient{
		Connection: client,
	}, nil
}

// ----------------------------------------
// Adventure callers getter
// ----------------------------------------

func (client *OnChainClient) getPondCaller() (*AdventurePondCaller, error) {
	return NewAdventurePondCaller(common.HexToAddress(constants.ADVENTURE_POND_CONTRACT), client.Connection)
}
func (client *OnChainClient) getStreamCaller() (*AdventureStreamCaller, error) {
	return NewAdventureStreamCaller(common.HexToAddress(constants.ADVENTURE_STREAM_CONTRACT), client.Connection)
}
func (client *OnChainClient) getSwampCaller() (*AdventureSwampCaller, error) {
	return NewAdventureSwampCaller(common.HexToAddress(constants.ADVENTURE_SWAMP_CONTRACT), client.Connection)
}
func (client *OnChainClient) getRiverCaller() (*AdventureRiverCaller, error) {
	return NewAdventureRiverCaller(common.HexToAddress(constants.ADVENTURE_RIVER_CONTRACT), client.Connection)
}
func (client *OnChainClient) getForestCaller() (*AdventureForestCaller, error) {
	return NewAdventureForestCaller(common.HexToAddress(constants.ADVENTURE_FOREST_CONTRACT), client.Connection)
}
func (client *OnChainClient) getGreatLakeCaller() (*AdventureGreatLakeCaller, error) {
	return NewAdventureGreatLakeCaller(common.HexToAddress(constants.ADVENTURE_GREAT_LAKE_CONTRACT), client.Connection)
}

func (client *OnChainClient) getAdventureCaller(adventure constants.Adventure) (ZoneContract, error) {
	switch adventure {
	case constants.AdventurePond:
		return client.getPondCaller()
	case constants.AdventureStream:
		return client.getStreamCaller()
	case constants.AdventureSwamp:
		return client.getSwampCaller()
	case constants.AdventureRiver:
		return client.getRiverCaller()
	case constants.AdventureForest:
		return client.getForestCaller()
	case constants.AdventureGreatLake:
		return client.getGreatLakeCaller()
	}

	// should never happen
	return nil, errors.New("invalid adventure")
}

// ----------------------------------------
// Contract read wrappers
// ----------------------------------------

func (client *OnChainClient) GetTotalBaseShare(adventure constants.Adventure) (*big.Int, error) {
	caller, err := client.getAdventureCaller(adventure)

	if err != nil {
		return big.NewInt(0), err
	}

	return caller.TotalBaseShare(nil)
}
