// 013 functions many returns

package main

import "fmt"

func vals() (int, int) {
	x := [5]int {1,2,3,4,5}
	return x[2], x[3]
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}