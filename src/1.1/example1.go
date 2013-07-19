// process each character in a time
package main

import (
    "fmt"
    "strings"
)

func main() {
    s := "hello world!"
    // range is useful for slice, map, string, channel 
    for index, char := range s {
        fmt.Printf("%v %c\n", index,  char)
    }

    strPlus := func(r rune) rune {
        return r + 1
    }
    fmt.Println(strings.Map(strPlus, s))

    s = "你好"
    for _, char := range s {
        fmt.Printf("%c", char)
    }
}

