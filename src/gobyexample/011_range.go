// 011_range.go

package main

import "fmt"

func main() {
	n := [] int {2, 3, 4}
	s := 0

	// neglecting the index
	for _, v := range n {
		s += v;
	}

	fmt.Println("sum :", s)

	// using the index
	for i, v := range n {
		if v == 3 {
			fmt.Println("index :", i)
		}
	}

	kvs := map[string] string {"a": "apple", "b" : "banana"}
	for k, v := range kvs {
		fmt.Printf( "%s -> %s \n", k, v)
	}

	for k := range kvs {
		fmt.Println("key :", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}