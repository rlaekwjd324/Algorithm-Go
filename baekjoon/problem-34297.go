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

   var x, y, z int
   fmt.Fscan(reader, &x, &y, &z)

   fmt.Fprintf(writer, "%v\n", x*z)
}
