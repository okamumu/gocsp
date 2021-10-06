package csp

// import (
// 	// 	"fmt"
// 	"github.com/pocke/go-minisat"
// )

// func MakeModel(codes []SATClause, vars map[interface{}]SATCode) {
// 	s := minisat.NewSolver(0.1)
// 	itov := make([]interface{}, len(vars))
// 	msvars := make([]*minisat.Var, len(vars))
// 	for k, c := range vars {
// 		itov[c] = k
// 		msvars[c] = s.NewVar()
// 	}
// 	tmp := make([]*minisat.Var, 0)
// 	for _, clause := range codes {
// 		tmp = tmp[:0]
// 		for _, c := range clause {
// 			if c > 0 {
// 				tmp = append(tmp, msvars[c])
// 			} else if c < 0 {
// 				tmp = append(tmp, msvars[c].Not())
// 			}
// 		}
// 		s.AddClause(tmp...)
// 	}
// 	return s
// }
