package div2

import (
	"NumberMethods/equation"
	"NumberMethods/utils"
	"github.com/shopspring/decimal"
)

func Solve(verbose bool, accuracy decimal.Decimal, equation *equation.MemberEquation) []decimal.Decimal {
	log := equation.Log()

	intervals := equation.Intervals()
	keys := make([]decimal.Decimal, 0)

	for i := 0; i < len(intervals)-1; i++ {
		x1 := intervals[i]
		x2 := intervals[i+1]

		var xn, xp decimal.Decimal

		for j := 0; j < 150; j++ {
			xn = x1.Add(x2).Div(utils.Two)
			f := equation.Value(xn)

			if verbose {
				log.Printf("xn = (%v + %v) / 2 = %v; f(xn) = %v\n", x1, x2, xn, f)
			}

			if f.LessThan(xn) {
				x1 = xn
			} else {
				x2 = xn
			}

			delta := xp.Sub(xn).Abs()
			if delta.LessThan(accuracy) {
				if verbose {
					log.Printf("Delta (%v, %v) -> %v is less than accuracy. Break\n", xp, xn, delta)
				}

				break
			}
		}

		keys = append(keys, xn)
	}

	return keys
}
