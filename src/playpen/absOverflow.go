package main

import (
	"fmt"
	"math"
)

func main() {
	// var n int8 = math.MinInt8
	// abs := func(x int) int {
	// 	fmt.Printf("x bef : %d\n", x)
	// 	if x < 0 {
	// 		x = -x
	// 	}
	// 	fmt.Printf("x aft : %d\n", x)
	// 	return x
	// }
	//fmt.Printf("%d %d \n", n, abs(n))
	// t := (abs(math.MinInt) == math.MinInt)
	// var x int8 = math.MinInt8
	// var nx int8 = -math.MinInt8
	// fmt.Printf("%d %d\n", x, nx)

	abs := func(x int) uint {
		var ux uint
		if x < 0 {
			ux = uint(-x)
		} else {
			ux = uint(x)
		}
		return ux
	}

	var n int = math.MinInt
	fmt.Printf("%d %d \n", n, abs(n))
}
