// 022 channels

package main

import "fmt"


func f(from string) {
	for i := 0; i < 20; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("go routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going anon")

	fmt.Scanln()
	fmt.Println("done")
}