package main

import "fmt"

type newtype int

var x newtype
var y int

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
