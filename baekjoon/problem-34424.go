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

   var a, b int
   fmt.Fscan(reader, &a, &b)

   fmt.Fprintf(writer, "%v\n", (a-1)*b)
}
