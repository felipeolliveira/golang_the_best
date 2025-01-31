package main

import (
	"fmt"
	"runtime"
	"time"
)

// Pode ocorrer um leak de goroutines quando o numero de goroutines for maior que
// a quantidade de valores passados para o canal.
// Como é unbuffered, as goroutines ficam bloqueadas porque foi lido apenas o primeiro valor: <-ch
func unbuffered() {
	const numberOfRoutines = 30
	ch := make(chan int)

	for i := 0; i <= numberOfRoutines; i++ {
		go func(ch chan<- int, i int) {
			ch <- i
			fmt.Println("I am goroutine:", i)
		}(ch, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Got value:", <-ch)

	for {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Go routines running:", runtime.NumGoroutine()-1)
	}
}

// Agora com um canal tem buffer (espaços na memória), do tamanho correspondente a quantidade de goroutines,
// nenhuma delas ficará blocked, pois vão preencher todos os espaços de buffer do channel:
// - [1] [2] [3] ... [30]
// - cada goroutine consegue escrever no buffer e nenhuma fica bloqueada
func buffered() {
	const numberOfRoutines = 30
	ch := make(chan int, numberOfRoutines)

	for i := 0; i <= numberOfRoutines; i++ {
		go func(ch chan<- int, i int) {
			ch <- i
			fmt.Println("I am goroutine:", i)
		}(ch, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Got value:", <-ch)

	for {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("Go routines running:", runtime.NumGoroutine()-1)
	}
}

func main() {
	// unbuffered()
	buffered()
}
