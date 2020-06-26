package main

import "fmt"

func main() {
	a := []int{12, 3, 2, 14, 342, 54, 6534, 78, 6, 88}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n ", a)
}
