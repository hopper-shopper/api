package user

import (
	"strings"

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

	UserEarningsHistoryFilter struct {
		User      string              `validate:"required,eth_addr"`
		Adventure constants.Adventure `validate:"oneof=0 1 2 3 4 5"`
		GroupBy   GroupTransfersBy    `validate:"required,oneof=1 2"`
	}

	GroupTransfersBy int
)

const (
	GroupTransfersByUnknown GroupTransfersBy = iota
	GroupTransfersByDay
	GroupTransfersByMonth
)

func GroupTransfersByFromString(groupBy string) GroupTransfersBy {
	lowerCased := strings.ToLower(groupBy)

	switch lowerCased {
	case "day":
		return GroupTransfersByDay
	case "month":
		return GroupTransfersByMonth
	default:
		return GroupTransfersByUnknown
	}
}
func (groupTransfersBy GroupTransfersBy) String() string {
	switch groupTransfersBy {
	case GroupTransfersByDay:
		return "day"
	case GroupTransfersByMonth:
		return "month"
	default:
		return "unknown"
	}
}

func ValidateFilter(filter any) error {
	validate := validator.New()
	return validate.Struct(filter)
}
