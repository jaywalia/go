// 030 range channels

package main

import "fmt"

func main() {
    q := make( chan string, 2)

    q <- "one"
    q <- "two"
    close(q)

    for elem := range q {
        fmt.Println(elem)
    }
}