package transfers

import "github.com/steschwa/hopper-analytics-collector/constants"

type (
	ContractHandle int
	ContractMethod int
)

const (
	UnknownContract ContractHandle = iota
	PondContract
	StreamContract
	SwampContract
	RiverContract
	ForestContract
	GreatLakeContract
	VeFlyContract
)

const (
	MethodUnknown ContractMethod = iota
	MethodLevelUp
	MethodClaim
	MethodVeFlyVote
	MethodVeFlyDeposit
	MethodBreedingEnter
)

func (handle ContractHandle) String() string {
	switch handle {
	case PondContract:
		return "pond"
	case StreamContract:
		return "stream"
	case SwampContract:
		return "swamp"
	case RiverContract:
		return "river"
	case ForestContract:
		return "forest"
	case GreatLakeContract:
		return "great-lake"
	default:
		return "unknown"
	}
}

func (method ContractMethod) String() string {
	switch method {
	case MethodLevelUp:
		return "levelup"
	case MethodClaim:
		return "claim"
	case MethodVeFlyVote:
		return "vefly-vote"
	case MethodVeFlyDeposit:
		return "vefly-deposit"
	case MethodBreedingEnter:
		return "breeding"
	default:
		return "unknown"
	}
}

func GetContractHandle(contractAddr string) ContractHandle {
	switch contractAddr {
	case constants.ADVENTURE_POND_CONTRACT:
		return PondContract
	case constants.ADVENTURE_STREAM_CONTRACT:
		return StreamContract
	case constants.ADVENTURE_SWAMP_CONTRACT:
		return SwampContract
	case constants.ADVENTURE_RIVER_CONTRACT:
		return RiverContract
	case constants.ADVENTURE_FOREST_CONTRACT:
		return ForestContract
	case constants.ADVENTURE_GREAT_LAKE_CONTRACT:
		return GreatLakeContract
	case constants.VE_FLY_CONTRACT:
		return VeFlyContract
	default:
		return UnknownContract
	}
}

func GetContractMethod(methodId string) ContractMethod {
	if len(methodId) != 10 {
		return MethodUnknown
	}

	switch methodId {
	case "0x4e71d92d":
		return MethodClaim
	case "0x0c679fa0":
		return MethodLevelUp
	case "0xc9d27afe":
		return MethodVeFlyVote
	case "0xb6b55f25":
		return MethodVeFlyDeposit
	case "0xa59f3e0c":
		return MethodBreedingEnter
	default:
		return MethodUnknown
	}
}
