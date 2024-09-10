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
      var a, b, d float64
      fmt.Fscan(reader, &a, &b)

      d = a-b
      fmt.Fprintln(writer, fmt.Sprintf("%.1f", math.Abs(d)))
   }
}
