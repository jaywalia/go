// 050 interfaces.go

package main

import (
	"fmt"
	"math"
)

type abser interface {
	abs() float64
}

type iFloat float64
type vertex struct {
	x, y float64
}

func (f iFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a abser
	f := iFloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f
	a = &v

	// following own't compile
	// a = v
	fmt.Println(a.abs())

}
