package main

import (
	"fmt"
	"time"
)

// É uma boa pratica especificar o que o channel pode fazer, ler ou escrever
func runner(ch chan<- int, id int) {
	fmt.Printf("Iniciada Go Routine runner: %d\n", id)
	time.Sleep(5 * time.Second)

	for {
		fmt.Printf("[Go Routine Runner %d] Tentando enviar...\n", id)
		ch <- id
		fmt.Printf("[Go Routine Runner %d] Enviado com sucesso\n", id)
	}
}

func MoreExamples() {
	ch := make(chan int, 4)

	go runner(ch, 1)
	go runner(ch, 2)
	go runner(ch, 3)
	go runner(ch, 4)
	go runner(ch, 5)

	for {
		num := <-ch
		fmt.Printf("[MAIN] Received a value: %d\n", num)

		time.Sleep(3 * time.Second)
	}
}
