package main

import "fmt"

func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	for i := 0; i < 16; i++ {
		a := make([]int, 0, i)
		for j := i; j > 0; j-- {
			a = append(a, j)
		}
		fmt.Println(BubbleSort(a))
	}
}
