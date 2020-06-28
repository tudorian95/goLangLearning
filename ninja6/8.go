package main

import (
	"fmt"
)

func main() {
	s := func() func() {
		return func() {
			fmt.Println("noice")
		}
	}
	s()()
}
