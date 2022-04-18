package supply

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	SupplyType int

	SupplyFilter struct {
		Type SupplyType `validate:"required,oneof=1"`
	}
)

const (
	UnknownSupplyType SupplyType = iota
	FlySupplyType
)

func SupplyTypeFromString(supplyType string) SupplyType {
	lowerCased := strings.ToLower(supplyType)

	switch lowerCased {
	case "fly":
		return FlySupplyType
	default:
		return UnknownSupplyType
	}
}
func (supplyType SupplyType) String() string {
	switch supplyType {
	case FlySupplyType:
		return "fly"
	default:
		return "unknown"
	}
}

func ValidateFilter(filter any) error {
	validate := validator.New()
	return validate.Struct(filter)
}
