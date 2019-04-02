// 58 command line args

package main

import "fmt"
import "os"

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]
	p := fmt.Println

	p(argsWithProg)
	p(argsWithoutProg)
	p(arg)
}
