package csp

import (
	"fmt"
	"github.com/okamumu/gocsp/pkg/sat"
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
			for i := 0; i < v.domain.Size(); i++ {
				if i == v.domain.Size()-1 {
					fmt.Println(v, v.domain.x[i])
					break
				} else if a := assigns[c.baseCode[v.id]+i]; a == true {
					fmt.Println(v, v.domain.x[i])
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
