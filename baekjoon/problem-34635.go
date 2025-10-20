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
   
   var r, g, b, cr, cg, cb, crg, cgb int
   fmt.Fscan(reader, &r, &g, &b, &cr, &cg, &cb, &crg, &cgb)
   
   r -= cr
   g -= cg
   b -= cb

   c := 0
   if r > 0 {
      crg -= r
      c += r
   }
   if b > 0 {
      cgb -= b
      c += b
   }
   if crg < 0 || cgb < 0 {
      fmt.Fprintf(writer, "%v\n", -1)
      return
   }
   
   if g > 0 {
      c += g
      if crg+cgb < g {
         fmt.Fprintf(writer, "%v\n", -1)
         return
      }
   }

   fmt.Fprintf(writer, "%v\n", c)
}
