package csp

import (
	"fmt"
	"testing"
)

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
