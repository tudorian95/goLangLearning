package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	wg.Add(101)
	//this will create a race condition,as we don't
	//set rules about priority over the one variable that multiple
	//functions are using concurrently
	for i := 0; i <= 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			fmt.Println(v)
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
}
