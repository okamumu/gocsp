package csp

type CSPClause struct {
	x []CSPLiteral
}

func NewCSPClause(x []CSPLiteral) *CSPClause {
	return &CSPClause{
		x: x,
	}
}

// func (l *CSPClause) add(c CSPLiteral) {
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
		cs := c.flattenOr(make([]CSPLiteral, 0))
		lt := make([]CSPLiteral, 0)
		lits := make([]CSPLiteral, 0)
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
	lt := NewCSPClause([]CSPLiteral{c})
	return append(cnf, *lt), auxvars
}

func (b *BoolVar) tocnf(cnf []CSPClause, auxvars []*BoolVar) ([]CSPClause, []*BoolVar) {
	lt := NewCSPClause([]CSPLiteral{b})
	return append(cnf, *lt), auxvars
}

// flattenOr: The function is to flatten list of literals for two or more OR operations.
// Example: Or(Or(a,b,c), AND(d, e)) -> [a, b, c, AND(d,e)]
func (c *CSPOperator) flattenOr(result []CSPLiteral) []CSPLiteral {
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

func (c *CSPComparator) flattenOr(result []CSPLiteral) []CSPLiteral {
	return append(result, c)
}

func (b *BoolVar) flattenOr(result []CSPLiteral) []CSPLiteral {
	return append(result, b)
}

// testin: Testin transform
func (c *CSPOperator) testin(first []CSPLiteral, result []CSPLiteral, auxvars []*BoolVar) ([]CSPLiteral, []CSPLiteral, []*BoolVar) {
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

func (c *CSPComparator) testin(first []CSPLiteral, result []CSPLiteral, auxvars []*BoolVar) ([]CSPLiteral, []CSPLiteral, []*BoolVar) {
	return append(first, c), result, auxvars
}

func (b *BoolVar) testin(first []CSPLiteral, result []CSPLiteral, auxvars []*BoolVar) ([]CSPLiteral, []CSPLiteral, []*BoolVar) {
	return append(first, b), result, auxvars
}

// simplify
func (c CSPClause) isSimple() bool {
	count := 0
	for _, l := range c.x {
		if !l.isSimple() {
			if count++; count > 1 {
				return false
			}
		}
	}
	return true
}

func (c *CSPOperator) isSimple() bool {
	panic("")
}

func (c *CSPComparator) isSimple() bool {
	return c.s.Size() <= 1
}

func (b *BoolVar) isSimple() bool {
	return true
}

func Simplify(c CSPLiteral) ([]CSPClause, []*BoolVar) {
	cnf := make([]CSPClause, 0)
	simplecnf := make([]CSPClause, 0)
	vars := make([]*BoolVar, 0)
	cnf, vars = c.tocnf(cnf, vars)
	for _, clause := range cnf {
		if clause.isSimple() {
			simplecnf = append(simplecnf, clause)
		} else {
			first := make([]CSPLiteral, 0, len(clause.x))
			for _, lit := range clause.x {
				if lit.isSimple() {
					first = append(first, lit)
				} else {
					p := NewAuxBoolVar(false)
					vars = append(vars, p)
					first = append(first, p)
					simplecnf = append(simplecnf, *NewCSPClause([]CSPLiteral{p.Not(), lit}))
				}
			}
			simplecnf = append(simplecnf, *NewCSPClause(first))
		}
	}
	return simplecnf, vars
}
