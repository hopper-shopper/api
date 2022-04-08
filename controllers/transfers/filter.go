package transfers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	TransferDirection int

	TransfersFilter struct {
		Direction TransferDirection `validate:"oneof=0 1"`
		User      string            `validate:"required,eth_addr"`
	}
)

const (
	FromUser TransferDirection = iota
	ToUser
)

func TransferDirectionFromString(direction string) TransferDirection {
	lowerCased := strings.ToLower(direction)

	switch lowerCased {
	case "out", "from":
		return FromUser
	case "in", "to":
		return ToUser
	default:
		return FromUser
	}
}
func (direction TransferDirection) String() string {
	switch direction {
	case FromUser:
		return "out"
	case ToUser:
		return "in"
	default:
		return "out"
	}
}

func ValidateFilter(filter TransfersFilter) error {
	validate := validator.New()
	return validate.Struct(filter)
}
