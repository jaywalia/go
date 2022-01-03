package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"time"
)

var philosphers = []string{"Jay", "Li", "Lilac", "Coral", "Kevin"}
var diners sync.WaitGroup

const servings = 3
const think = time.Second / 100
const eat = time.Second / 100

func dine(p string, leftFork, rightFork *sync.Mutex) {
	// random sleep function
	h := fnv.New64a()
	h.Write([]byte(p))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		// fmt.Println("Fixed duration: ", t/2)
		d := t/2 + time.Duration(rg.Int63n(int64(t)))
		// fmt.Println("Sleeping for ", d)
		time.Sleep(d)
	}

	// dine
	fmt.Println(p, " is seated.")

	for s := servings; s > 0; s-- {
		fmt.Println(p, " wants a serving.")
		// pick up forks
		leftFork.Lock()
		rightFork.Lock()
		// eat
		fmt.Println(p, " is eating.")
		rSleep(eat)
		// put down forks
		rightFork.Unlock()
		leftFork.Unlock()
		// think
		fmt.Println(p, " is thinking.")
		rSleep(think)
	}

	fmt.Println(p, " ate all his servings.")
	diners.Done()
	fmt.Println(p, " left the table.")
}

func main() {
	fmt.Println("Welcome to Annual Philosphers' Dinner!")
	diners.Add(5)

	// first philospher's left fork
	fork0 := &sync.Mutex{}
	forkLeft := fork0

	for i, p := range philosphers {
		var forkRight *sync.Mutex
		if i != len(philosphers) {
			forkRight = &sync.Mutex{}
		} else {
			// last philospher's right fork is first philsopher's left fork
			forkRight = fork0
		}
		go dine(p, forkLeft, forkRight)
		forkLeft = forkRight
	}

	diners.Wait()
	fmt.Println("Hope you all enjoyed your food! See you next year.")
}
