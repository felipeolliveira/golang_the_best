package main

import (
	"fmt"
	"time"
)

/*
* Select permite com que você espere multiplas operações de canal para ter um controle de fluxo das goroutines e canais.
* Funciona semelhante ao switch
* */

func Select() {
	example1()
}

func example1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
