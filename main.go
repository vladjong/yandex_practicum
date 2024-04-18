package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Test")
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		i := i
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
