// 030 pointers

package main

import "fmt"

func main() {
	i, j := 42, 2701

	// pointer to i
	p := &i
	fmt.Println(*p)

	// updating the value pointed to
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
