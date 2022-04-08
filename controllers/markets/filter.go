package markets

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	MarketsFilter struct {
		TokenIds []string `validate:"required,min=1,dive,numeric,min=1,max=4"`
	}
)

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
