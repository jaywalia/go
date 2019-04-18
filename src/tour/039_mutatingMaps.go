// 39 mutating maps

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 0
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// test if val exists
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
