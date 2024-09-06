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

   var t int64
   fmt.Fscan(reader, &t)
   for i := int64(0); i<t; i++ {
      var n, c int64
      fmt.Fscan(reader, &n)
      c = n*n
      fmt.Fprintln(writer, c)
   }
}
