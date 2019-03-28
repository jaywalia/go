//024_channel buffering

package main

import "fmt"

func main() {
    messages := make (chan string, 2)

    go func() { messages <- "hello" }()
    go func() { messages <- "world" }()

    fmt.Println(<- messages)
    fmt.Println(<- messages)
}