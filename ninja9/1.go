package main

import (
	"fmt"
	"sync"
)

var wt sync.WaitGroup

func main() {
	wt.Add(2)
	go func() {
		fmt.Println("aloha")
		wt.Done()
	}()

	go func() {
		fmt.Println("mi amigo")
		wt.Done()
	}()
	wt.Wait()
}
