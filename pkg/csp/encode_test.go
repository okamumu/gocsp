package csp

import (
	"fmt"
	"testing"
)

func toDimacs(sat [][]int) {
	max := 0
	for _, c := range sat {
		for _, x := range c {
			if x > 0 {
				if max < x {
					max = x
				}
			} else {
				if max < -x {
					max = -x
				}
			}
		}
	}
	fmt.Println("p cnf", max, len(sat))
	for _, c := range sat {
		for _, x := range c {
			fmt.Print(x, " ")
		}
		fmt.Println(0)
	}
}

func TestEncode1(t *testing.T) {
	baseCode := make(map[uint]int)
	code := 1
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = newBoolVar(uint(i), false)
		baseCode[x[i].id] = code
		code++
	}
	c := CSPOr(CSPAnd(x[0], x[1], CSPOr(x[2], x[3])), x[4], x[5], CSPAnd(x[6], x[7]))
	cnf, vars := simplify(c, make([]CSPClause, 0), make([]*BoolVar, 0), make([]CSPClause, 0))
	for i := 0; i < len(vars); i++ {
		vars[i].id = uint(len(x) + i)
		baseCode[vars[i].id] = code
		code++
	}
	fmt.Println(cnf)
	sat := make([][]int, 0)
	fmt.Println(baseCode)
	for _, x := range cnf {
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			sat = append(sat, tmp...)
		}
	}
	fmt.Println(sat)
}

func TestEncode2(t *testing.T) {
	// 3x + 5y -14 <= 0
	baseCode := make(map[uint]int)
	code := 1
	x := newIntVar(0, DomainSet{[]int{0, 1, 2, 3, 4, 5}})
	baseCode[x.id] = code
	code += x.domain.Size() - 1
	y := newIntVar(1, DomainSet{[]int{0, 1, 2, 3}})
	baseCode[y.id] = code
	code += y.domain.Size() - 1
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 3, y: 5}, -14))
	cnf, vars := c1.tocnf(make([]CSPClause, 0), make([]*BoolVar, 0))
	for i := 0; i < len(vars); i++ {
		vars[i].id = uint(i + 2)
		baseCode[vars[i].id] = code
		code++
	}
	fmt.Println(cnf, vars)
	// 1: x <= 0
	// 2: x <= 1
	// 3: x <= 2
	// 4: x <= 3
	// 5: x <= 4
	// 6: y <= 0
	// 7: y <= 1
	// 8: y <= 2
	sat := make([][]int, 0)
	for _, x := range cnf {
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			sat = append(sat, tmp...)
		}
	}
	fmt.Println(sat)
}

func TestEncode3(t *testing.T) {
	// x + y -20 <= 0
	baseCode := make(map[uint]int)
	code := 1
	x := newIntVar(0, DomainSet{[]int{0, 10, 20}})
	baseCode[x.id] = code
	code += x.domain.Size() - 1
	y := newIntVar(1, DomainSet{[]int{0, 10, 20}})
	baseCode[y.id] = code
	code += y.domain.Size() - 1
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 1, y: 1}, -20))
	cnf, vars := c1.tocnf(make([]CSPClause, 0), make([]*BoolVar, 0))
	for i := 0; i < len(vars); i++ {
		vars[i].id = uint(i + 2)
		baseCode[vars[i].id] = code
		code++
	}
	fmt.Println(cnf, vars)
	// 1: x <= 0
	// 2: x <= 10
	// 3: y <= 0
	// 4: y <= 10
	sat := make([][]int, 0)
	for _, x := range cnf {
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			sat = append(sat, tmp...)
		}
	}
	fmt.Println(sat)
}

func TestEncode4(t *testing.T) {
	// x + y -2z - 20 <= 0
	baseCode := make(map[uint]int)
	code := 1
	x := newIntVar(0, DomainSet{[]int{0, 10, 20}})
	baseCode[x.id] = code
	code += x.domain.Size() - 1
	y := newIntVar(1, DomainSet{[]int{0, 10, 20}})
	baseCode[y.id] = code
	code += y.domain.Size() - 1
	z := newIntVar(2, DomainSet{[]int{0, 10, 20}})
	baseCode[z.id] = code
	code += z.domain.Size() - 1
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 1, y: 1, z: -2}, -20))
	cnf, vars := c1.tocnf(make([]CSPClause, 0), make([]*BoolVar, 0))
	for i := 0; i < len(vars); i++ {
		vars[i].id = uint(i + 3)
		baseCode[vars[i].id] = code
		code++
	}
	fmt.Println(cnf, vars)
	// 1: x <= 0
	// 2: x <= 10
	// 3: y <= 0
	// 4: y <= 10
	// 5: z <= 0
	// 6: z <= 10
	sat := make([][]int, 0)
	for _, x := range cnf {
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			sat = append(sat, tmp...)
		}
	}
	fmt.Println(sat)
}

