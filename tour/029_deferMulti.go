// 29 stacking defers

package main

import "fmt"

func main() {
	fmt.Println("counting")

	// defer functions are pushed onto a stack.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
