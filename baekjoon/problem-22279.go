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
   
   sum := 0.0
   for i := 0; i < n; i++ {
      var q, l float64
      fmt.Fscan(reader, &q, &l)
      sum += q*l
   }
   
   fmt.Fprintf(writer, "%.3f\n", sum)
}
