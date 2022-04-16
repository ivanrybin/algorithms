package main

import "fmt"

func Fib(n int) uint64 {
	if n == 0 {
		return 0
	}
	prev, curr := uint64(0), uint64(1)
	for i := 1; i < n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func main() {
	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", Fib(i))
	}
}
