package main

import (
	"fmt"
	"time"
)

func takesTooLong(ch chan<- int) {
	time.Sleep(30 * time.Second)
	ch <- 10
}

func takesNotSoLong(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 2
}

// É possível parar um processo antes das goroutines darem deadlock por falta de leitores nos channels
// Colocando um `stopper` criado com um time.After para ser selecionado no select
//
// Essa maneira é antiga, antes da criação do Context
func main() {
	// a função time.After retorna um canal que manda um valor de time depois da quantidade de tempo especificado
	// Nesse caso, o stop acontece antes das goroutines darem deadlock!
	stop := time.After(31 * time.Second)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go takesTooLong(ch1)   // Demora 10 segundos
	go takesNotSoLong(ch2) // Demora 2 segundos

	for {
		select {
		case <-ch1:
			fmt.Println("too long finished")
		case <-ch2:
			fmt.Println("not so long finished")
		case <-stop:
			fmt.Println("this job is taking too long to finish... aborting")
			return

		}
	}
}
