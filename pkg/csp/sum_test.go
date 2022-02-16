package csp

// import (
// 	"fmt"
// 	"testing"
// )

// func TestSum01(t *testing.T) {
// 	d := DomainSet{1, 4, 6, 7, 19}
// 	x1 := newIntVar(1, d)
// 	x2 := newIntVar(2, d)
// 	s1 := NewSum(map[*IntVar]int{x1: 1}, 0)
// 	s2 := NewSum(map[*IntVar]int{x2: 1}, 0)
// 	s1.MulConst(3)
// 	s1.Add(s2)
// 	s1.AddConst(5)

// 	s3 := NewSum(map[*IntVar]int{x1: 3, x2: 1}, 5)
// 	fmt.Println(s1)
// 	fmt.Println(s3)
// 	if s1.Size() != s3.Size() {
// 		t.Error("Different length")
// 	}
// 	for k, v := range s1.coef {
// 		if s3.coef[k] != v {
// 			t.Error("s1:", v, "s3:", s3.coef[k])
// 		}
// 	}
// 	if s1.b != s3.b {
// 		t.Error("s1:", s1.b, "s3:", s3.b)
// 	}
// }

// func TestSum02(t *testing.T) {
// 	d := DomainSet{1, 4, 6, 7, 19}
// 	x1 := newIntVar(1, d)
// 	x2 := newIntVar(2, d)
// 	s1 := NewSum(map[*IntVar]int{x1: 1}, 0)
// 	s2 := NewSum(map[*IntVar]int{x2: 1}, 0)
// 	s1.MulConst(-5)
// 	s1.Sub(s2)
// 	s1.AddConst(-5)
// 	s1.Neg()

// 	s3 := NewSum(map[*IntVar]int{x1: 5, x2: 1}, 5)
// 	// fmt.Println(s1)
// 	// fmt.Println(s3)
// 	if s1.Size() != s3.Size() {
// 		t.Error("Different length")
// 	}
// 	for k, v := range s1.coef {
// 		if s3.coef[k] != v {
// 			t.Error("s1:", v, "s3:", s3.coef[k])
// 		}
// 	}
// 	if s1.b != s3.b {
// 		t.Error("s1:", s1.b, "s3:", s3.b)
// 	}
// }
