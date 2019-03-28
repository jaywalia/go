//017 pointers

package main
import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	v := 1

	fmt.Println("initial: ", v)

	zeroval(v)
	fmt.Println("zeroval :", v)

	zeroptr(&v)
	fmt.Println("zeroptr :", v)
}