package csp

import (
	// "log"
	"strconv"
)

type VarLabel string
type CSPOperatorType int

const (
	CSPOperatorAnd CSPOperatorType = iota + 1
	CSPOperatorOr
	CSPOperatorLeZero
	CSPOperatorGeZero
	CSPOperatorEqZero
	CSPOperatorNeZero
)

var tmpDomainSet []int // this is used for domain
var auxcount int       // this is used for numbering of auxvars
var tmpCNF []CSPClause // this is used in simplify
var satBaseCodes map[VarLabel]SATCode
var satcount SATCode

func init() {
	tmpDomainSet = make([]int, 0)
	tmpCNF = make([]CSPClause, 0)
	auxcount = 0
	satBaseCodes = make(map[VarLabel]SATCode)
	satcount = 1
}

type BoolVar struct {
	label VarLabel
	neg   bool
}

func NewBoolVar(label VarLabel, neg bool) *BoolVar {
	return &BoolVar{
		label: label,
		neg:   neg,
	}
}

func (b *BoolVar) getSATCodeBase() SATCode {
	if base, ok := satBaseCodes[b.label]; ok {
		return base
	} else {
		base = satcount
		satBaseCodes[b.label] = base
		satcount++
		return base
	}
}

func NewAuxBoolVar(neg bool) *BoolVar {
	auxcount++
	return &BoolVar{
		label: VarLabel("auxbool" + strconv.Itoa(auxcount)),
		neg:   neg,
	}
}

type IntVar struct {
	label  VarLabel
	domain *DomainSet
}

func (v *IntVar) getSATCodeBase() SATCode {
	if base, ok := satBaseCodes[v.label]; ok {
		return base
	} else {
		base = satcount
		satBaseCodes[v.label] = base
		satcount += SATCode(v.domain.Size() - 1)
		return base
	}
}

func NewIntVar(label VarLabel, domain *DomainSet) *IntVar {
	return &IntVar{
		label:  label,
		domain: domain,
	}
}

func NewAuxIntVar(domain *DomainSet) *IntVar {
	auxcount++
	return &IntVar{
		label:  VarLabel("auxint" + strconv.Itoa(auxcount)),
		domain: domain,
	}
}

type CSPLiteral interface {
	Not() CSPLiteral
	ToLeZero() CSPLiteral
	Decomp([]*IntVar) (CSPLiteral, []*IntVar)
	tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar)
	flattenOr(result []CSPLiteral) []CSPLiteral
	testin(first []CSPLiteral, result []CSPLiteral, auxvars []*BoolVar) ([]CSPLiteral, []CSPLiteral, []*BoolVar)
	isSimple() bool
	encode(codes []SATClause) []SATClause
}

type CSPComparator struct {
	op CSPOperatorType
	s  *Sum
}

type CSPOperator struct {
	op   CSPOperatorType
	args []CSPLiteral
}

func CSPAnd(args ...CSPLiteral) *CSPOperator {
	return &CSPOperator{
		op:   CSPOperatorAnd,
		args: args,
	}
}

func CSPOr(args ...CSPLiteral) *CSPOperator {
	return &CSPOperator{
		op:   CSPOperatorOr,
		args: args,
	}
}

func CSPImp(x, y CSPLiteral) CSPLiteral {
	return CSPOr(x.Not(), y)
}

func CSPIff(x, y CSPLiteral) CSPLiteral {
	return CSPAnd(CSPOr(x.Not(), y), CSPOr(x, y.Not()))
}

func CSPLeZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorLeZero,
		s:  s,
	}
}

func CSPGeZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorGeZero,
		s:  s,
	}
}

func CSPEqZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorEqZero,
		s:  s,
	}
}

func CSPNeZero(s *Sum) *CSPComparator {
	return &CSPComparator{
		op: CSPOperatorNeZero,
		s:  s,
	}
}

// not

func (c *CSPOperator) Not() CSPLiteral {
	switch c.op {
	case CSPOperatorAnd:
		newargs := make([]CSPLiteral, 0, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.Not())
		}
		return CSPOr(newargs...)
	case CSPOperatorOr:
		newargs := make([]CSPLiteral, 0, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.Not())
		}
		return CSPAnd(newargs...)
	default:
		panic("")
	}
}

func (c *CSPComparator) Not() CSPLiteral {
	s := c.s.copy()
	switch c.op {
	case CSPOperatorLeZero:
		s.AddConst(-1)
		return CSPGeZero(s)
	case CSPOperatorGeZero:
		s.AddConst(1)
		return CSPLeZero(s)
	case CSPOperatorEqZero:
		return CSPNeZero(s)
	case CSPOperatorNeZero:
		return CSPEqZero(s)
	default:
		panic("")
	}
}

func (b *BoolVar) Not() CSPLiteral {
	nb := NewBoolVar(b.label, b.neg)
	nb.neg = !nb.neg
	return nb
}

// ToLeZero
func (c *CSPComparator) ToLeZero() CSPLiteral {
	switch c.op {
	case CSPOperatorEqZero:
		s1 := c.s.copy()
		s2 := c.s.copy()
		return CSPAnd(CSPLeZero(s1), CSPLeZero(s2.Neg()))
	case CSPOperatorNeZero:
		s1 := c.s.copy()
		s2 := c.s.copy()
		return CSPOr(CSPLeZero(s1.AddConst(1)), CSPLeZero(s2.Neg().AddConst(1)))
	case CSPOperatorGeZero:
		s1 := c.s.copy()
		return CSPLeZero(s1.Neg())
	case CSPOperatorLeZero:
		s1 := c.s.copy()
		return CSPLeZero(s1)
	default:
		panic("")
	}
}

func (c *CSPOperator) ToLeZero() CSPLiteral {
	switch c.op {
	case CSPOperatorAnd:
		newargs := make([]CSPLiteral, 0, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.ToLeZero())
		}
		return CSPAnd(newargs...)
	case CSPOperatorOr:
		newargs := make([]CSPLiteral, 0, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.ToLeZero())
		}
		return CSPOr(newargs...)
	default:
		panic("")
	}
}

func (b *BoolVar) ToLeZero() CSPLiteral {
	return NewBoolVar(b.label, b.neg)
}
