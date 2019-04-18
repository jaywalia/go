// 59 cmd line flags

package main

import "flag"
import "fmt"

func main() {
	p := fmt.Println

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	p("word:", *wordPtr)
	p("numb:", *numbPtr)
	p("fork:", *boolPtr)
	p("svar:", svar)
	p("tail:", flag.Args())

}
