package main

import "fmt"

// QuickSort sort in place.
//  O(nlogn) best case
//  O(n^2)   worst case
func QuickSort(a []int) {
	quickSort(0, len(a)-1, a)
}

// quickSort recursive sort.
func quickSort(l, r int, a []int) {
	if l < r {
		m := partition(l, r, a)
		quickSort(l, m-1, a)
		quickSort(m+1, r, a)
	}
}

// partition
// m shows position of a[m]: a[m] <= p
// - a[i] <= p: i <= m
// - a[i]  > p: i  > m
func partition(l, r int, a []int) int {
	p := a[r]
	m := l
	for i := l; i < r; i++ {
		if a[i] <= p {
			a[i], a[m] = a[m], a[i]
			m++
		}
	}
	a[r], a[m] = a[m], a[r]
	return m
}

func main() {
	xs := []int{1}
	QuickSort(xs)
	fmt.Println(xs)

	xs = []int{2, 1}
	QuickSort(xs)
	fmt.Println(xs)

	xs = []int{0, 1, 1, 1, 0, 1, 1, 1, 0}
	QuickSort(xs)
	fmt.Println(xs)

	xs = []int{1, 0, 2, 3, 5, 4, 8}
	QuickSort(xs)
	fmt.Println(xs)

	xs = []int{1, 7, 10, -1, 0, 5, 6, 9, 8, 3, 2, 4}
	QuickSort(xs)
	fmt.Println(xs)
}
