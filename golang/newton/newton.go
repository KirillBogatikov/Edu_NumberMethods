package newton

import (
	"NumberMethods/equation"
	"github.com/shopspring/decimal"
)

const (
	maxIterations = 150
)

func Solve(verbose bool, accuracy decimal.Decimal, x0 decimal.Decimal, equation equation.IMemberEquation) decimal.Decimal {
	fx := equation.Value
	d := equation.Derivative()
	if verbose {
		equation.GetLog().Printf("Derivative is %s\n", d.String())
	}
	dfx := d.Value


	var x1 decimal.Decimal

	for j := 0; j < maxIterations; j++ {
		x1 = x0.Sub(fx(x0).Div(dfx(x0)))

		if x1.Sub(x0).Abs().LessThanOrEqual(accuracy) {
			break;
		}

		x0 = x1
	}

	return x1
}