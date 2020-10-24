package csp

import (
	"fmt"
	"testing"
)

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
