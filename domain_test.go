package csp

import (
	"fmt"
	"testing"
)

func TestDomainContains(t *testing.T) {
	d := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	for i := 0; i < 100; i++ {
		pos, ok := d.Contains(i)
		fmt.Printf("%d : pos %d %t\n", i, pos, ok)
	}
}

func TestDomainBound(t *testing.T) {
	d := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d.Bound(0, 10)
	fmt.Println(d)
}

func TestDomainCap1(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d2 := DomainSet{
		x: []int{4, 5, 10, 19},
	}
	d1.Cap(&d2)
	fmt.Println(d1)
}

func TestDomainCap2(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d2 := DomainInterval{
		lower: 4,
		upper: 8,
	}
	d1.Cap(&d2)
	fmt.Println(d1)
}

func TestDomainCap3(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d2 := DomainSet{
		x: []int{2, 5, 10, 20},
	}
	d1.Cap(&d2)
	fmt.Println(d1)
}

func TestDomainCup1(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d2 := DomainSet{
		x: []int{4, 5, 10, 25},
	}
	d1.Cup(&d2)
	fmt.Println(d1)
}

func TestDomainNeg(t *testing.T) {
	d := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d.Neg()
	fmt.Println(d)
}

func TestDomainAdd1(t *testing.T) {
	d1 := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d2 := DomainSet{
		x: []int{4, 5, 10, 25},
	}
	d1.Add(&d2)
	fmt.Println(d1)
}

func TestDomainFloorDiv(t *testing.T) {
	d := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d.FloorDiv(4)
	fmt.Println(d)
}

func TestDomainCeilDiv(t *testing.T) {
	d := DomainSet{
		x: []int{1, 4, 6, 7, 19},
	}
	d.CeilDiv(4)
	fmt.Println(d)
}
