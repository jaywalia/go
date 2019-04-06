// 008 named results

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	x, y := split(19)
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}
