package csp

import (
	// 	"log"
	"sort"
)

type DomainInterface interface {
	Size() int
	Get(i int) int
}

type DomainSet struct {
	x []int
}

func (d *DomainSet) copy() *DomainSet {
	x := make([]int, len(d.x))
	for i, v := range d.x {
		x[i] = v
	}
	return &DomainSet{
		x: x,
	}
}

func (d *DomainSet) Size() int {
	return len(d.x)
}

func (d *DomainSet) Get(i int) int {
	return d.x[i]
}

func (d *DomainSet) Lb() int {
	return d.x[0]
}

func (d *DomainSet) Ub() int {
	return d.x[len(d.x)-1]
}

// The function to give us the position of a given value with a binary search
// The second return value indicates whehter a value cotaints or not.
// If the second value is false, the first value indicates the following conditions
// for a given value x:
//   x < lb  -> -1
//   lb <= x <= ub -> 0
//   ub < x -> 1
func (d *DomainSet) Contains(value int) (int, bool) {
	l := 0
	u := len(d.x) - 1
	if value < d.x[l] {
		return l, false
	}
	if d.x[u] < value {
		return u, false
	}
	if value == d.x[l] {
		return l, true
	}
	if d.x[u] == value {
		return u, true
	}
	for u-l > 1 {
		m := (l + u) / 2
		// 		log.Printf("l:%d m:%d u:%d", l, m, u)
		if value == d.x[m] {
			return m, true
		} else if value < d.x[m] {
			u = m
		} else {
			l = m
		}
	}
	return l, false
}

func (d *DomainSet) Bound(lb, ub int) {
	// 	x := make([]int, 0, len(d.x))
	// 	for _, v := range d.x {
	// 		if lb <= v && v <= ub {
	// 			x = append(x, v)
	// 		}
	// 	}
	// 	d.x = x
	x := d.x
	for i, v := range d.x {
		if v < lb {
			continue
		}
		x = x[i:]
		break
	}
	for i, v := range x {
		if v < ub {
			continue
		}
		x = x[:i]
		break
	}
	d.x = x
}

func (d *DomainSet) Cap(other DomainInterface) {
	x := make([]int, 0)
	i := 0
	j := 0
	for i != d.Size() && j != other.Size() {
		if d.x[i] == other.Get(j) {
			x = append(x, d.x[i])
			i++
			j++
		} else if d.x[i] < other.Get(j) {
			i++
		} else {
			j++
		}
	}
	d.x = x
}

func (d *DomainSet) Cup(other DomainInterface) {
	x := make([]int, 0)
	i := 0
	j := 0
	for i != d.Size() && j != other.Size() {
		if d.x[i] == other.Get(j) {
			x = append(x, d.x[i])
			i++
			j++
		} else if d.x[i] < other.Get(j) {
			x = append(x, d.x[i])
			i++
		} else {
			x = append(x, other.Get(j))
			j++
		}
	}
	for ; i != d.Size(); i++ {
		x = append(x, d.x[i])
	}
	for ; j != other.Size(); j++ {
		x = append(x, other.Get(j))
	}
	d.x = x
}

func (d *DomainSet) Neg() {
	x := make([]int, len(d.x))
	for i, v := range d.x {
		x[len(d.x)-i-1] = -v
	}
	d.x = x
}

func (d *DomainSet) Func(other *DomainSet, f func(int, int) int) {
	tmpDomainSet = tmpDomainSet[:0]
	for _, v1 := range d.x {
		for _, v2 := range other.x {
			tmpDomainSet = append(tmpDomainSet, f(v1, v2))
		}
	}
	sort.Slice(tmpDomainSet, func(i, j int) bool {
		return tmpDomainSet[i] < tmpDomainSet[j]
	})
	x := make([]int, 0)
	var prev int
	for i, v := range tmpDomainSet {
		if i == 0 || v != prev {
			x = append(x, v)
			prev = v
		}
	}
	d.x = x
}

func (d *DomainSet) Add(other *DomainSet) {
	// tmpDomainSet = tmpDomainSet[:0]
	// for _, v1 := range d.x {
	// 	for _, v2 := range other.x {
	// 		tmpDomainSet = append(tmpDomainSet, v1+v2)
	// 	}
	// }
	// sort.Slice(tmpDomainSet, func(i, j int) bool {
	// 	return tmpDomainSet[i] < tmpDomainSet[j]
	// })
	// x := make([]int, 0)
	// var prev int
	// for i, v := range tmpDomainSet {
	// 	if i == 0 || v != prev {
	// 		x = append(x, v)
	// 		prev = v
	// 	}
	// }
	// d.x = x
	d.Func(other, func(x, y int) int {
		return x + y
	})
}

