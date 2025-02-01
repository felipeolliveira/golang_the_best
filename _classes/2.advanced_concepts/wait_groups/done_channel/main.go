package main

import (
	"fmt"
	"time"
)

// O Done Channel é uma forma primitiva de esperar uma GoRoutine acabar,
// pois a leitura do channel força a main (goRoutine principal) esperar pelo resultado,
// já que há um writer e um reader desse channel
func main() {
	done := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("GoRoutine terminou")
		done <- 0
	}()

	<-done
}
