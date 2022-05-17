package controllers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	ResponseFormat string

	ResponseFormatOptions struct {
		Format ResponseFormat `validate:"oneof=json csv"`
	}
)

const (
	RESPONSE_JSON ResponseFormat = "json"
	RESPONSE_CSV  ResponseFormat = "csv"
)

func ResponseFormatFromString(format string) ResponseFormat {
	lowerCased := strings.ToLower(format)

	switch lowerCased {
	case "json":
		return RESPONSE_JSON
	case "csv":
		return RESPONSE_CSV
	default:
		return RESPONSE_JSON
	}
}

func ValidateFilter(formatOptions ResponseFormatOptions) error {
	validate := validator.New()
	return validate.Struct(formatOptions)
}
