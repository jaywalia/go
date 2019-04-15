// 043 fib

package main

import "fmt"

func fib() func() int {
	f, f0, f1 := 0, 0, 1
	return func() int {
		f, f0, f1 = f0, f1, f0+f1
		return f
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
