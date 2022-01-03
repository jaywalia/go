package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	parts     sync.WaitGroup
	partList  = []string{"Front Wheel", "Rear Wheel", "Handle Bar", "Frame", "Drive", "Brakes", "Seat"}
	noOfBikes = 5
)

func worker(part string) {
	log.Println(part, ": begins")
	time.Sleep(time.Duration(rand.Int63n((1e6))))
	log.Println(part, ": finished")
	parts.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for b := 0; b < noOfBikes; b++ {
		// log start
		log.Println("Assembling Bike: ", b+1)
		// set up how many parts to wait for
		parts.Add(len(partList))
		// hand out parts to build
		for _, part := range partList {
			go worker(part)
		}
		// wait for workers to finish their parts
		parts.Wait()
		// log completion
		log.Println("Assembly Completed", b+1)
	}
}
