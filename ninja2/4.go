package main

import "fmt"

const a int = 5

func main() {
	const b int = a << 1
	fmt.Printf("%d\n %b\n %#x\n", a, a, a)
	fmt.Printf("%d\n %b\n %#x", b, b, b)
}
