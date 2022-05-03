package prices

import (
	"github.com/go-playground/validator/v10"
)

type (
	HistoricalPriceFilter struct {
		Times    []string `validate:"required,min=1,dive,datetime=2006-01-02T15:04:05Z07:00"`
		Coin     string   `validate:"required,oneof=avax fly"`
		Currency string   `validate:"required,oneof=usd eur"`
	}
)

func ValidateFilter(filter HistoricalPriceFilter) error {
	validate := validator.New()
	return validate.Struct(filter)
}
