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

   var a, b, c int
   fmt.Fscan(reader, &a, &b, &c)

   answer := a*3 + b*4 + c*5

   fmt.Fprintln(writer, answer)
}
