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

   var n, s int
   fmt.Fscan(reader, &n)
   fmt.Fscan(reader, &s)
   
   c := 0
   for i := 1; i < n; i++ {
      var t int
      fmt.Fscan(reader, &t)
      
      if s >= t {
         c++
      }
   }

   fmt.Fprintf(writer, "%v\n", c)
}
