package seidel

import (
	"NumberMethods/equation"
	"NumberMethods/matrix"
	"fmt"
	"github.com/shopspring/decimal"
)

type System struct {
	*equation.System
}

func NewSystem(m *matrix.Matrix, a float64) *System {
	return &System{
		System: &equation.System{
			Matrix:         m,
			TargetAccuracy: decimal.NewFromFloat(a),
		},
	}
}

func (e *System) Solve() []decimal.Decimal {
	previousX := make([]decimal.Decimal, e.Height())
	currentX := make([]decimal.Decimal, e.Height())
	for i := 0; i < e.Height(); i++ {
		currentX[i] = e.Get(i, e.Width()-1)
	}

	i := 0
	for i = 0; i < 1_000; i++ {
		c := 0

		for j, x := range currentX {
			result := x

			for k := 0; k < e.Width(); k++ {
				result = result.Mul(e.Get(j, k))
			}

			currentX[j] = result
			if result.Equals(previousX[j]) {
				c++
			} else {
				previousX[j] = result
			}
		}

		if c == e.Height() {
			break
		}
	}

	fmt.Println(i)

	return currentX
}
