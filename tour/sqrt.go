package main

import (
	"fmt"
	"math"
)

func NewtSqrt(x float64) float64 {
	var z float64
	var d float64 // delta
	var nz float64 // new z

	d = 0.000000000000001
	// init z to x
	z = x + 2 * d
	nz = x

	var i int
	i = 0

//	for i < 10  {
	for ( z - nz ) > d {
		z = nz
		nz = z - (( (z*z) - x ) / (2*z) )
		i = i + 1
	}

	fmt.Println("# of loops : ", i)

	return z
}


func main() {
	var x float64
	x = 2.0

	for ( x < 10 ) {
		fmt.Println("math vs newt\n", x, math.Sqrt(x), NewtSqrt(x))
		x++
	}
}

