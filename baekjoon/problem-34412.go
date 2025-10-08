package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

    var a int
    fmt.Fscan(reader, &a)

    if a >= 13 {
       a++
    }

    fmt.Fprintf(writer, "%v\n", a)
}
