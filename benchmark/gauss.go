package benchmark

import (
	"NumberMethods/gauss"
	"NumberMethods/matrix"
)

type GaussProvider struct {
	Provider
}

func (m *GaussProvider) Count() int64 {
	return 1_000_000
}

func (m *GaussProvider) Operator() func() {
	return func() {
		gauss.NewSystem(matrix.NewMatrix(4, 3, [][]float64{
			{5.4, -2.3, 3.4, -3.5},
			{4.2, 1.7, -2.3, 2.7},
			{3.4, 2.4, 7.4, 1.9},
		}))
	}
}
