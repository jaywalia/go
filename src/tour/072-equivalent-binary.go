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

func walkTree(t *Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go walkTree(t1, c1)
	go walkTree(t2, c2)

	// compare items in c1 to c2
	for i := range c1 {
		j := <-c2
		if i != j {
			return false
		}
	}

	// items left in c2
	// then not equal
	for _ = range c2 {
		return false
	}
	// else equal

	return true
}

func main() {
	fmt.Println("solve tree problem")
	t1 := New(1)
	t2 := New(3)

	fmt.Println(t1)
	fmt.Println(t2)

	if same(t1, t2) {
		fmt.Println("same trees")
	} else {
		fmt.Println("different trees")
	}
}
