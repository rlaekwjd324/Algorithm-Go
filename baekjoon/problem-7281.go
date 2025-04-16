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
   p := 0
   for i := 0; i < n; i++ {
      var t, m int
      fmt.Fscan(reader, &t, &m)
      
      if m == 1 {
         if max < (t - p) {
            max = t - p
         }
         p = t
      }
   }
   fmt.Fprintf(writer, "%v\n", max)
}
