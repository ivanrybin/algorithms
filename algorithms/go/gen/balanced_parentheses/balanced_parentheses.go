package main

import "fmt"

// Parentheses generates balanced parentheses of len = n*2.
func Parentheses(n int) []string {
	return parentheses("", 0, 0, n)
}

func parentheses(s string, l, r, n int) []string {
	var ps []string
	if l == n && r == n {
		ps = append(ps, s)
	} else {
		if l < n {
			ps = append(ps, parentheses(s+"(", l+1, r, n)...)
		}
		if r < l {
			ps = append(ps, parentheses(s+")", l, r+1, n)...)
		}
	}
	return ps
}

func main() {
	for n := 1; n <= 4; n++ {
		fmt.Println(Parentheses(n))
	}
}
