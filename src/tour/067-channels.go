// 67 channels

package main

import "fmt"

func sum(a []int, c chan int) {
    s := 0
    for _, v := range a {
        s += v
    }
    c <- s
}

func main() {
    a := []int{1,2,3,4,5,6,7,8,9}
    c := make(chan int)

    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)

    x,y := <-c, <-c
    fmt.Println(x,y,x+y)
}