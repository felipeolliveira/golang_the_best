package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, name string) {
	ticker := time.NewTicker(250 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println(name, "tick")
		case <-ctx.Done():
			fmt.Println("Finishing due to:", ctx.Err().Error())
			fmt.Println(name, "Received a signal to finish. Exiting...")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go worker(ctx, "Worker 1")
	time.Sleep(5 * time.Second)
	go worker(ctx, "Worker 2")

	time.Sleep(20 * time.Second)
	fmt.Println("main exiting")
}
