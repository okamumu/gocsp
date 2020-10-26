package csp

// Order encode: http://bach.istc.kobe-u.ac.jp/sugar/docs/order-encoding/

import (
	// "fmt"
	// "log"
	"sort"
)

type SATCode int
type SATClause []SATCode

// SATcode for a*x <= b, x in dom
func le(base SATCode, dom *DomainSet, a, b int) (SATCode, bool) {
	if a > 0 {
		c := floorDiv(b, a)
		pos, ok := dom.Contains(c)
		if pos == 0 && ok == false {
			return 0, false // falselit
		} else if pos == dom.Size()-1 {
			return 1, false // truelit
		}
		return base + SATCode(pos), true
	} else {
		c := ceilDiv(b, a) - 1
		pos, ok := dom.Contains(c)
		if pos == 0 && ok == false {
			return 1, false // truelit
		} else if pos == dom.Size()-1 {
			return 0, false // falselit
		}
		return -(base + SATCode(pos)), true
	}
}

func applyOr(codes []SATClause, p SATCode) []SATClause {
	if len(codes) == 0 {
		return []SATClause{SATClause{p}}
	}
	for i, _ := range codes {
		codes[i] = append(codes[i], p)
	}
	return codes
}

func Encode(cnf []CSPClause) []SATClause {
	codes := []SATClause{}
	for _, lit := range cnf {
		for _, c := range lit.encode() {
			codes = append(codes, c)
		}
	}
	return codes
}

func (c CSPClause) encode() []SATClause {
	codes := []SATClause{}
	for _, lit := range c {
		codes = lit.encode(codes)
	}
	return codes
}

func (b *BoolVar) encode(codes []SATClause) []SATClause {
	base := b.getSATCodeBase()
	if b.neg == false {
		return applyOr(codes, base)
	} else {
		return applyOr(codes, -base)
	}
}

func (c *CSPOperator) encode(codes []SATClause) []SATClause {
	panic("SAT encoding cannot be applied to CSPOperator")
}

func (c *CSPComparator) encode(codes []SATClause) []SATClause {
	if c.op == CSPOperatorLeZero {
		vars := make([]*IntVar, 0, len(c.s.coef))
		for k, _ := range c.s.coef {
			vars = append(vars, k)
		}
		sort.Slice(vars, func(i, j int) bool {
			s1 := vars[i].domain.Size()
			s2 := vars[j].domain.Size()
			if s1 == s2 {
				k1 := abs(c.s.coef[vars[i]])
				k2 := abs(c.s.coef[vars[j]])
				return k1 > k2
			} else {
				return s1 < s2
			}
		})
		return encodeIntVar(codes, vars, c.s, -c.s.b)
	} else {
		panic("SAT encoding can be applied to LeZero")
	}
}

// Encoding based on the formula: a_1 x_1 + ... a_n x_n <= b
func encodeIntVar(codes []SATClause, vars []*IntVar, s *Sum, rhs int) []SATClause {
	a := s.coef[vars[0]]
	base := vars[0].getSATCodeBase()
	dom := vars[0].domain
	if len(vars) == 1 {
		p, ok := le(base, dom, a, rhs)
		if ok {
			return applyOr(codes, p)
		} else if p == 1 { // truelit
			return []SATClause{SATClause{}}
		} else { // falselit
			return codes
		}
	}
	if a > 0 {
		satcode := []SATClause{}
		for i, b := range dom.x {
			// assume x_1 >= b. Get clauses sum a_ix_i <= c - a_1b. Put literal x_1 <= b-1 to the clauses
			if i == 0 {
				tmp := encodeIntVar(codes, vars[1:], s, rhs-a*b)
				for _, c := range tmp {
					if len(c) != 0 {
						satcode = append(satcode, c)
					}
				}
			} else {
				p := base + SATCode(i-1) // literal x <= b - 1
				tmp := encodeIntVar(codes, vars[1:], s, rhs-a*b)
				for _, c := range applyOr(tmp, p) {
					satcode = append(satcode, c)
				}
			}
		}
		if len(satcode) == 0 {
			return codes
		} else {
			return satcode
		}
	} else { // assume a < 0
		satcode := []SATClause{}
		for i, b := range dom.x {
			// assume x_1 <= b. Get clauses sum a_ix_i <= c - a_1b. Put literal x_1 > b to the clauses
			if i != dom.Size()-1 {
				p := -(base + SATCode(i)) // literal x > b
				tmp := encodeIntVar(codes, vars[1:], s, rhs-a*b)
				for _, c := range applyOr(tmp, p) {
					satcode = append(satcode, c)
				}
			} else {
				tmp := encodeIntVar(codes, vars[1:], s, rhs-a*b)
				for _, c := range tmp {
					if len(c) != 0 {
						satcode = append(satcode, c)
					}
				}
			}
		}
		if len(satcode) == 0 {
			return codes
		} else {
			return satcode
		}
	}
}
