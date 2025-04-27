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
   
   var z int
   fmt.Fscan(reader, &z)
   
   for i := 0; i < z; i++ {
      var x1, y1, z1 float64
      var x2, y2, z2 float64
      fmt.Fscan(reader, &x1, &y1, &z1, &x2, &y2, &z2)

      var m1, m2 float64
      m1 += (x1*y2 + y1*z2 + z1*x2)
      m2 += (x2*y1 + y2*z1 + z2*x1)
      if m1 > m2 {
         fmt.Fprintf(writer, "%v\n", "ADAM")
         continue
      }
      if m1 < m2 {
         fmt.Fprintf(writer, "%v\n", "GOSIA")
         continue
      }
      fmt.Fprintf(writer, "%v\n", "=")
   }
}
