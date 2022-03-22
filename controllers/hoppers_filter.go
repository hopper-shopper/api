package controllers

import "strings"

type (
	AdventureFilter int
	MarketFilter    int

	HoppersFilter struct {
		Adventure AdventureFilter
		Market    MarketFilter
	}
)

const (
	AnyAdventure AdventureFilter = iota
	RiverAdventure
	ForestAdventure
	GreatLakeAdventure

	AnyMarket MarketFilter = iota
	OnMarket
	OffMarket
)

func AdventureFilterFromString(adventure string) AdventureFilter {
	lowerCased := strings.ToLower(adventure)

	switch lowerCased {
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
