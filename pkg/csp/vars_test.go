package csp

import (
	"fmt"
	"testing"
)

func TestVars01(t *testing.T) {
	b1 := newBoolVar(1, false)
	b2 := newAuxBoolVar(2, true)
	fmt.Println(b1)
	fmt.Println(b2)
}

func TestVars02(t *testing.T) {
	x1 := newIntVar(1, DomainSet{[]int{1, 2, 3, 4}})
	x2 := NewAuxIntVar(2, DomainSet{[]int{3, 4, 5}})
	fmt.Println(x1)
	fmt.Println(x2)
}

func TestVars03(t *testing.T) {
	d := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	x1 := newIntVar(0, d)
	fmt.Println(x1)
	fmt.Println(x1.domain)
}
