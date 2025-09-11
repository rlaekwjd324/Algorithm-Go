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

   var x, y, z float64
   fmt.Fscan(reader, &x, &y, &z)

   v := x*0.25+y*0.25+z*0.5
   if v >= 90 {
      fmt.Fprintf(writer, "%v\n", "A")
      return
   }
   if v >= 80 {
      fmt.Fprintf(writer, "%v\n", "B")
      return
   }
   if v >= 70 {
      fmt.Fprintf(writer, "%v\n", "C")
      return
   }
   if v >= 60 {
      fmt.Fprintf(writer, "%v\n", "D")
      return
   }
   fmt.Fprintf(writer, "%v\n", "F")
}
