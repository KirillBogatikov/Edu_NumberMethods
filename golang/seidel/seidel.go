package seidel

import (
	"NumberMethods/equation"
	"NumberMethods/matrix"
	"github.com/shopspring/decimal"
)

const (
	MaxIterations = 100
)

type System struct {
	CompareAccuracy decimal.Decimal
	*equation.System
}

func NewSystem(m *matrix.Matrix, a, b float64) *System {
	return &System{
		CompareAccuracy: decimal.NewFromFloat(b),
		System: &equation.System{
			Matrix:         m,
			TargetAccuracy: decimal.NewFromFloat(a),
		},
	}
}

func (e *System) Solve() ([]decimal.Decimal, int) {
	unknownWidth := e.Width() - 1

	previous := make([]decimal.Decimal, e.Height())

	freeValues := make([]decimal.Decimal, e.Height())
	for i := 0; i < e.Height(); i++ {
		freeValues[i] = e.Get(i, unknownWidth)
	}

	i := 0
	for i = 0; i < MaxIterations; i++ {
		for r := 0; r < e.Height(); r++ {
			k := e.Get(r, unknownWidth)
			x := k

			for c := 0; c < unknownWidth; c++ {
				if r == c {
					continue
				}

				negativeValue := e.Get(r, c).Mul(decimal.NewFromInt(-1))
				x = x.Add(negativeValue.Mul(freeValues[c]))
			}

			freeValues[r] = x
		}

		if e.equals(previous, freeValues) {
			break
		}

		for i, _ := range previous {
			previous[i] = freeValues[i]
		}
	}

	return freeValues, i
}

func (s *System) equals(a []decimal.Decimal, b []decimal.Decimal) bool {
	for i, v := range a {
		if b[i].Sub(v).Abs().GreaterThan(s.CompareAccuracy) {
			return false
		}
	}

	return true
}
