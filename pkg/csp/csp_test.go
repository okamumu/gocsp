package csp

import (
	"fmt"
	"testing"
)

func TestCSP1(t *testing.T) {
	c := NewCSP()
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = c.NewBoolVar(false)
	}
	y := make([]*IntVar, 10)
	for i, _ := range x {
		y[i] = c.NewIntVarWithRange(0, 5)
	}
	c1 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[3]: 10, y[9]: 8}, 1))
	c2 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[5]: 10, y[6]: -10, y[9]: 8}, 2))
	c3 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[2]: 1, y[8]: -20, y[9]: 8}, 3))
	c4 := CSPOr(CSPOr(c1, x[1], CSPOr(x[2], c2)), x[4], x[5], CSPAnd(c3, x[7]))
	c.AddConstraint(c4, true)
	c.CNF()
	fmt.Println(c.cnf)
	c.genBase()
	c.Encode()
	fmt.Println(c.sat)
}
