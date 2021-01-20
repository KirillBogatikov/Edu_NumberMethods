package matrix

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func (m *Matrix) Div(d decimal.Decimal) *Matrix {
	return m.process(func(input decimal.Decimal) decimal.Decimal {
		return input.Div(d)
	})
}

func (m *Matrix) Mul(d decimal.Decimal) *Matrix {
	return m.process(func(input decimal.Decimal) decimal.Decimal {
		return input.Mul(d)
	})
}

func (m *Matrix) MulMatrix(o *Matrix) *Matrix {
	var newData [][]decimal.Decimal

	for r := 0; r < m.height; r++ {
		var newRow []decimal.Decimal

		for c := 0; c < o.width; c++ {
			s := ""
			v := decimal.NewFromInt(0)

			for i := 0; i < m.width; i++ {
				s += fmt.Sprintf("%v * %v + ", m.data[r][i], o.data[i][c])
				v = v.Add(m.data[r][i].Mul(o.data[i][c]))
			}

			fmt.Printf("%s = %v\n", s, v)
			newRow = append(newRow, v)
		}

		newData = append(newData, newRow)
	}

	return &Matrix{newData, o.width, m.height, defaultCache()}
}