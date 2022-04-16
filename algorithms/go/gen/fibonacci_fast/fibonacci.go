package main

import "fmt"

func Fib(n int) uint64 {
	if n == 0 {
		return 0
	}
	m := [][]uint64{
		{1, 1},
		{1, 0},
	}
	p := 1
	for ; p*2 <= n; p *= 2 {
		m = squareMatrix(m)
	}
	curr, prev := m[0][0], m[0][1]
	for ; p < n; p++ {
		prev, curr = curr, prev+curr
	}
	return prev
}

func squareMatrix(m [][]uint64) [][]uint64 {
	s := make([][]uint64, 0, len(m))
	for i := 0; i < len(m); i++ {
		s = append(s, make([]uint64, len(m)))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(s); k++ {
				s[i][j] += m[i][k] * m[k][j]
			}
		}
	}
	return s
}

func main() {
	for i := 0; i <= 32; i++ {
		fmt.Printf("%d ", Fib(i))
	}
}
