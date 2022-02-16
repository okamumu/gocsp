package csp

import (
	_ "bufio"
	_ "bytes"
	"fmt"
	"github.com/okamumu/gocsp/pkg/sat"
	_ "io"
	_ "os"
	_ "strconv"
	_ "strings"
	"testing"
)

func TestCSP1(t *testing.T) {
	c := NewCSP()
	x := make([]*IntVar, 4)
	for i, _ := range x {
		x[i] = c.NewIntVarWithRange(0, 5)
	}
	c1 := CSPEqZero(NewSum(map[*IntVar]int{x[0]: 4, x[1]: -3, x[2]: 2}, -1))
	c2 := CSPLeZero(NewSum(map[*IntVar]int{x[0]: 1, x[1]: 2, x[3]: 4}, -19))
	c3 := CSPLeZero(NewSum(map[*IntVar]int{x[0]: 4, x[1]: 2, x[3]: 4}, -5))
	c4 := CSPAnd(c1, c2, c3)
	c.AddConstraint(c4, true)
	c.CNF()
	fmt.Println(c.cnf)
	c.genBase()
	c.Encode()

	s := sat.NewSolver()
	assigns := make(map[int]bool)
	options := sat.DefaultSolverOptions()
	for _, x := range c.sat {
		s.AddClauseFromCode(x, options)
	}
	s.Simplify()
	s.Solve(assigns, options)

	fmt.Println(assigns)
	if len(assigns) == 0 {
		fmt.Println("UNSAT")
	} else {
		fmt.Println("SAT")
		for _, v := range c.intVars {
			for i := 0; i < v.domain.size(); i++ {
				if i == v.domain.size()-1 {
					fmt.Println(v, v.domain[i])
					break
				} else if a := assigns[c.baseCode[v.id]+i]; a == true {
					fmt.Println(v, v.domain[i])
					break
				}
			}
		}
		for _, v := range c.boolVars {
			if a := assigns[c.baseCode[v.id]]; a == true {
				fmt.Println(v, "T")
			} else {
				fmt.Println(v, "F")
			}
		}
	}
}

func TestCSP2(t *testing.T) {
	c := NewCSP()
	x := make([]*BoolVar, 10)
	for i, _ := range x {
		x[i] = c.NewBoolVar()
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

	s := sat.NewSolver()
	assigns := make(map[int]bool)
	options := sat.DefaultSolverOptions()
	for _, x := range c.sat {
		s.AddClauseFromCode(x, options)
	}
	s.Simplify()
	s.Solve(assigns, options)

	fmt.Println(assigns)
	if len(assigns) == 0 {
		fmt.Println("UNSAT")
	} else {
		fmt.Println("SAT")
		for _, v := range c.intVars {
			for i := 0; i < v.domain.size(); i++ {
				if i == v.domain.size()-1 {
					fmt.Println(v, v.domain[i])
					break
				} else if a := assigns[c.baseCode[v.id]+i]; a == true {
					fmt.Println(v, v.domain[i])
					break
				}
			}
		}
		for _, v := range c.boolVars {
			if a := assigns[c.baseCode[v.id]]; a == true {
				fmt.Println(v, "T")
			} else {
				fmt.Println(v, "F")
			}
		}
	}
}

func TestCSP3(t *testing.T) {
	// map coloring
	c := NewCSP()
	yamaguchi := c.NewIntVarWithRange(0, 2) // 3 colors
	hiroshima := c.NewIntVarWithRange(0, 2)
	okayama := c.NewIntVarWithRange(0, 2)
	shimane := c.NewIntVarWithRange(0, 2)
	tottori := c.NewIntVarWithRange(0, 2)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{yamaguchi: 1, shimane: -1}, 0)), true)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{yamaguchi: 1, hiroshima: -1}, 0)), true)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{hiroshima: 1, shimane: -1}, 0)), true)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{hiroshima: 1, tottori: -1}, 0)), true)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{hiroshima: 1, okayama: -1}, 0)), true)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{okayama: 1, tottori: -1}, 0)), true)
	c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{shimane: 1, tottori: -1}, 0)), true)
	c.CNF()
	fmt.Println(c.cnf)
	c.genBase()
	c.Encode()

	s := sat.NewSolver()
	assigns := make(map[int]bool)
	options := sat.DefaultSolverOptions()
	for _, x := range c.sat {
		s.AddClauseFromCode(x, options)
	}
	s.Simplify()
	s.Solve(assigns, options)

	fmt.Println(assigns)
	if len(assigns) == 0 {
		fmt.Println("UNSAT")
	} else {
		fmt.Println("SAT")
		for _, v := range c.intVars {
			for i := 0; i < v.domain.size(); i++ {
				if i == v.domain.size()-1 {
					fmt.Println(v, v.domain[i])
					break
				} else if a := assigns[c.baseCode[v.id]+i]; a == true {
					fmt.Println(v, v.domain[i])
					break
				}
			}
		}
		for _, v := range c.boolVars {
			if a := assigns[c.baseCode[v.id]]; a == true {
				fmt.Println(v, "T")
			} else {
				fmt.Println(v, "F")
			}
		}
	}
}

