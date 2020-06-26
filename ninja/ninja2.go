package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
}
