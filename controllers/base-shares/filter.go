package baseshares

import (
	"github.com/go-playground/validator/v10"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	BaseShareByAdventureFilter struct {
		Adventure constants.Adventure `validate:"required,oneof=0 1 2 3 4 5"`
	}
)

func ValidateFilter(filter any) error {
	validate := validator.New()
	return validate.Struct(filter)
}
