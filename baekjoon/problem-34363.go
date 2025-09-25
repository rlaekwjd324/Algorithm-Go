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

   var s, d, t float64
   fmt.Fscan(reader, &s, &d, &t)
   
   a := s/3600*t
   b := d/5280

   if a >= b {
      fmt.Fprintf(writer, "%v\n", "MADE IT")
      return
   }
   fmt.Fprintf(writer, "%v\n", "FAILED TEST")
}
