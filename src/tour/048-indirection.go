// 048 indirection

package main

import "fmt"

type vertex struct {
	x, y float64
}

func (v *vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func scalef(v *vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := vertex{3, 4}
	v.scale(2)
	scalef(&v, 10)

	p := &vertex{4, 3}
	p.scale(3)
	scalef(p, 8)

	fmt.Println(v, p)
}
