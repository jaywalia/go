// 033 slices

package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func testSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func main() {
	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)

	sl := a[1:2]
	fmt.Println(sl)

	// slice literal
	s2 := []int{3, 4, 5}
	fmt.Println(s2)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(st)

	// slice defaults : low bound : zero
	s3 := a[:3]
	fmt.Println(s3)

	// slice defaults : high bound : len
	s4 := a[1:]
	fmt.Println(s4)

	// slice defaults
	s5 := a[:]
	fmt.Println(s5)

	testSlice()
}
