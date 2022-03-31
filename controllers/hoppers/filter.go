package hoppers

import "strings"

type (
	AdventureFilter int
	MarketFilter    int
	PermitFilter    int

	HoppersFilter struct {
		Adventure AdventureFilter
		Market    MarketFilter
		Permit    PermitFilter
	}
)

const (
	AnyAdventure AdventureFilter = iota
	PondAdventure
	StreamAdventure
	SwampAdventure
	RiverAdventure
	ForestAdventure
	GreatLakeAdventure

	AnyMarket MarketFilter = iota
	OnMarket
	OffMarket

	AnyPermit PermitFilter = iota
	RiverPermit
	ForestPermit
	GreatLakePermit
)

func AdventureFilterFromString(adventure string) AdventureFilter {
	lowerCased := strings.ToLower(adventure)

	switch lowerCased {
	case "pond":
		return PondAdventure
	case "stream":
		return StreamAdventure
	case "swamp":
		return SwampAdventure
	case "river":
		return RiverAdventure
	case "forest":
		return ForestAdventure
	case "great-lake":
		return GreatLakeAdventure
	default:
		return AnyAdventure
	}
}
func (adventureFilter AdventureFilter) String() string {
	switch adventureFilter {
	case PondAdventure:
		return "pond"
	case StreamAdventure:
		return "stream"
	case SwampAdventure:
		return "swamp"
	case RiverAdventure:
		return "river"
	case ForestAdventure:
		return "forest"
	case GreatLakeAdventure:
		return "great-lake"
	default:
		return "any"
	}
}

func MarketFilterFromString(market string) MarketFilter {
	lowerCased := strings.ToLower(market)

	switch lowerCased {
	case "1", "true", "on", "yes":
		return OnMarket
	case "0", "false", "off", "no":
		return OffMarket
	default:
		return AnyMarket
	}
}

func (marketFilter MarketFilter) String() string {
	switch marketFilter {
	case OnMarket:
		return "yes"
	case OffMarket:
		return "no"
	default:
		return "any"
	}
}

func PermitFilterFromString(permit string) PermitFilter {
	lowerCased := strings.ToLower(permit)

	switch lowerCased {
	case "river":
		return RiverPermit
	case "forest":
		return ForestPermit
	case "great-lake":
		return GreatLakePermit
	default:
		return AnyPermit
	}
}

func (permitFilter PermitFilter) String() string {
	switch permitFilter {
	case RiverPermit:
		return "river"
	case ForestPermit:
		return "forest"
	case GreatLakePermit:
		return "great-lake"
	default:
		return "any"
	}
}
