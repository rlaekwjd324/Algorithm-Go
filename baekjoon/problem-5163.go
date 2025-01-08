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

   var n int
   fmt.Fscan(reader, &n)
   
   for i := 0; i < n; i++ {
      var b int
      var w float64
      fmt.Fscan(reader, &b, &w)
      
      if b == 0 {
         fmt.Fprintf(writer, "Data Set %v:\n%s\n\n", i+1, "No")
         continue
      }

      sum := float64(0)
      for j := 0; j < b; j++ {
         var r float64
         fmt.Fscan(reader, &r)

         if r == float64(0) {
            continue
         }
         sum += float64(4.0/3.0*math.Pow(r/10.0, 3)*math.Pi)
      }

      if w <= sum {
         fmt.Fprintf(writer, "Data Set %v:\n%s\n\n", i+1, "Yes")
      } else {
         fmt.Fprintf(writer, "Data Set %v:\n%s\n\n", i+1, "No")
      }
   }
}
