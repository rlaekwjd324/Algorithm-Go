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

   var n int
   fmt.Fscan(reader, &n)

   count := 0
   for i := 0; i < n; i++ {
      var t, a int
      fmt.Fscan(reader, &t, &a)
      if t == 1 {
         count += a
      } else {
         count -= a
         if count < 0 {
            fmt.Fprintln(writer, "Adios")
            return
         }
      }
   }
   fmt.Fprintln(writer, "See you next month")
}
