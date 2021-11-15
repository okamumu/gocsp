package csp

// Order encode: http://bach.istc.kobe-u.ac.jp/sugar/docs/order-encoding/

import (
	_ "errors"
	_ "fmt"
	"log"
	"sort"
)

// var (
// 	baseBool    map[uint]int
// 	baseInt     map[uint]int
// 	baseAuxBool map[uint]int
// 	baseAuxInt  map[uint]int
// )

type Clause []int

// func makeBase(x []*IntVar, auxx []*IntVar, b []*BoolVar, auxb []*BoolVar) {
// 	baseBool = make(map[uint]int)
// 	baseInt = make(map[uint]int)
// 	baseAuxBool = make(map[uint]int)
// 	baseAuxInt = make(map[uint]int)
// 	c := 1
// 	for _, v := range x {
// 		baseInt[v.id] = c
// 		c += v.domain.Size() - 1
// 	}
// 	for _, v := range auxx {
// 		baseAuxInt[v.id] = c
// 		c += v.domain.Size() - 1
// 	}
// 	for _, v := range b {
// 		baseBool[v.id] = c
// 		c += 1
// 	}
// 	for _, v := range auxb {
// 		baseAuxBool[v.id] = c
// 		c += 1
// 	}
// }

// func (x *BoolVar) getSATCodeBase(b map[uint]int) int {
// 	if x.aux == false {
// 		return baseBool[x.id]
// 	} else {
// 		return baseAuxBool[x.id]
// 	}
// }

// func (x *IntVar) getSATCodeBase() int {
// 	return baseInt[x.id]
// }

func Encode(c CSPClause, baseCode map[uint]int) ([]Clause, bool) {
	cs := []Clause{Clause{}}
	unsat := false
	var ok bool
	for _, p := range c {
		cs, ok = p.encode(cs, baseCode)
		unsat = unsat || ok
	}
	return cs, unsat
}

func applyOr(cs []Clause, p ...int) []Clause {
	for i, _ := range cs {
		cs[i] = append(cs[i], p...)
	}
	return cs
}

func (b *BoolVar) encode(codes []Clause, baseCode map[uint]int) ([]Clause, bool) {
	base := baseCode[b.id]
	var p int
	if b.neg == false {
		p = base
	} else {
		p = -base
	}
	return applyOr(codes, p), true
}

func (c *CSPComparator) encode(codes []Clause, baseCode map[uint]int) ([]Clause, bool) {
	if c.op == CSPOperatorLeZero {
		vars := make([]*IntVar, 0, len(c.s.coef))
		for k, _ := range c.s.coef {
			vars = append(vars, k)
		}
		sort.Slice(vars, func(i, j int) bool {
			s1 := vars[i].domain.Size()
			s2 := vars[j].domain.Size()
			if s1 == s2 {
				k1 := abs(c.s.coef[vars[i]])
				k2 := abs(c.s.coef[vars[j]])
				return k1 > k2
			} else {
				return s1 < s2
			}
		})
		if cs, ok := encodeIntVar([]Clause{}, make(Clause, 0, len(vars)), vars, c.s, -c.s.b, baseCode); ok {
			switch {
			case len(cs) == 0: // trueLit
				return []Clause{}, true
			case len(cs) == 1: // simple
				return applyOr(codes, cs[0]...), true
			case len(codes) == 1:
				return applyOr(cs, codes[0]...), true
			default:
				result := make([]Clause, 0, len(cs)*len(codes))
				for _, x := range codes {
					for _, y := range cs {
						p := make(Clause, 0, len(x)+len(y))
						p = append(p, x...)
						p = append(p, y...)
						result = append(result, p)
					}
				}
				return result, true
			}
		} else {
			// falseLit
			return codes, false
		}
	} else {
		log.Fatal("SAT encoding can be applied to LeZero")
		return codes, false
	}
}

// Encoding based on the formula: a_1 x_1 + ... a_n x_n <= b
func encodeIntVar(cs []Clause, lit Clause, vars []*IntVar, s *Sum, rhs int, baseCode map[uint]int) ([]Clause, bool) {
	// log.Println("encodeIntVar: start", "cs", cs, "vars", vars, "rhs", rhs)
	a := s.coef[vars[0]]
	base := baseCode[vars[0].id]
	dom := vars[0].domain
	var ok bool
	if len(vars) == 1 {
		return encodeLe(cs, lit, base, dom, a, rhs)
	}
	if a > 0 {
		b := dom.x[0]
		cs, ok = encodeIntVar(cs, lit, vars[1:], s, rhs-a*b, baseCode)
		if ok == false {
			return cs, ok
		}
		for i, b := range dom.x[1:] {
			cs, ok = encodeIntVar(cs, append(lit, base+i), vars[1:], s, rhs-a*b, baseCode)
			if ok == false {
				return cs, ok
			}
		}
	} else { // assume a < 0
		for i, b := range dom.x[:dom.Size()-1] {
			cs, ok = encodeIntVar(cs, append(lit, -(base+i)), vars[1:], s, rhs-a*b, baseCode)
			if ok == false {
				return cs, ok
			}
		}
		b := dom.x[dom.Size()-1]
		cs, ok = encodeIntVar(cs, lit, vars[1:], s, rhs-a*b, baseCode)
		if ok == false {
			return cs, ok
		}
	}
	return cs, true
}

// SATcode for a*x <= b, x in dom
func encodeLe(cs []Clause, lit Clause, base int, dom DomainSet, a, b int) ([]Clause, bool) {
	if a > 0 {
		c := floorDiv(b, a)
		// log.Println("base", base, "a", a, "b", b, "c", c, "dom", dom.x)
		pos, ok := dom.Contains(c)
		switch {
		case pos == ErrDomainLower && ok == false: // falseLit
			if len(lit) == 0 {
				return cs, false
			} else {
				return append(cs, clauseCopy(lit)), true
			}
		case pos == ErrDomainUpper && ok == false: // trueLit
			return cs, true
		default:
			lit = append(lit, base+pos)
			return append(cs, clauseCopy(lit)), true
		}
	} else {
		c := ceilDiv(b, a) - 1
		// log.Println("base", base, "a", a, "b", b, "c", c, "dom", dom.x)
		pos, ok := dom.Contains(c)
		switch {
		case pos == ErrDomainLower && ok == false:
			return cs, true
		case pos == ErrDomainUpper && ok == false:
			if len(lit) == 0 {
				return cs, false
			} else {
				return append(cs, clauseCopy(lit)), true
			}
		default:
			lit = append(lit, -(base + pos))
			return append(cs, clauseCopy(lit)), true
		}
	}
}

func clauseCopy(s Clause) Clause {
	t := make([]int, len(s))
	copy(t, []int(s))
	return Clause(t)
}
