package main

import "fmt"

func main() {
	a := []int{12, 3, 2, 14, 342, 54, 6534, 78, 6, 88}
	y := []int{56, 57, 58, 59, 60}
	a = append(a, 52)
	fmt.Println(a)
	a = append(a, 53, 54, 55)
	fmt.Println(a)
	a = append(a, y[:]...) //y... is the same as y[:]...
	fmt.Println(a)
}
