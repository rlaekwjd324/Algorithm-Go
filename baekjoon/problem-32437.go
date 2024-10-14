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

   var n int
   fmt.Fscanln(reader, &n)

   a := 1
   b := 1
   for i := 0; i < n; i++ {
      temp := a
      a = b
      b = temp+b
   }
   
   fmt.Fprintln(writer, a)
}
