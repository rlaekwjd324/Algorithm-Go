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

   var a, b, c int
   fmt.Fscan(reader, &a, &b, &c)
   var d, e, f int
   fmt.Fscan(reader, &d, &e, &f)
   
   one := a+b*2+c*3
   two := d+e*2+f*3

   if one == two {
      fmt.Fprintln(writer, 0)
      return
   }
   if one > two {
      fmt.Fprintln(writer, 1)
      return
   }
   fmt.Fprintln(writer, 2)
}
