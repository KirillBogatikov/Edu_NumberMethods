package equation

import (
	"NumberMethods/matrix"
	"fmt"
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

type System struct {
	*matrix.Matrix
	TargetAccuracy decimal.Decimal
}

func NewSystem(m *matrix.Matrix, a decimal.Decimal) *System {
	return &System{m, a}
}

func (s *System) CheckSolution(sol []decimal.Decimal) (bool, []decimal.Decimal) {
	if length := len(sol); length > s.Height() || length < s.Height() {
		panic(fmt.Sprintf("Required %d unknowns. Found %d", s.Height(), length))
	}

	ok := true
	accuracy := make([]decimal.Decimal, len(sol))

	for r := 0; r < s.Height(); r++ {
		var result decimal.Decimal

		for c := 0; c < s.Width()-1; c++ {
			result = result.Add(s.Get(r, c).Mul(sol[c]))
		}

		accuracy[r] = result.Abs().Sub(s.Get(r, s.Width()-1).Abs()).Abs()

		if ok && accuracy[r].GreaterThan(s.TargetAccuracy) {
			ok = false
		}
	}

	return ok, accuracy
}
