package main

import "fmt"

func main() {
	permutations("Hello World")
}

// prints string permutations
func permutations (s string) {
	if( len(s) == 1 ) {
		fmt.Printf(s)
	} else {
		permutations(s[1:]);
	}
}

