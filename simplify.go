package csp

type CSPClause struct {
	x []CSPConstraint
}

func NewCSPClause(x []CSPConstraint) *CSPClause {
	return &CSPClause{
		x: x,
	}
}

// func (l *CSPClause) add(c CSPConstraint) {
// 	l.x = append(l.x, c)
// }

// tocnf: The function is to create CNF
func (c *CSPOperator) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	switch c.op {
	case CSPOperatorAnd:
		for _, x := range c.args {
			cnf, auxvars = x.tocnf(cnf, auxvars)
		}
		return cnf, auxvars
	case CSPOperatorOr:
		cs := c.flattenOr(make([]CSPConstraint, 0))
		lt := make([]CSPConstraint, 0)
		lits := make([]CSPConstraint, 0)
		for _, x := range cs {
			lt, lits, auxvars = x.testin(lt, lits, auxvars)
		}
		cnf = append(cnf, CSPClause{x: lt})
		for _, x := range lits {
			cnf, auxvars = x.tocnf(cnf, auxvars)
		}
		return cnf, auxvars
	default:
		panic("")
	}
}

func (c *CSPComparator) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	lt := NewCSPClause([]CSPConstraint{c})
	return append(cnf, *lt), auxvars
}

func (b *BoolVar) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	lt := NewCSPClause([]CSPConstraint{b})
	return append(cnf, *lt), auxvars
}

// flattenOr: The function is to flatten list of literals for two or more OR operations.
// Example: Or(Or(a,b,c), AND(d, e)) -> [a, b, c, AND(d,e)]
func (c *CSPOperator) flattenOr(result []CSPConstraint) []CSPConstraint {
	switch c.op {
	case CSPOperatorAnd:
		return append(result, c)
	case CSPOperatorOr:
		for _, x := range c.args {
			result = x.flattenOr(result)
		}
		return result
	default:
		panic("")
	}
}

func (c *CSPComparator) flattenOr(result []CSPConstraint) []CSPConstraint {
	return append(result, c)
}

func (b *BoolVar) flattenOr(result []CSPConstraint) []CSPConstraint {
	return append(result, b)
}

// testin: Testin transform
func (c *CSPOperator) testin(first []CSPConstraint, result []CSPConstraint, auxvars []*BoolVar) ([]CSPConstraint, []CSPConstraint, []*BoolVar) {
	switch c.op {
	case CSPOperatorAnd:
		p := NewAuxBoolVar(false)
		auxvars = append(auxvars, p)
		first = append(first, p)
		for _, x := range c.args {
			result = append(result, CSPOr(p.Not(), x))
		}
		return first, result, auxvars
	case CSPOperatorOr:
		p := NewAuxBoolVar(false)
		auxvars = append(auxvars, p)
		first = append(first, p)
		c.args = append(c.args, p.Not())
		result = append(result, CSPOr(c.args...))
		return first, result, auxvars
	default:
		panic("")
	}
}

func (c *CSPComparator) testin(first []CSPConstraint, result []CSPConstraint, auxvars []*BoolVar) ([]CSPConstraint, []CSPConstraint, []*BoolVar) {
	return append(first, c), result, auxvars
}

func (b *BoolVar) testin(first []CSPConstraint, result []CSPConstraint, auxvars []*BoolVar) ([]CSPConstraint, []CSPConstraint, []*BoolVar) {
	return append(first, b), result, auxvars
}
