package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done!\n", id)
}

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %s starting\n", url)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %s done!\n", url)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}

	fmt.Println("goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("goroutines:", runtime.NumGoroutine())
}
