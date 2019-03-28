//016 recursion

package main
import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}

func fib(n int) int{
	if n < 0 {
		return 0
	} else if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println("factorial(5): ", factorial(5))
}