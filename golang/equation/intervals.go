package equation

import (
	"NumberMethods/utils"
	_ "fmt"
	"github.com/shopspring/decimal"
	"sort"
	"strings"
)

type MemberEquation struct {
	members      []*Member
	discriminant *decimal.Decimal
	keys         []decimal.Decimal
	verbose      bool
	log          *utils.LocalLog
}

func NewMemberEquation(members ...*Member) *MemberEquation {
	return &MemberEquation{
		members: members,
		log:     utils.NewLocalLog(),
	}
}

func (m *MemberEquation) String() string {
	var b strings.Builder

	for i, m := range m.members {
		if i > 0 && m.factor.GreaterThanOrEqual(decimal.Zero) {
			b.WriteString(" + ")
		}

		b.WriteString(m.String())
	}

	b.WriteString(" = 0")
	return b.String()
}

func (m *MemberEquation) Verbose(v bool) {
	m.verbose = v
}

func (m *MemberEquation) abc() (a, b, c decimal.Decimal, has3 bool) {
	for _, d := range m.members {
		if d.degree.IsZero() {
			c = d.factor
		} else if d.degree.Equals(utils.One) {
			b = d.factor
		} else if d.degree.Equals(utils.Two) {
			a = d.factor
		} else {
			has3 = true
		}
	}

	return
}

func (m *MemberEquation) Discriminant() decimal.Decimal {
	if m.discriminant != nil {
		if m.verbose {
			m.log.Printf("Cached discriminant: %s\n", m.discriminant.StringFixed(3))
		}

		return *m.discriminant
	}

	a, b, c, _ := m.abc()

	d := b.Pow(utils.Two).Sub(a.Mul(c).Mul(decimal.NewFromInt(4)))

	if m.verbose {
		m.log.Printf("A = %s, ", a.StringFixed(3))
		m.log.Printf("B = %s, ", b.StringFixed(3))
		m.log.Printf("C = %s\n", c.StringFixed(3))

		m.log.Printf("D = b^2 - 4ac = %s^2 - 4 * %s * %s = %s\n", b.StringFixed(3), a.StringFixed(3), c.StringFixed(3), d.StringFixed(3))
	}

	m.discriminant = &d

	return d
}

func (m *MemberEquation) Derivative() *MemberEquation {
	derivatives := make([]*Member, 0)

	for _, member := range m.members {
		d := member.Derivative()

		if m.verbose {
			m.log.Printf("(%s)` = %s", member.String(), d.String())
		}

		if d.degree.IsNegative() {
			if m.verbose {
				m.log.Println(": skip")
			}

			continue
		}

		if m.verbose {
			m.log.Print("\n")
		}
		derivatives = append(derivatives, d)
	}

	equation := NewMemberEquation(derivatives...)
	equation.verbose = m.verbose
	equation.log = m.log

	return equation
}

func (m *MemberEquation) Intervals() []decimal.Decimal {
	m.keys = make([]decimal.Decimal, 0)

	derivative := m.Derivative()
	keys := derivative.Keys()

	if m.verbose {
		m.log.Printf("Keys = %v\n", keys)
	}

	leftX := keys[0].Sub(utils.One)
	rightX := keys[len(keys)-1].Add(utils.One)

	intervals := append(keys, leftX, rightX)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].LessThan(intervals[j])
	})

	return intervals
}

func (m *MemberEquation) Keys() []decimal.Decimal {
	if m.keys != nil {
		if m.verbose {
			m.log.Printf("Cached keys = %v\n", m.keys)
		}

		return m.keys
	}

	a, b, _, needSimplify := m.abc()

	tempMembers := make([]*Member, 0)

	if needSimplify {
		if m.verbose {
			m.log.Println("Need simplify")
		}

		for i, d := range tempMembers {
			tempMembers[i] = NewMember(d.factor, d.degree.Sub(utils.One))

			if d.degree.Equals(utils.Three) {
				if m.verbose {
					m.log.Println("Key found: 0.000")
				}

				m.keys = append(m.keys, decimal.Zero)
			}

			if d.degree.Equals(utils.One) {
				b = d.factor
			} else if d.degree.Equals(utils.Two) {
				a = d.factor
			}
		}

		if m.verbose {
			m.log.Printf("After simplify: %s\n", m.String())
		}
	}

	d := m.Discriminant()
	sqrtD := utils.Sqrt(d)

	x1 := b.Neg().Add(sqrtD).Div(a.Mul(utils.Two))
	x2 := b.Neg().Sub(sqrtD).Div(a.Mul(utils.Two))
	if m.verbose {
		m.log.Printf("Found keys: %v, %v\n", x1.StringFixed(3), x2.StringFixed(3))
	}

	m.keys = append(m.keys, x1, x2)
	sort.Slice(m.keys, func(i, j int) bool {
		return m.keys[i].LessThan(m.keys[j])
	})

	return m.keys
}

func (m *MemberEquation) Value(x decimal.Decimal) decimal.Decimal {
	r := decimal.Zero
	for _, member := range m.members {
		r = r.Add(member.Calc(x))
	}

	return r
}

func (m *MemberEquation) Log() *utils.LocalLog {
	return m.log
}
