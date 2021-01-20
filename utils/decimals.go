package utils

import "github.com/shopspring/decimal"

func Mul(values... decimal.Decimal) decimal.Decimal {
	result := decimal.NewFromInt(1)

	for _, value := range values {
		result = result.Mul(value)
	}

	return result
}