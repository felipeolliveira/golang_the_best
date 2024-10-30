package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println(counter(50))
	fmt.Println(counterWithAtomic(50))
	fmt.Println(counterWithMutex(50))
}

func counter(v int) int {
	wg := sync.WaitGroup{}
	var count int

	for range v {
		wg.Add(1)

		go func() {
			value := count
			runtime.Gosched()
			count = value + 1
			wg.Done()
		}()
	}

	wg.Wait()

	return count
}

func counterWithAtomic(v int) int64 {
	wg := sync.WaitGroup{}
	var count int64

	for range v {
		wg.Add(1)

		go func() {
			runtime.Gosched()
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	return count
}

func counterWithMutex(v int) int {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	var count int

	for range v {
		wg.Add(1)

		go func() {
			m.Lock()
			count++
			m.Unlock()
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()

	return count
}
