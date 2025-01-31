package main

import (
	"log"
	"time"
)

func main() {
	stopper := time.After(10 * time.Second)
	ticker := time.NewTicker(250 * time.Millisecond)

	log.Println("Start")
	defer log.Println("Finish")

	for {
		select {
		case <-ticker.C:
			log.Println("Updating...")
		case <-stopper:
			log.Println("Stop")
			return
		}
	}
}
