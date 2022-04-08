package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func ToDecimal(value string, decimals int) decimal.Decimal {
	v := new(big.Int)
	v.SetString(value, 10)

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(v.String())
	result := num.Div(mul)

	return result
}
