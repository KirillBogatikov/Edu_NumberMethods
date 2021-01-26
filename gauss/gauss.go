package gauss

import (
	"NumberMethods/equation"
	"NumberMethods/matrix"
	"github.com/shopspring/decimal"
)

type System struct {
	*equation.System
}

func NewSystem(m *matrix.Matrix) *System {
	return &System{
		System: &equation.System{
			Matrix:         m,
			TargetAccuracy: decimal.Zero,
		},
	}
}

func (e *System) Solve() []decimal.Decimal {
	x := make([]decimal.Decimal, e.Height())

	for k := 1; k < e.Height(); k++ {
		for j := k; j < e.Height(); j++ {
			m := e.Get(j, k-1).Div(e.Get(k-1, k-1))

			for i := 0; i < e.Width(); i++ {
				e.Set(j, i, e.Get(j, i).Sub(m.Mul(e.Get(k-1, i))))
			}
		}

		for i := e.Height() - 1; i >= 0; i-- {
			x[i] = e.Get(i, e.Height()).Div(e.Get(i, i))

			for c := e.Height() - 1; c > i; c-- {
				x[i] = x[i].Sub(e.Get(i, c).Mul(x[c]).Div(e.Get(i, i)))
			}
		}
	}

	return x
}
