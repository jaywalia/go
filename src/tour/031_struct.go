// 031 structs

package main

import (
	"errors"
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

type line struct {
	start, end point
}

type circle struct {
	center point
	radius float64
}

func (l line) slope() (float64, error) {
	n := l.end.y - l.start.y
	d := l.end.x - l.start.x
	if d == 0 {
		return 0, errors.New("math: can't divide by zero")
	}

	return n / d, nil
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	p := point{x: 3, y: 5}
	fmt.Println(p)

	s := point{x: 0, y: 0}
	e := point{x: 1, y: 1}

	l := line{start: s, end: e}
	fmt.Println(l)
	fmt.Println(l.slope())

	l2 := line{start: point{x: 0, y: 0}, end: point{x: 0, y: 2}}
	fmt.Println(l2)
	fmt.Println(l2.slope())

	c := circle{p, 2.0}
	fmt.Println(c.area())

}
