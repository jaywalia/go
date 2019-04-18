// 64 images

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Print(m.At(i, j).RGBA())
		}
		fmt.Println()
	}
}
