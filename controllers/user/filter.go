package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	UserFilter struct {
		User      string              `validate:"required,eth_addr"`
		Adventure constants.Adventure `validate:"oneof=0 1 2 3 4 5"`
	}
)

func ValidateFilter(filter UserFilter) error {
	validate := validator.New()
	return validate.Struct(filter)
}
