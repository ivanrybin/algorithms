package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MergeSort non-recursive.
// T: O(nlogn) best/worst case
// M: O(nlogn) best/worst case
func MergeSort(a []int) []int {
	for size := 1; size <= len(a); size *= 2 {
		for l := 0; l < len(a); l += size * 2 {
			ll := l
			lr := min(len(a), l+size)
			rl := min(len(a), l+size)
			rr := min(len(a), l+2*size)
			copy(a[ll:rr], merge(a[ll:lr], a[rl:rr]))
		}
	}
	return a
}

// merge merges two sorted slices.
func merge(l, r []int) []int {
	m := make([]int, 0, len(l)+len(r))
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			m = append(m, l[i])
			i++
		} else {
			m = append(m, r[j])
			j++
		}
	}
	m = append(m, l[i:]...)
	m = append(m, r[j:]...)
	return m
}

func main() {
	for j := 0; j <= 32; j++ {
		a := make([]int, 0, j)
		for i := j; i >= 0; i-- {
			a = append(a, i)
		}
		fmt.Println(MergeSort(a))
	}
}
