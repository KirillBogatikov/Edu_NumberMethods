package utils

import "github.com/shopspring/decimal"

type ValueProcessor func(i decimal.Decimal) decimal.Decimal

type CellProcessor func(other decimal.Decimal, origin decimal.Decimal) decimal.Decimal
