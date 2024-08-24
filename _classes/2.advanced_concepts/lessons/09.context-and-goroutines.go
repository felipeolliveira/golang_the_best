package lessons

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

/*
	- Go routines são uma forma de executar funções de forma concorrente ou paralela, caso o processador tenha mais de um núcleo.
		Pararelismo: Execução de várias tarefas ao mesmo tempo.
		Concorrência: Execução de várias tarefas em um único núcleo, porém de forma alternada. Quando uma tarefa está ociosa, outra pode ser executada, quando a primeira tarefa volta a ser executada, a segunda é pausada. Quando uma tarefa de fato é finalizada, a próxima tarefa é executada e assim por diante.

	- Por padrão, o Go executa o código de forma sequencial, ou seja, linha por linha. Para executar uma função de forma concorrente, basta adicionar a palavra-chave "go" antes da chamada da função.
		`go func() { ... }()`
*/

func requestGoogle(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	res, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)
	defer res.Body.Close()
}

func runNormal() {
	start := time.Now()
	for range 10 {
		requestGoogle(nil)
	}
	elapsed := time.Since(start)
	fmt.Println("Elapsed time: ", elapsed)
	fmt.Println("==============")
}

func runWithConcurrence() {
	start := time.Now()
	for range 10 {
		go requestGoogle(nil) // Executa a função de forma concorrente
	}
	fmt.Println("Concurrence, but no wait, so no output")
	elapsed := time.Since(start)
	fmt.Println("Concurrence Elapsed time: ", elapsed)
	fmt.Println("==============")
}

/*
- Todos os retornos de funções executadas de forma concorrente são descartados. Para obter um retorno, é necessário utilizar channels, wait groups.
*/

func runWithWaitGroup() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n) // Adiciona n goroutines ao wait group, ou seja, n funções que serão executadas de forma concorrente
	for range n {
		// Foi necessário passar o wait group como parâmetro para a função, para que a mesma possa chamar o método Done() ao finalizar
		go requestGoogle(&wg)
	}
	wg.Wait() // Aguarda todas as goroutines finalizarem
	elapsed := time.Since(start)
	fmt.Println("Wait Group Elapsed time: ", elapsed)
	fmt.Println("==============")
}

/*
  - Contexts são utilizados para controlar o que acontece dentro das goroutines. É possível cancelar a execução de uma goroutine, por exemplo.
*/

func runWithContext() {
	const n = 10
	const timetoutInSeconds = 10
	const responseInSeconds = 5

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()
	ctx, cancelTimeout := context.WithTimeout(ctx, timetoutInSeconds*time.Second)
	defer cancelTimeout()

	testServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(responseInSeconds * time.Second)
			fmt.Println("Server return")
		},
	))

	for range n {
		go func(ctx context.Context) {
			defer wg.Done()

			request, err := http.NewRequestWithContext(ctx, "GET", testServer.URL, nil)

			if err != nil {
				panic(err)
			}

			res, err := http.DefaultClient.Do(request)

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("Context timeout")
					return
				}
				panic(err)
			}

			defer res.Body.Close()
		}(ctx)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Context Elapsed time: ", elapsed)
	fmt.Println("==============")
}

func ContextAndGoRoutines() {
	runNormal()
	runWithConcurrence()
	runWithWaitGroup()
	runWithContext()
}
