// 012 functions

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func addAdd(a, b, c int) int {
	return add(a, add(b,c))
}

func main() {
	r := add(1,2)
	fmt.Println("1+2 =", r)

	r = addAdd(1,2,3)
	fmt.Println("1+2+3=", r)
}