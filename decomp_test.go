package csp

import (
	"fmt"
	"testing"
)

// test func

func (x *IntVar) String() string {
	return string(x.label)
}

func (s *Sum) String() string {
	str := ""
	for k, v := range s.coef {
		str += fmt.Sprintf("+(%d)*%s", v, k.label)
	}
	str += fmt.Sprintf("+(%d)", s.b)
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
	str := ""
	switch c.op {
	case CSPOperatorAnd:
		for _, x := range c.args {
			str += fmt.Sprintf("&& %s", x)
		}
		return str
	case CSPOperatorOr:
		for _, x := range c.args {
			str += fmt.Sprintf("|| %s", x)
		}
		return str
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

func TestDecomp1(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	x := NewIntVar("x", &d1)
	d2 := DomainSet{
		x: []int{1, 2, 3, 4, 5},
	}
	y := NewIntVar("y", &d2)
	f := NewSum(map[*IntVar]int{x: 3, y: 4}, -1)
	fmt.Println(f)
}

func TestDecomp(t *testing.T) {
	x := NewIntVarWithRange("x", 0, 10)
	y := NewIntVarWithRange("y", 0, 10)
	z := NewIntVarWithRange("z", 0, 10)
	k := NewIntVarWithRange("k", 0, 10)
	f := NewSum(map[*IntVar]int{x: 3, y: 4, z: 10, k: -5}, -1)
	c := CSPEqZero(f)
	fmt.Println(f)
	v, auxvars := c.Decomp(make([]*IntVar, 0))
	fmt.Println(v)
	fmt.Println(auxvars)
}
