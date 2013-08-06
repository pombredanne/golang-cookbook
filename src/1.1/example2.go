package main

import (
    "fmt"
    "strings"
)

func stringJoin(s1 string, s2 string) string {
    joinStr := ""
    for _, char := range s1 {
        if strings.ContainsRune(s2, char) {
            joinStr += string(char)
        }
    }
    return joinStr
}

func main() {
    s1 := "abcdefg"
    s2 := "testajdflkajdfke"
    fmt.Println(stringJoin(s1, s2))
}

