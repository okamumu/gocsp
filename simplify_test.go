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
	cs := make([]CSPLiteral, 0)
	c := CSPOr(CSPOr(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
	cs = c.flattenOr(cs)
	fmt.Println(cs)
}

func TestCNF1(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	c := CSPOr(CSPOr(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
	cnf, vars = c.tocnf(cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
}

func TestSimplify1(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	c := CSPOr(CSPOr(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
	cnf, vars := Simplify(c)
	fmt.Println(cnf)
	fmt.Println(vars)
}

func TestSimplify2(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	y := make([]*IntVar, 10)
	for i, _ := range x {
		y[i] = NewIntVarWithRange(VarLabel(fmt.Sprintf("int%d", i)), 0, 10)
	}
	c1 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[3]: 10, y[9]: 8}, 1))
	c2 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[5]: 10, y[6]: -100, y[9]: 8}, 2))
	c3 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[2]: 1, y[8]: -200, y[9]: 8}, 3))
	c := CSPOr(CSPOr(c1, x[1], CSPOr(x[2], c2)), x[4], x[5], CSPAnd(c3, x[7]))
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	cnf, vars = c.tocnf(cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
}

func TestSimplify3(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	y := make([]*IntVar, 10)
	for i, _ := range x {
		y[i] = NewIntVarWithRange(VarLabel(fmt.Sprintf("int%d", i)), 0, 10)
	}
	c1 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[3]: 10, y[9]: 8}, 1))
	c2 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[5]: 10, y[6]: -100, y[9]: 8}, 2))
	c3 := CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[2]: 1, y[8]: -200, y[9]: 8}, 3))
	c := CSPOr(CSPOr(c1, x[1], CSPOr(x[2], c2)), x[4], x[5], CSPAnd(c3, x[7]))
	cnf, vars := Simplify(c)
	fmt.Println(cnf)
	fmt.Println(vars)
}
