// 63 signals

package main

import "os"
import "fmt"
import "os/signal"
import "syscall"

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signals")
	<-done
	fmt.Println("exiting")
}
