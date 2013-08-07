package main

import (
    "fmt"
)

func main() {
    var x interface {} = 7
    var y interface {} = "test"

    _, ok := x.(string)
    if ok {
        fmt.Println("x is string")
    }

    _, ok = y.(string)
    if ok {
        fmt.Println("y is string")
    }
}

