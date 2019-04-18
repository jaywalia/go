//010_underscore.go

package main

import "fmt"

func main() {
	x := []int { 1, 2, 3, 4, 5, 6}
	sum := 0

	// for i, val := range x {
	// 	sum += val
	// }

	// _ is used to throw away the value
	// also look at the example above, 
	// golang doesn't allow to create variable that is not used
	// from go book
	// The blank identifier may be used whenever syntax requires a variable name 
	// but program logic does not, 
	// for instance to discard an unwanted loop index 
	// when we require only the element value.
	for _, val := range x {
		sum += val
	}

	fmt.Println("sum: ", sum)
}