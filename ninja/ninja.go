package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
}
