package matrix

import "github.com/shopspring/decimal"

func (m *Matrix) Transpose() *Matrix {
	var newData [][]decimal.Decimal

	for r := 0; r < m.height; r++ {
		var newRow []decimal.Decimal

		for c := 0; c < m.width; c++ {
			newRow = append(newRow, m.data[c][r])
		}

		newData = append(newData, newRow)
	}

	return &Matrix{ newData, m.width, m.height, defaultCache() }
}
