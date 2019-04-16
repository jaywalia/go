// 59 errors

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// implement error interface { Error() string }
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"not working......",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
