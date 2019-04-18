// 24 loops

package main

import "fmt"
import "math"

func sqrt(x float64) float64 {
	z := 1.0
	// z = x, x/2 etc.
	//z = x / 2
	zi := z
	eps := 0.0000000000001
	delta := 0.0
	//for i := 0; i < 10; i++ {
	for {
		zi = z
		z -= (z*z - x) / (2 * z)

		delta = math.Abs(z - zi)
		//fmt.Println(z)
		fmt.Println("[zi:", zi, "][z", z, "][delta :", delta, "]")
		if delta < eps {
			break
		}
		//fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(256))
}
