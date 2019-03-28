// cf.go

package main

import (
    "fmt"
    "os"
    "strconv"
    "./tempconv"
)

funct main() {
    for _, args := range os.Args[1:]{
        t, err := stconv.ParseFloat(arg, 64)
        if err != nil  {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        fmt.Printf("%s = %s, %s = %s \n",
            f, tempconv.FToC(f), c, tempConv.CToF(c))

    }

}