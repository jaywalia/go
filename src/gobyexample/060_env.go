// 60 env

package main

import "os"
import "strings"
import "fmt"

func main() {
	p := fmt.Println
	f := fmt.Printf

	os.Setenv("FOO", "1")
	p("FOO :", os.Getenv("FOO"))
	p("BAR :", os.Getenv("BAR"))

	p()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		f("%s >>>>> %s\n\n", pair[0], pair[1])
	}

}
