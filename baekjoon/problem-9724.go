package main

import (
	"fmt"
	"bufio"
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
      var a, b float64
      fmt.Fscan(reader, &a, &b)

      c := 0
      for j := 1.0; j <= math.Cbrt(b); j++ {
         if j*j*j >= a {
            c++
         }
      }

      fmt.Fprintf(writer, "Case #%v: %v\n", i+1, c)
   }
}
