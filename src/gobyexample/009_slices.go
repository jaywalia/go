package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty : ", s)
	
	s[0] = "asdf"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: " , s)
	fmt.Println("get: ", s[0])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "ele", "fant")
	fmt.Println("append : ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice[low:high] doesn't include high
	l := s[2:5]
	fmt.Println("slice 1 :", l)

	l = s[:5]
	fmt.Println("slice 2 :", l)

	t := []string{"one", "two", "Three"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD :", twoD)
	 
}