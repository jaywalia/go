// 36 pic

package main

import "fmt"

//https://stackoverflow.com/questions/7703251/slice-of-slices-types
func Pic(dx, dy int) [][]uint8 {

	//fmt.Printf("%d, %d\n", dx, dy)

	pic := make([][]uint8, dy, dy*dx)
	//fmt.Println(pic)

	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8((i + j) / 2)
		}
	}

	//fmt.Println(pic)

	return pic
}

func main() {
	fmt.Println(Pic(3, 4))
}
