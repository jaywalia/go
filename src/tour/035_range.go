package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
