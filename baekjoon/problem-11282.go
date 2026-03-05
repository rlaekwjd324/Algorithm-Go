package main

import "fmt"

func main() {
    var N int
    fmt.Scan(&N)

    fmt.Printf("%c\n", rune(0xAC00 + N - 1))
}
