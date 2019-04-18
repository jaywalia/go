// 66 go routines

package main

import "fmt"
import "time"

func thinkBeforeYouSayIt(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go thinkBeforeYouSayIt("world!")
	thinkBeforeYouSayIt("hello")
}
