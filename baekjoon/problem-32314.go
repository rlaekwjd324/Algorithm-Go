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

   var a int
   fmt.Fscan(reader, &a)
   var w, v int
   fmt.Fscan(reader, &w, &v)
   
   if a <= w/v {
      fmt.Fprintln(writer, 1)
      return
   }
   fmt.Fprintln(writer, 0)
}
