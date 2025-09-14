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

   var l, a, h int
   fmt.Fscan(reader, &l, &a, &h)
   
   for i := 0; i < l; i++ {
      var hi int
      fmt.Fscan(reader, &hi)

      d := hi-h
      if d > a {
         fmt.Fprintf(writer, "%v\n", "BUG REPORT")
         return
      }
      h = hi
   }
   fmt.Fprintf(writer, "%v\n", "POSSIBLE")
}
