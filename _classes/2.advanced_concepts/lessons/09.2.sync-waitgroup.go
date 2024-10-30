package lessons

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	GoRoutine é a maneira que uma função pode ser colocada de maneira concorrente no Go, para ser executada com paralelismo ou não.
	Essas GoRoutines entram em processo dentro das `green threads` do Go, que são threads virtuais menores que as threads reais dos processadores

	Por padrão a Goroutine não faz o código esperar, por isso é necessário usar o sync.WaitGroup
*/

func GoRoutinesWithSyncWaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("(n cpu): ", runtime.NumCPU())
	fmt.Println("(n goroutines): ", runtime.NumGoroutine())

	// Executando com concorrencia no código
	fmt.Println("======= com goroutine e waitgroup =======")
	go loop01(&wg)
	go loop02(&wg)

	fmt.Println("(n goroutines): ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("=========================================")
	fmt.Println()
}

func loop01(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("loop1: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Println()
}

func loop02(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("loop2: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Println()
}
