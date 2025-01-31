package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// O padrão de convergencia faz com que um canal especifico fique responsável por concentrar todos os valores de outros canals
func ConvergenceRobPike() {
	work1 := work("American", "Yes")
	work2 := work("Latino", "Si")

	works := converge(work1, work2)

	for i := 0; i < 15; i++ {
		fmt.Println(<-works)
	}
}

func work(name, message string) chan string {
	c := make(chan string)

	go func(c chan string, n string, m string) {
		for {
			c <- fmt.Sprintf("Work %v say: %v", n, m)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(c, name, message)

	return c
}

func converge(channelA, channelB chan string) chan string {
	wg := sync.WaitGroup{}
	channelC := make(chan string)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range channelA {
			channelC <- v
		}
	}()
	go func() {
		defer wg.Done()
		for v := range channelB {
			channelC <- v
		}
	}()

	return channelC
}
