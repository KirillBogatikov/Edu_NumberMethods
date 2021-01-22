package gauss

import (
	"github.com/shopspring/decimal"
)

type Equation struct {
	values []decimal.Decimal
	result decimal.Decimal
}

func NewEquation(values []float64, result float64) *Equation {
	newValues := make([]decimal.Decimal, 0)
	for _, value := range values {
		newValues = append(newValues, decimal.NewFromFloat(value))
	}

	return &Equation{newValues, decimal.NewFromFloat(result)}
}

func (e *Equation) Values() []decimal.Decimal {
	return e.values
}

func (e *Equation) Result() decimal.Decimal {
	return e.result
}
