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
	for size := 1; size <= len(a)/2+1; size++ {
		// [l:l+size)[l+size:l+size*2)
		for l := 0; l < len(a); l += size * 2 {
			lr := min(l+size, len(a))
			rr := min(l+size*2, len(a))
			if lr != len(a) {
				copy(a[l:rr], merge(a[l:lr], a[lr:rr]))
			}
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
	for j := 0; j <= 16; j++ {
		a := make([]int, 0, j)
		for i := 0; i < j; i++ {
			a = append(a, i)
		}
		fmt.Println(MergeSort(a))
	}
}
