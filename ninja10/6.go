package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	populate(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}
