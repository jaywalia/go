package main

import (
	"fmt"
)

func PrintPermutations(rs []rune, pi int, ps *[]string) {

	if pi == (len(rs) - 1) {
		s := string(rs)
		*ps = append(*ps, s)
		// fmt.Println(s)
	}

	for i := pi; i < len(rs); i++ {
		rs[pi], rs[i] = rs[i], rs[pi]
		PrintPermutations(rs, pi+1, ps)
		rs[pi], rs[i] = rs[i], rs[pi]
	}

}

func main() {
	s := "abc"
	rs := []rune(s)
	ps := []string{}
	PrintPermutations(rs, 0, &ps)
	fmt.Println(ps)

}
