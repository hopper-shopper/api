package hoppers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	AdventureFilter string
	MarketFilter    string
	PermitFilter    string

	HoppersFilter struct {
		Adventure AdventureFilter `validate:"oneof=any pond stream swamp river forest great-lake"`
		Market    MarketFilter    `validate:"oneof=any yes no"`
		Permit    PermitFilter    `validate:"oneof=any river forest great-lake"`
		TokenIds  []string        `validate:"omitempty,min=0,dive,numeric,min=1,max=4"`
		Owner     string          `validate:"omitempty,eth_addr"`
	}
)

const (
	AnyAdventure       AdventureFilter = "any"
	PondAdventure      AdventureFilter = "pond"
	StreamAdventure    AdventureFilter = "stream"
	SwampAdventure     AdventureFilter = "swamp"
	RiverAdventure     AdventureFilter = "river"
	ForestAdventure    AdventureFilter = "forest"
	GreatLakeAdventure AdventureFilter = "great-lake"
)

const (
	AnyMarket MarketFilter = "any"
	OnMarket  MarketFilter = "yes"
	OffMarket MarketFilter = "no"
)

const (
	AnyPermit       PermitFilter = "any"
	RiverPermit     PermitFilter = "river"
	ForestPermit    PermitFilter = "forest"
	GreatLakePermit PermitFilter = "great-lake"
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

func MarketFilterFromString(market string) MarketFilter {
	lowerCased := strings.ToLower(market)

	switch lowerCased {
	case "1", "true", "yes":
		return OnMarket
	case "0", "false", "no":
		return OffMarket
	default:
		return AnyMarket
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
