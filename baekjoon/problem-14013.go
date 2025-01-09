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

   var x, y float64
   var n int
   fmt.Fscan(reader, &x, &y, &n)
   
   for i := 0; i < n; i++ {
      var z float64
      var q string
      fmt.Fscan(reader, &z, &q)
      
      var a float64
      if q == "A" {
         a = z * y/x
      } else {
         a = z * x/y
      }

      if a == 0 {
         fmt.Fprintln(writer, "0.000000000000000")
         continue
      }
      fmt.Fprintln(writer, a)
   }
}
