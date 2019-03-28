package main

import "fmt"

func main() {


    if 7 % 2 == 0 {
        fmt.Println("7 is even?")
    } else {
        fmt.Println("7 is odd")
	}
	
	if 8 % 4 == 0 {
		fmt.Println("8 / 4")
	}

	if n := 9; n < 0 {
		fmt.Println(n, " is negative")
	} else if n < 10 {
		fmt.Println( n, " has one digit")
	} else {
		fmt.Println(n, "has many diigts")
	}
}