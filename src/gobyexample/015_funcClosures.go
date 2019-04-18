//015_funcClosures

package main

import "fmt"

func intSeq() func() int {
	i := 0

	//nested func declarations not allowed

	// following is allowed
	increment := func () int {
		i++;
		return i;
	}

	// return anon function also works
	// return func() int {
	// 	i++ 
	// 	return i
	// }

	return increment
}

func main() {
	fmt.Println("hellllllloooooooooooo....")

	//nextInt is a function
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
}