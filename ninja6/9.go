package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func cb(f func(a int, b int) int, c int, d int) int {
	return f(c, d) + 2
}

func main() {
	r := cb(sum, 5, 7)

	fmt.Println(r)
}

/*
func main() {

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5, 6})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
*/
