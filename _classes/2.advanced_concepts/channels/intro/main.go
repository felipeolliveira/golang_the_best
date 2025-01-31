package main

import "fmt"

// Channels
//
// Quando o channel está preparado?
//
// Writer
// - Quando tem alguem pronto pra ouvir
// - Buffer, caso exista, não esteja cheio
// Reader
// - Quando tem alguém pronto pra escrever
// - Caso tenha dados no Buffer
// - caso seja close e tenha buffer, pode ser quantas vezes quiser
//
// Contraints
// - Apenas leitura..........<-chan T
// - Apenas escrita..........chan<- T
// - Qualquer operação.......chan T

func main() {
	ch := make(chan int, 1)

	ch <- 0
	n, ok := <-ch
	fmt.Println("Got value:", n, "| is this channel is closed?", ok)

	close(ch)

	n, ok = <-ch
	fmt.Println("Got value:", n, "| is this channel is closed?", ok)
}
