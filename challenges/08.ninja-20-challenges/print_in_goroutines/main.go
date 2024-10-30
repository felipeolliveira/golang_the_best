package main

import (
	"fmt"
	"sync"
)

func main() {
	printInGoroutines(100)
}

func printInGoroutines(i int) {
	wg := sync.WaitGroup{}
	wg.Add(i)

	for value := range i {
		go func() {
			fmt.Println("Goroutine: ", value+1)
			wg.Done()
		}()

	}

	wg.Wait()
}
