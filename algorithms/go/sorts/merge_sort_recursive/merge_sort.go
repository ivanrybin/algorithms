package main

import "fmt"

// MergeSort
// T: O(nlogn) best/worst case
// M: O(nlogn) best/worst case
func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	return merge(MergeSort(a[:len(a)/2]), MergeSort(a[len(a)/2:]))
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
	fmt.Println(MergeSort([]int{1}))
	fmt.Println(MergeSort([]int{2, 1}))
	fmt.Println(MergeSort([]int{1, 1, 1, 1, 1, 0}))
	fmt.Println(MergeSort([]int{1, 7, 10, -1, 0, 5, 6, 9, 8, 3, 2, 4}))
}
