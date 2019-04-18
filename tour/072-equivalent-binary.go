// 072 equivalent-binary

package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func walk(t *Tree, ch chan int) {

}

func same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go walk(t1, c1)
	go walk(t2, c2)

	for {
		n1 := <-c1
		n2 := <-c2

		if n1 != n2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("solve tree problem")
	t1 := New(10)
	t2 := New(10)

	fmt.Println(t1)
	fmt.Println(t2)
}
