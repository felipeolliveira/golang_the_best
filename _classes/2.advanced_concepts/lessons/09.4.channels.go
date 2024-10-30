package lessons

import "fmt"

// Channels são usados para trocar informações entre goroutines

func Channels() {
	simpleChannel()
	bufferedChannel()
}

func simpleChannel() {
	channel := make(chan int)

	go func() {
		// Está passando um valor para o canal por dentro da goroutine
		channel <- 42
	}()

	// está retirando o valor de dentro do canal
	fmt.Println(<-channel)
}

// Os canais com buffers não são recomendados em casos de uso geral
func bufferedChannel() {
	// Indica que o canal precisa de uma quantidade máxima no buffer
	channel := make(chan int, 1)

	go func() {
		// Está passando um valor para o canal por dentro da goroutine
		channel <- 42

		// Por conta do tamanho do buffer ser 1, apenas a primeira mensagem é passada para o channel
		// O próximo valor é ignorado
		channel <- 43
	}()

	// está retirando o valor de dentro do canal
	fmt.Println(<-channel)
}
