package csp

import (
	"fmt"
	"testing"
)

func TestDecomp1(t *testing.T) {
	d1 := DomainSet{1, 4, 6, 7, 19}
	x := newIntVar(1, d1)
	d2 := DomainSet{1, 2, 3, 4, 5}
	y := newIntVar(2, d2)
	f := NewSum(map[*IntVar]int{x: 3, y: 4}, -1)
	fmt.Println(f)
}

func TestDecomp2(t *testing.T) {
	d := DomainSet{1, 4, 6, 7, 19}
	x := newIntVar(0, d)
	y := newIntVar(1, d)
	z := newIntVar(2, d)
	k := newIntVar(3, d)
	f := NewSum(map[*IntVar]int{x: 3, y: 4, z: 10, k: -5}, -1)
	c := CSPEqZero(f)
	fmt.Println(f)
	v, auxvars := c.Decomp(make([]*IntVar, 0))
	fmt.Println(v)
	fmt.Println(auxvars)
}
