// 57 stringer

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	a := person{"Von Neumann", 41}
	z := person{"Zaphod Beebleborx", 9001}

	fmt.Println(a, z)
}
