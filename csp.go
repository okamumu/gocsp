package csp

type CSPOperatorType int

const (
	CSPOperatorAnd CSPOperatorType = iota + 1
	CSPOperatorOr
	CSPOperatorLeZero
	CSPOperatorGeZero
	CSPOperatorEqZero
	CSPOperatorNeZero
)

type CSPConstraint interface {
	Not() CSPConstraint
	ToLeZero() CSPConstraint
	Decomp([]*IntVar) (CSPConstraint, []*IntVar)
}

type CSPComparator struct {
	op CSPOperatorType
	s  *Sum
}

type CSPOperator struct {
	op   CSPOperatorType
	args []CSPConstraint
}

func CSPAnd(args ...CSPConstraint) *CSPOperator {
	return &CSPOperator{
		op:   CSPOperatorAnd,
		args: args,
	}
}

func CSPOr(args ...CSPConstraint) *CSPOperator {
	return &CSPOperator{
		op:   CSPOperatorOr,
		args: args,
	}
}

func CSPImp(x, y CSPConstraint) CSPConstraint {
	return CSPOr(x.Not(), y)
}

func CSPIff(x, y CSPConstraint) CSPConstraint {
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

func (c *CSPOperator) Not() CSPConstraint {
	switch c.op {
	case CSPOperatorAnd:
		newargs := make([]CSPConstraint, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.Not())
		}
		return CSPOr(newargs...)
	case CSPOperatorOr:
		newargs := make([]CSPConstraint, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.Not())
		}
		return CSPAnd(newargs...)
	default:
		panic("")
	}
}

func (c *CSPComparator) Not() CSPConstraint {
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

// ToLeZero

func (c *CSPComparator) ToLeZero() CSPConstraint {
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

func (c *CSPOperator) ToLeZero() CSPConstraint {
	switch c.op {
	case CSPOperatorAnd:
		newargs := make([]CSPConstraint, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.ToLeZero())
		}
		return CSPAnd(newargs...)
	case CSPOperatorOr:
		newargs := make([]CSPConstraint, len(c.args))
		for _, x := range c.args {
			newargs = append(newargs, x.ToLeZero())
		}
		return CSPAnd(newargs...)
	default:
		panic("")
	}
}
