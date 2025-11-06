package main

import (
    "bufio"
    "fmt"
    "os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
    defer writer.Flush()

    var s, t string
    fmt.Fscan(reader, &s, &t)
	
	fmt.Fprintf(writer, "%v%v\n", string(s[0]), t)
}
