package main

import "fmt"

// PowerSet
func PowerSet(a []int) [][]int {
	ps := [][]int{{}}
	for i := 0; i < len(a); i++ {
		r := len(ps)
		for _, s := range ps[:r] {
			x := append([]int{}, s...)
			x = append(x, a[i])
			ps = append(ps, x)
		}
	}
	return ps
}

func main() {
	fmt.Println(PowerSet([]int{1, 2, 3, 4}))
}
