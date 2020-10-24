package csp

import (
	"fmt"
	"testing"
)

// test func

func (x *IntVar) String() string {
	return string(x.label)
}

func (x *BoolVar) String() string {
	if x.neg {
		return "!" + string(x.label)
	} else {
		return string(x.label)
	}
}

func (lits CSPClause) String() string {
	str := "["
	for _, v := range lits.x {
		str += fmt.Sprintf("%s,", v)
	}
	return str + "]"
}

func (s *Sum) String() string {
	str := "["
	for k, v := range s.coef {
		str += fmt.Sprintf("+(%d)*%s", v, k.label)
	}
	str += fmt.Sprintf("+(%d)]", s.b)
	return str
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
		panic("")
	}
}

func (c *CSPOperator) String() string {
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
		panic("")
	}
}

func NewIntVarWithRange(label VarLabel, lb, ub int) *IntVar {
	x := make([]int, ub-lb+1)
	for i, _ := range x {
		x[i] = lb + i
	}
	d := &DomainSet{x: x}
	return NewIntVar(label, d)
}

func TestCSPConstraint(t *testing.T) {
	x := NewBoolVar("x", true)
	y := NewBoolVar("y", true)
	z := CSPAnd(x, y)
	fmt.Println(z)
}

func TestCSPVar1(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d2 := DomainSet{
		x: []int{1, 2},
	}
	x := NewIntVar("x", &d1)
	y := NewIntVar("y", &d2)
	s1 := NewSumFromInt(x)
	s2 := NewSumFromInt(y)
	s1.MulConst(10).Add(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
