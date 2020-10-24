package csp

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i := 0; i < 10; i++ {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	cs := make([]CSPConstraint, 0)
	c := CSPOr(CSPOr(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
	cs = c.flattenOr(cs)
	fmt.Println(cs)
}

func TestCNF1(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i := 0; i < 10; i++ {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	c := CSPOr(CSPOr(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
	cnf, vars = c.tocnf(cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
}
