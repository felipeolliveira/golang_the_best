package lessons

import "fmt"

/*
	GoRoutine é a maneira que uma função pode ser colocada de maneira concorrente no Go, para ser executada com paralelismo ou não.
	Essas GoRoutines entram em processo dentro das `green threads` do Go, que são threads virtuais menores que as threads reais dos processadores

	Por padrão a Goroutine não faz o código esperar, por isso é necessário usar o sync.WaitGroup
*/

func GoRoutines() {
	// Executando de maneira sequencial no código
	fmt.Println("======= sem goroutine =======")
	loop1()
	loop2()
	fmt.Println("=============================")
	fmt.Println()

	// Executando com concorrencia no código
	fmt.Println("======= com goroutine no loop1 =======")
	go loop1()
	loop2()
	fmt.Println("======================================")
	fmt.Println()
}

func loop1() {
	fmt.Printf("loop1: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Println()
}

func loop2() {
	fmt.Printf("loop2: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Println()
}
