package main

import "fmt"

type newtype int

var x newtype

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 42
	fmt.Println(x)
}
