package lessons

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
  Race Condition:
  É quando duas ou mais goroutines acessam uma variável ao mesmo tempo e tentam alterar o seu valor.
  Isso pode causar problemas de concorrência, pois o valor final da variável pode não ser o esperado.

  Solução:
  - Atomic: pacote que fornece funções para manipular variáveis de forma segura.
    - É recomendado para variáveis simples como int, string, etc, por ser mais performático que Mutex para dados imples.
  - Mutex(Mutual Exclusion): pacote que fornece funções para criar locks e garantir que apenas uma goroutine acesse a variável por vez.
    - É recomendado para variáveis complexas como structs, maps, etc, mas é mais lento que Atomic.

	Go tem uma ferramenta para saber se o código tem race condition:
	$ go run -race <file>.go
*/

func AtomicAndMutexRaceCondition() {
	problem()
	solutionWithAtomic()
	solutionWithMutex()
}

func problem() {
	count := 0
	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)

		go func() {
			for range 1000 {
				count++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("problem -> %v\n", count)
}

func solutionWithAtomic() {
	var count int64 = 0
	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)

		go func() {
			for range 1000 {
				atomic.AddInt64(&count, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("solution: atomic -> %v\n", count)
}

func solutionWithMutex() {
	var count int64 = 0
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for range 10 {
		wg.Add(1)

		go func() {
			for range 1000 {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("solution: mutex -> %v\n", count)
}
