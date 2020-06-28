package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	vi := []int{4, 5, 6, 7}
	defer fmt.Println(foo(xi...))
	fmt.Println(bar(vi))
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
