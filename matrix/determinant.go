package matrix

import (
	"NumberMethods/utils"
	"fmt"
	"github.com/shopspring/decimal"
)

func (m *Matrix) Determinant() decimal.Decimal {
	if m.width != m.height {
		panic("Required quad matrix")
	}

	if m.cache.determinant != nil {
		return *m.cache.determinant
	}

	determinant := decimal.NewFromInt(0)

	if m.width == 2 {
		determinant = m.data[0][0].Mul(m.data[1][1]).Sub(m.data[0][1].Mul(m.data[1][0]))
	} else if m.width == 3 {
		determinant = utils.Mul(m.data[0][0], m.data[1][1], m.data[2][2]).Add(
			utils.Mul(m.data[2][0], m.data[0][1], m.data[1][2])).Add(
			utils.Mul(m.data[1][0], m.data[2][1], m.data[0][2])).Sub(
			utils.Mul(m.data[0][2], m.data[1][1], m.data[2][0])).Sub(
			utils.Mul(m.data[0][0], m.data[1][2], m.data[2][1])).Sub(
			utils.Mul(m.data[0][1], m.data[1][0], m.data[2][2]))

		fmt.Printf("Calc determinant with triangles rule: (%v * %v * %v) + (%v * %v * %v) + (%v * %v * %v) - (%v * %v * %v) - (%v * %v * %v) - (%v * %v * %v) = %v\n",
			m.data[0][0], m.data[1][1], m.data[2][2], m.data[2][0], m.data[0][1], m.data[1][2], m.data[1][0], m.data[2][1], m.data[0][2],
			m.data[0][2], m.data[1][1], m.data[2][0], m.data[0][0], m.data[1][2], m.data[2][1], m.data[0][1], m.data[1][0], m.data[2][2], determinant)
	} else {
		var k decimal.Decimal

		for i := 0; i < m.width; i++ {
			if i % 2 == 0 {
				k = PositiveK
			} else {
				k = NegativeK
			}

			minorDeterminant := m.MinorFor(i, 0).Determinant()
			determinant = determinant.Add(k.Mul(m.data[0][i]).Mul(minorDeterminant))
		}
	}

	m.cache.determinant = &determinant
	return determinant
}
