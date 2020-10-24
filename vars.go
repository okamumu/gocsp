package csp

type VarLabel string

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

func NewAuxBoolVar(neg bool) *BoolVar {
	return &BoolVar{
		neg: neg,
	}
}

func (b *BoolVar) Not() CSPConstraint {
	nb := NewBoolVar(b.label, b.neg)
	nb.neg = !nb.neg
	return nb
}

func (b *BoolVar) ToLeZero() CSPConstraint {
	return b
}

func (b *BoolVar) Decomp(auxvars []*IntVar) (CSPConstraint, []*IntVar) {
	return b, auxvars
}

type IntVar struct {
	label  VarLabel
	domain *DomainSet
}

func NewIntVar(label VarLabel, domain *DomainSet) *IntVar {
	return &IntVar{
		label:  label,
		domain: domain,
	}
}

func NewAuxIntVar(domain *DomainSet) *IntVar {
	return &IntVar{
		label:  "aux",
		domain: domain,
	}
}
