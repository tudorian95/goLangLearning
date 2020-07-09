/*package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)// buffered channel
	c <- 42
	fmt.Println(<-c)
}*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//anonymous function
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
