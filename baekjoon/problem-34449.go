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
   
   var d, w, n float64
   fmt.Fscan(reader, &d, &w, &n)
   
   v := w*n

   if v > 3.14159*d {
      fmt.Fprintf(writer, "%v\n", "NO")
      return
   }

   fmt.Fprintf(writer, "%v\n", "YES")
}
