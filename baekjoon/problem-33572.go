package main

import (
   "bufio"
   "fmt"
   "os"
   "math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
   defer writer.Flush()
   
   var n, m float64
   fmt.Fscan(reader, &n, &m)

   t := int(math.Ceil(m / n))
   for i := 0; i < int(n); i++ {
      var a int
      fmt.Fscan(reader, &a)

      if t >= a {
         fmt.Fprintln(writer, "OUT")
         return
      } 
   }
   fmt.Fprintln(writer, "DIMI")
}
