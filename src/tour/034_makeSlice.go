// 035 make a slice

package main

import "fmt"

func main() {

	// make allocates a zero array
	// and returns a slice
	a := make([]int, 5)
	fmt.Println(a)
	b := make([]int, 0, 5)
	fmt.Println(b)
	b = b[:cap(b)]
	fmt.Println(cap(b))
	fmt.Println(b)
	b = b[1:]
	fmt.Println(b)
}
