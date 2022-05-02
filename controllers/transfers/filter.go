package transfers

import (
	"github.com/go-playground/validator/v10"
	"github.com/steschwa/hopper-analytics-collector/constants"
)

type (
	TransfersFilter struct {
		Direction constants.TransferDirection `validate:"oneof=0 1"`
		User      string                      `validate:"required,eth_addr"`
		Method    constants.TransferMethod    `validate:"oneof=0 1 2 3 4 5 6 7 8 9 10 11"`
	}
)

func ValidateFilter(filter TransfersFilter) error {
	validate := validator.New()
	return validate.Struct(filter)
}
