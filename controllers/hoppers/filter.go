package hoppers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	AdventureFilter int
	MarketFilter    int
	PermitFilter    int

	HoppersFilter struct {
		Adventure AdventureFilter `validate:"oneof=0 1 2 3 4 5 6"`
		Market    MarketFilter    `validate:"oneof=0 1 2"`
		Permit    PermitFilter    `validate:"oneof=0 1 2 3"`
		TokenIds  []string        `validate:"omitempty,min=0,dive,numeric,min=1,max=4"`
		Owner     string          `validate:"omitempty,eth_addr"`
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
)

const (
	AnyMarket MarketFilter = iota
	OnMarket
	OffMarket
)

const (
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

func TokenIdsFilterFromString(text string) []string {
	splittableText := strings.Replace(text, " ", "", -1)
	if len(splittableText) == 0 {
		return []string{}
	}

	tokenIds := strings.Split(splittableText, ",")

	for i, tokenId := range tokenIds {
		tokenIds[i] = strings.Replace(tokenId, " ", "", -1)
	}

	return tokenIds
}

func ValidateFilter(filter HoppersFilter) error {
	validate := validator.New()
	return validate.Struct(filter)
}
