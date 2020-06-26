package main

import "fmt"

func main() {
	for i := 65; i <= 65+27; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
