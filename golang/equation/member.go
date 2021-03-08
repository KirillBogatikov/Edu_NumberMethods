package equation

import (
	"NumberMethods/utils"
	"fmt"
	"github.com/shopspring/decimal"
)

type Member struct {
	factor decimal.Decimal
	degree decimal.Decimal
}

func NewMember(f, d decimal.Decimal) *Member {
	return &Member{f, d}
}

func NewMemberFromInt(f, d int64) *Member {
	return &Member{decimal.NewFromInt(f), decimal.NewFromInt(d)}
}

func NewMemberFromFloat(f, d float64) *Member {
	return &Member{decimal.NewFromFloat(f), decimal.NewFromFloat(d)}
}

func (d *Member) Factor() decimal.Decimal {
	return d.factor
}

func (d *Member) Degree() decimal.Decimal {
	return d.degree
}

func (d *Member) IsFree() bool {
	return d.degree.Abs().Equals(utils.One)
}

func (d *Member) Derivative() *Member {
	return &Member{d.factor.Mul(d.degree), d.degree.Sub(utils.One)}
}

func (d *Member) Calc(value decimal.Decimal) decimal.Decimal {
	return value.Pow(d.degree).Mul(d.factor)
}

func (d *Member) String() string {
	return fmt.Sprintf("%sx^%s", d.factor.StringFixed(3), d.degree.StringFixed(0))
}
