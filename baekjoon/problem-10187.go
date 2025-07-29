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
   
   for i := 0; i < n; i++ {
      var a, b float64
      fmt.Fscan(reader, &a, &b)
      
      if a < b {
         t := a
         a = b
         b = t
      }
      d := a/b
      if d <= 1.61803399*1.01 && d >= 1.61803399*0.99 {
         fmt.Fprintf(writer, "%v\n", "golden")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "not")
   }
}
