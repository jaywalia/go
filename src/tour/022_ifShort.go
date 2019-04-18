package main

import (
	"fmt"
	"math"
)

// normal way
func pow2(x, n, lim float64) float64 {
	v := math.Pow(x, n)

	if v < lim {
		return v
	}
	return lim
}

// shorted statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// vi s not visible here
	return lim
}

func main() {

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
