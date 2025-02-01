package main

import (
	"fmt"
	"sync"
)

// Mutexes (Mutual Exclusions) é a pratica de sincronização de dados,
// impedindo que uma valor seja alterado em uma condição de corrida,
// garantindo que o valor seja manipulado de forma correta por apenas um agente por vez
// - mu := sync.Mutex
// - mu.Lock() // Reservou o acesso para a rotina em execução
// - mu.Unlock() // Liberou o acesso a outras rotinas
//
// É possivel usar channels para bloquear as manipulações, usando um buffer channel de 1.
// Sendo assim, o channel fica blocked e aguarda a saída de uma ação para realizar outra
// - ch := make(chan, struct{}, 1)
// - ch <- struct{}{} // Reservou o espaço do buffer e parou outros acessos
// - <-ch // Leu o valor struct{} do buffer e liberou o acesso para outra rotina

var (
	sema        = make(chan struct{}, 1)
	balance int = 0
	mutex   sync.Mutex
)

func withBufferedChannel() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Bloqueando o valor. Deposito de:", i)
			sema <- struct{}{}

			balance += i

			fmt.Println("Liberando o valor. Total:", balance)
			<-sema
		}()
	}

	wg.Wait()
}

func withMutex() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Bloqueando o valor. Deposito de:", i)
			mutex.Lock()

			balance += i

			fmt.Println("Liberando o valor. Total:", balance)
			mutex.Unlock()
		}()
	}

	wg.Wait()
}

func main() {
	withBufferedChannel()
	fmt.Println("====================")
	withMutex()
}
