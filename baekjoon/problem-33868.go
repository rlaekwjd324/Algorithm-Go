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
   
   max := 1
   min := 5000
   for i := 0; i < n; i++ {
      var t, b int
      fmt.Fscan(reader, &t, &b)
      
      if max < t {
         max = t
      }
      if min > b {
         min = b
      }
   }

   fmt.Fprintf(writer, "%v\n", (min*max)%7+1)
}
