// 045 methods funcs

package main

import (
	"fmt"
	"math"
)

type McFloatyFloat float64

func (m McFloatyFloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
