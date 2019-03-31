// 34 : rate limiting

package main

import "fmt"
import "time"

func main() {

	// requests := make(chan int, 5)
	// for i := 1; i <= 5; i++ {
	// 	requests <- i
	// }
	// close(requests)

	// limiter := time.Tick(1000 * time.Millisecond)

	// for req := range requests {
	// 	<-limiter
	// 	fmt.Println("request", req, time.Now())
	// }

	// allow three bursts
	burstLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}

	close(burstRequests)

	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request", req, time.Now())
	}

}
