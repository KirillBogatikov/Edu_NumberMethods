package utils

import (
	"github.com/shopspring/decimal"
	"math"
)

var (
	One   = decimal.NewFromInt(1)
	Two   = decimal.NewFromInt(2)
	Three = decimal.NewFromInt(3)
)

func Mul(values ...decimal.Decimal) decimal.Decimal {
	result := decimal.NewFromInt(1)

	for _, value := range values {
		result = result.Mul(value)
	}

	return result
}

func Sqrt(v decimal.Decimal) decimal.Decimal {
	f, _ := v.Float64()
	return decimal.NewFromFloat(math.Sqrt(f))
}
