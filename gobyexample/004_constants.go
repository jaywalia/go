package main

import "fmt"
import "math"

const s string = "the quick brown fox jumped over the lazy dog"

func main() {
    fmt.Println(s)

    const pi = 3.14

    const n = 500000000
    const d = 3e20 / n

    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
}