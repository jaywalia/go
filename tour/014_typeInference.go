// 14 type inference

package main

import "fmt"

func main() {
	v := 42 // change me!
	v = 42.000123

	fmt.Printf("v is of type %T %v\n", v, v)
}
