package gauss

import (
	"NumberMethods/matrix"
	"github.com/shopspring/decimal"
)

type EquationSystem struct {
	data *matrix.Matrix
}

func NewEquationSystem(matrix *matrix.Matrix) *EquationSystem {
	return &EquationSystem{matrix}
}

func (e *EquationSystem) Solve() []decimal.Decimal {
	x := make([]decimal.Decimal, e.data.Height())

	for k := 1; k < e.data.Height(); k++ {
		for j := k; j < e.data.Height(); j++ {
			m := e.data.Get(j, k-1).Div(e.data.Get(k-1, k-1))

			for i := 0; i < e.data.Width(); i++ {
				e.data.Set(j, i, e.data.Get(j, i).Sub(m.Mul(e.data.Get(k-1, i))))
			}
		}

		for i := e.data.Height() - 1; i >= 0; i-- {
			x[i] = e.data.Get(i, e.data.Height()).Div(e.data.Get(i, i))

			for c := e.data.Height() - 1; c > i; c-- {
				x[i] = x[i].Sub(e.data.Get(i, c).Mul(x[c]).Div(e.data.Get(i, i)))
			}
		}
	}

	return x
}
