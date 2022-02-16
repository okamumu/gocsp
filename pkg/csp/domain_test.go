package csp

// import (
// 	_ "fmt"
// 	"testing"
// )

// func TestDomainContains(t *testing.T) {
// 	d := DomainSet{1, 4, 6, 7, 11}
// 	result_pos := []int{}
// 	result_ok := []bool{}
// 	for i := 0; i <= 15; i++ {
// 		pos, ok := d.contains(i)
// 		result_pos = append(result_pos, pos)
// 		result_ok = append(result_ok, ok)
// 	}
// 	expected_pos := []int{0, 0, 0, 0, 1, 1, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}
// 	expected_ok := []bool{false, true, false, false, true, false, true, true, false, false, false, true, false, false, false, false}
// 	for i := 0; i <= 15; i++ {
// 		if expected_pos[i] != result_pos[i] {
// 			t.Errorf("pos: Expected %d, Result %d\n", expected_pos[i], result_pos[i])
// 		}
// 		if expected_ok[i] != result_ok[i] {
// 			t.Errorf("ok: Expected %t, Result %t\n", expected_ok[i], result_ok[i])
// 		}
// 	}
// }

// func TestDomainFilter(t *testing.T) {
// 	x := []int{1, 2, 3, 4, 4, 5, 6, 8}
// 	fmt.Println(x)
// 	x = filter(x, func(x int) bool {
// 		return x > 4
// 	})
// 	fmt.Println(x)
// }

// func TestDomainBound(t *testing.T) {
// 	d := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d.Bound(3, 10)
// 	expected := []int{4, 6, 7}
// 	for i, x := range expected {
// 		if d.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d.x[i])
// 		}
// 	}
// }

// func TestDomainCap1(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainSet{
// 		x: []int{4, 5, 6, 10, 19},
// 	}
// 	d1.Cap(d2)
// 	expected := []int{4, 6, 19}
// 	if len(expected) != len(d1.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d1.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d1.x[i])
// 		}
// 	}
// }

// func TestDomainCap2(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainSet{
// 		x: []int{2, 5, 10, 20},
// 	}
// 	d1.Cap(d2)
// 	expected := []int{}
// 	if len(expected) != len(d1.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d1.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d1.x[i])
// 		}
// 	}
// }

// func TestDomainCap2(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainInterval{
// 		lower: 4,
// 		upper: 8,
// 	}
// 	d1.Cap(d2)
// 	fmt.Println(d1)
// }

// func TestDomainCup1(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainSet{
// 		x: []int{4, 5, 10, 25},
// 	}
// 	d1.Cup(d2)
// 	expected := []int{1, 4, 5, 6, 7, 10, 19, 25}
// 	if len(expected) != len(d1.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d1.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d1.x[i])
// 		}
// 	}
// }

// func TestDomainNeg1(t *testing.T) {
// 	d := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d.Neg()
// 	expected := []int{-19, -7, -6, -4, -1}
// 	if len(expected) != len(d.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d.x[i])
// 		}
// 	}
// }

// func TestDomainNeg2(t *testing.T) {
// 	d := DomainSet{
// 		x: []int{1, 4, 7, 19},
// 	}
// 	d.Neg()
// 	expected := []int{-19, -7, -4, -1}
// 	if len(expected) != len(d.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d.x[i])
// 		}
// 	}
// }

// func TestDomainAdd1(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainSet{
// 		x: []int{4, 5, 10, 25},
// 	}
// 	d1.Add(d2)
// 	expected := []int{5, 6, 8, 9, 10, 11, 12, 14, 16, 17, 23, 24, 26, 29, 31, 32, 44}
// 	if len(expected) != len(d1.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d1.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d1.x[i])
// 		}
// 	}
// }

// func TestDomainSub1(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainSet{
// 		x: []int{4, 5, 10, 25},
// 	}
// 	d1.Sub(d2)
// 	expected := []int{-24, -21, -19, -18, -9, -6, -4, -3, -1, 0, 1, 2, 3, 9, 14, 15}
// 	if len(expected) != len(d1.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d1.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d1.x[i])
// 		}
// 	}
// }

// func TestDomainMul1(t *testing.T) {
// 	d1 := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d2 := DomainSet{
// 		x: []int{4, 5, 10, 25},
// 	}
// 	d1.Mul(d2)
// 	expected := []int{4, 5, 10, 16, 20, 24, 25, 28, 30, 35, 40, 60, 70, 76, 95, 100, 150, 175, 190, 475}
// 	if len(expected) != len(d1.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d1.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d1.x[i])
// 		}
// 	}
// }

// func TestDomainFloorDiv(t *testing.T) {
// 	d := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d.FloorDiv(4)
// 	expected := []int{0, 1, 4}
// 	if len(expected) != len(d.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d.x[i])
// 		}
// 	}
// }

// func TestDomainCeilDiv(t *testing.T) {
// 	d := DomainSet{
// 		x: []int{1, 4, 6, 7, 19},
// 	}
// 	d.CeilDiv(4)
// 	expected := []int{1, 2, 5}
// 	if len(expected) != len(d.x) {
// 		t.Errorf("Length is different")
// 	}
// 	for i, x := range expected {
// 		if d.x[i] != x {
// 			t.Errorf("Expected %d, Result %d\n", x, d.x[i])
// 		}
// 	}
// }