func TestEncode5(t *testing.T) {
	// x + y + 1 <= 0
	baseCode := make(map[uint]int)
	code := 1
	x := newIntVar(0, DomainSet{[]int{0, 10, 20}})
	baseCode[x.id] = code
	code += x.domain.Size() - 1
	y := newIntVar(1, DomainSet{[]int{0, 10, 20}})
	baseCode[y.id] = code
	code += y.domain.Size() - 1
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 1, y: 1}, 1))
	cnf, vars := c1.tocnf(make([]CSPClause, 0), make([]*BoolVar, 0))
	for i := 0; i < len(vars); i++ {
		vars[i].id = uint(i + 2)
		baseCode[vars[i].id] = code
		code++
	}
	fmt.Println(cnf, vars)
	// 1: x <= 0
	// 2: x <= 10
	// 3: y <= 0
	// 4: y <= 10
	sat := make([][]int, 0)
	for _, x := range cnf {
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			sat = append(sat, tmp...)
		}
	}
	fmt.Println(sat)
}

func TestEncode6(t *testing.T) {
	// x + y - 1000 <= 0
	baseCode := make(map[uint]int)
	code := 1
	x := newIntVar(0, DomainSet{[]int{0, 10, 20}})
	baseCode[x.id] = code
	code += x.domain.Size() - 1
	y := newIntVar(1, DomainSet{[]int{0, 10, 20}})
	baseCode[y.id] = code
	code += y.domain.Size() - 1
	c1 := CSPLeZero(NewSum(map[*IntVar]int{x: 1, y: 1}, -1000))
	cnf, vars := c1.tocnf(make([]CSPClause, 0), make([]*BoolVar, 0))
	for i := 0; i < len(vars); i++ {
		vars[i].id = uint(i + 2)
		baseCode[vars[i].id] = code
		code++
	}
	fmt.Println(cnf, vars)
	// 1: x <= 0
	// 2: x <= 10
	// 3: y <= 0
	// 4: y <= 10
	sat := make([][]int, 0)
	for _, x := range cnf {
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			sat = append(sat, tmp...)
		}
	}
	fmt.Println(sat)
}

func TestEncode7(t *testing.T) {
	baseCode := make(map[uint]int)
	code := 1
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = newBoolVar(uint(i), false)
		baseCode[x[i].id] = code
		code++
	}
	y := make([]*IntVar, 10)
	d := DomainSet{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	for i, _ := range x {
		y[i] = newIntVar(uint(len(x)+i), d)
		baseCode[y[i].id] = code
		code += y[i].domain.Size() - 1
	}
	var c1, c2, c3 CSPConstraint
	auxx := make([]*IntVar, 0)
	c1, auxx = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[3]: 10, y[9]: 8}, 1)).Decomp(auxx)
	c2, auxx = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[5]: 10, y[6]: -10, y[9]: 8}, 2)).Decomp(auxx)
	c3, auxx = CSPEqZero(NewSum(map[*IntVar]int{y[0]: 4, y[1]: -3, y[2]: 1, y[8]: -20, y[9]: 8}, 3)).Decomp(auxx)
	c := CSPOr(CSPOr(c1, x[1], CSPOr(x[2], c2)), x[4], x[5], CSPAnd(c3, x[7])).ToLeZero()
	for i := 0; i < len(auxx); i++ {
		auxx[i].id = uint(len(x) + len(y) + i)
		baseCode[auxx[i].id] = code
		code += auxx[i].domain.Size() - 1
	}
	fmt.Print(c)
	cnf := make([]CSPClause, 0)
	auxb := make([]*BoolVar, 0)
	cnf, auxb = simplify(c, cnf, auxb, []CSPClause{})
	for i := 0; i < len(auxb); i++ {
		auxb[i].id = uint(len(x) + len(y) + len(auxx) + i)
		baseCode[auxb[i].id] = code
		code++
	}
	fmt.Println(auxx)
	fmt.Println(auxb)
	sat := make([][]int, 0)
	for _, x := range cnf {
		fmt.Println(x)
		if tmp, ok := Encode(x, baseCode); ok == false {
			fmt.Println("UNSAT")
		} else {
			fmt.Println(tmp)
			sat = append(sat, tmp...)
		}
	}
	// fmt.Println(sat)

	toDimacs([][]int(sat))
}
