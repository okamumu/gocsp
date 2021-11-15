package csp

import (
	"fmt"
	"log"
)

type CSPOperatorType int

const (
	CSPOperatorAnd CSPOperatorType = iota + 1
	CSPOperatorOr
	CSPOperatorLeZero
	CSPOperatorGeZero
	CSPOperatorEqZero
	CSPOperatorNeZero
)

type CSPConstraint interface {
	Not() CSPConstraint
	ToLeZero() CSPConstraint
	Decomp([]*IntVar) (CSPConstraint, []*IntVar)
	tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar)
	flattenOr(cs []CSPConstraint) []CSPConstraint
	tseitin(first CSPClause, cs []CSPConstraint, auxvars []*BoolVar) (CSPClause, []CSPConstraint, []*BoolVar)
}

// CSPComparator
// This is one constraint using one of the comparators <= (Le), >= (Ge), == (Eq), != (Ne).
// The fundamental form becomes
//     a_1 x_1 + a_2 x_2 + ... + a_n x_n (comparator) 0
// where the right-hand side is Sum in the program
//
type CSPComparator struct {
	op CSPOperatorType
	s  *Sum
}

func CSPLeZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorLeZero,
		s:  s,
	}
}

func CSPGeZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorGeZero,
		s:  s,
	}
}

func CSPEqZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorEqZero,
		s:  s,
	}
}

func CSPNeZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorNeZero,
		s:  s,
	}
}

func (c *CSPComparator) String() string {
	switch c.op {
	case CSPOperatorEqZero:
		return fmt.Sprintf("EqZero(%s)", c.s)
	case CSPOperatorNeZero:
		return fmt.Sprintf("NeZero(%s)", c.s)
	case CSPOperatorLeZero:
		return fmt.Sprintf("LeZero(%s)", c.s)
	case CSPOperatorGeZero:
		return fmt.Sprintf("GeZero(%s)", c.s)
	default:
		log.Fatal("Error: operator does not exist")
		return ""
	}
}

// CSPOperator
// This is one constraint using logical operators (AND, OR, IMP, IFF)
//

type CSPOperator struct {
	op   CSPOperatorType
	args []CSPConstraint
}

func CSPAnd(args ...CSPConstraint) *CSPOperator {
	return &CSPOperator{
		op:   CSPOperatorAnd,
		args: args,
	}
}

func CSPOr(args ...CSPConstraint) *CSPOperator {
	return &CSPOperator{
		op:   CSPOperatorOr,
		args: args,
	}
}

func CSPImp(x, y CSPConstraint) CSPConstraint {
	return CSPOr(x.Not(), y)
}

func CSPIff(x, y CSPConstraint) CSPConstraint {
	return CSPAnd(CSPOr(x.Not(), y), CSPOr(x, y.Not()))
}

func (c CSPOperator) String() string {
	switch c.op {
	case CSPOperatorAnd:
		str := "AND("
		for _, x := range c.args {
			str += fmt.Sprintf("%s,", x)
		}
		return str + ")"
	case CSPOperatorOr:
		str := "OR("
		for _, x := range c.args {
			str += fmt.Sprintf("%s,", x)
		}
		return str + ")"
	default:
		log.Fatal("Error: operator does not exist")
		return ""
	}
}

// Not
// The method is to take the negation
func (c *CSPComparator) Not() CSPConstraint {
	s := c.s.copy()
	switch c.op {
	case CSPOperatorLeZero:
		s.AddConst(-1)
		return CSPGeZero(s)
	case CSPOperatorGeZero:
		s.AddConst(1)
		return CSPLeZero(s)
	case CSPOperatorEqZero:
		return CSPNeZero(s)
	case CSPOperatorNeZero:
		return CSPEqZero(s)
	default:
		log.Fatal("CSP Operator is invalid.")
		return nil
	}
}

func (c *CSPOperator) Not() CSPConstraint {
	newargs := make([]CSPConstraint, len(c.args))
	for i, x := range c.args {
		newargs[i] = x.Not()
	}
	switch c.op {
	case CSPOperatorAnd:
		return CSPOr(newargs...)
	case CSPOperatorOr:
		return CSPAnd(newargs...)
	default:
		log.Fatal("CSP Operator is invalid.")
		return nil
	}
}

func (b *BoolVar) Not() CSPConstraint {
	return &BoolVar{b.id, !b.neg, b.aux}
}

// ToLeZero
// The method is to change the CSPComparator to the forms using only Sum <= 0
func (c *CSPComparator) ToLeZero() CSPConstraint {
	switch c.op {
	case CSPOperatorEqZero:
		s1 := c.s.copy()
		s2 := c.s.copy()
		return CSPAnd(CSPLeZero(s1), CSPLeZero(s2.Neg()))
	case CSPOperatorNeZero:
		s1 := c.s.copy()
		s2 := c.s.copy()
		return CSPOr(CSPLeZero(s1.AddConst(1)), CSPLeZero(s2.Neg().AddConst(1)))
	case CSPOperatorGeZero:
		s1 := c.s.copy()
		return CSPLeZero(s1.Neg())
	case CSPOperatorLeZero:
		s1 := c.s.copy()
		return CSPLeZero(s1)
	default:
		log.Fatal("CSP Operator is invalid.")
		return nil
	}
}

func (c *CSPOperator) ToLeZero() CSPConstraint {
	newargs := make([]CSPConstraint, 0, len(c.args))
	for _, x := range c.args {
		newargs = append(newargs, x.ToLeZero())
	}
	switch c.op {
	case CSPOperatorAnd:
		return CSPAnd(newargs...)
	case CSPOperatorOr:
		return CSPOr(newargs...)
	default:
		log.Fatal("CSP Operator is invalid.")
		return nil
	}
}

func (b *BoolVar) ToLeZero() CSPConstraint {
	return &BoolVar{b.id, b.neg, b.aux}
}