func (d *DomainSet) Sub(other *DomainSet) {
	d.Func(other, func(x, y int) int {
		return x - y
	})
	// tmpDomainSet = tmpDomainSet[:0]
	// for _, v1 := range d.x {
	// 	for _, v2 := range other.x {
	// 		tmpDomainSet = append(tmpDomainSet, v1-v2)
	// 	}
	// }
	// sort.Slice(tmpDomainSet, func(i, j int) bool {
	// 	return tmpDomainSet[i] < tmpDomainSet[j]
	// })
	// x := make([]int, 0)
	// var prev int
	// for i, v := range tmpDomainSet {
	// 	if i == 0 || v != prev {
	// 		x = append(x, v)
	// 		prev = v
	// 	}
	// }
	// d.x = x
}

func (d *DomainSet) Mul(other *DomainSet) {
	d.Func(other, func(x, y int) int {
		return x * y
	})
	// tmpDomainSet = tmpDomainSet[:0]
	// for _, v1 := range d.x {
	// 	for _, v2 := range other.x {
	// 		tmpDomainSet = append(tmpDomainSet, v1*v2)
	// 	}
	// }
	// sort.Slice(tmpDomainSet, func(i, j int) bool {
	// 	return tmpDomainSet[i] < tmpDomainSet[j]
	// })
	// x := make([]int, 0)
	// var prev int
	// for i, v := range tmpDomainSet {
	// 	if i == 0 || v != prev {
	// 		x = append(x, v)
	// 		prev = v
	// 	}
	// }
	// d.x = x
}

func floorDiv(b, a int) int {
	if a > 0 {
		if b >= 0 {
			return b / a
		} else {
			return (b - a + 1) / a
		}
	} else if b >= 0 {
		return (b - a - 1) / a
	} else {
		return b / a
	}
}

func ceilDiv(b, a int) int {
	if a > 0 {
		if b >= 0 {
			return (b + a - 1) / a
		} else {
			return b / a
		}
	} else if b >= 0 {
		return b / a
	} else {
		return (b + a + 1) / a
	}
}

func (d *DomainSet) FloorDiv(a int) {
	tmpDomainSet = tmpDomainSet[:0]
	for _, v1 := range d.x {
		tmpDomainSet = append(tmpDomainSet, floorDiv(v1, a))
	}
	sort.Slice(tmpDomainSet, func(i, j int) bool {
		return tmpDomainSet[i] < tmpDomainSet[j]
	})
	x := make([]int, 0)
	var prev int
	for i, v := range tmpDomainSet {
		if i == 0 || v != prev {
			x = append(x, v)
			prev = v
		}
	}
	d.x = x
}

func (d *DomainSet) CeilDiv(a int) {
	tmpDomainSet = tmpDomainSet[:0]
	for _, v1 := range d.x {
		tmpDomainSet = append(tmpDomainSet, ceilDiv(v1, a))
	}
	sort.Slice(tmpDomainSet, func(i, j int) bool {
		return tmpDomainSet[i] < tmpDomainSet[j]
	})
	x := make([]int, 0)
	var prev int
	for i, v := range tmpDomainSet {
		if i == 0 || v != prev {
			x = append(x, v)
			prev = v
		}
	}
	d.x = x
}

// function floorDiv(x::AbstractDomain{Ti}, y::Ti) where Ti
//     domain(union(sort([floorDiv(s, y) for s = x.range])))
// end

// function ceilDiv(x::AbstractDomain{Ti}, y::Ti) where Ti
//     domain(union(sort([ceilDiv(s, y) for s = x.range])))
// end

// function Base.:+(x::DomainSet{Ti}, y::Ti) where Ti
//     domain([u + y for u = x.range])
// end

// function Base.:+(x::DomainInterval{Ti}, y::Ti) where Ti
//     domain((x.range.start+y):(x.range.stop+y))
// end

// function Base.:+(x::Ti, y::AbstractDomain{Ti}) where Ti
//     Base.:+(y, x)
// end

// function Base.:+(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
//     domain(union(sort([s1 + s2 for s1 = x.range for s2 = y.range])))
// end

type DomainInterval struct {
	lower int
	upper int
}

func (d *DomainInterval) Size() int {
	return d.upper - d.lower + 1
}

func (d *DomainInterval) Get(i int) int {
	return d.lower + i
}

func (d *DomainInterval) Lb() int {
	return d.lower
}

func (d *DomainInterval) Ub() int {
	return d.upper
}

func (d *DomainInterval) Contains(value int) bool {
	return d.lower <= value && value <= d.upper
}

func (d *DomainInterval) Bound(lb, ub int) {
	if d.lower < lb {
		d.lower = lb
	}
	if ub < d.upper {
		d.upper = ub
	}
}

// """
// cap
// """

// function cap(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
//     domain(intersect(x.range, y.range))
// end

// function cap(x::DomainInterval{Ti}, y::DomainInterval{Ti}) where Ti
//     s1 = max(x.range.start, y.range.start)
//     s2 = min(x.range.stop, y.range.stop)
//     if s1 <= s2
//         domain(s1:s2)
//     else
//         domain(Ti[])
//     end
// end

