package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("the remainder of %v divided by 4 is %v\n", i, i%4)		
		}

	}
}
