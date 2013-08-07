package main

import (
    "fmt"
    "strconv"
)

func main() {
    // rune is int32
    c := 'b'
    fmt.Println(c)

    // atoi and itoa in strconv
    value, _ := strconv.Atoi("155")
    fmt.Println(value)
}

