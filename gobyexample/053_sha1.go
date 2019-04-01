// 53 sha hashes

package main

import "crypto/sha1"
import "fmt"

func main() {
	p := fmt.Println
	f := fmt.Printf
	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	p(s)
	f("%x\n", bs)
}
