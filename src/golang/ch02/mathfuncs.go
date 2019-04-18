package main

import "fmt"

func main() {
	fmt.Printf("Hello World! \n")
	//fmt.Printf("fib (5) :%d \n",fib(9))
	//fmt.Printf("fib2 (5) :%d \n",fib2(9))
	fmt.Printf("gcd(4, 8) : %d \n", gcd(16,32))
}

// 1 1 2 3 5 8 13
func fib(n int) int {
	if ( n == 0 || n == 1 ) {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fib2(n int) int {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}