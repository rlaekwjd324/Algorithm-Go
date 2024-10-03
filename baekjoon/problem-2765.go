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

   i := 0
   for {
      var r,c,t float64
      fmt.Fscan(reader, &r, &c, &t)
      
      if c == 0 {
         return
      }
      i++
      m := r*c*3.1415927/12.0/5280.0
      v := m / (t/3600.0)
      fmt.Fprintln(writer, fmt.Sprintf("Trip #%v: %.2f %.2f", i, m, v))
   }

}
