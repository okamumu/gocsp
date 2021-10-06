package csp

import (
	// "fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func sortVars(coef map[*IntVar]int) []*IntVar {
	vars := make([]*IntVar, 0, len(coef))
	for k, _ := range coef {
		vars = append(vars, k)
	}
	sort.Slice(vars, func(i, j int) bool {
		s1 := vars[i].domain.Size()
		s2 := vars[j].domain.Size()
		if s1 == s2 {
			k1 := abs(coef[vars[i]])
			k2 := abs(coef[vars[j]])
			return k1 > k2
		} else {
			return s1 < s2
		}
	})
	return vars
}

// This function is to decompose a linear constraint so that they become up to three terms.
func (s *Sum) decompsum(auxvars []*IntVar) ([]CSPLiteral, []*IntVar) {
	lits := make([]CSPLiteral, 0)
	for s.Size() > 3 {
		vars := sortVars(s.coef)
		x, y := vars[0], vars[1]
		a, b := s.coef[x], s.coef[y]
		d := x.domain.copy()
		d.Func(y.domain, func(x, y int) int {
			return a*x + b*y
		})
		z := NewAuxIntVar(d)
		auxvars = append(auxvars, z)
		coef := map[*IntVar]int{x: a, y: b, z: -1}
		f := NewSum(coef, 0)
		s.Sub(f)
		lits = append(lits, CSPEqZero(f))
	}
	return lits, auxvars
}

func (c *CSPComparator) Decomp(auxvars []*IntVar) (CSPLiteral, []*IntVar) {
	var lits []CSPLiteral
	switch c.op {
	case CSPOperatorEqZero:
		lits, auxvars = c.s.decompsum(auxvars)
		lits = append(lits, CSPEqZero(c.s))
	case CSPOperatorNeZero:
		lits, auxvars = c.s.decompsum(auxvars)
		lits = append(lits, CSPNeZero(c.s))
	case CSPOperatorLeZero:
		lits, auxvars = c.s.decompsum(auxvars)
		lits = append(lits, CSPLeZero(c.s))
	case CSPOperatorGeZero:
		lits, auxvars = c.s.decompsum(auxvars)
		lits = append(lits, CSPGeZero(c.s))
	default:
		panic("")
	}
	if len(lits) == 1 {
		return lits[0], auxvars
	} else {
		return CSPAnd(lits...), auxvars
	}
}

func (c *CSPOperator) Decomp(auxvars []*IntVar) (CSPLiteral, []*IntVar) {
	switch c.op {
	case CSPOperatorAnd:
		args := make([]CSPLiteral, len(c.args))
		for i, x := range c.args {
			args[i], auxvars = x.Decomp(auxvars)
		}
		return CSPAnd(args...), auxvars
	case CSPOperatorOr:
		args := make([]CSPLiteral, len(c.args))
		for i, x := range c.args {
			args[i], auxvars = x.Decomp(auxvars)
		}
		return CSPOr(args...), auxvars
	default:
		panic("")
	}
}

func (b *BoolVar) Decomp(auxvars []*IntVar) (CSPLiteral, []*IntVar) {
	return b, auxvars
}