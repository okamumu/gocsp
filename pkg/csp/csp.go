package csp

import (
	_ "fmt"
	"log"
)

type CSP struct {
	varId           uint
	boolVars        []*BoolVar
	auxBoolVars     []*BoolVar
	intVars         []*IntVar
	auxIntVars      []*IntVar
	constraints     []CSPConstraint
	cnf             []CSPClause
	cnfDone         []bool      // indicator whether the simplify is done or note
	cnfStart        []int       // the index to start CSPClause for the corresponding CSP constraint
	cnfStartAuxBool []int       // the index to start auxbool for the corresponding CSP constraint
	tmpCNF          []CSPClause // this is tempolary used in simplify

	baseCode map[uint]int // SAT code base
	sat      [][]int      // SAT code
}

func NewCSP() *CSP {
	return &CSP{
		varId:           0,
		boolVars:        make([]*BoolVar, 0),
		auxBoolVars:     make([]*BoolVar, 0),
		intVars:         make([]*IntVar, 0),
		auxIntVars:      make([]*IntVar, 0),
		constraints:     make([]CSPConstraint, 0),
		cnf:             make([]CSPClause, 0),
		cnfDone:         make([]bool, 0),
		cnfStart:        make([]int, 0),
		cnfStartAuxBool: make([]int, 0),
		tmpCNF:          make([]CSPClause, 0),

		baseCode: make(map[uint]int),
		sat:      make([][]int, 0),
	}
}

func (c *CSP) NewBoolVar(neg bool) *BoolVar {
	v := newBoolVar(c.varId, neg)
	c.boolVars = append(c.boolVars, v)
	c.varId++
	return v
}

func (c *CSP) NewIntVarWithRange(lb, ub int) *IntVar {
	x := make([]int, ub-lb+1)
	for i, _ := range x {
		x[i] = lb + i
	}
	d := DomainSet{x: x}
	v := newIntVar(c.varId, d)
	c.intVars = append(c.intVars, v)
	c.varId++
	return v
}

// AddConstraint
// The method to add a CSP constraint; Comparators, Operators and Bool
// The argument `decomp` indicates whether the CSP constraint is decomposed to up to three terms or not
// for all the linear functions in the CSP constraint.
func (c *CSP) AddConstraint(x CSPConstraint, decomp bool) {
	if decomp {
		var cs CSPConstraint
		start := len(c.auxIntVars)
		cs, c.auxIntVars = x.Decomp(c.auxIntVars)
		// rewrite id
		for k := start; k < len(c.auxIntVars); k++ {
			c.auxIntVars[k].id = c.varId
			c.varId++
		}
		c.constraints = append(c.constraints, cs.ToLeZero())
		c.cnfDone = append(c.cnfDone, false)
		c.cnfStart = append(c.cnfStart, 0)
		c.cnfStartAuxBool = append(c.cnfStartAuxBool, 0)
	} else {
		c.constraints = append(c.constraints, x.ToLeZero())
	}
}

// To save the current states (constraints and codes)
func (c *CSP) Save() int {
	// TODO: should be implemented
	return 0
}

// To load the status
func (c *CSP) Load(num int) {
	// TODO: should be implemented
}

func (c *CSP) CNF() {
	for i, cs := range c.constraints {
		if c.cnfDone[i] == false {
			c.cnfStart[i] = len(c.cnf)
			c.cnfStartAuxBool[i] = len(c.auxBoolVars)
			c.cnf, c.auxBoolVars = simplify(cs, c.cnf, c.auxBoolVars, c.tmpCNF)
			// rewrite id
			for k := c.cnfStartAuxBool[i]; k < len(c.auxBoolVars); k++ {
				c.auxBoolVars[k].id = c.varId
				c.varId++
			}
			c.cnfDone[i] = true
		}
	}
}

func (c *CSP) genBase() {
	code := 1
	for _, v := range c.intVars {
		c.baseCode[v.id] = code
		// for k := 0; k < v.domain.Size()-1; k++ {
		// 	log.Println("int", v, "<=", v.domain.x[k], "code", code+k)
		// }
		for k := 0; k < v.domain.Size()-2; k++ {
			c.sat = append(c.sat, []int{-(code + k), code + k + 1})
		}
		code += v.domain.Size() - 1
	}
	for _, v := range c.auxIntVars {
		c.baseCode[v.id] = code
		// for k := 0; k < v.domain.Size()-1; k++ {
		// 	log.Println("aux int", v, "<=", v.domain.x[k], "code", code+k)
		// }
		for k := 0; k < v.domain.Size()-2; k++ {
			c.sat = append(c.sat, []int{-(code + k), code + k + 1})
		}
		code += v.domain.Size() - 1
	}
	for _, v := range c.boolVars {
		// log.Println("bool", v, "code", code)
		c.baseCode[v.id] = code
		code += 1
	}
	for _, v := range c.auxBoolVars {
		// log.Println("aux bool", v, "code", code)
		c.baseCode[v.id] = code
		code += 1
	}
}

func (c *CSP) Encode() {
	for _, x := range c.cnf {
		if tmp, ok := Encode(x, c.baseCode); ok == false {
			log.Fatal("UNSAT", x)
		} else {
			// log.Println("CNF", x)
			// log.Println("Code", tmp)
			c.sat = append(c.sat, tmp...)
		}
	}
}
