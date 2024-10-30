package main

import (
	"fmt"
	"sync"
)

// O padrão de convergencia faz com que um canal especifico fique responsável por concentrar todos os valores de outros canals
func ConvergenceTodd() {
	even := make(chan int)
	odd := make(chan int)
	converge := make(chan int)

	go sender(even, odd)
	go receiver(even, odd, converge)

	for v := range converge {
		fmt.Println("Valor recebido no converge:", v)
	}
}

func sender(even, odd chan int) {
	for n := range 100 {
		if n%2 == 0 {
			even <- n
		} else {
			odd <- n
		}
	}
	close(even)
	close(odd)
}

func receiver(even, odd, converge chan int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range even {
			converge <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range odd {
			converge <- v
		}
	}()

	wg.Wait()
	close(converge)
}
