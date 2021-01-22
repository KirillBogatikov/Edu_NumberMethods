package matrix

import (
	"NumberMethods/utils"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

func NewMatrix(width int, height int, data [][]float64) *Matrix {
	pkg := make([][]decimal.Decimal, height)
	for i := 0; i < height; i++ {
		pkg[i] = make([]decimal.Decimal, width)
		for j := 0; j < width; j++ {
			pkg[i][j] = decimal.NewFromFloat(data[i][j])
		}
	}

	return &Matrix{pkg, width, height, defaultCache()}
}

func defaultCache() Cache {
	return Cache{nil}
}

func (m *Matrix) process(processor utils.ValueProcessor) *Matrix {
	var newData [][]decimal.Decimal

	for _, row := range m.data {
		var newRow []decimal.Decimal

		for _, value := range row {
			newRow = append(newRow, processor(value))
		}

		newData = append(newData, newRow)
	}

	return &Matrix{newData, m.width, m.height, defaultCache()}
}

func (m *Matrix) MinorFor(x, y int) *Matrix {
	var newData [][]decimal.Decimal

	for r := 0; r < m.height; r++ {
		if r == y {
			continue
		}

		var newRow []decimal.Decimal

		for c := 0; c < m.width; c++ {
			if c == x {
				continue
			}

			newRow = append(newRow, m.data[r][c])
		}

		newData = append(newData, newRow)
	}

	return &Matrix{newData, m.width - 1, m.height - 1, defaultCache()}
}

func (m *Matrix) ToString() string {
	var builder strings.Builder

	for _, row := range m.data {
		for _, value := range row {
			floatValue, _ := value.Float64()
			builder.WriteString(strconv.FormatFloat(floatValue, 'f', 3, 64))
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
