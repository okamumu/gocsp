package csp

type Sum struct {
	coef map[*IntVar]int
	b    int
}

func NewSum(coef map[*IntVar]int, b int) *Sum {
	return &Sum{
		coef: coef,
		b:    b,
	}
}

func NewSumFromInt(v *IntVar) *Sum {
	coef := make(map[*IntVar]int)
	coef[v] = 1
	return &Sum{
		coef: coef,
		b:    0,
	}
}

func (s *Sum) Size() int {
	return len(s.coef)
}

func (s *Sum) Func(other *Sum, f func(int, int) int) *Sum {
	for k, v := range other.coef {
		if tmp := f(s.coef[k], v); tmp != 0 {
			s.coef[k] = tmp
		} else {
			delete(s.coef, k)
		}
	}
	s.b = f(s.b, other.b)
	return s
}

func (s *Sum) Neg() *Sum {
	for k, v := range s.coef {
		s.coef[k] = -v
	}
	s.b = -s.b
	return s
}

func (s *Sum) AddConst(b int) *Sum {
	s.b += b
	return s
}

func (s *Sum) Add(other *Sum) *Sum {
	return s.Func(other, func(x, y int) int {
		return x + y
	})
}

func (s *Sum) Sub(other *Sum) *Sum {
	return s.Func(other, func(x, y int) int {
		return x - y
	})
}

func (s *Sum) MulConst(a int) *Sum {
	if a != 0 {
		for k, v := range s.coef {
			s.coef[k] = a * v
		}
		s.b *= a
	} else {
		s.coef = make(map[*IntVar]int)
		s.b = 0
	}
	return s
}

func (s *Sum) copy() *Sum {
	coef := make(map[*IntVar]int)
	for k, v := range s.coef {
		coef[k] = v
	}
	return &Sum{
		coef: coef,
		b:    s.b,
	}
}
