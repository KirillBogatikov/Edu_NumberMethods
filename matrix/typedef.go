package matrix

import "github.com/shopspring/decimal"

var (
	PositiveK = decimal.NewFromInt(1)
	NegativeK = decimal.NewFromInt(-1)
)

type Matrix struct {
	data          [][]decimal.Decimal
	width, height int
	cache         Cache
}

func (m *Matrix) Get(row, column int) decimal.Decimal {
	return m.data[row][column]
}

func (m *Matrix) Set(row, column int, value decimal.Decimal) decimal.Decimal {
	tmp := m.data[row][column]
	m.data[row][column] = value
	return tmp
}

func (m *Matrix) Width() int {
	return m.width
}

func (m *Matrix) Height() int {
	return m.height
}

func (m *Matrix) Copy() *Matrix {
	data := make([][]decimal.Decimal, m.height)

	for r := 0; r < m.height; r++ {
		row := make([]decimal.Decimal, m.width)

		for c, value := range m.data[r] {
			row[c] = value
		}

		data[r] = row
	}

	return &Matrix{data, m.width, m.height, Cache{}}
}

type Cache struct {
	determinant *decimal.Decimal
}
