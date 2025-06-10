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

   var d1, d2, d3 float64
   fmt.Fscan(reader, &d1, &d2, &d3)
   
   a := (d1+d2+d3)/2.0-d3
   b := (d1+d2+d3)/2.0-d2
   c := (d1+d2+d3)/2.0-d1
  
   if a <= 0 || b <= 0 || c <= 0 {
      fmt.Fprintf(writer, "%v\n", -1)
      return
   }
   fmt.Fprintf(writer, "%v\n", 1)
   fmt.Fprintf(writer, "%.1f %.1f %.1f\n", a, b, c)
}
