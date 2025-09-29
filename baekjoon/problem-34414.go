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

   var t int
   fmt.Fscan(reader, &t)
   
   for i := 0; i < t; i++ {
      var a int
      fmt.Fscan(reader, &a)
      if a < 48 {
         fmt.Fprintf(writer, "%v\n", "False")
         return
      }
   }
   fmt.Fprintf(writer, "%v\n", "True")
}
