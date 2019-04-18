// 70 select

package main

import "fmt"

func fib(c, q chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("quit it now!")
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//send message on quit channel
		q <- 0
	}()

	fib(c, q)
}
