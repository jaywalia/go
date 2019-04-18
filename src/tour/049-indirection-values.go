// 49 indirection values

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func absfunc(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := vertex{3, 4}

	fmt.Println(v.abs())
	fmt.Println(absfunc(v))

	p := &vertex{4, 3}
	fmt.Println(p.abs())
	fmt.Println(absfunc(*p))
}