// """
// cup
// """

// function cup(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
//     domain(union(x.range, y.range))
// end

// """
// neg
// """

// function Base.:-(x::DomainInterval{Ti}) where Ti
//     domain((-x.range.stop):(-x.range.start))
// end

// function Base.:-(x::DomainSet{Ti}) where Ti
//     domain(reverse(-x.range))
// end

// # """
// # abs
// # """

// # function Base.abs(x::DomainInterval{Ti}) where Ti
// #     x.range.start < Ti(0) && x.range.stop < Ti(0) && return domain((-x.range.stop):(-x.range.start))
// #     x.range.start < Ti(0) && x.range.stop >= Ti(0) && return domain(0:max(-x.range.start, x.range.stop))
// #     domain(x.range)
// # end

// # function Base.abs(x::DomainSet{Ti}) where Ti
// #     domain(unique(sort(abs.(x.range))))
// # end

// """
// add
// """

// function Base.:+(x::DomainSet{Ti}, y::Ti) where Ti
//     domain([u + y for u = x.range])
// end

// function Base.:+(x::DomainInterval{Ti}, y::Ti) where Ti
//     domain((x.range.start+y):(x.range.stop+y))
// end

// function Base.:+(x::Ti, y::AbstractDomain{Ti}) where Ti
//     Base.:+(y, x)
// end

// function Base.:+(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
//     domain(union(sort([s1 + s2 for s1 = x.range for s2 = y.range])))
// end

// """
// sub
// """

// function Base.:-(x::DomainSet{Ti}, y::Ti) where Ti
//     domain([u - y for u = x.range])
// end

// function Base.:-(x::DomainInterval{Ti}, y::Ti) where Ti
//     domain((x.range.start-y):(x.range.stop-y))
// end

// function Base.:-(x::Ti, y::AbstractDomain{Ti}) where Ti
//     Base.:+(x, Base.:-(y))
// end

// function Base.:-(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
//     domain(union(sort([s1 - s2 for s1 = x.range for s2 = y.range])))
// end

// """
// mul
// """

// function Base.:*(x::AbstractDomain{Ti}, y::Ti) where Ti
//     domain(union(sort([s * y for s = x.range])))
// end

// function Base.:*(x::Ti, y::AbstractDomain{Ti}) where Ti
//     y * x
// end

// # function Base.:*(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
// #     domain(union(sort([s1 * s2 for s1 = x.range for s2 = y.range])))
// # end

// """
// div
// """

// function floorDiv(b::Ti, a::Ti)::Ti where Ti
//     a > 0 && begin
//         b >= 0 && return div(b, a)
//         return div(b-a+1, a)
//     end
//     b >= 0 && return div(b-a-1, a)
//     return div(b, a)
// end

// function ceilDiv(b::Ti, a::Ti)::Ti where Ti
//     a > 0 && begin
//         b >= 0 && return div(b+a-1, a)
//         return div(b, a)
//     end
//     b >= 0 && return div(b, a)
//     return div(b+a+1, a)
// end

// function floorDiv(x::AbstractDomain{Ti}, y::Ti) where Ti
//     domain(union(sort([floorDiv(s, y) for s = x.range])))
// end

// function ceilDiv(x::AbstractDomain{Ti}, y::Ti) where Ti
//     domain(union(sort([ceilDiv(s, y) for s = x.range])))
// end

// # """
// # mod
// # """

// # function Base.:%(x::AbstractDomain{Ti}, y::Ti) where Ti
// #     domain(union(sort([rem(s, y) for s = x.range])))
// # end

// # function Base.:%(x::Ti, y::AbstractDomain{Ti}) where Ti
// #     domain(union(sort([rem(x, s) for s = y.range])))
// # end

// # function Base.:%(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
// #     domain(union(sort([rem(s1, s2) for s1 = x.range for s2 = y.range])))
// # end

// # """
// # pow
// # """

// # function Base.:^(x::AbstractDomain{Ti}, y::Ti) where Ti
// #     domain(union(sort([s^y for s = x.range])))
// # end

// # """
// # min
// # """

// # function Base.min(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
// #     lb = min(minimum(x.range), minimum(y.range))
// #     ub = min(maximum(x.range), maximum(y.range))
// #     v = [s for s = x.range if lb <= s <= ub]
// #     append!(v, [s for s = y.range if lb <= s <= ub])
// #     domain(union(sort(v)))
// # end

// # """
// # max
// # """

// # function Base.max(x::AbstractDomain{Ti}, y::AbstractDomain{Ti}) where Ti
// #     lb = max(minimum(x.range), minimum(y.range))
// #     ub = max(maximum(x.range), maximum(y.range))
// #     v = [s for s = x.range if lb <= s <= ub]
// #     append!(v, [s for s = y.range if lb <= s <= ub])
// #     domain(union(sort(v)))
// # end
