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
   
   v := 0.0
   for i := 0; i < n; i++ {
      var a float64
      fmt.Fscan(reader, &a)
      a *= 0.01
		v = 1.0 - (1.0 - v)*(1.0 - a)
      fmt.Fprintf(writer, "%.6f\n", v * 100)
   }
   
}
