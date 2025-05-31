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
   
   max := 0
   for i := 0; i < n; i++ {
      var a int
      fmt.Fscan(reader, &a)
      if max > a {
         continue
      }
      max = a
      fmt.Fprintf(writer, "%v ", a)
   }
}
