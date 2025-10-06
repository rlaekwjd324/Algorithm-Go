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

   var f, t int
   fmt.Fscan(reader, &f, &t)
   
   for i := 0; i < t; i++ {
      var c string
      var n int
      fmt.Fscan(reader, &c, &n)
      
      switch c {
      case "+":
         f += n
      case "-":
         f -= n
      }
   }

   fmt.Fprintf(writer, "%v\n", f)
}
