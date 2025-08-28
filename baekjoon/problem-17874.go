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

   var n, h, v int64
   fmt.Fscan(reader, &n, &h, &v)
   
   if h < n-h {
      h = n-h
   }
   if v < n-v {
      v = n-v
   }

   fmt.Fprintf(writer, "%v\n", h*v*4)
}
