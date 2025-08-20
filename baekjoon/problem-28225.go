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
   var f float64
   fmt.Fscan(reader, &n, &f)
   
   min := 10000.0
   idx := 0
   for i := 1; i <= n; i++ {
      var x, v float64
      fmt.Fscan(reader, &x, &v)

      d := (f-x)/v
      if min > d {
         min = d
         idx = i
      }
   }
   
   fmt.Fprintf(writer, "%v\n", idx)
}
