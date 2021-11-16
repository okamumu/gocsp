package csp

import (
	"fmt"
	"log"
)

type CSPLiteral interface {
	isSimple() bool
	encode([][]int, map[uint]int) ([][]int, bool)
}

type CSPClause []CSPLiteral

func (c CSPClause) String() string {
	str := "["
	for _, v := range c {
		str += fmt.Sprintf("%s,", v)
	}
	return str + "]"
}

// tocnf: The function is to create CNF
func (c *CSPOperator) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	switch c.op {
	case CSPOperatorAnd:
		for _, x := range c.args {
			cnf, auxvars = x.tocnf(cnf, auxvars)
		}
		return cnf, auxvars
	case CSPOperatorOr:
		flattencs := c.flattenOr(make([]CSPConstraint, 0))
		clause := make(CSPClause, 0)
		cs := make([]CSPConstraint, 0)
		for _, x := range flattencs {
			clause, cs, auxvars = x.tseitin(clause, cs, auxvars)
		}
		cnf = append(cnf, clause)
		for _, x := range cs {
			cnf, auxvars = x.tocnf(cnf, auxvars)
		}
		return cnf, auxvars
	default:
		log.Fatal("Error: operator does not exist")
		return cnf, auxvars
	}
}

func (c *CSPComparator) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	return append(cnf, CSPClause{c}), auxvars
}

func (b *BoolVar) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	return append(cnf, CSPClause{b}), auxvars
}

// flattenOr: The function is to flatten list of literals for two or more OR operations.
// Example: Or(Or(a,b,c), AND(d, e)) -> [a, b, c, AND(d,e)]
func (c *CSPOperator) flattenOr(cs []CSPConstraint) []CSPConstraint {
	switch c.op {
	case CSPOperatorAnd:
		return append(cs, c)
	case CSPOperatorOr:
		for _, x := range c.args {
			cs = x.flattenOr(cs)
		}
		return cs
	default:
		log.Fatal("Error: operator does not exist")
		return cs
	}
}

func (c *CSPComparator) flattenOr(cs []CSPConstraint) []CSPConstraint {
	return append(cs, c)
}

func (b *BoolVar) flattenOr(cs []CSPConstraint) []CSPConstraint {
	return append(cs, b)
}

// tseitin: Tsetin transform
func (c *CSPOperator) tseitin(first CSPClause, cs []CSPConstraint, auxvars []*BoolVar) (CSPClause, []CSPConstraint, []*BoolVar) {
	p := newAuxBoolVar(uint(len(auxvars)), false)
	auxvars = append(auxvars, p)
	first = append(first, p)
	switch c.op {
	case CSPOperatorAnd:
		for _, x := range c.args {
			cs = append(cs, CSPOr(x, p.Not()))
		}
	case CSPOperatorOr:
		c.args = append(c.args, p.Not())
		cs = append(cs, CSPOr(c.args...))
	default:
		log.Fatal("Error: operator does not exist")
	}
	return first, cs, auxvars
}

func (c *CSPComparator) tseitin(first CSPClause, cs []CSPConstraint, auxvars []*BoolVar) (CSPClause, []CSPConstraint, []*BoolVar) {
	return append(first, c), cs, auxvars
}

func (b *BoolVar) tseitin(first CSPClause, cs []CSPConstraint, auxvars []*BoolVar) (CSPClause, []CSPConstraint, []*BoolVar) {
	return append(first, b), cs, auxvars
}

// simplify
func isSimple(c CSPClause) bool {
	count := 0
	for _, l := range c {
		if !l.isSimple() {
			if count++; count > 1 {
				return false
			}
		}
	}
	return true
}

func (c *CSPComparator) isSimple() bool {
	return c.s.Size() <= 1
}

func (b *BoolVar) isSimple() bool {
	return true
}

func simplify(c CSPConstraint, cnf []CSPClause, auxvars []*BoolVar, tmp []CSPClause) ([]CSPClause, []*BoolVar) {
	tmp, auxvars = c.tocnf(tmp, auxvars)
	for _, clause := range tmp {
		if isSimple(clause) {
			cnf = append(cnf, clause)
		} else {
			first := make([]CSPLiteral, 0, len(clause))
			for _, lit := range clause {
				if lit.isSimple() {
					first = append(first, lit)
				} else {
					p := newAuxBoolVar(uint(len(auxvars)), false)
					q, _ := p.Not().(*BoolVar)
					auxvars = append(auxvars, p)
					first = append(first, p)
					cnf = append(cnf, CSPClause{lit, q})
				}
			}
			cnf = append(cnf, CSPClause(first))
		}
	}
	return cnf, auxvars
}
