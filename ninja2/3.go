package main

import "fmt"

const a int = 5

func main() {
	const b = "hi"
	fmt.Printf("%T\n %T\n", a, b)
}
