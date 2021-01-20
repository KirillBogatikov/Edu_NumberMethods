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

type Cache struct {
	determinant *decimal.Decimal
}

type ValueProcessor func(i decimal.Decimal) decimal.Decimal

type CellProcessor func(origin decimal.Decimal, victim decimal.Decimal) decimal.Decimal