func TestCSP4(t *testing.T) {
	// map coloring
	c := NewCSP()
	x := make([]*IntVar, 47)
	m := 3 // 3 colors
	for i, _ := range x {
		x[i] = c.NewIntVarWithRange(0, m-1)
	}
	adj := [][]int{
		[]int{1},
		[]int{0, 2, 4},
		[]int{1, 3, 4},
		[]int{2, 4, 5, 6},
		[]int{1, 2, 3, 5},
		[]int{3, 4, 6, 14},
		[]int{3, 5, 7, 8, 9, 14},
		[]int{6, 8, 10, 11},
		[]int{6, 7, 9, 10},
		[]int{6, 8, 10, 14, 19},
		[]int{7, 8, 9, 11, 12, 19, 18},
		[]int{7, 10, 12, 13},
		[]int{10, 11, 13, 18},
		[]int{11, 12, 18, 21},
		[]int{5, 6, 9, 15, 19},
		[]int{14, 16, 19, 20},
		[]int{15, 17, 20},
		[]int{16, 20, 24, 25},
		[]int{10, 12, 13, 19, 21},
		[]int{9, 10, 14, 15, 18, 20, 21, 22},
		[]int{15, 16, 17, 19, 22, 23, 24},
		[]int{13, 18, 19, 22},
		[]int{19, 20, 21, 23},
		[]int{20, 22, 24, 25, 28, 29},
		[]int{17, 20, 23, 25},
		[]int{17, 23, 24, 26, 27, 28},
		[]int{25, 27, 28, 29},
		[]int{25, 26, 30, 32, 35},
		[]int{23, 25, 26, 29},
		[]int{23, 26, 28},
		[]int{27, 31, 32, 33},
		[]int{30, 33, 34},
		[]int{27, 30, 33, 36},
		[]int{30, 31, 32, 34, 37},
		[]int{31, 33, 39},
		[]int{27, 36, 37, 38},
		[]int{32, 35, 37},
		[]int{33, 35, 36, 38},
		[]int{35, 37},
		[]int{34, 40, 42, 43},
		[]int{39, 41},
		[]int{40},
		[]int{39, 43, 44, 45},
		[]int{39, 42, 44},
		[]int{42, 43, 45},
		[]int{42, 44, 46},
		[]int{45},
	}
	for i, a := range adj {
		for _, j := range a {
			if i < j {
				c.AddConstraint(CSPNeZero(NewSum(map[*IntVar]int{x[i]: 1, x[j]: -1}, 0)), true)
			}
		}
	}
	c.CNF()
	fmt.Println(c.cnf)
	c.genBase()
	c.Encode()

	s := sat.NewSolver()
	assigns := make(map[int]bool)
	options := sat.DefaultSolverOptions()
	for _, x := range c.sat {
		s.AddClauseFromCode(x, options)
	}
	s.Simplify()
	s.Solve(assigns, options)

	fmt.Println(assigns)
	if len(assigns) == 0 {
		fmt.Println("UNSAT")
	} else {
		fmt.Println("SAT")
		for _, v := range c.intVars {
			for i := 0; i < v.domain.size(); i++ {
				if i == v.domain.size()-1 {
					fmt.Println(v, v.domain[i])
					break
				} else if a := assigns[c.baseCode[v.id]+i]; a == true {
					fmt.Println(v, v.domain[i])
					break
				}
			}
		}
		for _, v := range c.boolVars {
			if a := assigns[c.baseCode[v.id]]; a == true {
				fmt.Println(v, "T")
			} else {
				fmt.Println(v, "F")
			}
		}
	}
}

// func TestCSP5(t *testing.T) {
// 	// map coloring
// 	c := NewCSP()
// 	x := make([]*IntVar, 800)
// 	m := 2 // 2 colors
// 	K := 10
// 	all := make(map[*IntVar]int)
// 	for i, _ := range x {
// 		x[i] = c.NewIntVarWithRange(0, m-1)
// 		all[x[i]] = 1
// 	}
// 	total := c.NewIntVarWithRange(0, 800)
// 	all[total] = -1
// 	c.AddConstraint(CSPEqZero(NewSum(all, 0)), true)
// 	c.AddConstraint(CSPGeZero(NewSum(map[*IntVar]int{total: 1}, -K)), true)
// 	file, _ := os.Open("../../brock800-4.mtx")
// 	defer file.Close()
// 	b, err := io.ReadAll(file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	in := bytes.NewBuffer(b)
// 	sc := bufio.NewScanner(in)
// 	for sc.Scan() {
// 		a := strings.Fields(sc.Text())
// 		if len(a) == 2 {
// 			i, _ := strconv.Atoi(a[0])
// 			j, _ := strconv.Atoi(a[1])
// 			c.AddConstraint(CSPLeZero(NewSum(map[*IntVar]int{x[i-1]: 1, x[j-1]: 1}, -1)), true)
// 		}
// 	}
// 	c.CNF()
// 	// fmt.Println(c.cnf)
// 	c.genBase()
// 	c.Encode()

// 	toDimacs(c.sat)

// 	// s := sat.NewSolver()
// 	// assigns := make(map[int]bool)
// 	// options := sat.DefaultSolverOptions()
// 	// for _, x := range c.sat {
// 	// 	s.AddClauseFromCode(x, options)
// 	// }
// 	// s.Simplify()
// 	// s.Solve(assigns, options)

// 	// if len(assigns) == 0 {
// 	// 	fmt.Println("UNSAT")
// 	// } else {
// 	// 	fmt.Println("SAT")
// 	// 	for _, v := range c.intVars {
// 	// 		for i := 0; i < v.domain.size(); i++ {
// 	// 			if i == v.domain.size()-1 {
// 	// 				fmt.Println(v, v.domain[i])
// 	// 				break
// 	// 			} else if a := assigns[c.baseCode[v.id]+i]; a == true {
// 	// 				fmt.Println(v, v.domain[i])
// 	// 				break
// 	// 			}
// 	// 		}
// 	// 	}
// 	// 	for _, v := range c.boolVars {
// 	// 		if a := assigns[c.baseCode[v.id]]; a == true {
// 	// 			fmt.Println(v, "T")
// 	// 		} else {
// 	// 			fmt.Println(v, "F")
// 	// 		}
// 	// 	}
// 	// }
// }
