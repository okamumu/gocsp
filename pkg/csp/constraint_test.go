package csp

import (
	"fmt"
	_ "log"
	"testing"
)

// test func

func TestCSPConstraint(t *testing.T) {
	x := newBoolVar(0)
	y := newBoolVar(1)
	z := CSPAnd(x, y)
	fmt.Println(z)
}

func TestToLeZero1(t *testing.T) {
	d := DomainSet{1, 4, 6, 7, 19}
	y := make([]*IntVar, 10)
	for i, _ := range y {
		y[i] = newIntVar(i, d)
	}
	var c1 CSPConstraint
	c1 = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[3]: 10, y[9]: 8}, 1))
	fmt.Println(c1)
	fmt.Println(c1.ToLeZero())
}
