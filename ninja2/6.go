package main

import "fmt"

const (
	_ = iota
	b = 2020 + iota
	c = 2020 + iota
	d = 2020 + iota

	e = 2020 + iota
)

func main() {
	fmt.Println(b, c, d, e)
}
