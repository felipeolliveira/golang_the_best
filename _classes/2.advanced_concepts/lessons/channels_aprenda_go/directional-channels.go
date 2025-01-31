package main

import "fmt"

/*
* Channels são comunicações bidirecionais, ou seja, qualquer um dos lados do canal pode receber e enviar informações
* Porém, há canais unidirecionais que podem apenas fazer uma função de cada vez:
*
* - `chain int` : Cria um canal bidirecional para transportar inteiros
* - `chain<- int`: Cria um canal que pode apenas **mandar** inteiros
* - `<-chain int`: Cria um canal que pode apenas **receber** inteiros
*
* */

func DirectionChannels() {
	c := make(chan int)

	// As funções que recebem um canal unidirecionais, aceitam canais feitos de forma bidirecional, o contrario não é possível
	go sendChannel(c)
	receiveChannel(c)

	assertionOrConversion()
}

// Caso tente usar o <-chan, apontará o erro na proxima linha que o canal seria para receber e não poderia enviar o valor 42
// o erro é `invalid operation`
func sendChannel(s chan<- int) {
	s <- 42
}

// O mesmo acontece aqui caso tente usar um chan<-, mostrará que é um canal para receber e não enviar
func receiveChannel(r <-chan int) {
	fmt.Println("O valor do canal é:", <-r)
}

func assertionOrConversion() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Println("------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// Canais específicos para os gerais(bidirecional) não funcionam na conversão
	// fmt.Println("------")
	// fmt.Printf("c\t%T\n", (chan int)(cr)) // type error
	// fmt.Printf("c\t%T\n", (chan int)(cs)) // type error
}
