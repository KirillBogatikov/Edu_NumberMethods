package equation

import (
	"NumberMethods/utils"
	"github.com/shopspring/decimal"
)

func (e *Equation) process(vp utils.ValueProcessor) *Equation {
	newValues := make([]decimal.Decimal, 0)
	for _, value := range e.values {
		newValues = append(newValues, vp(value))
	}

	return &Equation{newValues, vp(e.result)}
}

func (e *Equation) processEquation(other *Equation, cp utils.CellProcessor) *Equation {
	newValues := make([]decimal.Decimal, 0)
	for i, value := range e.values {
		newValues = append(newValues, cp(other.values[i], value))
	}

	return &Equation{newValues, cp(other.result, e.result)}
}

func (e *Equation) Div(d decimal.Decimal) *Equation {
	return e.process(func(i decimal.Decimal) decimal.Decimal {
		return i.Div(d)
	})
}

func (e *Equation) Mul(d decimal.Decimal) *Equation {
	return e.process(func(i decimal.Decimal) decimal.Decimal {
		return i.Mul(d)
	})
}

func (e *Equation) Add(other *Equation) *Equation {
	return e.processEquation(other, func(other, origin decimal.Decimal) decimal.Decimal {
		return origin.Add(other)
	})
}

func (e *Equation) Sub(other *Equation) *Equation {
	return e.processEquation(other, func(other, origin decimal.Decimal) decimal.Decimal {
		return origin.Sub(other)
	})
}
