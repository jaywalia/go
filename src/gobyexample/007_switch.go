package main

import "fmt"
import "time"

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")

	// looks like break is default
	switch i {
	case 1 : fmt.Println("one")
	case 2 : fmt.Println("two")
	case 3 : fmt.Println("three")
	}

	// commmas for multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("weeeeeekendddddddddd")
	default :
		fmt.Println("workkkkkkkk")
	}

	// switch without expression (if)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("still morning...")
	default : 
		fmt.Println("parrrrty... ")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type){
		case bool : fmt.Println("boool")
		case int : fmt.Println("int")
		default : fmt.Printf("Unknonw type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}