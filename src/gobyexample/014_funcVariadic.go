// 014_funcVariadic

package main
import "fmt"

func sum(nz ...int) int {
	fmt.Print(nz, " ")
	t := 0
	for _, n := range nz {
		t += n
	}
	fmt.Println(t)
	return t
}

func main() {
	sum(1,2)
	sum(1,2, 3)
	
	x := []int{1,2,3,4,5}
	sum(x...)
}