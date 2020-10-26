package csp

import (
	"fmt"
	"testing"
)

// func TestEncode1(t *testing.T) {
// 	x := make([]*BoolVar, 10)
// 	for i, _ := range x {
// 		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
// 	}
// 	c := CSPOr(CSPOr(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
// 	cnf := make([]CSPClause, 0)
// 	vars := make([]*BoolVar, 0)
// 	cnf, vars = Simplify(c, cnf, vars)
// 	fmt.Println(cnf)
// 	fmt.Println(cnf[0])
// 	fmt.Println(vars)
// 	fmt.Println(cnf[0].encode())
// 	fmt.Println(Encode(cnf))
// }

func TestEncode2(t *testing.T) {
	x := NewIntVarWithRange("x", 0, 5)
	y := NewIntVarWithRange("y", 0, 3)
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 3, y: 5}, -14))
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	cnf, vars = c1.tocnf(cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
	fmt.Println(Encode(cnf))
	fmt.Println(satBaseCodes)
}

func TestEncode3(t *testing.T) {
	x := NewIntVarWithRange("x", 0, 5)
	y := NewIntVarWithRange("y", 0, 3)
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: -3, y: -5}, 14))
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	cnf, vars = c1.tocnf(cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
	fmt.Println(Encode(cnf))
	fmt.Println(satBaseCodes)
}

func TestEncode4(t *testing.T) {
	x := NewIntVarWithSet("x", []int{0, 10, 20})
	y := NewIntVarWithSet("y", []int{0, 10, 20})
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 1, y: 1}, -20))
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	cnf, vars = c1.tocnf(cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
	fmt.Println(Encode(cnf))
	fmt.Println(satBaseCodes)
}

func TestEncode5(t *testing.T) {
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = NewBoolVar(VarLabel(fmt.Sprintf("bool%d", i)), false)
	}
	y := make([]*IntVar, 10)
	for i, _ := range x {
		y[i] = NewIntVarWithRange(VarLabel(fmt.Sprintf("int%d", i)), 0, 10)
	}
	var c1, c2, c3 CSPLiteral
	c1, y = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[3]: 10, y[9]: 8}, 1)).Decomp(y)
	c2, y = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[5]: 10, y[6]: -10, y[9]: 8}, 2)).Decomp(y)
	c3, y = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[2]: 1, y[8]: -20, y[9]: 8}, 3)).Decomp(y)
	c := CSPOr(CSPOr(c1, x[1], CSPOr(x[2], c2)), x[4], x[5], CSPAnd(c3, x[7])).ToLeZero()
	fmt.Print(c)
	cnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	cnf, vars = Simplify(c, cnf, vars)
	fmt.Println(cnf)
	fmt.Println(vars)
	fmt.Println(Encode(cnf))
	fmt.Println(satBaseCodes)
}
