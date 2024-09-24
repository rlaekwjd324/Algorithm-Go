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

   var s int
   fmt.Fscan(reader, &s)

   for {
      var v int
      fmt.Fscan(reader, &v)
      if s <= v {
         fmt.Fprintln(writer, s)
         return
      }
      s += v
   }
}
