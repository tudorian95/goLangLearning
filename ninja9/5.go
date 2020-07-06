package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64
	for i := 0; i < 100; i++ {
		go func() {
			//even easier than with the mutex
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
		}()
	}
}
