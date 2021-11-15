package csp

import (
	// "fmt"
	"log"
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

// decompSum
// The function is to decompose a linear constraint to up to three terms
func decompSum(s *Sum, auxvars []*IntVar) ([]CSPConstraint, []*IntVar) {
	cs := make([]CSPConstraint, 0)
	for s.Size() > 3 {
		vars := sortVars(s.coef)
		x, y := vars[0], vars[1]
		a, b := s.coef[x], s.coef[y]
		d := x.domain.Copy()
		d.CrossApplyFunc(y.domain, func(x, y int) int {
			return a*x + b*y
		})
		z := NewAuxIntVar(uint(len(auxvars)), d)
		auxvars = append(auxvars, z)
		coef := map[*IntVar]int{x: a, y: b, z: -1}
		f := NewSum(coef, 0)
		s.Sub(f)
		cs = append(cs, CSPEqZero(f))
	}
	return cs, auxvars
}

func (c *CSPComparator) Decomp(auxvars []*IntVar) (CSPConstraint, []*IntVar) {
	var cs []CSPConstraint
	switch c.op {
	case CSPOperatorEqZero:
		cs, auxvars = decompSum(c.s, auxvars)
		cs = append(cs, CSPEqZero(c.s))
	case CSPOperatorNeZero:
		cs, auxvars = decompSum(c.s, auxvars)
		cs = append(cs, CSPNeZero(c.s))
	case CSPOperatorLeZero:
		cs, auxvars = decompSum(c.s, auxvars)
		cs = append(cs, CSPLeZero(c.s))
	case CSPOperatorGeZero:
		cs, auxvars = decompSum(c.s, auxvars)
		cs = append(cs, CSPGeZero(c.s))
	default:
		log.Fatal("Error: operator does not exist")
		return nil, auxvars
	}
	if len(cs) == 1 {
		return cs[0], auxvars
	} else {
		return CSPAnd(cs...), auxvars
	}
}

func (c *CSPOperator) Decomp(auxvars []*IntVar) (CSPConstraint, []*IntVar) {
	switch c.op {
	case CSPOperatorAnd:
		args := make([]CSPConstraint, len(c.args))
		for i, x := range c.args {
			args[i], auxvars = x.Decomp(auxvars)
		}
		return CSPAnd(args...), auxvars
	case CSPOperatorOr:
		args := make([]CSPConstraint, len(c.args))
		for i, x := range c.args {
			args[i], auxvars = x.Decomp(auxvars)
		}
		return CSPOr(args...), auxvars
	default:
		log.Fatal("Error: operator does not exist")
		return nil, auxvars
	}
}

func (b *BoolVar) Decomp(auxvars []*IntVar) (CSPConstraint, []*IntVar) {
	return b, auxvars
}
