// concurrency producer consumer
// https://www.golangprograms.com/go-language/concurrency.html

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type Consumer struct {
	msgs *chan int
}

func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

func (c *Consumer) consume() {
	fmt.Println("consumer : started")
	for {
		msg := <-*c.msgs
		fmt.Println("consume : received", msg)
	}
}

type Producer struct {
	msgs *chan int
	done *chan bool
}

func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

func (p *Producer) produce(max int) {
	fmt.Println("produce: started")
	for i := 0; i < max; i++ {
		fmt.Println("produce: Sending ", i)
		*p.msgs <- i
	}

	*p.done <- true
	fmt.Println("produce: Done")
}

func main() {
	// profile flags
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")

	// get the maximum number of messages from flags
	max := flag.Int("n", 5, "defines the number of messages")

	flag.Parse()

	// utilize the max num of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	// CPU Profile
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var msgs = make(chan int)  // channel to send messages
	var done = make(chan bool) // channel to control when production is done

	// Start a goroutine for Produce.produce
	go NewProducer(&msgs, &done).produce(*max)

	// Start a goroutine for Consumer.consume
	go NewConsumer(&msgs).consume()

	// Finish the program when the production is done
	<-done

	// Memory Profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
