package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i, ch := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i+2) * time.Second) // 1 -> 2 segundos, 2 -> 3 segundos
				ch <- i + 1
			}
		}(i, ch)
	}

	// Sem select, o canal que demora mais tempo acaba bloqueando o outro de realizar a sua função
	// Assim, os valores printados podem acabar se repetindo porque no momento do print, o outro canal ainda não enviou um valor
	// e o valor que ele tem acesso pode acabar sendo o mesmo da iteração anterior
	for i := 0; i < 5; i++ {
		v1 := <-chans[0]
		fmt.Println("Got a value, on channel 1", v1)
		v2 := <-chans[1]
		fmt.Println("Got a value, on channel 2", v2)
	}

	fmt.Println("=====================================")

	// Com o select, o canal que demora mais tempo não bloqueia o outro de realizar a sua função
	for i := 0; i < 5; i++ {
		select {
		case v1 := <-chans[0]:
			fmt.Println("Got a value, on channel 1", v1)
		case v2 := <-chans[1]:
			fmt.Println("Got a value, on channel 2", v2)
		}
	}
}
