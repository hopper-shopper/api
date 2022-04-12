package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	UserCapFilter struct {
		User      string              `validate:"required,eth_addr"`
		Adventure constants.Adventure `validate:"oneof=0 1 2 3 4 5"`
	}

	UserEarningsFilter struct {
		User      string              `validate:"required,eth_addr"`
		Adventure constants.Adventure `validate:"oneof=0 1 2 3 4 5"`
	}
)

func ValidateFilter(filter any) error {
	validate := validator.New()
	return validate.Struct(filter)
}
