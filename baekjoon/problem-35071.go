package main

import (
    "bufio"
    "fmt"
    "os"
    "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

    var d float64
    fmt.Fscan(reader, &d)

    v := math.Pi*d/2.0-d

    fmt.Fprintf(writer, "%v\n", v)
}
