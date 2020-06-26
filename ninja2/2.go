package main

import "fmt"

func main() {
	x := 7
	y := 8
	a := fmt.Sprintf("%v", x == y)
	b := fmt.Sprintf("%v", x <= y)
	c := fmt.Sprintf("%v", x >= y)
	d := fmt.Sprintf("%v", x > y)
	e := fmt.Sprintf("%v", x < y)
	f := fmt.Sprintf("%v", x != y)
	fmt.Println(a, b, c, d, e, f)
}
