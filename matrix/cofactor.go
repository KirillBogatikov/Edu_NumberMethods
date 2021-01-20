package matrix

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

func (m *Matrix) Cofactor() *Matrix {
	var newData [][]decimal.Decimal

	for r := 0; r < m.height; r++ {
		newRow := make([]decimal.Decimal, m.width)

		for c := 0; c < m.width; c++ {
			i := math.Pow(-1, float64(r + c))
			minor := m.MinorFor(c, r)

			fmt.Printf("Minor for (%d, %d): %s\n", c + 1, r + 1, "\n" + minor.ToString())

			minorDeterminant := minor.Determinant()
			newRow[c] = minorDeterminant.Mul(decimal.NewFromFloat(i))
			fmt.Printf("A(%d, %d) = (-1 ^ %v) * D(%d, %d) = %v * %v = %v\n", r + 1, c + 1, r + c + 2, c + 1, r + 1, i, minorDeterminant, newRow[c])
		}

		newData = append(newData, newRow)
	}

	return &Matrix{newData, m.width, m.height, defaultCache()}
}
