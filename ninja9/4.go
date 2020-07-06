package main

import (
	"fmt"
	"sync"
)

var mtx sync.Mutex

func main() {
	counter := 0
	for i := 0; i < 100; i++ {
		go func() {
			//no more race conditions, as the mutex assures
			//only one routine can access the shared variable at a time
			mtx.Lock()
			v := counter
			v++
			fmt.Println(v)
			counter = v
			mtx.Unlock()
		}()
	}
}
