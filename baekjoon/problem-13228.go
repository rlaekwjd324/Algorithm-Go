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
   
   var t int
   fmt.Fscan(reader, &t)
   for i := 0; i < t; i++ {
      var x1, y1, f1, x2, y2, f2 float64
      fmt.Fscan(reader, &x1, &y1, &f1, &x2, &y2, &f2)

      a := math.Abs(x1-x2)+math.Abs(y1-y2)+f1+f2

      fmt.Fprintln(writer, int(a))
   }
}
