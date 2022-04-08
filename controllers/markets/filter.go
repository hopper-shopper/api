package markets

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	SoldFilter    int
	MarketsFilter struct {
		TokenIds []string   `validate:"required,min=1,dive,numeric,min=1,max=4"`
		Sold     SoldFilter `validate:"oneof=0 1 2"`
	}
)

const (
	AnySold SoldFilter = iota
	AlreadySold
	NotSold
)

func SoldFilterFromString(sold string) SoldFilter {
	lowerCased := strings.ToLower(sold)

	switch lowerCased {
	case "1", "true", "yes":
		return AlreadySold
	case "0", "false", "no":
		return NotSold
	default:
		return AnySold
	}
}

func (soldFilter SoldFilter) String() string {
	switch soldFilter {
	case AlreadySold:
		return "yes"
	case NotSold:
		return "no"
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

func ValidateFilter(filter MarketsFilter) error {
	validate := validator.New()
	return validate.Struct(filter)
}
