// 033 slices

package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)

	sl := a[1:2]
	fmt.Println(sl)
}
