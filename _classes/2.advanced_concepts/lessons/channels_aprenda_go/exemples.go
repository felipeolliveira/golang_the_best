package main

import "fmt"

func Examples() {
	c := make(chan int)

	go loopWithChannels(10, c)
	printByReceiverChannel(c)
}

func loopWithChannels(i int, s chan<- int) {
	for v := range i {
		s <- v
	}

	// É importante que apos o uso da comunicação entre canais
	// Haja o close(chan) para que o canal não fique aberto e ocorra o fatal error: all goroutines are asleep - deadlock!
	close(s)
}

func printByReceiverChannel(r <-chan int) {
	for v := range r {
		fmt.Println("iteração por canais", v)
	}
}
