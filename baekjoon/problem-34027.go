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
      var a float64
      fmt.Fscan(reader, &a)
      
      if math.Sqrt(a)-float64(int(math.Sqrt(a))) == 0.0 {
         fmt.Fprintf(writer, "%v\n", 1)
         continue
      }
      
      fmt.Fprintf(writer, "%v\n", 0)
   }
}
