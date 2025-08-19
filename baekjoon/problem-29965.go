package main

import (
	"fmt"
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

   var n int
   fmt.Fscan(reader, &n)
   
   m, j := 0.0, 0.0
   mc, jc := 0.0, 0.0
   for i := 0; i < n; i++ {
      var a string
      var p float64
      fmt.Fscan(reader, &a, &p)

      switch a {
      case "M":
         m += p
         mc++
      case "J":
         j += p
         jc++
      }
   }
   ma, ja := 0.0, 0.0
   if mc > 0.0 {
      ma = m/mc
   }
   if jc > 0.0 {
      ja = j/jc
   }
   if ma > ja {
      fmt.Fprintf(writer, "%v\n", "M")
      return
   }
   if ma < ja {
      fmt.Fprintf(writer, "%v\n", "J")
      return
   }
   fmt.Fprintf(writer, "%v\n", "V")
}
