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

    var n, c, p int
    fmt.Fscan(reader, &n, &c, &p)

    if c*p >= n {
       fmt.Fprintf(writer, "%v\n", "yes")
       return
    }

    fmt.Fprintf(writer, "%v\n", "no")
}
