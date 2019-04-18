// 042 function closures

package main

import "fmt"

func adder() func(int) int {
	s := 0
	return func(x int) int {
		s += x
		return s
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(i, pos(i), neg(-2*i))
	}
}
