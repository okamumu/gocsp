package csp

import (
	// "log"
	"strconv"
)

// BoolVar

type BoolVar struct {
	id  uint
	neg bool
	aux bool
}

func newBoolVar(id uint, neg bool) *BoolVar {
	return &BoolVar{
		id:  id,
		neg: neg,
		aux: false,
	}
}

func newAuxBoolVar(id uint, neg bool) *BoolVar {
	return &BoolVar{
		id:  id,
		neg: neg,
		aux: true,
	}
}

func (x BoolVar) String() string {
	var s string
	if x.aux == false {
		s = "b" + strconv.Itoa(int(x.id))
	} else {
		s = "ab" + strconv.Itoa(int(x.id))
	}
	if x.neg {
		return "!" + s
	} else {
		return s
	}
}

// IntVar

type IntVar struct {
	id     uint
	domain DomainSet
	aux    bool
}

func newIntVar(id uint, domain DomainSet) *IntVar {
	return &IntVar{
		id:     id,
		domain: domain,
		aux:    false,
	}
}

func NewAuxIntVar(id uint, domain DomainSet) *IntVar {
	return &IntVar{
		id:     id,
		domain: domain,
		aux:    true,
	}
}

func (x IntVar) String() string {
	var s string
	if x.aux == false {
		s = "x" + strconv.Itoa(int(x.id))
	} else {
		s = "ax" + strconv.Itoa(int(x.id))
	}
	return s
}

// func NewIntVarWithRange(id uint, lb, ub int) *IntVar {
// 	x := make([]int, ub-lb+1)
// 	for i, _ := range x {
// 		x[i] = lb + i
// 	}
// 	return NewIntVar(id, DomainSet{x: x})
// }

// func NewIntVarWithSet(id uint, s []int) *IntVar {
// 	return NewIntVar(id, DomainSet{x: s})
// }
